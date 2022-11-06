package dns_test

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	. "../v2fly_core/app/dns"
	"../v2fly_core/common"
	"../v2fly_core/common/net"
	dns_feature "../v2fly_core/features/dns"
)

func TestDoHLocalNameServer(t *testing.T) {
	url, err := url.Parse("https+local://1.1.1.1/dns-query")
	common.Must(err)

	s := NewDoHLocalNameServer(url)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	ips, err := s.QueryIP(ctx, "google.com", net.IP(nil), dns_feature.IPOption{
		IPv4Enable: true,
		IPv6Enable: true,
	}, false)
	cancel()
	common.Must(err)
	if len(ips) == 0 {
		t.Error("expect some ips, but got 0")
	}
}

func TestDoHLocalNameServerWithCache(t *testing.T) {
	url, err := url.Parse("https+local://1.1.1.1/dns-query")
	common.Must(err)

	s := NewDoHLocalNameServer(url)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	ips, err := s.QueryIP(ctx, "google.com", net.IP(nil), dns_feature.IPOption{
		IPv4Enable: true,
		IPv6Enable: true,
	}, false)
	cancel()
	common.Must(err)
	if len(ips) == 0 {
		t.Error("expect some ips, but got 0")
	}

	ctx2, cancel := context.WithTimeout(context.Background(), time.Second*5)
	ips2, err := s.QueryIP(ctx2, "google.com", net.IP(nil), dns_feature.IPOption{
		IPv4Enable: true,
		IPv6Enable: true,
	}, true)
	cancel()
	common.Must(err)
	if r := cmp.Diff(ips2, ips); r != "" {
		t.Fatal(r)
	}
}
