// Package account contains resource configurations for Vast.ai account-level
// resources such as API keys, SSH keys, teams, sub-accounts, environment
// variables, and instance templates.
package account

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

func Configure(p *ujconfig.Provider) {
	for tfName, kind := range map[string]string{
		"vastai_api_key":              "APIKey",
		"vastai_ssh_key":              "SSHKey",
		"vastai_subaccount":           "Subaccount",
		"vastai_team":                 "Team",
		"vastai_team_role":            "TeamRole",
		"vastai_team_member":          "TeamMember",
		"vastai_environment_variable": "EnvironmentVariable",
		"vastai_template":             "InstanceTemplate",
	} {
		tfName, kind := tfName, kind
		p.AddResourceConfigurator(tfName, func(r *ujconfig.Resource) {
			r.ShortGroup = "account"
			r.Kind = kind
		})
	}
}
