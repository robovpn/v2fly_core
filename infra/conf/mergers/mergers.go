package mergers

//go:generate go run github.com/robovpn/v2fly_core/common/errors/errorgen

import (
	"strings"

	core "github.com/robovpn/v2fly_core"
	"github.com/robovpn/v2fly_core/common"
	"github.com/robovpn/v2fly_core/infra/conf/json"
)

func init() {
	common.Must(RegisterMerger(makeMerger(
		core.FormatJSON,
		[]string{".json", ".jsonc"},
		nil,
	)))
	common.Must(RegisterMerger(makeMerger(
		core.FormatTOML,
		[]string{".toml"},
		json.FromTOML,
	)))
	common.Must(RegisterMerger(makeMerger(
		core.FormatYAML,
		[]string{".yml", ".yaml"},
		json.FromYAML,
	)))
	common.Must(RegisterMerger(
		&Merger{
			Name:       core.FormatAuto,
			Extensions: nil,
			Merge:      Merge,
		}),
	)
}

// Merger is a configurable format merger for v2fly config files.
type Merger struct {
	Name       string
	Extensions []string
	Merge      MergeFunc
}

// MergeFunc is a utility to merge v2fly config from external source into a map and returns it.
type MergeFunc func(input interface{}, m map[string]interface{}) error

var (
	mergersByName = make(map[string]*Merger)
	mergersByExt  = make(map[string]*Merger)
)

// RegisterMerger add a new Merger.
func RegisterMerger(format *Merger) error {
	if _, found := mergersByName[format.Name]; found {
		return newError(format.Name, " already registered.")
	}
	mergersByName[format.Name] = format

	for _, ext := range format.Extensions {
		lext := strings.ToLower(ext)
		if f, found := mergersByExt[lext]; found {
			return newError(ext, " already registered to ", f.Name)
		}
		mergersByExt[lext] = format
	}

	return nil
}
