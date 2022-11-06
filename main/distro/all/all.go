package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/app/dispatcher"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/app/proxyman/inbound"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/app/commander"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/app/log/command"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/app/proxyman/command"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/app/stats/command"

	// Developer preview services
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/app/instman/command"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/app/observatory/command"

	// Other optional features.
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/app/dns"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/app/dns/fakedns"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/app/log"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/app/policy"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/app/reverse"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/app/router"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/transport/internet/tagged/taggedimpl"

	// Developer preview features
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/app/instman"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/app/observatory"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/app/restfulapi"

	// Inbound and outbound proxies.
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/proxy/blackhole"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/proxy/dns"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/proxy/dokodemo"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/proxy/freedom"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/proxy/http"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/proxy/shadowsocks"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/proxy/socks"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/proxy/trojan"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/proxy/vless/inbound"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/proxy/vless/outbound"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/proxy/vmess/inbound"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/proxy/vmess/outbound"

	// Developer preview proxies
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/proxy/vlite/inbound"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/proxy/vlite/outbound"

	// Transports
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/transport/internet/domainsocket"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/transport/internet/grpc"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/transport/internet/http"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/transport/internet/kcp"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/transport/internet/quic"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/transport/internet/tcp"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/transport/internet/tls"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/transport/internet/udp"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/transport/internet/websocket"

	// Transport headers
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/transport/internet/headers/http"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/transport/internet/headers/noop"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/transport/internet/headers/srtp"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/transport/internet/headers/tls"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/transport/internet/headers/utp"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/transport/internet/headers/wechat"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/transport/internet/headers/wireguard"

	// Geo loaders
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/infra/conf/geodata/memconservative"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/infra/conf/geodata/standard"

	// JSON, TOML, YAML config support. (jsonv4) This disable selective compile
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/main/formats"

	// commands
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/main/commands/all"

	// engineering commands
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/main/commands/all/engineering"

	// Commands that rely on jsonv4 format This disable selective compile
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/main/commands/all/api/jsonv4"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/main/commands/all/jsonv4"

	// V5 version of json configure file parser
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/infra/conf/v5cfg"

	// Simplified config
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/proxy/http/simplified"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/proxy/shadowsocks/simplified"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/proxy/socks/simplified"
	_ "github.com/robovpn/github.com/robovpn/v2fly_core/proxy/trojan/simplified"
)
