package scenarios

import (
	"os"
	"runtime"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/anypb"

	core "github.com/robovpn/v2fly_core"
	"github.com/robovpn/v2fly_core/app/log"
	"github.com/robovpn/v2fly_core/app/proxyman"
	"github.com/robovpn/v2fly_core/common"
	clog "github.com/robovpn/v2fly_core/common/log"
	"github.com/robovpn/v2fly_core/common/net"
	"github.com/robovpn/v2fly_core/common/protocol"
	"github.com/robovpn/v2fly_core/common/serial"
	"github.com/robovpn/v2fly_core/common/uuid"
	"github.com/robovpn/v2fly_core/proxy/dokodemo"
	"github.com/robovpn/v2fly_core/proxy/freedom"
	"github.com/robovpn/v2fly_core/proxy/vmess"
	"github.com/robovpn/v2fly_core/proxy/vmess/inbound"
	"github.com/robovpn/v2fly_core/proxy/vmess/outbound"
	"github.com/robovpn/v2fly_core/testing/servers/tcp"
	"github.com/robovpn/v2fly_core/testing/servers/udp"
	"github.com/robovpn/v2fly_core/transport/internet"
	"github.com/robovpn/v2fly_core/transport/internet/domainsocket"
	"github.com/robovpn/v2fly_core/transport/internet/headers/http"
	"github.com/robovpn/v2fly_core/transport/internet/headers/wechat"
	"github.com/robovpn/v2fly_core/transport/internet/quic"
	tcptransport "github.com/robovpn/v2fly_core/transport/internet/tcp"
)

func TestHTTPConnectionHeader(t *testing.T) {
	tcpServer := tcp.Server{
		MsgProcessor: xor,
	}
	dest, err := tcpServer.Start()
	common.Must(err)
	defer tcpServer.Close()

	userID := protocol.NewID(uuid.New())
	serverPort := tcp.PickPort()
	serverConfig := &core.Config{
		Inbound: []*core.InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortRange: net.SinglePortRange(serverPort),
					Listen:    net.NewIPOrDomain(net.LocalHostIP),
					StreamSettings: &internet.StreamConfig{
						TransportSettings: []*internet.TransportConfig{
							{
								Protocol: internet.TransportProtocol_TCP,
								Settings: serial.ToTypedMessage(&tcptransport.Config{
									HeaderSettings: serial.ToTypedMessage(&http.Config{}),
								}),
							},
						},
					},
				}),
				ProxySettings: serial.ToTypedMessage(&inbound.Config{
					User: []*protocol.User{
						{
							Account: serial.ToTypedMessage(&vmess.Account{
								Id: userID.String(),
							}),
						},
					},
				}),
			},
		},
		Outbound: []*core.OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&freedom.Config{}),
			},
		},
	}

	clientPort := tcp.PickPort()
	clientConfig := &core.Config{
		Inbound: []*core.InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortRange: net.SinglePortRange(clientPort),
					Listen:    net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address: net.NewIPOrDomain(dest.Address),
					Port:    uint32(dest.Port),
					NetworkList: &net.NetworkList{
						Network: []net.Network{net.Network_TCP},
					},
				}),
			},
		},
		Outbound: []*core.OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&outbound.Config{
					Receiver: []*protocol.ServerEndpoint{
						{
							Address: net.NewIPOrDomain(net.LocalHostIP),
							Port:    uint32(serverPort),
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
				SenderSettings: serial.ToTypedMessage(&proxyman.SenderConfig{
					StreamSettings: &internet.StreamConfig{
						TransportSettings: []*internet.TransportConfig{
							{
								Protocol: internet.TransportProtocol_TCP,
								Settings: serial.ToTypedMessage(&tcptransport.Config{
									HeaderSettings: serial.ToTypedMessage(&http.Config{}),
								}),
							},
						},
					},
				}),
			},
		},
	}

	servers, err := InitializeServerConfigs(serverConfig, clientConfig)
	common.Must(err)
	defer CloseAllServers(servers)

	if err := testTCPConn(clientPort, 1024, time.Second*2)(); err != nil {
		t.Error(err)
	}
}

func TestDomainSocket(t *testing.T) {
	if runtime.GOOS == "windows" || runtime.GOOS == "android" {
		t.Skip("Not supported on windows and android")
		return
	}
	tcpServer := tcp.Server{
		MsgProcessor: xor,
	}
	dest, err := tcpServer.Start()
	common.Must(err)
	defer tcpServer.Close()

	const dsPath = "/tmp/ds_scenario"
	os.Remove(dsPath)

	userID := protocol.NewID(uuid.New())
	serverPort := tcp.PickPort()
	serverConfig := &core.Config{
		Inbound: []*core.InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortRange: net.SinglePortRange(serverPort),
					Listen:    net.NewIPOrDomain(net.LocalHostIP),
					StreamSettings: &internet.StreamConfig{
						Protocol: internet.TransportProtocol_DomainSocket,
						TransportSettings: []*internet.TransportConfig{
							{
								Protocol: internet.TransportProtocol_DomainSocket,
								Settings: serial.ToTypedMessage(&domainsocket.Config{
									Path: dsPath,
								}),
							},
						},
					},
				}),
				ProxySettings: serial.ToTypedMessage(&inbound.Config{
					User: []*protocol.User{
						{
							Account: serial.ToTypedMessage(&vmess.Account{
								Id: userID.String(),
							}),
						},
					},
				}),
			},
		},
		Outbound: []*core.OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&freedom.Config{}),
			},
		},
	}

	clientPort := tcp.PickPort()
	clientConfig := &core.Config{
		Inbound: []*core.InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortRange: net.SinglePortRange(clientPort),
					Listen:    net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address: net.NewIPOrDomain(dest.Address),
					Port:    uint32(dest.Port),
					NetworkList: &net.NetworkList{
						Network: []net.Network{net.Network_TCP},
					},
				}),
			},
		},
		Outbound: []*core.OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&outbound.Config{
					Receiver: []*protocol.ServerEndpoint{
						{
							Address: net.NewIPOrDomain(net.LocalHostIP),
							Port:    uint32(serverPort),
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
				SenderSettings: serial.ToTypedMessage(&proxyman.SenderConfig{
					StreamSettings: &internet.StreamConfig{
						Protocol: internet.TransportProtocol_DomainSocket,
						TransportSettings: []*internet.TransportConfig{
							{
								Protocol: internet.TransportProtocol_DomainSocket,
								Settings: serial.ToTypedMessage(&domainsocket.Config{
									Path: dsPath,
								}),
							},
						},
					},
				}),
			},
		},
	}

	servers, err := InitializeServerConfigs(serverConfig, clientConfig)
	common.Must(err)
	defer CloseAllServers(servers)

	if err := testTCPConn(clientPort, 1024, time.Second*2)(); err != nil {
		t.Error(err)
	}
}

func TestVMessQuic(t *testing.T) {
	tcpServer := tcp.Server{
		MsgProcessor: xor,
	}
	dest, err := tcpServer.Start()
	common.Must(err)
	defer tcpServer.Close()

	userID := protocol.NewID(uuid.New())
	serverPort := udp.PickPort()
	serverConfig := &core.Config{
		App: []*anypb.Any{
			serial.ToTypedMessage(&log.Config{
				Error: &log.LogSpecification{Level: clog.Severity_Debug, Type: log.LogType_Console},
			}),
		},
		Inbound: []*core.InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortRange: net.SinglePortRange(serverPort),
					Listen:    net.NewIPOrDomain(net.LocalHostIP),
					StreamSettings: &internet.StreamConfig{
						ProtocolName: "quic",
						TransportSettings: []*internet.TransportConfig{
							{
								ProtocolName: "quic",
								Settings: serial.ToTypedMessage(&quic.Config{
									Header: serial.ToTypedMessage(&wechat.VideoConfig{}),
									Security: &protocol.SecurityConfig{
										Type: protocol.SecurityType_NONE,
									},
								}),
							},
						},
					},
				}),
				ProxySettings: serial.ToTypedMessage(&inbound.Config{
					User: []*protocol.User{
						{
							Account: serial.ToTypedMessage(&vmess.Account{
								Id:      userID.String(),
								AlterId: 0,
							}),
						},
					},
				}),
			},
		},
		Outbound: []*core.OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&freedom.Config{}),
			},
		},
	}

	clientPort := tcp.PickPort()
	clientConfig := &core.Config{
		App: []*anypb.Any{
			serial.ToTypedMessage(&log.Config{
				Error: &log.LogSpecification{Level: clog.Severity_Debug, Type: log.LogType_Console},
			}),
		},
		Inbound: []*core.InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortRange: net.SinglePortRange(clientPort),
					Listen:    net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address: net.NewIPOrDomain(dest.Address),
					Port:    uint32(dest.Port),
					NetworkList: &net.NetworkList{
						Network: []net.Network{net.Network_TCP},
					},
				}),
			},
		},
		Outbound: []*core.OutboundHandlerConfig{
			{
				SenderSettings: serial.ToTypedMessage(&proxyman.SenderConfig{
					StreamSettings: &internet.StreamConfig{
						ProtocolName: "quic",
						TransportSettings: []*internet.TransportConfig{
							{
								ProtocolName: "quic",
								Settings: serial.ToTypedMessage(&quic.Config{
									Header: serial.ToTypedMessage(&wechat.VideoConfig{}),
									Security: &protocol.SecurityConfig{
										Type: protocol.SecurityType_NONE,
									},
								}),
							},
						},
					},
				}),
				ProxySettings: serial.ToTypedMessage(&outbound.Config{
					Receiver: []*protocol.ServerEndpoint{
						{
							Address: net.NewIPOrDomain(net.LocalHostIP),
							Port:    uint32(serverPort),
							User: []*protocol.User{
								{
									Account: serial.ToTypedMessage(&vmess.Account{
										Id:      userID.String(),
										AlterId: 0,
										SecuritySettings: &protocol.SecurityConfig{
											Type: protocol.SecurityType_AES128_GCM,
										},
									}),
								},
							},
						},
					},
				}),
			},
		},
	}

	servers, err := InitializeServerConfigs(serverConfig, clientConfig)
	if err != nil {
		t.Fatal("Failed to initialize all servers: ", err.Error())
	}
	defer CloseAllServers(servers)

	var errg errgroup.Group
	for i := 0; i < 10; i++ {
		errg.Go(testTCPConn(clientPort, 10240*1024, time.Second*40))
	}

	if err := errg.Wait(); err != nil {
		t.Error(err)
	}
}
