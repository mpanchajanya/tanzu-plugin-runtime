// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package featureflags_test

import (
	"github.com/onsi/ginkgo/v2"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/types"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/legacyclientconfig"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/executer"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/featureflags"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

var _ = ginkgo.Describe("Cross-version Feature Flags APIs compatibility tests", func() {
	ginkgo.GinkgoWriter.Println("IsFeatureEnabled, SetFeature, DeleteFeature methods are tested for cross-version API compatibility with supported Runtime versions v0.11.6, v0.25.4, v0.28.0, latest")

	ginkgo.BeforeEach(func() {
		// Setup mock temporary config files for testing
		_, cleanup := core.SetupTempCfgFiles()
		ginkgo.DeferCleanup(func() {
			cleanup()
		})
	})

	ginkgo.Context("using default single feature flag", func() {

		ginkgo.It("Run SetFeature latest - DeleteFeature v0.28.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.Version_latest))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.28.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version0_28))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature latest - DeleteFeature v0.90.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.Version_latest))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version0_90))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature latest - DeleteFeature v1.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.Version_latest))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version1_0))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetFeature v1.0 - DeleteFeature v0.28.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.Version1_0))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.28.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version0_28))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v1.0 - DeleteFeature v0.90.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.Version1_0))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version0_90))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v1.0 - DeleteFeature latest", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.Version1_0))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version_latest))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetFeature v0.90.0 - DeleteFeature v0.28.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.Version0_90))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.28.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version0_28))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v0.90.0 - DeleteFeature latest", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.Version0_90))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature latest Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version_latest))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v0.90.0 - DeleteFeature v1.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.Version0_90))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature latest Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version1_0))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetFeature v0.28.0 - DeleteFeature latest", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime v0.28.0
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.Version0_28))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature latest Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version_latest))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v0.28.0 - DeleteFeature v1.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime v0.28.0
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.Version0_28))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version1_0))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v0.28.0 - DeleteFeature v0.90.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime v0.28.0
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.Version0_28))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version0_90))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run StoreClientConfig v0.25.4 - DeleteFeature latest", func() {
			// Build test case with commands

			testCase := core.NewTestCase()

			// Add StoreClientConfig Command for Runtime v0.25.4
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0_25, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11))

			// Add DeleteFeature latest Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version_latest))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run StoreClientConfig v0.25.4 - DeleteFeature v1.0", func() {
			// Build test case with commands

			testCase := core.NewTestCase()

			// Add StoreClientConfig Command for Runtime v0.25.4
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0_25, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11))

			// Add DeleteFeature latest Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version1_0))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run StoreClientConfig v0.25.4 - DeleteFeature v0.90.0", func() {
			// Build test case with commands

			testCase := core.NewTestCase()

			// Add StoreClientConfig Command for Runtime v0.25.4
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0_25, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version0_90))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run StoreClientConfig v0.25.4 - DeleteFeature v0.28.0", func() {
			// Build test case with commands

			testCase := core.NewTestCase()

			// Add StoreClientConfig Command for Runtime v0.25.4
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0_25, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11))

			// Add DeleteFeature latest Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version0_28))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})

	})

	ginkgo.Context("using default multiple feature flag", func() {

		ginkgo.It("Run SetFeature latest - DeleteFeature v0.28.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version_latest))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig latest Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.28.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version0_28))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature latest - DeleteFeature v0.90.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version_latest))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig latest Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version0_90))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature latest - DeleteFeature v1.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version_latest))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig latest Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version1_0))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetFeature v1.0 - DeleteFeature v0.28.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version1_0))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig latest Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.28.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version0_28))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v1.0 - DeleteFeature v0.90.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version1_0))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig latest Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version0_90))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v1.0 - DeleteFeature latest", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version1_0))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig latest Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version_latest))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetFeature v0.90.0 - DeleteFeature v0.28.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version0_90))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig latest Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.28.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version0_28))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v0.90.0 - DeleteFeature v1.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version0_90))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig latest Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature latest Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version1_0))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v0.90.0 - DeleteFeature latest", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version0_90))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig latest Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature latest Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version_latest))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetFeature v0.28.0 - DeleteFeature latest", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime v0.28.0
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version0_28))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig v0.28.0 Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature latest Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version_latest))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v0.28.0 - DeleteFeature v1.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime v0.28.0
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version0_28))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig v0.28.0 Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature latest Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version1_0))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v0.28.0 - DeleteFeature v0.90.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime v0.28.0
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version0_28))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig v0.28.0 Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version0_90))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version_latest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version1_0, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_90, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_28, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_25, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0_11, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})

	})

})
