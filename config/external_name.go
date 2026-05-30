package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
//
// Vast.ai resources are created server-side with provider-generated numeric
// IDs returned in Terraform state's `id` field, so every entry uses
// IdentifierFromProvider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// account
	"vastai_api_key":              config.IdentifierFromProvider,
	"vastai_ssh_key":              config.IdentifierFromProvider,
	"vastai_subaccount":           config.IdentifierFromProvider,
	"vastai_team":                 config.IdentifierFromProvider,
	"vastai_team_role":            config.IdentifierFromProvider,
	"vastai_team_member":          config.IdentifierFromProvider,
	"vastai_environment_variable": config.IdentifierFromProvider,
	"vastai_template":             config.IdentifierFromProvider,

	// compute
	"vastai_instance":       config.IdentifierFromProvider,
	"vastai_cluster":        config.IdentifierFromProvider,
	"vastai_cluster_member": config.IdentifierFromProvider,
	"vastai_worker_group":   config.IdentifierFromProvider,
	"vastai_endpoint":       config.IdentifierFromProvider,

	// network
	"vastai_overlay":        config.IdentifierFromProvider,
	"vastai_overlay_member": config.IdentifierFromProvider,

	// storage
	"vastai_volume":         config.IdentifierFromProvider,
	"vastai_network_volume": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, 0, len(ExternalNameConfigs))
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l = append(l, name+"$")
	}
	return l
}
