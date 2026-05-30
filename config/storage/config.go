// Package storage contains resource configurations for Vast.ai volume
// resources (local volumes and network volumes).
package storage

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures the storage group.
func Configure(p *ujconfig.Provider) {
	for tfName, kind := range map[string]string{
		"vastai_volume":         "Volume",
		"vastai_network_volume": "NetworkVolume",
	} {
		tfName, kind := tfName, kind
		p.AddResourceConfigurator(tfName, func(r *ujconfig.Resource) {
			r.ShortGroup = "storage"
			r.Kind = kind
		})
	}
}
