// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package contextserver_test

import (
	"github.com/onsi/ginkgo/v2"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/context"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/executer"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/legacyclientconfig"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/server"
)

var _ = ginkgo.Describe("Combination Tests for Context - Server APIs", func() {
	// Description on the Tests
	ginkgo.GinkgoWriter.Println("Get/Set/Delete Context, CurrentContext, Server and CurrentServer API methods are tested for cross-version API compatibility with supported Runtime versions v0.25.4, v0.28.0, latest")

	// Setup Data
	ginkgo.BeforeEach(func() {
		// Setup mock temporary config files for testing
		_, cleanup := core.SetupTempCfgFiles()
		ginkgo.DeferCleanup(func() {
			cleanup()
		})
	})

	ginkgo.Context("using single context- Server", func() {

		ginkgo.It("Set Context@latest - Set Server@v1.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetCurrentContextCommand())

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version1_0)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version1_0)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@latest - Set Server@v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetCurrentContextCommand())

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_90)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@latest - Set Server@v0.25.4", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetCurrentContextCommand())

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_25)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_25)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@latest - Set Server@v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetCurrentContextCommand())

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_28)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})

		ginkgo.It("Set Context@v1.0 - Set Server@latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version1_0)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0)))

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetCurrentServerCommand())

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v1.0 - Set Server@v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version1_0)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_90)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v1.0 - Set Server@v0.25.4", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version1_0)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_25)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_25)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v1.0 - Set Server@v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version1_0)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_28)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})

		ginkgo.It("Set Context@v0.90.0 - Set Server@latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90)))

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetCurrentServerCommand())

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.90.0 - Set Server@v1.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version1_0)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version1_0)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.90.0 - Set Server@v0.25.4", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_25)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_25)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.90.0 - Set Server@v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_28)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})

		ginkgo.It("Set Context@v0.28.0 - Set Server@latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetCurrentServerCommand())

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.28.0 - Set Server@v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_90)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.28.0 - Set Server@v0.25.4", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_25)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_25)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.28.0 - Set Server@v1.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version1_0)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version1_0)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})

		ginkgo.It("Set Context@v0.25.4 - Set Server@latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_25)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetCurrentServerCommand())

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.25.4 - Set Server@v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_25)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_90)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.25.4 - Set Server@v1.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_25)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version1_0)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version1_0)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.25.4 - Set Server@v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_25)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_28)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
	})

	ginkgo.Context("using two different contexts- Servers", func() {

		ginkgo.It("Set Context@v0.90.0 - Set Server@latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90)))

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand())

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.90.0 - Set Server@v1.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version1_0)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version1_0), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version1_0)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.90.0 - Set Server@v0.25.4", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_25)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_25), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_25)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.90.0 - Set Server@v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_28), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_28)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})

		ginkgo.It("Set Context@latest - Set Server@latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand())

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand())

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@latest - Set Server@v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand())

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_90), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_90)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@latest - Set Server@v0.25.4", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand())

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_25)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_25), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_25)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@latest - Set Server@v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand())

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_28), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_28)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})

		ginkgo.It("Set Context@v0.28.0 - Set Server@latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand())

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.28.0 - Set Server@v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_90), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_90)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.28.0 - Set Server@v0.25.4", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_25)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_25), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_25)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.28.0 - Set Server@v1.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version1_0)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version1_0), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version1_0)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})

		ginkgo.It("Set Context@v0.25.4 - Set Server@latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_25)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand())

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.25.4 - Set Server@v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_25)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_90), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_90)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.25.4 - Set Server@v1.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_25)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version1_0)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version1_0), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version1_0)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.25.4 - Set Server@v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_25)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0_28), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_28)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
	})

})

func addTestCasesToVerifyContextAndServer(testCase *core.TestCase) {
	testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithDefaultContextAndServer(core.Version_latest)))
	testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithDefaultContextAndServer(core.Version1_0)))
	testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithDefaultContextAndServer(core.Version0_90)))
	testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultContextAndServer(core.Version0_28)))
	testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithDefaultContextAndServer(core.Version0_25)))
	testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithDefaultContextAndServer(core.Version0_11)))

	testCase.Add(context.GetContextCommand())
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0)))
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90)))
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28)))
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25)))

	testCase.Add(context.GetCurrentContextCommand())
	testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0)))
	testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90)))
	testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))
	testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

	testCase.Add(server.GetServerCommand())
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version1_0)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0_90)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0_28)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0_25)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0_11)))

	testCase.Add(server.GetCurrentServerCommand())
	testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version1_0)))
	testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_90)))
	testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_28)))
	testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_25)))
	testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_11)))
}

func addTestCasesToVerifyTwoContextsAndServers(testCase *core.TestCase) {
	testCase.Add(context.GetContextCommand())
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0)))
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90)))
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28)))
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25)))

	testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithContextName(common.CompatibilityTestTwo)))
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo)))
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo)))
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))

	testCase.Add(context.GetCurrentContextCommand())
	testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0)))
	testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90)))
	testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))
	testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

	testCase.Add(server.GetServerCommand())
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version1_0)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0_90)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0_28)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0_25)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0_11)))

	testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version1_0), server.WithServerName(common.CompatibilityTestTwo)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0_90), server.WithServerName(common.CompatibilityTestTwo)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0_28), server.WithServerName(common.CompatibilityTestTwo)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0_25), server.WithServerName(common.CompatibilityTestTwo)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0_11), server.WithServerName(common.CompatibilityTestTwo)))

	testCase.Add(server.GetCurrentServerCommand())
	testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version1_0)))
	testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_90)))
	testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_28)))
	testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_25)))
	testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0_11)))
}
