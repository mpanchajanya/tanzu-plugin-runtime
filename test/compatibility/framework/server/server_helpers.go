// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"fmt"

	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/types"
)

type CfgServerArgs struct {
	RuntimeAPIVersion     *core.RuntimeAPIVersion
	ServerName            string // required
	Type                  types.ServerType
	SetCurrentServer      bool // required
	GlobalOpts            *types.GlobalServerOpts
	ManagementClusterOpts *types.ManagementClusterServerOpts
	DiscoverySources      []types.PluginDiscoveryOpts
	ValidationStrategy    core.ValidationStrategy
	Error                 bool
}

type CfgServerArgsOption func(args *CfgServerArgs)

func WithServerName(name string) CfgServerArgsOption {
	return func(c *CfgServerArgs) {
		c.ServerName = name
	}
}

func WithRuntimeVersion(version core.RuntimeVersion) CfgServerArgsOption {
	return func(c *CfgServerArgs) {
		c.RuntimeAPIVersion = &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		}
	}
}

func WithError() CfgServerArgsOption {
	return func(c *CfgServerArgs) {
		c.Error = true
	}
}

// SetServerCommand creates a core.Command instance for setting server
// configurations. By default, the function initializes configurations
// with preset values, and then applies any user-provided options.
// The command setup depends on the selected runtime API version.
//
// Parameters:
// opts ...CfgServerArgsOption: Variadic list of options to configure the
//                              server arguments.
//
// Returns:
// *core.Command: A pointer to the configured core.Command instance.
func SetServerCommand(opts ...CfgServerArgsOption) *core.Command {
	// Default configuration values.
	args := &CfgServerArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.Version_latest,
		},
		ServerName: common.CompatibilityTestOne,
		Type:       types.ManagementClusterServerType,
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint,
		},
	}

	// Apply any additional provided options.
	for _, opt := range opts {
		opt(args)
	}

	// Placeholders for input and output options. They will be
	// configured based on the runtime API version.
	var inputOpts *SetServerInputOptions
	var outputOpts *SetServerOutputOptions

	// Configuring the input options based on the provided runtime API version.
	switch args.RuntimeAPIVersion.RuntimeVersion {
	case core.Version_latest, core.Version1_0, core.Version0_90, core.Version0_28, core.Version0_25, core.Version0_11:
		inputOpts = &SetServerInputOptions{
			RuntimeAPIVersion: args.RuntimeAPIVersion,
			ServerOpts: &types.ServerOpts{
				Name:       args.ServerName,
				Type:       args.Type,
				GlobalOpts: args.GlobalOpts,
			},
		}
	}

	// Create a new server command with the specified input and output options.
	cmd, err := NewSetServerCommand(inputOpts, outputOpts)
	// Expect no error during command creation.
	gomega.Expect(err).To(gomega.BeNil())

	return cmd
}

// GetServerCommand creates a core.Command instance for retrieving server
// configurations. The function initializes the command with default values
// and subsequently applies any user-provided options. Depending on the runtime
// API version, different configurations and error messages can be set.
//
// Parameters:
// opts ...CfgServerArgsOption: Variadic list of options to configure the
//                              server arguments.
//
// Returns:
// *core.Command: A pointer to the configured core.Command instance.
func GetServerCommand(opts ...CfgServerArgsOption) *core.Command {
	// Default configuration values.
	args := &CfgServerArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.Version_latest,
		},
		ServerName: common.CompatibilityTestOne,
		Type:       types.ManagementClusterServerType,
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint,
		},
		Error: false,
	}

	// Apply any additional provided options.
	for _, opt := range opts {
		opt(args)
	}

	// Define input options based on the provided server arguments.
	inputOpts := &GetServerInputOptions{
		RuntimeAPIVersion: args.RuntimeAPIVersion,
		ServerName:        args.ServerName,
	}

	// Placeholder for output options which will be
	// configured based on the runtime API version and potential errors.
	var outputOpts *GetServerOutputOptions

	// Configuring the output options based on the provided runtime API version.
	switch args.RuntimeAPIVersion.RuntimeVersion {
	case core.Version_latest, core.Version1_0, core.Version0_90, core.Version0_28:
		if args.Error {
			outputOpts = &GetServerOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("could not find server \"%v\"", args.ServerName),
			}
		} else {
			outputOpts = &GetServerOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				ServerOpts: &types.ServerOpts{
					Name:       args.ServerName,
					Type:       args.Type,
					GlobalOpts: args.GlobalOpts,
				},
				ValidationStrategy: core.ValidationStrategyStrict,
			}
		}

	case core.Version0_25, core.Version0_11:
		if args.Error {
			outputOpts = &GetServerOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("could not find server \"%v\"", args.ServerName),
			}
		} else {
			outputOpts = &GetServerOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				ServerOpts: &types.ServerOpts{
					Name:       args.ServerName,
					Type:       args.Type,
					GlobalOpts: args.GlobalOpts,
				},
			}
		}
	}

	// Create a new command for getting server details with the specified input and output options.
	cmd, err := NewGetServerCommand(inputOpts, outputOpts)
	// Expect no error during command creation.
	gomega.Expect(err).To(gomega.BeNil())

	return cmd
}

// DeleteServerCommand creates a core.Command instance for deleting server
// configurations. The function initializes the command with default values
// and subsequently applies any user-provided options. Depending on the runtime
// API version and potential errors, different configurations and error messages can be set.
//
// Parameters:
// opts ...CfgServerArgsOption: Variadic list of options to configure the
//                              server arguments.
//
// Returns:
// *core.Command: A pointer to the configured core.Command instance.
//
//nolint:dupl
func DeleteServerCommand(opts ...CfgServerArgsOption) *core.Command {
	// Default configuration values.
	args := &CfgServerArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.Version_latest,
		},
		ServerName: common.CompatibilityTestOne,
		Type:       types.ManagementClusterServerType,
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint,
		},
	}

	// Apply any additional provided options.
	for _, opt := range opts {
		opt(args)
	}

	// Define input options based on the provided server arguments.
	inputOpts := &DeleteServerInputOptions{
		RuntimeAPIVersion: args.RuntimeAPIVersion,
		ServerName:        args.ServerName,
	}

	// Placeholder for output options which will be
	// configured based on potential errors and runtime API version.
	var outputOpts *DeleteServerOutputOptions

	// Configuring the output options based on potential errors and runtime API version.
	if args.Error {
		switch args.RuntimeAPIVersion.RuntimeVersion {
		case core.Version_latest, core.Version1_0, core.Version0_90, core.Version0_28, core.Version0_25, core.Version0_11:
			outputOpts = &DeleteServerOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("context %v not found", args.ServerName),
			}
		}
	}

	// Create a new command for deleting server details with the specified input and output options.
	cmd, err := NewDeleteServerCommand(inputOpts, outputOpts)
	// Expect no error during command creation.
	gomega.Expect(err).To(gomega.BeNil())

	return cmd
}

// RemoveCurrentServerCommand creates a core.Command instance for removing
// the current server configurations. The function initializes the command
// with default values and subsequently applies any user-provided options.
// Depending on the runtime API version and potential errors, different
// configurations and error messages can be set.
//
// Parameters:
// opts ...CfgServerArgsOption: Variadic list of options to configure the
//                              server arguments.
//
// Returns:
// *core.Command: A pointer to the configured core.Command instance.
//
//nolint:dupl
func RemoveCurrentServerCommand(opts ...CfgServerArgsOption) *core.Command {
	// Default configuration values.
	args := &CfgServerArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.Version_latest,
		},
		ServerName: common.CompatibilityTestOne,
		Type:       types.ManagementClusterServerType,
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint,
		},
	}

	// Apply any additional provided options.
	for _, opt := range opts {
		opt(args)
	}

	// Define input options based on the provided server arguments.
	inputOpts := &RemoveCurrentServerInputOptions{
		RuntimeAPIVersion: args.RuntimeAPIVersion,
		ServerName:        args.ServerName,
	}

	// Placeholder for output options which will be
	// configured based on potential errors and runtime API version.
	var outputOpts *RemoveCurrentServerOutputOptions

	// Configuring the output options based on potential errors and runtime API version.
	if args.Error {
		switch args.RuntimeAPIVersion.RuntimeVersion {
		case core.Version_latest, core.Version1_0, core.Version0_90, core.Version0_28, core.Version0_25, core.Version0_11:
			outputOpts = &RemoveCurrentServerOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("context %v not found", args.ServerName),
			}
		}
	}

	// Create a new command for removing the current server details with the specified input and output options.
	cmd, err := NewRemoveCurrentServerCommand(inputOpts, outputOpts)
	// Expect no error during command creation.
	gomega.Expect(err).To(gomega.BeNil())

	return cmd
}

// GetCurrentServerCommand constructs and returns a core.Command instance
// intended for retrieving details about the current server configurations.
// It first sets up default configurations for the server and then overwrites
// them if user-specified options are provided. Depending on the runtime
// API version and potential errors, different configurations and error
// messages are set.
//
// Parameters:
// opts ...CfgServerArgsOption: Variadic list of options that allow customization
//                              of the server arguments.
//
// Returns:
// *core.Command: A pointer to a configured core.Command instance.
func GetCurrentServerCommand(opts ...CfgServerArgsOption) *core.Command {
	// Default configuration values.
	args := &CfgServerArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.Version_latest,
		},
		ServerName: common.CompatibilityTestOne,
		Type:       types.ManagementClusterServerType,
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint,
		},
	}

	// Apply additional user-provided options.
	for _, opt := range opts {
		opt(args)
	}

	// Placeholder for input and output options.
	var inputOpts *GetCurrentServerInputOptions
	var outputOpts *GetCurrentServerOutputOptions

	// Define input options based on the provided arguments.
	inputOpts = &GetCurrentServerInputOptions{
		RuntimeAPIVersion: args.RuntimeAPIVersion,
	}

	// Configure the output options based on the runtime API version and potential error flags.
	switch args.RuntimeAPIVersion.RuntimeVersion {
	case core.Version_latest, core.Version1_0, core.Version0_90, core.Version0_28, core.Version0_25, core.Version0_11:
		if args.Error {
			outputOpts = &GetCurrentServerOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             "current server \"\" not found in tanzu config",
			}
		} else {
			outputOpts = &GetCurrentServerOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				ServerOpts: &types.ServerOpts{
					Name:       args.ServerName,
					Type:       args.Type,
					GlobalOpts: args.GlobalOpts,
				},
			}
		}
	}

	// Create a command for retrieving the current server's details with the specified input and output options.
	cmd, err := NewGetCurrentServerCommand(inputOpts, outputOpts)
	// Ensure there are no errors during command creation.
	gomega.Expect(err).To(gomega.BeNil())

	return cmd
}

// SetCurrentServerCommand constructs and returns a core.Command instance
// intended for setting the current server configurations.
// It starts with default configurations for the server and may overwrite
// them based on user-specified options.
//
// Parameters:
// opts ...CfgServerArgsOption: Variadic list of options that allow customization
//                              of the server arguments.
//
// Returns:
// *core.Command: A pointer to a configured core.Command instance.
func SetCurrentServerCommand(opts ...CfgServerArgsOption) *core.Command {
	// Default configuration values.
	args := &CfgServerArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.Version_latest,
		},
		ServerName: common.CompatibilityTestOne,
		Type:       types.ManagementClusterServerType,
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint,
		},
	}

	// Apply additional user-provided options.
	for _, opt := range opts {
		opt(args)
	}

	// Define input options based on the provided arguments.
	inputOpts := &SetCurrentServerInputOptions{
		RuntimeAPIVersion: args.RuntimeAPIVersion,
		ServerName:        args.ServerName,
	}

	// Placeholder for output options. Note that as of now, the output options aren't set
	// in this function and are passed as nil to the command.
	var outputOpts *SetCurrentServerOutputOptions

	// Create a command for setting the current server's details with the specified input options.
	cmd, err := NewSetCurrentServerCommand(inputOpts, outputOpts)
	// Ensure there are no errors during command creation.
	gomega.Expect(err).To(gomega.BeNil())

	return cmd
}
