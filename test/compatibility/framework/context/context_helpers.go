// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package context contains all the cross version api compatibility tests for context apis
package context

import (
	"fmt"

	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/types"
)

// SetContextCommand creates a new instance of core.Command based on the given options.
// It initializes default configuration arguments, applies any provided options,
// and constructs the appropriate input options based on the RuntimeAPIVersion.
//
// Parameters:
// opts ...CfgContextArgsOption : A variadic list of optional configuration functions to be applied.
//
// Returns:
// *core.Command : A pointer to a constructed core.Command instance.
func SetContextCommand(opts ...CfgContextArgsOption) *core.Command {
	// Initialize default configuration arguments for the context.
	args := &CfgContextArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.Version_latest, // By default, set the runtime version to the latest.
		},
		ContextName: common.CompatibilityTestOne, // Default context name.
		Target:      types.TargetK8s,             // Default target.
		Type:        types.CtxTypeK8s,            // Default context type.
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint, // Default endpoint.
		},
	}

	// Apply any provided options to the default configuration.
	for _, opt := range opts {
		opt(args)
	}

	// Define pointers for input and output options.
	// These will be used to create the actual command.
	var inputOpts *SetContextInputOptions
	var outputOpts *SetContextOutputOptions

	// Construct the input options based on the given or default RuntimeAPIVersion.
	switch args.RuntimeAPIVersion.RuntimeVersion {
	case core.Version_latest:
		// For the latest version, include all fields in the context options.
		inputOpts = &SetContextInputOptions{
			RuntimeAPIVersion: args.RuntimeAPIVersion,
			ContextOpts: &types.ContextOpts{
				Name:        args.ContextName,
				Target:      args.Target,
				ContextType: types.ContextTypeK8s,
				GlobalOpts:  args.GlobalOpts,
			},
		}
	case core.Version1_0, core.Version0_90, core.Version0_28:
		// For these versions, exclude the ContextType field.
		inputOpts = &SetContextInputOptions{
			RuntimeAPIVersion: args.RuntimeAPIVersion,
			ContextOpts: &types.ContextOpts{
				Name:       args.ContextName,
				Target:     args.Target,
				GlobalOpts: args.GlobalOpts,
			},
		}
	case core.Version0_25:
		// For version 0.25, exclude the Target field.
		inputOpts = &SetContextInputOptions{
			RuntimeAPIVersion: args.RuntimeAPIVersion,
			ContextOpts: &types.ContextOpts{
				Name:       args.ContextName,
				Type:       args.Type,
				GlobalOpts: args.GlobalOpts,
			},
		}
	}

	// Create the command with the constructed input and output options.
	cmd, err := NewSetContextCommand(inputOpts, outputOpts)

	// Assert that there are no errors during the command creation.
	gomega.Expect(err).To(gomega.BeNil())

	return cmd // Return the constructed command.
}

// GetContextCommand creates a new instance of core.Command to get context information
// based on the provided or default configuration options.
//
// Parameters:
// opts ...CfgContextArgsOption : A variadic list of optional configuration functions to be applied.
//
// Returns:
// *core.Command : A pointer to a constructed core.Command instance.
func GetContextCommand(opts ...CfgContextArgsOption) *core.Command {
	// Initialize default configuration arguments for the context.
	args := &CfgContextArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.Version_latest, // By default, set the runtime version to the latest.
		},
		ContextName: common.CompatibilityTestOne, // Default context name.
		Target:      types.TargetK8s,             // Default target.
		Type:        types.CtxTypeK8s,            // Default context type.
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint, // Default endpoint.
		},
		Error: false, // Default error flag set to false.
	}

	// Apply any provided options to the default configuration.
	for _, opt := range opts {
		opt(args)
	}

	// Construct input options with the runtime API version and context name.
	inputOpts := &GetContextInputOptions{
		RuntimeAPIVersion: args.RuntimeAPIVersion,
		ContextName:       args.ContextName,
	}

	var outputOpts *GetContextOutputOptions // Pointer for the output options.

	// Construct the output options based on the given or default RuntimeAPIVersion.
	switch args.RuntimeAPIVersion.RuntimeVersion {
	case core.Version_latest:
		if args.Error {
			// If error flag is true, set error message indicating context not found.
			outputOpts = &GetContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("context %v not found", args.ContextName),
			}
		} else {
			// Otherwise, set the context options with all fields.
			outputOpts = &GetContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				ContextOpts: &types.ContextOpts{
					Name:        args.ContextName,
					Target:      args.Target,
					ContextType: types.ContextTypeK8s,
					GlobalOpts:  args.GlobalOpts,
				},
				ValidationStrategy: core.ValidationStrategyStrict, // Validation strategy set to strict.
			}
		}
	case core.Version1_0, core.Version0_90, core.Version0_28:
		if args.Error {
			// If error flag is true, set error message indicating context not found.
			outputOpts = &GetContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("context %v not found", args.ContextName),
			}
		} else {
			// Otherwise, set the context options excluding the ContextType field.
			outputOpts = &GetContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				ContextOpts: &types.ContextOpts{
					Name:       args.ContextName,
					Target:     args.Target,
					GlobalOpts: args.GlobalOpts,
				},
				ValidationStrategy: core.ValidationStrategyStrict, // Validation strategy set to strict.
			}
		}
	case core.Version0_25:
		if args.Error {
			// Set error message specific to version 0.25 indicating context could not be found.
			outputOpts = &GetContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("could not find context \"%v\"", args.ContextName),
			}
		} else {
			// Set context options excluding the Target field.
			outputOpts = &GetContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				ContextOpts: &types.ContextOpts{
					Name:       args.ContextName,
					Type:       args.Type,
					GlobalOpts: args.GlobalOpts,
				},
			}
		}
	}

	// Create the command with the constructed input and output options.
	cmd, err := NewGetContextCommand(inputOpts, outputOpts)
	// Assert that there are no errors during the command creation.
	gomega.Expect(err).To(gomega.BeNil())

	return cmd // Return the constructed command.
}

// DeleteContextCommand creates a new instance of core.Command intended to delete a context.
// It initializes the default configuration arguments, applies any provided options,
// and constructs the appropriate input and output options.
//
// Parameters:
// opts ...CfgContextArgsOption : A variadic list of optional configuration functions to be applied.
//
// Returns:
// *core.Command : A pointer to a constructed core.Command instance.
//
//nolint:dupl
func DeleteContextCommand(opts ...CfgContextArgsOption) *core.Command {
	// Initialize default configuration arguments for the context.
	args := &CfgContextArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.Version_latest, // By default, set the runtime version to the latest.
		},
		ContextName: common.CompatibilityTestOne, // Default context name.
		Target:      types.TargetK8s,             // Default target.
		Type:        types.CtxTypeK8s,            // Default context type.
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint, // Default endpoint.
		},
	}

	// Apply any provided options to the default configuration.
	for _, opt := range opts {
		opt(args)
	}

	// Construct input options with the runtime API version and context name.
	inputOpts := &DeleteContextInputOptions{
		RuntimeAPIVersion: args.RuntimeAPIVersion,
		ContextName:       args.ContextName,
	}

	var outputOpts *DeleteContextOutputOptions // Pointer for the output options.

	// If the error flag is set in the arguments, construct the output options based on the RuntimeAPIVersion.
	if args.Error {
		switch args.RuntimeAPIVersion.RuntimeVersion {
		case core.Version_latest, core.Version1_0, core.Version0_90, core.Version0_28:
			// For these versions, set error message indicating context not found.
			outputOpts = &DeleteContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("context %v not found", args.ContextName),
			}
		case core.Version0_25:
			// For version 0.25, set a specific error message indicating the context could not be found.
			outputOpts = &DeleteContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("could not find context \"%v\"", args.ContextName),
			}
		}
	}

	// Create the command with the constructed input and output options.
	cmd, err := NewDeleteContextCommand(inputOpts, outputOpts)
	// Assert that there are no errors during the command creation.
	gomega.Expect(err).To(gomega.BeNil())

	return cmd // Return the constructed command.
}

// RemoveCurrentContextCommand creates a new instance of core.Command to remove the current context.
// The function initializes with default configuration arguments, applies any provided options,
// and constructs the appropriate input and output options based on the given or default settings.
//
// Parameters:
// opts ...CfgContextArgsOption : A variadic list of optional configuration functions to be applied.
//
// Returns:
// *core.Command : A pointer to a constructed core.Command instance.
//
//nolint:dupl
func RemoveCurrentContextCommand(opts ...CfgContextArgsOption) *core.Command {
	// Initialize default configuration arguments for the context.
	args := &CfgContextArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.Version_latest, // By default, set the runtime version to the latest.
		},
		ContextName: common.CompatibilityTestOne, // Default context name.
		Target:      types.TargetK8s,             // Default target.
		Type:        types.CtxTypeK8s,            // Default context type.
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint, // Default endpoint.
		},
	}

	// Apply any provided options to the default configuration.
	for _, opt := range opts {
		opt(args)
	}

	// Construct input options with the runtime API version and target.
	inputOpts := &RemoveCurrentContextInputOptions{
		RuntimeAPIVersion: args.RuntimeAPIVersion,
		Target:            args.Target,
	}

	var outputOpts *RemoveCurrentContextOutputOptions // Pointer for the output options.

	// If the error flag is set in the arguments, construct the output options based on the RuntimeAPIVersion.
	if args.Error {
		switch args.RuntimeAPIVersion.RuntimeVersion {
		case core.Version_latest, core.Version1_0, core.Version0_90, core.Version0_28:
			// For these versions, set error message indicating no current context set for the specified target.
			outputOpts = &RemoveCurrentContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("no current context set for target \"%v\"", args.Target),
			}
		case core.Version0_25:
			// For version 0.25, set a specific error message indicating no current context set for the specified type.
			outputOpts = &RemoveCurrentContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("no current context set for type \"%v\"", args.Type),
			}
		}
	}

	// Create the command with the constructed input and output options.
	cmd, err := NewRemoveCurrentContextCommand(inputOpts, outputOpts)
	// Assert that there are no errors during the command creation.
	gomega.Expect(err).To(gomega.BeNil())

	return cmd // Return the constructed command.
}

// GetCurrentContextCommand creates a core.Command instance for obtaining the
// current context configuration. It initializes with default values for
// configurations, and then applies any user-provided options. Based on the
// selected runtime API version and potential error scenarios, it constructs
// appropriate input and output options.
//
// Parameters:
// opts ...CfgContextArgsOption: Variadic list of options to configure the
//                               context arguments.
//
// Returns:
// *core.Command: A pointer to the configured core.Command instance.
func GetCurrentContextCommand(opts ...CfgContextArgsOption) *core.Command {
	args := &CfgContextArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.Version_latest,
		},
		ContextName: common.CompatibilityTestOne,
		Target:      types.TargetK8s,
		Type:        types.CtxTypeK8s,
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint,
		},
	}

	for _, opt := range opts {
		opt(args)
	}

	var inputOpts *GetCurrentContextInputOptions
	var outputOpts *GetCurrentContextOutputOptions

	switch args.RuntimeAPIVersion.RuntimeVersion {
	case core.Version_latest:
		inputOpts = &GetCurrentContextInputOptions{
			RuntimeAPIVersion: args.RuntimeAPIVersion,
			Target:            args.Target,
		}
		if args.Error {
			outputOpts = &GetCurrentContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("no current context set for target \"%v\"", args.Target),
			}
		} else {
			outputOpts = &GetCurrentContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				ContextOpts: &types.ContextOpts{
					Name:        args.ContextName,
					Target:      args.Target,
					ContextType: types.ContextTypeK8s,
					GlobalOpts:  args.GlobalOpts,
				},
				ValidationStrategy: core.ValidationStrategyStrict,
			}
		}
	case core.Version1_0, core.Version0_90, core.Version0_28:
		inputOpts = &GetCurrentContextInputOptions{
			RuntimeAPIVersion: args.RuntimeAPIVersion,
			Target:            args.Target,
		}

		if args.Error {
			outputOpts = &GetCurrentContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("no current context set for target \"%v\"", args.Target),
			}
		} else {
			outputOpts = &GetCurrentContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				ContextOpts: &types.ContextOpts{
					Name:       args.ContextName,
					Target:     args.Target,
					GlobalOpts: args.GlobalOpts,
				},
				ValidationStrategy: core.ValidationStrategyStrict,
			}
		}

	case core.Version0_25:
		inputOpts = &GetCurrentContextInputOptions{
			RuntimeAPIVersion: args.RuntimeAPIVersion,
			ContextType:       args.Type,
		}

		if args.Error {
			outputOpts = &GetCurrentContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("no current context set for type \"%v\"", args.Type),
			}
		} else {
			outputOpts = &GetCurrentContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				ContextOpts: &types.ContextOpts{
					Name:       args.ContextName,
					Type:       args.Type,
					GlobalOpts: args.GlobalOpts,
				},
			}
		}
	}

	cmd, err := NewGetCurrentContextCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())

	return cmd
}

// SetCurrentContextCommand creates a core.Command instance to designate a specific
// context as the "current" one. By default, the function initializes configurations
// with preset values, and subsequently applies any user-provided options.
//
// Parameters:
// opts ...CfgContextArgsOption: Variadic list of options to configure the
//                               context arguments.
//
// Returns:
// *core.Command: A pointer to the configured core.Command instance.
func SetCurrentContextCommand(opts ...CfgContextArgsOption) *core.Command {
	args := &CfgContextArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.Version_latest,
		},
		ContextName: common.CompatibilityTestOne,
		Target:      types.TargetK8s,
		Type:        types.CtxTypeK8s,
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint,
		},
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &SetCurrentContextInputOptions{
		RuntimeAPIVersion: args.RuntimeAPIVersion,
		ContextName:       args.ContextName,
	}

	var outputOpts *SetCurrentContextOutputOptions

	cmd, err := NewSetCurrentContextCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())

	return cmd
}
