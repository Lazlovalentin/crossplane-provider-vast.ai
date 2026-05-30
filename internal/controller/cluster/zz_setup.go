// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0


package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	cluster "gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider/internal/controller/cluster/compute/cluster"
overlaymember "gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider/internal/controller/cluster/network/overlaymember"
volume "gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider/internal/controller/cluster/storage/volume"
environmentvariable "gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider/internal/controller/cluster/account/environmentvariable"
instancetemplate "gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider/internal/controller/cluster/account/instancetemplate"
team "gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider/internal/controller/cluster/account/team"
teamrole "gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider/internal/controller/cluster/account/teamrole"
clustermember "gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider/internal/controller/cluster/compute/clustermember"
workergroup "gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider/internal/controller/cluster/compute/workergroup"
overlay "gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider/internal/controller/cluster/network/overlay"
providerconfig "gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider/internal/controller/cluster/providerconfig"
apikey "gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider/internal/controller/cluster/account/apikey"
sshkey "gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider/internal/controller/cluster/account/sshkey"
subaccount "gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider/internal/controller/cluster/account/subaccount"
teammember "gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider/internal/controller/cluster/account/teammember"
endpoint "gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider/internal/controller/cluster/compute/endpoint"
instance "gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider/internal/controller/cluster/compute/instance"
networkvolume "gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider/internal/controller/cluster/storage/networkvolume"

)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		apikey.Setup,
		environmentvariable.Setup,
		instancetemplate.Setup,
		sshkey.Setup,
		subaccount.Setup,
		team.Setup,
		teammember.Setup,
		teamrole.Setup,
		cluster.Setup,
		clustermember.Setup,
		endpoint.Setup,
		instance.Setup,
		workergroup.Setup,
		overlay.Setup,
		overlaymember.Setup,
		providerconfig.Setup,
		networkvolume.Setup,
		volume.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		apikey.SetupGated,
		environmentvariable.SetupGated,
		instancetemplate.SetupGated,
		sshkey.SetupGated,
		subaccount.SetupGated,
		team.SetupGated,
		teammember.SetupGated,
		teamrole.SetupGated,
		cluster.SetupGated,
		clustermember.SetupGated,
		endpoint.SetupGated,
		instance.SetupGated,
		workergroup.SetupGated,
		overlay.SetupGated,
		overlaymember.SetupGated,
		providerconfig.SetupGated,
		networkvolume.SetupGated,
		volume.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}