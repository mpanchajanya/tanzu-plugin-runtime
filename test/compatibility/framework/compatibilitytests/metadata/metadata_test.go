// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package metadata_test

import (
	"github.com/onsi/ginkgo/v2"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/executer"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/metadata"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/types"
)

var _ = ginkgo.Describe("Cross-version ConfigMetadata Flags APIs compatibility tests", func() {
	ginkgo.GinkgoWriter.Println("SetConfigMetadataPatchStrategy, SetConfigMetadataSetting, GetMetadata, GetConfigMetadata, GetConfigMetadataPatchStrategy, GetConfigMetadataSettings, GetConfigMetadataSetting,IsConfigMetadataSettingsEnabled, UseUnifiedConfig,  DeleteConfigMetadataSetting methods are tested for cross-version API compatibility with supported Runtime versions v0.28.0, latest")

	ginkgo.BeforeEach(func() {
		// Setup mock temporary config files for testing
		_, cleanup := core.SetupTempCfgFiles()
		ginkgo.DeferCleanup(func() {
			cleanup()
		})
	})

	ginkgo.Context("using default metadata", func() {

		ginkgo.It("Run SetConfigMetadataPatchStrategy, SetConfigMetadataSetting latest  DeleteConfigMetadataSetting v0.28.0", func() {
			// Build test case with commands

			// Add SetConfigMetadata Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(metadata.DefaultSetConfigMetadataPatchStrategyCommand(core.Version_latest))
			testCase.Add(metadata.DefaultSetConfigMetadataSettingCommand(core.Version_latest))

			verifyAfterSet(testCase)

			// Add DeleteConfigMetadata v0.28.0 Command
			testCase.Add(metadata.DefaultDeleteConfigMetadataSettingCommand(core.Version0_28))

			verifyAfterDelete(testCase)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetConfigMetadataPatchStrategy, SetConfigMetadataSetting latest  DeleteConfigMetadataSetting v0.90.0", func() {
			// Build test case with commands

			// Add SetConfigMetadata Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(metadata.DefaultSetConfigMetadataPatchStrategyCommand(core.Version_latest))
			testCase.Add(metadata.DefaultSetConfigMetadataSettingCommand(core.Version_latest))

			verifyAfterSet(testCase)

			// Add DeleteConfigMetadata v0.90.0 Command
			testCase.Add(metadata.DefaultDeleteConfigMetadataSettingCommand(core.Version0_90))

			verifyAfterDelete(testCase)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetConfigMetadataPatchStrategy, SetConfigMetadataSetting latest  DeleteConfigMetadataSetting v1.0", func() {
			// Build test case with commands

			// Add SetConfigMetadata Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(metadata.DefaultSetConfigMetadataPatchStrategyCommand(core.Version_latest))
			testCase.Add(metadata.DefaultSetConfigMetadataSettingCommand(core.Version_latest))

			verifyAfterSet(testCase)

			// Add DeleteConfigMetadata v0.90.0 Command
			testCase.Add(metadata.DefaultDeleteConfigMetadataSettingCommand(core.Version1_0))

			verifyAfterDelete(testCase)
			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetConfigMetadataPatchStrategy, SetConfigMetadataSetting v1.0  DeleteConfigMetadataSetting v0.28.0", func() {
			// Build test case with commands

			// Add SetConfigMetadata Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(metadata.DefaultSetConfigMetadataPatchStrategyCommand(core.Version1_0))
			testCase.Add(metadata.DefaultSetConfigMetadataSettingCommand(core.Version1_0))

			verifyAfterSet(testCase)

			// Add DeleteConfigMetadata v0.28.0 Command
			testCase.Add(metadata.DefaultDeleteConfigMetadataSettingCommand(core.Version0_28))

			verifyAfterDelete(testCase)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetConfigMetadataPatchStrategy, SetConfigMetadataSetting v1.0  DeleteConfigMetadataSetting v0.90.0", func() {
			// Build test case with commands

			// Add SetConfigMetadata Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(metadata.DefaultSetConfigMetadataPatchStrategyCommand(core.Version1_0))
			testCase.Add(metadata.DefaultSetConfigMetadataSettingCommand(core.Version1_0))

			verifyAfterSet(testCase)

			// Add DeleteConfigMetadata v0.90.0 Command
			testCase.Add(metadata.DefaultDeleteConfigMetadataSettingCommand(core.Version0_90))

			verifyAfterDelete(testCase)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetConfigMetadataPatchStrategy, SetConfigMetadataSetting v1.0  DeleteConfigMetadataSetting latest", func() {
			// Build test case with commands

			// Add SetConfigMetadata Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(metadata.DefaultSetConfigMetadataPatchStrategyCommand(core.Version1_0))
			testCase.Add(metadata.DefaultSetConfigMetadataSettingCommand(core.Version1_0))

			verifyAfterSet(testCase)

			// Add DeleteConfigMetadata v0.90.0 Command
			testCase.Add(metadata.DefaultDeleteConfigMetadataSettingCommand(core.Version_latest))

			verifyAfterDelete(testCase)
			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetConfigMetadataPatchStrategy, SetConfigMetadataSetting v0.90.0  DeleteConfigMetadataSetting v0.28.0", func() {
			// Build test case with commands

			// Add SetConfigMetadata Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(metadata.DefaultSetConfigMetadataPatchStrategyCommand(core.Version0_90))
			testCase.Add(metadata.DefaultSetConfigMetadataSettingCommand(core.Version0_90))

			verifyAfterSet(testCase)

			// Add DeleteConfigMetadata v0.28.0 Command
			testCase.Add(metadata.DefaultDeleteConfigMetadataSettingCommand(core.Version0_28))

			verifyAfterDelete(testCase)
			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetConfigMetadataPatchStrategy, SetConfigMetadataSetting v0.90.0  DeleteConfigMetadataSetting v1.0", func() {
			// Build test case with commands

			// Add SetConfigMetadata Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(metadata.DefaultSetConfigMetadataPatchStrategyCommand(core.Version0_90))
			testCase.Add(metadata.DefaultSetConfigMetadataSettingCommand(core.Version0_90))

			verifyAfterSet(testCase)

			// Add DeleteConfigMetadata latest Command
			testCase.Add(metadata.DefaultDeleteConfigMetadataSettingCommand(core.Version1_0))

			verifyAfterDelete(testCase)
			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetConfigMetadataPatchStrategy, SetConfigMetadataSetting v0.90.0  DeleteConfigMetadataSetting latest", func() {
			// Build test case with commands

			// Add SetConfigMetadata Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(metadata.DefaultSetConfigMetadataPatchStrategyCommand(core.Version0_90))
			testCase.Add(metadata.DefaultSetConfigMetadataSettingCommand(core.Version0_90))

			verifyAfterSet(testCase)

			// Add DeleteConfigMetadata latest Command
			testCase.Add(metadata.DefaultDeleteConfigMetadataSettingCommand(core.Version_latest))

			verifyAfterDelete(testCase)
			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetConfigMetadataPatchStrategy, SetConfigMetadataSetting v0.28.0  DeleteConfigMetadataSetting latest", func() {
			// Build test case with commands
			testCase := core.NewTestCase()

			// Add SetConfigMetadata Commands of Runtime v0.28.0
			testCase.Add(metadata.DefaultSetConfigMetadataPatchStrategyCommand(core.Version0_28))
			testCase.Add(metadata.DefaultSetConfigMetadataSettingCommand(core.Version0_28))

			verifyAfterSet(testCase)

			// Add DeleteConfigMetadataSetting latest Command
			testCase.Add(metadata.DefaultDeleteConfigMetadataSettingCommand(core.Version_latest))

			verifyAfterDelete(testCase)
			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetConfigMetadataPatchStrategy, SetConfigMetadataSetting v0.28.0  DeleteConfigMetadataSetting v1.0", func() {
			// Build test case with commands
			testCase := core.NewTestCase()

			// Add SetConfigMetadata Commands of Runtime v0.28.0
			testCase.Add(metadata.DefaultSetConfigMetadataPatchStrategyCommand(core.Version0_28))
			testCase.Add(metadata.DefaultSetConfigMetadataSettingCommand(core.Version0_28))

			verifyAfterSet(testCase)

			// Add DeleteConfigMetadataSetting latest Command
			testCase.Add(metadata.DefaultDeleteConfigMetadataSettingCommand(core.Version1_0))

			verifyAfterDelete(testCase)
			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetConfigMetadataPatchStrategy, SetConfigMetadataSetting v0.28.0  DeleteConfigMetadataSetting v0.90.0", func() {
			// Build test case with commands
			testCase := core.NewTestCase()

			// Add SetConfigMetadata Commands of Runtime v0.28.0
			testCase.Add(metadata.DefaultSetConfigMetadataPatchStrategyCommand(core.Version0_28))
			testCase.Add(metadata.DefaultSetConfigMetadataSettingCommand(core.Version0_28))

			verifyAfterSet(testCase)

			// Add DeleteConfigMetadataSetting v0.90.0 Command
			testCase.Add(metadata.DefaultDeleteConfigMetadataSettingCommand(core.Version0_90))

			verifyAfterDelete(testCase)
			// Run all the commands
			executer.Execute(testCase)
		})

	})

})

func verifyAfterSet(testCase *core.TestCase) {
	// Add GetConfigMetadata latest, v0.90.0, v0.28.0 Commands
	testCase.Add(metadata.DefaultGetMetadataCommand(core.Version_latest))
	testCase.Add(metadata.DefaultGetMetadataCommand(core.Version1_0))
	testCase.Add(metadata.DefaultGetMetadataCommand(core.Version0_90))
	testCase.Add(metadata.DefaultGetMetadataCommand(core.Version0_28))

	testCase.Add(metadata.DefaultGetConfigMetadataCommand(core.Version_latest))
	testCase.Add(metadata.DefaultGetConfigMetadataCommand(core.Version1_0))
	testCase.Add(metadata.DefaultGetConfigMetadataCommand(core.Version0_90))
	testCase.Add(metadata.DefaultGetConfigMetadataCommand(core.Version0_28))

	testCase.Add(metadata.DefaultGetConfigMetadataPatchStrategyCommand(core.Version_latest))
	testCase.Add(metadata.DefaultGetConfigMetadataPatchStrategyCommand(core.Version1_0))
	testCase.Add(metadata.DefaultGetConfigMetadataPatchStrategyCommand(core.Version0_90))
	testCase.Add(metadata.DefaultGetConfigMetadataPatchStrategyCommand(core.Version0_28))

	testCase.Add(metadata.DefaultGetConfigMetadataSettingCommand(core.Version_latest))
	testCase.Add(metadata.DefaultGetConfigMetadataSettingCommand(core.Version1_0))
	testCase.Add(metadata.DefaultGetConfigMetadataSettingCommand(core.Version0_90))
	testCase.Add(metadata.DefaultGetConfigMetadataSettingCommand(core.Version0_28))

	testCase.Add(metadata.DefaultGetConfigMetadataSettingsCommand(core.Version_latest))
	testCase.Add(metadata.DefaultGetConfigMetadataSettingsCommand(core.Version1_0))
	testCase.Add(metadata.DefaultGetConfigMetadataSettingsCommand(core.Version0_90))
	testCase.Add(metadata.DefaultGetConfigMetadataSettingsCommand(core.Version0_28))

	testCase.Add(metadata.DefaultIsConfigMetadataSettingsEnabledCommand(core.Version_latest))
	testCase.Add(metadata.DefaultIsConfigMetadataSettingsEnabledCommand(core.Version1_0))
	testCase.Add(metadata.DefaultIsConfigMetadataSettingsEnabledCommand(core.Version0_90))
	testCase.Add(metadata.DefaultIsConfigMetadataSettingsEnabledCommand(core.Version0_28))

	testCase.Add(metadata.DefaultUseUnifiedConfigCommand(core.Version_latest))
	testCase.Add(metadata.DefaultUseUnifiedConfigCommand(core.Version1_0))
	testCase.Add(metadata.DefaultUseUnifiedConfigCommand(core.Version0_90))
	testCase.Add(metadata.DefaultUseUnifiedConfigCommand(core.Version0_28))
}

func verifyAfterDelete(testCase *core.TestCase) {
	metadataOpts := &types.MetadataOpts{
		ConfigMetadata: &types.ConfigMetadataOpts{
			Settings: map[string]string{},
		},
	}

	testCase.Add(metadata.DefaultGetMetadataCommand(core.Version_latest, metadata.WithMetadataOpts(metadataOpts)))
	testCase.Add(metadata.DefaultGetMetadataCommand(core.Version1_0, metadata.WithMetadataOpts(metadataOpts)))
	testCase.Add(metadata.DefaultGetMetadataCommand(core.Version0_90, metadata.WithMetadataOpts(metadataOpts)))
	testCase.Add(metadata.DefaultGetMetadataCommand(core.Version0_28, metadata.WithMetadataOpts(metadataOpts)))

	testCase.Add(metadata.DefaultGetConfigMetadataCommand(core.Version_latest, metadata.WithConfigMetadataOpts(metadataOpts.ConfigMetadata)))
	testCase.Add(metadata.DefaultGetConfigMetadataCommand(core.Version1_0, metadata.WithConfigMetadataOpts(metadataOpts.ConfigMetadata)))
	testCase.Add(metadata.DefaultGetConfigMetadataCommand(core.Version0_90, metadata.WithConfigMetadataOpts(metadataOpts.ConfigMetadata)))
	testCase.Add(metadata.DefaultGetConfigMetadataCommand(core.Version0_28, metadata.WithConfigMetadataOpts(metadataOpts.ConfigMetadata)))

	testCase.Add(metadata.DefaultGetConfigMetadataPatchStrategyCommand(core.Version_latest))
	testCase.Add(metadata.DefaultGetConfigMetadataPatchStrategyCommand(core.Version1_0))
	testCase.Add(metadata.DefaultGetConfigMetadataPatchStrategyCommand(core.Version0_90))
	testCase.Add(metadata.DefaultGetConfigMetadataPatchStrategyCommand(core.Version0_28))

	testCase.Add(metadata.DefaultGetConfigMetadataCommand(core.Version_latest, metadata.WithConfigMetadataOpts(metadataOpts.ConfigMetadata)))
	testCase.Add(metadata.DefaultGetConfigMetadataCommand(core.Version1_0, metadata.WithConfigMetadataOpts(metadataOpts.ConfigMetadata)))
	testCase.Add(metadata.DefaultGetConfigMetadataSettingCommand(core.Version0_90, metadata.WithError(common.ErrNotFound)))
	testCase.Add(metadata.DefaultGetConfigMetadataSettingCommand(core.Version0_28, metadata.WithError(common.ErrNotFound)))

	testCase.Add(metadata.DefaultIsConfigMetadataSettingsEnabledCommand(core.Version_latest, metadata.WithError(common.ErrNotFound)))
	testCase.Add(metadata.DefaultIsConfigMetadataSettingsEnabledCommand(core.Version1_0, metadata.WithError(common.ErrNotFound)))
	testCase.Add(metadata.DefaultIsConfigMetadataSettingsEnabledCommand(core.Version0_90, metadata.WithError(common.ErrNotFound)))
	testCase.Add(metadata.DefaultIsConfigMetadataSettingsEnabledCommand(core.Version0_28, metadata.WithError(common.ErrNotFound)))

	testCase.Add(metadata.DefaultUseUnifiedConfigCommand(core.Version_latest, metadata.WithError(common.ErrNotFound)))
	testCase.Add(metadata.DefaultUseUnifiedConfigCommand(core.Version1_0, metadata.WithError(common.ErrNotFound)))
	testCase.Add(metadata.DefaultUseUnifiedConfigCommand(core.Version0_90, metadata.WithError(common.ErrNotFound)))
	testCase.Add(metadata.DefaultUseUnifiedConfigCommand(core.Version0_28, metadata.WithError(common.ErrNotFound)))
}
