// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package envflags_test

import (
	"github.com/onsi/ginkgo/v2"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/envflags"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/executer"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

var _ = ginkgo.Describe("Cross-version Env Flags APIs compatibility tests", func() {
	ginkgo.GinkgoWriter.Println("GetEnv, GetEnvConfigurations, SetEnv, DeleteEnv methods are tested for cross-version API compatibility with supported Runtime versions v0.11.6, v0.25.4, v0.28.0, latest")
	var multipleTestEnvs map[string]string
	ginkgo.BeforeEach(func() {
		multipleTestEnvs = map[string]string{
			envflags.CompatibilityTestsEnvZero: envflags.CompatibilityTestsEnvVal,
			envflags.CompatibilityTestsEnvOne:  envflags.CompatibilityTestsEnvVal,
		}
		// Setup mock temporary config files for testing
		_, cleanup := core.SetupTempCfgFiles()
		ginkgo.DeferCleanup(func() {
			cleanup()
		})
	})

	ginkgo.Context("using single env flag", func() {

		ginkgo.It("Run SetEnv latest - DeleteEnv v0.90.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.Version_latest))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.90.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv latest - DeleteEnv v1.0 ", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.Version_latest))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.90.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv latest - DeleteEnv v0.28.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.Version_latest))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.28.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetEnv v1.0 - DeleteEnv v0.90.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.Version1_0))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.90.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v1.0 - DeleteEnv latest", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.Version1_0))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.90.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v1.0 - DeleteEnv v0.28.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.Version1_0))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.28.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetEnv v0.90.0 - DeleteEnv latest", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.Version0_90))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv latest Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v0.90.0 - DeleteEnv v1.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.Version0_90))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv latest Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v0.90.0 - DeleteEnv v0.28.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.Version0_90))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.28.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetEnv v0.28.0 - DeleteEnv latest", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime v0.28.0
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.Version0_28))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv latest Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version_latest))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v0.28.0 - DeleteEnv v1.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime v0.28.0
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.Version0_28))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv latest Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version1_0))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v0.28.0 - DeleteEnv v0.90.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.Version0_28))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv latest Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version0_90))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
	})

	ginkgo.Context("using multiple env flags", func() {

		ginkgo.It("Run SetEnv latest - DeleteEnv v0.28.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.Version_latest))
			testCase.Add(envflags.DefaultSetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.28.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv latest - DeleteEnv v0.90.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.Version_latest))
			testCase.Add(envflags.DefaultSetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.90.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv latest - DeleteEnv v1.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.Version_latest))
			testCase.Add(envflags.DefaultSetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.90.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetEnv v1.0 - DeleteEnv v0.28.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.Version1_0))
			testCase.Add(envflags.DefaultSetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.28.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v1.0 - DeleteEnv v0.90.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.Version1_0))
			testCase.Add(envflags.DefaultSetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.90.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v1.0 - DeleteEnv latest", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.Version1_0))
			testCase.Add(envflags.DefaultSetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.90.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetEnv v0.90.0 - DeleteEnv v0.28.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.Version0_90))
			testCase.Add(envflags.DefaultSetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.28.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v0.90.0 - DeleteEnv latest", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.Version0_90))
			testCase.Add(envflags.DefaultSetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv latest Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v0.90.0 - DeleteEnv v1.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.Version0_90))
			testCase.Add(envflags.DefaultSetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv latest Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetEnv v0.28.0 - DeleteEnv latest", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime v0.28.0
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.Version0_28))
			testCase.Add(envflags.DefaultSetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.28.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v0.28.0 - DeleteEnv v1.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime v0.28.0
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.Version0_28))
			testCase.Add(envflags.DefaultSetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.28.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v0.28.0 - DeleteEnv v0.90.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime v0.28.0
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.Version0_28))
			testCase.Add(envflags.DefaultSetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.90.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.Version_latest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version1_0, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_90, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0_28, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_25, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0_11, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
	})
})
