package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	vastaiprovider "github.com/realnedsanders/terraform-provider-vastai/provider"

	"github.com/Lazlovalentin/crossplane-provider-vast.ai/config/account"
	"github.com/Lazlovalentin/crossplane-provider-vast.ai/config/compute"
	"github.com/Lazlovalentin/crossplane-provider-vast.ai/config/network"
	"github.com/Lazlovalentin/crossplane-provider-vast.ai/config/storage"
)

const (
	resourcePrefix = "vastai"
	modulePath     = "github.com/Lazlovalentin/crossplane-provider-vast.ai"

	// terraformProviderVersion must stay in sync with TERRAFORM_PROVIDER_VERSION
	// in the Makefile so that the Plugin Framework provider reports a
	// matching version string when upjet runs it in-process.
	terraformProviderVersion = "v0.3.3"
)

// vastaiPluginFrameworkInclude restricts upjet's Plugin Framework client
// generation to the Vast.ai resources. Every CRD in this provider is
// implemented via terraform-plugin-framework, so the match is permissive.
var vastaiPluginFrameworkInclude = []string{"vastai_.+$"}

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("vastai.crossplane.io"),
		ujconfig.WithIncludeList(nil),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithTerraformPluginFrameworkIncludeList(vastaiPluginFrameworkInclude),
		ujconfig.WithTerraformPluginFrameworkProvider(vastaiprovider.New(terraformProviderVersion)()))

	for _, configure := range []func(provider *ujconfig.Provider){
		account.Configure,
		compute.Configure,
		network.Configure,
		storage.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("vastai.m.crossplane.io"),
		ujconfig.WithIncludeList(nil),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}),
		ujconfig.WithTerraformPluginFrameworkIncludeList(vastaiPluginFrameworkInclude),
		ujconfig.WithTerraformPluginFrameworkProvider(vastaiprovider.New(terraformProviderVersion)()))

	for _, configure := range []func(provider *ujconfig.Provider){
		account.Configure,
		compute.Configure,
		network.Configure,
		storage.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
