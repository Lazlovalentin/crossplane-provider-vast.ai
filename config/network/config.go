// Package network contains resource configurations for Vast.ai network
// overlays and overlay membership.
package network

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures the network group.
func Configure(p *ujconfig.Provider) {
	for tfName, kind := range map[string]string{
		"vastai_overlay":        "Overlay",
		"vastai_overlay_member": "OverlayMember",
	} {
		tfName, kind := tfName, kind
		p.AddResourceConfigurator(tfName, func(r *ujconfig.Resource) {
			r.ShortGroup = "network"
			r.Kind = kind
		})
	}
}
