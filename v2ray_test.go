package core_test

import (
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	. "github.com/robovpn/v2fly_core"
	"github.com/robovpn/v2fly_core/app/dispatcher"
	"github.com/robovpn/v2fly_core/app/proxyman"
	"github.com/robovpn/v2fly_core/common"
	"github.com/robovpn/v2fly_core/common/net"
	"github.com/robovpn/v2fly_core/common/protocol"
	"github.com/robovpn/v2fly_core/common/serial"
	"github.com/robovpn/v2fly_core/common/uuid"
	"github.com/robovpn/v2fly_core/features/dns"
	"github.com/robovpn/v2fly_core/features/dns/localdns"
	_ "github.com/robovpn/v2fly_core/main/distro/all"
	"github.com/robovpn/v2fly_core/proxy/dokodemo"
	"github.com/robovpn/v2fly_core/proxy/vmess"
	"github.com/robovpn/v2fly_core/proxy/vmess/outbound"
	"github.com/robovpn/v2fly_core/testing/servers/tcp"
)

func TestV2RayDependency(t *testing.T) {
	instance := new(Instance)

	wait := make(chan bool, 1)
	instance.RequireFeatures(func(d dns.Client) {
		if d == nil {
			t.Error("expected dns client fulfilled, but actually nil")
		}
		wait <- true
	})
	instance.AddFeature(localdns.New())
	<-wait
}

func TestV2RayClose(t *testing.T) {
	port := tcp.PickPort()

	userID := uuid.New()
	config := &Config{
		App: []*anypb.Any{
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
		Inbound: []*InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortRange: net.SinglePortRange(port),
					Listen:    net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address: net.NewIPOrDomain(net.LocalHostIP),
					Port:    uint32(0),
					NetworkList: &net.NetworkList{
						Network: []net.Network{net.Network_TCP, net.Network_UDP},
					},
				}),
			},
		},
		Outbound: []*OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&outbound.Config{
					Receiver: []*protocol.ServerEndpoint{
						{
							Address: net.NewIPOrDomain(net.LocalHostIP),
							Port:    uint32(0),
							User: []*protocol.User{
								{
									Account: serial.ToTypedMessage(&vmess.Account{
										Id: userID.String(),
									}),
								},
							},
						},
					},
				}),
			},
		},
	}

	cfgBytes, err := proto.Marshal(config)
	common.Must(err)

	server, err := StartInstance(FormatProtobuf, cfgBytes)
	common.Must(err)
	server.Close()
}
