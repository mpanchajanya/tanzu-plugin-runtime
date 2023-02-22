module runtime-test-plugin-0-25

go 1.18

replace (
github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework => ./../../framework
github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/compatibility-test-plugins/helpers => ./../../compatibility-test-plugins/helpers
)


replace (
	sigs.k8s.io/cluster-api => sigs.k8s.io/cluster-api v1.0.1
	sigs.k8s.io/kind => sigs.k8s.io/kind v0.11.1
)


require (
	github.com/spf13/cobra v1.6.1
	github.com/vmware-tanzu/tanzu-framework v0.25.4
	github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework v0.0.0-00010101000000-000000000000
	github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/compatibility-test-plugins/helpers v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/aunum/log v0.0.0-20200821225356-38d2e2c8b489 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-logr/logr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/juju/fslock v0.0.0-20160525022230-4d5c94c67b4b // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/onsi/gomega v1.20.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/vmware-tanzu/tanzu-framework v0.25.4 // indirect
	github.com/vmware-tanzu/tanzu-framework/apis/cli v0.0.0-20230119181514-3c34115bc248 // indirect
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/apimachinery v0.24.2 // indirect
	k8s.io/klog/v2 v2.60.1 // indirect
	k8s.io/utils v0.0.0-20220210201930-3a6ce19ff2f9 // indirect
	sigs.k8s.io/controller-runtime v0.12.3 // indirect
	sigs.k8s.io/json v0.0.0-20211208200746-9f7c6b3444d2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.1 // indirect
)

