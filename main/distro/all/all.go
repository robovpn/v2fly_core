package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "../v2fly_core/app/dispatcher"
	_ "../v2fly_core/app/proxyman/inbound"
	_ "../v2fly_core/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "../v2fly_core/app/commander"
	_ "../v2fly_core/app/log/command"
	_ "../v2fly_core/app/proxyman/command"
	_ "../v2fly_core/app/stats/command"

	// Developer preview services
	_ "../v2fly_core/app/instman/command"
	_ "../v2fly_core/app/observatory/command"

	// Other optional features.
	_ "../v2fly_core/app/dns"
	_ "../v2fly_core/app/dns/fakedns"
	_ "../v2fly_core/app/log"
	_ "../v2fly_core/app/policy"
	_ "../v2fly_core/app/reverse"
	_ "../v2fly_core/app/router"
	_ "../v2fly_core/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "../v2fly_core/transport/internet/tagged/taggedimpl"

	// Developer preview features
	_ "../v2fly_core/app/instman"
	_ "../v2fly_core/app/observatory"
	_ "../v2fly_core/app/restfulapi"

	// Inbound and outbound proxies.
	_ "../v2fly_core/proxy/blackhole"
	_ "../v2fly_core/proxy/dns"
	_ "../v2fly_core/proxy/dokodemo"
	_ "../v2fly_core/proxy/freedom"
	_ "../v2fly_core/proxy/http"
	_ "../v2fly_core/proxy/shadowsocks"
	_ "../v2fly_core/proxy/socks"
	_ "../v2fly_core/proxy/trojan"
	_ "../v2fly_core/proxy/vless/inbound"
	_ "../v2fly_core/proxy/vless/outbound"
	_ "../v2fly_core/proxy/vmess/inbound"
	_ "../v2fly_core/proxy/vmess/outbound"

	// Developer preview proxies
	_ "../v2fly_core/proxy/vlite/inbound"
	_ "../v2fly_core/proxy/vlite/outbound"

	// Transports
	_ "../v2fly_core/transport/internet/domainsocket"
	_ "../v2fly_core/transport/internet/grpc"
	_ "../v2fly_core/transport/internet/http"
	_ "../v2fly_core/transport/internet/kcp"
	_ "../v2fly_core/transport/internet/quic"
	_ "../v2fly_core/transport/internet/tcp"
	_ "../v2fly_core/transport/internet/tls"
	_ "../v2fly_core/transport/internet/udp"
	_ "../v2fly_core/transport/internet/websocket"

	// Transport headers
	_ "../v2fly_core/transport/internet/headers/http"
	_ "../v2fly_core/transport/internet/headers/noop"
	_ "../v2fly_core/transport/internet/headers/srtp"
	_ "../v2fly_core/transport/internet/headers/tls"
	_ "../v2fly_core/transport/internet/headers/utp"
	_ "../v2fly_core/transport/internet/headers/wechat"
	_ "../v2fly_core/transport/internet/headers/wireguard"

	// Geo loaders
	_ "../v2fly_core/infra/conf/geodata/memconservative"
	_ "../v2fly_core/infra/conf/geodata/standard"

	// JSON, TOML, YAML config support. (jsonv4) This disable selective compile
	_ "../v2fly_core/main/formats"

	// commands
	_ "../v2fly_core/main/commands/all"

	// engineering commands
	_ "../v2fly_core/main/commands/all/engineering"

	// Commands that rely on jsonv4 format This disable selective compile
	_ "../v2fly_core/main/commands/all/api/jsonv4"
	_ "../v2fly_core/main/commands/all/jsonv4"

	// V5 version of json configure file parser
	_ "../v2fly_core/infra/conf/v5cfg"

	// Simplified config
	_ "../v2fly_core/proxy/http/simplified"
	_ "../v2fly_core/proxy/shadowsocks/simplified"
	_ "../v2fly_core/proxy/socks/simplified"
	_ "../v2fly_core/proxy/trojan/simplified"
)
