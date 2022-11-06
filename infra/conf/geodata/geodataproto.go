package geodata

import (
	"/v2fly_core/app/router/routercommon"
)

//go:generate go run /v2fly_core/common/errors/errorgen

type LoaderImplementation interface {
	LoadSite(filename, list string) ([]*routercommon.Domain, error)
	LoadIP(filename, country string) ([]*routercommon.CIDR, error)
}

type Loader interface {
	LoaderImplementation
	LoadGeoSite(list string) ([]*routercommon.Domain, error)
	LoadGeoSiteWithAttr(file string, siteWithAttr string) ([]*routercommon.Domain, error)
	LoadGeoIP(country string) ([]*routercommon.CIDR, error)
}
