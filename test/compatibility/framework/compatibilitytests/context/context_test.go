// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package context_test

import (
	"github.com/onsi/ginkgo/v2"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/context"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/executer"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/legacyclientconfig"
)

var _ = ginkgo.Describe("Cross-version Context APIs compatibility tests", func() {
	ginkgo.GinkgoWriter.Println("GetContext, SetContext, DeleteContext, GetCurrentContext, SetCurrentContext, RemoveCurrentContext methods are tested for cross-version API compatibility with supported Runtime versions v0.25.4, v0.28.0, v0.90.0, latest")

	ginkgo.BeforeEach(func() {
		// Setup mock temporary config files for testing
		_, cleanup := core.SetupTempCfgFiles()
		ginkgo.DeferCleanup(func() {
			cleanup()
		})
	})

	ginkgo.Context("using single context object on supported Runtime API versions", func() {

		ginkgo.It("Run SetContext, SetCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetCurrentContextCommand())

			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithDefaultContextAndServer(core.Version_latest)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version1_0, legacyclientconfig.WithDefaultContextAndServer(core.Version_latest)))
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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version0_28)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetCurrentContextCommand())

			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithDefaultContextAndServer(core.Version_latest)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_90, legacyclientconfig.WithDefaultContextAndServer(core.Version0_90)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_28, legacyclientconfig.WithDefaultContextAndServer(core.Version0_28)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_25, legacyclientconfig.WithDefaultContextAndServer(core.Version0_25)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0_11, legacyclientconfig.WithDefaultContextAndServer(core.Version0_11)))

			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version_latest, legacyclientconfig.WithDefaultContextAndServer(core.Version_latest)))
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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version0_90)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v1.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetCurrentContextCommand())

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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version1_0)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90)))

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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version0_28)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v1.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90)))

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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version1_0)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90)))

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

			testCase.Add(context.RemoveCurrentContextCommand())
			testCase.Add(context.DeleteContextCommand())

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))

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

			testCase.Add(context.RemoveCurrentContextCommand())
			testCase.Add(context.DeleteContextCommand())

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v1.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))

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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version1_0)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v0.90 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))

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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version0_90)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))

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

			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.GetContextCommand())
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			testCase.Add(context.GetCurrentContextCommand())
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0)))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_25)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.RemoveCurrentContextCommand(context.WithError()))
			testCase.Add(context.DeleteContextCommand(context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v1.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_25)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v0.90 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_25)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			executer.Execute(testCase)
		})

	})

	ginkgo.Context("using multiple context objects on supported Runtime API versions", func() {

		ginkgo.It("Run SetContext, SetCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand())

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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version0_90)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v1.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand())

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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version1_0)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand())

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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version0_28)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0 latest then DeleteContext, RemoveCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand())

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

			testCase.Add(context.RemoveCurrentContextCommand())
			testCase.Add(context.DeleteContextCommand())

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0 latest then DeleteContext, RemoveCurrentContext v1.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand())

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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version1_0)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0 latest then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90)))

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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version0_28)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0 latest then DeleteContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))

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

			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.GetContextCommand())
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand())
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0)))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90)))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0 latest then DeleteContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))

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

			testCase.Add(context.DeleteContextCommand())

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0 latest then DeleteContext v1.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))

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

			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version1_0)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0 latest then DeleteContext v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28)))

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

			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version0_90)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0 latest then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_25)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0 latest then DeleteContext, RemoveCurrentContext v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_25)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0 latest then DeleteContext, RemoveCurrentContext v1.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_25)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0 latest then DeleteContext, RemoveCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_25)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.RemoveCurrentContextCommand(context.WithError()))
			testCase.Add(context.DeleteContextCommand(context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0_25), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version1_0), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_90), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_28), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0_25)))

			executer.Execute(testCase)
		})

	})
})
