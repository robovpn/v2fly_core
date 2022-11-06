package v2binding

import (
	"google.golang.org/protobuf/types/known/anypb"

	core "github.com/robovpn/v2fly_core"
	"github.com/robovpn/github.com/robovpn/v2fly_core/app/commander"
	"github.com/robovpn/github.com/robovpn/v2fly_core/app/dispatcher"
	"github.com/robovpn/github.com/robovpn/v2fly_core/app/instman"
	"github.com/robovpn/github.com/robovpn/v2fly_core/app/instman/command"
	"github.com/robovpn/github.com/robovpn/v2fly_core/app/proxyman"
	"github.com/robovpn/github.com/robovpn/v2fly_core/app/router"
	"github.com/robovpn/github.com/robovpn/v2fly_core/common/net"
	"github.com/robovpn/github.com/robovpn/v2fly_core/common/serial"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/main/distro/all"
	"github.com/robovpn/github.com/robovpn/v2fly_core/proxy/blackhole"
	"github.com/robovpn/github.com/robovpn/v2fly_core/proxy/dokodemo"
)

type bindingInstance struct {
	started  bool
	instance *core.Instance
}

var binding bindingInstance

func (b *bindingInstance) startAPIInstance() {
	bindConfig := &core.Config{
		App: []*anypb.Any{
			serial.ToTypedMessage(&instman.Config{}),
			serial.ToTypedMessage(&commander.Config{
				Tag: "api",
				Service: []*anypb.Any{
					serial.ToTypedMessage(&command.Config{}),
				},
			}),
			serial.ToTypedMessage(&router.Config{
				Rule: []*router.RoutingRule{
					{
						InboundTag: []string{"api"},
						TargetTag: &router.RoutingRule_Tag{
							Tag: "api",
						},
					},
				},
			}),
		},
		Inbound: []*core.InboundHandlerConfig{
			{
				Tag: "api",
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortRange: net.SinglePortRange(10999),
					Listen:    net.NewIPOrDomain(net.AnyIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address:  net.NewIPOrDomain(net.LocalHostIP),
					Port:     uint32(10999),
					Networks: []net.Network{net.Network_TCP},
				}),
			},
		},
		Outbound: []*core.OutboundHandlerConfig{
			{
				Tag:           "default-outbound",
				ProxySettings: serial.ToTypedMessage(&blackhole.Config{}),
			},
		},
	}
	bindConfig = withDefaultApps(bindConfig)

	instance, err := core.New(bindConfig)
	if err != nil {
		panic(err)
	}
	err = instance.Start()
	if err != nil {
		panic(err)
	}
	b.instance = instance
}

func withDefaultApps(config *core.Config) *core.Config {
	config.App = append(config.App, serial.ToTypedMessage(&dispatcher.Config{}))
	config.App = append(config.App, serial.ToTypedMessage(&proxyman.InboundConfig{}))
	config.App = append(config.App, serial.ToTypedMessage(&proxyman.OutboundConfig{}))
	return config
}

func StartAPIInstance(basedir string) {
	if binding.started {
		return
	}
	binding.started = true
	binding.startAPIInstance()
}
