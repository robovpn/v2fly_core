package memconservative

import (
	"runtime"

	"github.com/robovpn/github.com/robovpn/v2fly_core/app/router/routercommon"
	"github.com/robovpn/github.com/robovpn/v2fly_core/infra/conf/geodata"
)

//go:generate go run github.com/robovpn/github.com/robovpn/v2fly_core/common/errors/errorgen

type memConservativeLoader struct {
	geoipcache   GeoIPCache
	geositecache GeoSiteCache
}

func (m *memConservativeLoader) LoadIP(filename, country string) ([]*routercommon.CIDR, error) {
	defer runtime.GC()
	geoip, err := m.geoipcache.Unmarshal(filename, country)
	if err != nil {
		return nil, newError("failed to decode geodata file: ", filename).Base(err)
	}
	return geoip.Cidr, nil
}

func (m *memConservativeLoader) LoadSite(filename, list string) ([]*routercommon.Domain, error) {
	defer runtime.GC()
	geosite, err := m.geositecache.Unmarshal(filename, list)
	if err != nil {
		return nil, newError("failed to decode geodata file: ", filename).Base(err)
	}
	return geosite.Domain, nil
}

func newMemConservativeLoader() geodata.LoaderImplementation {
	return &memConservativeLoader{make(map[string]*routercommon.GeoIP), make(map[string]*routercommon.GeoSite)}
}

func init() {
	geodata.RegisterGeoDataLoaderImplementationCreator("memconservative", newMemConservativeLoader)
}
