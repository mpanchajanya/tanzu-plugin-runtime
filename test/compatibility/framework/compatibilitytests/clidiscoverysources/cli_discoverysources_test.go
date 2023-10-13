// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package clidiscoverysources_test

import (
	"github.com/onsi/ginkgo/v2"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/legacyclientconfig"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/types"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/clidiscoverysources"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/executer"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

var _ = ginkgo.Describe("Cross-version CLI Discovery Source APIs Compatibility Tests for supported Runtime versions v0.11.6, v0.25.4, v0.28.0, v0.90.0, latest", func() {
	ginkgo.GinkgoWriter.Println("Get/Set/Delete CLI Discovery Source API methods are tested for cross-version API compatibility with supported Runtime versions v0.11.6, v0.25.4, v0.28.0, v0.90.0, latest")

	ginkgo.BeforeEach(func() {
		// Setup mock temporary config files for testing
		_, cleanup := core.SetupTempCfgFiles()
		ginkgo.DeferCleanup(func() {
			cleanup()
		})

	})

	ginkgo.Context("Run SetCLIDiscoverySource, GetCLIDiscoverySource, GetClientConfig, DeleteCLIDiscoverySource on all supported Runtime library versions latest, v1.0, v0.90.0, v0.28.0, v0.25.4, v0.11.6", func() {

		ginkgo.It("Run SetCLIDiscoverySource of Runtime latest then GetCLIDiscoverySource, GetClientConfig on all supported Runtime library versions and then DeleteCLIDiscoverySource of Runtime v0.28.0 then GetCLIDiscoverySource on all supported Runtime library versions", func() {
			// Add SetCLIDiscoverySource Commands of Runtime Latest version
			testCase := core.NewTestCase()
			testCase.Add(clidiscoverysources.DefaultSetCLIDiscoverySourceCommand(core.Version_latest))

			// Add GetCLIDiscoverySource Commands on all supported Runtime library versions
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version_latest, clidiscoverysources.WithStrictValidationStrategy()))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version1_0, clidiscoverysources.WithStrictValidationStrategy()))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_90, clidiscoverysources.WithStrictValidationStrategy()))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_28, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))

			// Add GetClientConfig Commands for latest
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithDefaultCLIDiscoverySource(core.Version_latest)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithDefaultCLIDiscoverySource(core.Version1_0)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithDefaultCLIDiscoverySource(core.Version0_90)))

			// Add GetClientConfig Commands for v0.25.4 and v0.11.6 with no discovery sources
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11))

			// Add DeleteCLIDiscoverySource v0.28.0 Command
			testCase.Add(clidiscoverysources.DefaultDeleteCLIDiscoverySourceCommand(core.Version0_28, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))

			// Add GetCLIDiscoverySource Commands on all supported Runtime library versions
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version_latest))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version1_0))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_90))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_28, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetCLIDiscoverySource of Config API v1.0 then GetCLIDiscoverySource, GetClientConfig on all supported Runtime library versions and then DeleteCLIDiscoverySource of Runtime v0.28.0 then GetCLIDiscoverySource on all supported Runtime library versions", func() {
			// Add SetCLIDiscoverySource Commands of Runtime Latest version
			testCase := core.NewTestCase()
			testCase.Add(clidiscoverysources.DefaultSetCLIDiscoverySourceCommand(core.Version1_0))

			// Add GetCLIDiscoverySource Commands on all supported Runtime library versions
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version_latest, clidiscoverysources.WithStrictValidationStrategy()))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version1_0, clidiscoverysources.WithStrictValidationStrategy()))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_90, clidiscoverysources.WithStrictValidationStrategy()))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_28, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))

			// Add GetClientConfig Commands for latest
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithDefaultCLIDiscoverySource(core.Version_latest)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithDefaultCLIDiscoverySource(core.Version1_0)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithDefaultCLIDiscoverySource(core.Version0_90)))

			// Add GetClientConfig Commands for v0.25.4 and v0.11.6 with no discovery sources
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11))

			// Add DeleteCLIDiscoverySource v0.28.0 Command
			testCase.Add(clidiscoverysources.DefaultDeleteCLIDiscoverySourceCommand(core.Version0_28, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))

			// Add GetCLIDiscoverySource Commands on all supported Runtime library versions
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version_latest))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version1_0))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_90))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_28, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetCLIDiscoverySource of Runtime v0.90.0 then GetCLIDiscoverySource, GetClientConfig on all supported Runtime library versions and then DeleteCLIDiscoverySource of Runtime v0.28.0 then GetCLIDiscoverySource on all supported Runtime library versions", func() {
			// Add SetCLIDiscoverySource Commands of Runtime v0.90.0 version
			testCase := core.NewTestCase()
			testCase.Add(clidiscoverysources.DefaultSetCLIDiscoverySourceCommand(core.Version0_90))

			// Add GetCLIDiscoverySource Commands on all supported Runtime library versions
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version_latest, clidiscoverysources.WithStrictValidationStrategy()))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version1_0, clidiscoverysources.WithStrictValidationStrategy()))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_90, clidiscoverysources.WithStrictValidationStrategy()))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_28, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))

			// Add GetClientConfig Commands for latest and v0.90.0
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithDefaultCLIDiscoverySource(core.Version_latest)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithDefaultCLIDiscoverySource(core.Version1_0)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithDefaultCLIDiscoverySource(core.Version0_90)))

			// Add GetClientConfig Commands for v0.25.4 and v0.11.6 with no discovery sources
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11))

			// Add DeleteCLIDiscoverySource v0.28.0 Command
			testCase.Add(clidiscoverysources.DefaultDeleteCLIDiscoverySourceCommand(core.Version0_28, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))

			// Add GetCLIDiscoverySource Commands on all supported Runtime library versions
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version_latest))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version1_0))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_90))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_28, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetCLIDiscoverySource of Runtime v0.28.0 then GetCLIDiscoverySource, GetClientConfig on all supported Runtime library versions and then DeleteCLIDiscoverySource of Runtime v0.11.6 then GetCLIDiscoverySource on all supported Runtime library versions", func() {
			// Setup discovery sources
			ociSource := &types.PluginDiscoveryOpts{
				OCI: &types.OCIDiscoveryOpts{
					Name:  clidiscoverysources.CompatibilityTestsSourceName,
					Image: clidiscoverysources.CompatibilityTestsSourceImage,
				},
			}

			defaultOCISource := &types.PluginDiscoveryOpts{
				OCI: &types.OCIDiscoveryOpts{
					Name:  "default",
					Image: "/:",
				},
				ContextType: types.CtxTypeK8s,
			}

			defaultOCISourceWithoutContextType := &types.PluginDiscoveryOpts{
				OCI: &types.OCIDiscoveryOpts{
					Name:  "default",
					Image: "/:",
				},
			}

			// Add SetCLIDiscoverySource Commands of Runtime v0.28.0 version
			testCase := core.NewTestCase().Add(clidiscoverysources.DefaultSetCLIDiscoverySourceCommand(core.Version0_28))

			// Add GetCLIDiscoverySource Commands on all supported Runtime library versions
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version_latest, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version1_0, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_90, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_28, clidiscoverysources.WithStrictValidationStrategy()))

			// Add GetClientConfig Commands for v0.28.0, v0.25.4, v0.11.6
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultCLIDiscoverySource(core.Version0_28)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithCLIDiscoverySources(core.Version0_25, []types.PluginDiscoveryOpts{
				*defaultOCISource, *ociSource,
			})))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithCLIDiscoverySources(core.Version0_11, []types.PluginDiscoveryOpts{
				*defaultOCISourceWithoutContextType, *ociSource,
			})))

			// Add DeleteCLIDiscoverySource Command for Runtime latest
			testCase.Add(clidiscoverysources.DefaultDeleteCLIDiscoverySourceCommand(core.Version_latest))

			// Add GetCLIDiscoverySource Commands on all supported Runtime library versions
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version_latest, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version1_0, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_90, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_28, clidiscoverysources.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})

	})

	ginkgo.Context("Run StoreClientConfig, GetClientConfig, GetCLIDiscoverySource, DeleteCLIDiscoverySource on all supported Runtime library versions latest, v1.0, v0.90.0, v0.28.0, 0.25.4, 0.11.6", func() {

		ginkgo.It("Run StoreClientConfig of Runtime v0.28.0 then GetClientConfig, GetCLIDiscoverySource on all supported Runtime library versions and then DeleteCLIDiscoverySource of Runtime latest then GetCLIDiscoverySource on all supported Runtime library versions", func() {
			// Setup discovery sources
			ociSource := &types.PluginDiscoveryOpts{
				OCI: &types.OCIDiscoveryOpts{
					Name:  clidiscoverysources.CompatibilityTestsSourceName,
					Image: clidiscoverysources.CompatibilityTestsSourceImage,
				},
			}

			defaultOCISource := &types.PluginDiscoveryOpts{
				OCI: &types.OCIDiscoveryOpts{
					Name:  "default",
					Image: "/:",
				},
				ContextType: types.CtxTypeK8s,
			}

			defaultOCISourceWithoutContextType := &types.PluginDiscoveryOpts{
				OCI: &types.OCIDiscoveryOpts{
					Name:  "default",
					Image: "/:",
				},
			}

			testCase := core.NewTestCase()

			// Add StoreClientConfig Commands of Runtime v0.28.0
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultCLIDiscoverySource(core.Version0_28)))

			// Add GetClientConfig Commands on all supported runtime versions
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultCLIDiscoverySource(core.Version0_28)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithCLIDiscoverySources(core.Version0_25, []types.PluginDiscoveryOpts{
				*defaultOCISource, *ociSource,
			})))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithCLIDiscoverySources(core.Version0_11, []types.PluginDiscoveryOpts{
				*defaultOCISourceWithoutContextType, *ociSource,
			})))

			// Add GetCLIDiscoverySource Commands on all supported Runtime library versions
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_28))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_90, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version1_0, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version_latest, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))

			// Add DeleteCLIDiscoverySource Command of Runtime latest
			testCase.Add(clidiscoverysources.DefaultDeleteCLIDiscoverySourceCommand(core.Version_latest))

			// Add GetCLIDiscoverySource Commands on all supported Runtime library versions
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version_latest, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version1_0, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_90, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_28))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run StoreClientConfig of Runtime v0.28.0 then GetClientConfig, GetCLIDiscoverySource on all supported Runtime library versions and then DeleteCLIDiscoverySource of Runtime v0.90.0 then GetCLIDiscoverySource on all supported Runtime library versions", func() {
			// Setup discovery sources
			ociSource := &types.PluginDiscoveryOpts{
				OCI: &types.OCIDiscoveryOpts{
					Name:  clidiscoverysources.CompatibilityTestsSourceName,
					Image: clidiscoverysources.CompatibilityTestsSourceImage,
				},
			}

			defaultOCISource := &types.PluginDiscoveryOpts{
				OCI: &types.OCIDiscoveryOpts{
					Name:  "default",
					Image: "/:",
				},
				ContextType: types.CtxTypeK8s,
			}

			defaultOCISourceWithoutContextType := &types.PluginDiscoveryOpts{
				OCI: &types.OCIDiscoveryOpts{
					Name:  "default",
					Image: "/:",
				},
			}

			testCase := core.NewTestCase()

			// Add StoreClientConfig Commands of Runtime v0.28.0
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultCLIDiscoverySource(core.Version0_28)))

			// Add GetClientConfig Commands on all supported runtime versions
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultCLIDiscoverySource(core.Version0_28)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithCLIDiscoverySources(core.Version0_25, []types.PluginDiscoveryOpts{
				*defaultOCISource, *ociSource,
			})))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithCLIDiscoverySources(core.Version0_11, []types.PluginDiscoveryOpts{
				*defaultOCISourceWithoutContextType, *ociSource,
			})))

			// Add GetCLIDiscoverySource Commands on all supported Runtime library versions
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_28))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_90, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version1_0, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version_latest, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))

			// Add DeleteCLIDiscoverySource Command of Runtime latest
			testCase.Add(clidiscoverysources.DefaultDeleteCLIDiscoverySourceCommand(core.Version0_90))

			// Add GetCLIDiscoverySource Commands on all supported Runtime library versions
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version_latest, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version1_0, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_90, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_28))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run StoreClientConfig of Runtime v0.25.4 then GetClientConfig, GetCLIDiscoverySource on all supported Runtime library versions and then DeleteCLIDiscoverySource of Runtime v0.28.0 then GetCLIDiscoverySource on all supported Runtime library versions", func() {
			// Setup discovery sources
			ociSource := &types.PluginDiscoveryOpts{
				OCI: &types.OCIDiscoveryOpts{
					Name:  clidiscoverysources.CompatibilityTestsSourceName,
					Image: clidiscoverysources.CompatibilityTestsSourceImage,
				},
			}

			defaultOCISource := &types.PluginDiscoveryOpts{
				OCI: &types.OCIDiscoveryOpts{
					Name:  "default",
					Image: "/:",
				},
				ContextType: types.CtxTypeK8s,
			}

			defaultOCISourceWithoutContextType := &types.PluginDiscoveryOpts{
				OCI: &types.OCIDiscoveryOpts{
					Name:  "default",
					Image: "/:",
				},
			}

			testCase := core.NewTestCase()

			// Add StoreClientConfig Commands of Runtime v0.25.4
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0_25, legacyclientconfig.WithDefaultCLIDiscoverySource(core.Version0_25)))

			// Add GetClientConfig Commands on all supported runtime versions
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultCLIDiscoverySource(core.Version0_28)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithCLIDiscoverySources(core.Version0_25, []types.PluginDiscoveryOpts{
				*defaultOCISource, *ociSource,
			})))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithCLIDiscoverySources(core.Version0_11, []types.PluginDiscoveryOpts{
				*defaultOCISourceWithoutContextType, *ociSource,
			})))

			// Add GetCLIDiscoverySource Commands on all supported Runtime library versions
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_28))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_90, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version1_0, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version_latest, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))

			// Add DeleteCLIDiscoverySource v0.28.0 Command
			testCase.Add(clidiscoverysources.DefaultDeleteCLIDiscoverySourceCommand(core.Version0_28))

			// Add GetCLIDiscoverySource Commands on all supported Runtime library versions
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version_latest, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version1_0, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_90, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_28, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run StoreClientConfig of Runtime v0.11.6 then GetClientConfig, GetCLIDiscoverySource on all supported Runtime library versions and then DeleteCLIDiscoverySource of Runtime v0.28.0 then GetCLIDiscoverySource on all supported Runtime library versions", func() {
			// Setup discovery sources
			ociSource := &types.PluginDiscoveryOpts{
				OCI: &types.OCIDiscoveryOpts{
					Name:  clidiscoverysources.CompatibilityTestsSourceName,
					Image: clidiscoverysources.CompatibilityTestsSourceImage,
				},
			}

			defaultOCISource := &types.PluginDiscoveryOpts{
				OCI: &types.OCIDiscoveryOpts{
					Name:  "default",
					Image: "/:",
				},
				ContextType: types.CtxTypeK8s,
			}

			defaultOCISourceWithoutContextType := &types.PluginDiscoveryOpts{
				OCI: &types.OCIDiscoveryOpts{
					Name:  "default",
					Image: "/:",
				},
			}

			testCase := core.NewTestCase()

			// Add StoreClientConfig Commands of Runtime v0.11.6
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0_11, legacyclientconfig.WithDefaultCLIDiscoverySource(core.Version0_11)))

			// Add GetClientConfig Commands on all supported runtime versions
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultCLIDiscoverySource(core.Version0_28)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithCLIDiscoverySources(core.Version0_25, []types.PluginDiscoveryOpts{
				*defaultOCISource, *ociSource,
			})))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithCLIDiscoverySources(core.Version0_11, []types.PluginDiscoveryOpts{
				*defaultOCISourceWithoutContextType, *ociSource,
			})))

			// Add GetCLIDiscoverySource Commands on all supported Runtime library versions
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_28))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_90, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version1_0, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version_latest, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))

			// Add DeleteCLIDiscoverySource v0.28.0 Command
			testCase.Add(clidiscoverysources.DefaultDeleteCLIDiscoverySourceCommand(core.Version0_28))

			// Add GetCLIDiscoverySource Commands on all supported Runtime library versions
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version_latest, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version1_0, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_90, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0_28, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})

	})

})
