// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	apikey "github.com/Lazlovalentin/crossplane-provider-vast.ai/internal/controller/cluster/account/apikey"
	environmentvariable "github.com/Lazlovalentin/crossplane-provider-vast.ai/internal/controller/cluster/account/environmentvariable"
	instancetemplate "github.com/Lazlovalentin/crossplane-provider-vast.ai/internal/controller/cluster/account/instancetemplate"
	sshkey "github.com/Lazlovalentin/crossplane-provider-vast.ai/internal/controller/cluster/account/sshkey"
	subaccount "github.com/Lazlovalentin/crossplane-provider-vast.ai/internal/controller/cluster/account/subaccount"
	team "github.com/Lazlovalentin/crossplane-provider-vast.ai/internal/controller/cluster/account/team"
	teammember "github.com/Lazlovalentin/crossplane-provider-vast.ai/internal/controller/cluster/account/teammember"
	teamrole "github.com/Lazlovalentin/crossplane-provider-vast.ai/internal/controller/cluster/account/teamrole"
	cluster "github.com/Lazlovalentin/crossplane-provider-vast.ai/internal/controller/cluster/compute/cluster"
	clustermember "github.com/Lazlovalentin/crossplane-provider-vast.ai/internal/controller/cluster/compute/clustermember"
	endpoint "github.com/Lazlovalentin/crossplane-provider-vast.ai/internal/controller/cluster/compute/endpoint"
	instance "github.com/Lazlovalentin/crossplane-provider-vast.ai/internal/controller/cluster/compute/instance"
	workergroup "github.com/Lazlovalentin/crossplane-provider-vast.ai/internal/controller/cluster/compute/workergroup"
	overlay "github.com/Lazlovalentin/crossplane-provider-vast.ai/internal/controller/cluster/network/overlay"
	overlaymember "github.com/Lazlovalentin/crossplane-provider-vast.ai/internal/controller/cluster/network/overlaymember"
	providerconfig "github.com/Lazlovalentin/crossplane-provider-vast.ai/internal/controller/cluster/providerconfig"
	networkvolume "github.com/Lazlovalentin/crossplane-provider-vast.ai/internal/controller/cluster/storage/networkvolume"
	volume "github.com/Lazlovalentin/crossplane-provider-vast.ai/internal/controller/cluster/storage/volume"
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
