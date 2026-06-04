// Package network contains resource configurations for Vast.ai network
// overlays and overlay membership.
package network

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

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

	p.AddResourceConfigurator("vastai_overlay", func(r *ujconfig.Resource) {
		r.References["cluster_id"] = ujconfig.Reference{
			TerraformName: "vastai_cluster",
		}
	})

	p.AddResourceConfigurator("vastai_overlay_member", func(r *ujconfig.Resource) {
		r.References["overlay_id"] = ujconfig.Reference{
			TerraformName: "vastai_overlay",
		}
		r.References["instance_id"] = ujconfig.Reference{
			TerraformName: "vastai_instance",
		}
	})
}
