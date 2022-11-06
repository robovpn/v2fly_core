package conf

import (
	"/v2fly_core/app/policy"
	"/v2fly_core/common"
	"/v2fly_core/common/bytes"
)

type Policy struct {
	@@ -12,6 +14,8 @@ type Policy struct {
	StatsUserUplink   bool    `json:"statsUserUplink"`
	StatsUserDownlink bool    `json:"statsUserDownlink"`
	BufferSize        *int32  `json:"bufferSize"`
	InboundSpeed      string  `json:"inboundSpeed"`
	OutboundSpeed     string  `json:"outboundSpeed"`
}

func (t *Policy) Build() (*policy.Policy, error) {
	@@ -47,6 +51,23 @@ func (t *Policy) Build() (*policy.Policy, error) {
		}
	}

	var err error
	var inboundSpeed uint64
	if t.InboundSpeed != "" {
		inboundSpeed, err = bytes.ToBytes(t.InboundSpeed)
		common.Must(err)
	}

	var outboundSpeed uint64
	if t.OutboundSpeed != "" {
		outboundSpeed, err = bytes.ToBytes(t.OutboundSpeed)
		common.Must(err)
	}
	p.Speed = &policy.Policy_Speed{
		Inbound:  inboundSpeed,
		Outbound: outboundSpeed,
	}

	return p, nil
}