package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider/config/account"
	"gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider/config/compute"
	"gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider/config/network"
	"gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider/config/storage"
)

const (
	resourcePrefix = "vastai"
	modulePath     = "gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("vastai.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

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
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

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
