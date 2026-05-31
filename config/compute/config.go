// Package compute contains resource configurations for Vast.ai compute
// resources such as GPU instances, clusters, worker groups, and endpoints.
package compute

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

func Configure(p *ujconfig.Provider) {
	for tfName, kind := range map[string]string{
		"vastai_instance":       "Instance",
		"vastai_cluster":        "Cluster",
		"vastai_cluster_member": "ClusterMember",
		"vastai_worker_group":   "WorkerGroup",
		"vastai_endpoint":       "Endpoint",
	} {
		tfName, kind := tfName, kind
		p.AddResourceConfigurator(tfName, func(r *ujconfig.Resource) {
			r.ShortGroup = "compute"
			r.Kind = kind
		})
	}
}
