// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package clidiscoverysources

import (
	"github.com/pkg/errors"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// Validate the setDiscoverySourceInputOptions as per runtime version i.e. check whether mandatory fields are set and throw error if missing
//
//nolint:dupl
func (opts *SetCLIDiscoverySourceInputOptions) Validate() (bool, error) {
	// Run Core Validators
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	switch opts.RuntimeVersion {
	case core.Version_latest, core.Version1_0, core.Version0_90, core.Version0_28, core.Version0_11:
		err = opts.PluginDiscoveryOpts.ValidPluginDiscovery()
		if err != nil {
			return false, err
		}
		return true, nil
	case core.Version0_25:
		err = opts.PluginDiscoveryOpts.ValidPluginDiscovery()
		if err != nil {
			return false, err
		}

		err = opts.PluginDiscoveryOpts.ValidContextType()
		if err != nil {
			return false, err
		}
		return true, nil
	default:
		return false, errors.New("SetCLIDiscoverySource API is not supported for the specified runtime version")
	}
}

// Validate the opts as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *GetCLIDiscoverySourceInputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	if opts.DiscoverySourceName == "" {
		return false, errors.New("discovery source name is required")
	}
	return true, nil
}

// Validate the opts as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
//
//nolint:dupl
func (opts *GetCLIDiscoverySourceOutputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	switch opts.RuntimeVersion {
	case core.Version_latest, core.Version1_0, core.Version0_90, core.Version0_28, core.Version0_11:
		err = opts.PluginDiscoveryOpts.ValidPluginDiscovery()
		if err != nil {
			return false, err
		}
		return true, nil
	case core.Version0_25:
		err = opts.PluginDiscoveryOpts.ValidPluginDiscovery()
		if err != nil {
			return false, err
		}

		err = opts.PluginDiscoveryOpts.ValidContextType()
		if err != nil {
			return false, err
		}
		return true, nil
	default:
		return false, errors.New("GetCLIDiscoverySource API is not supported for the specified runtime version")
	}
}

// Validate the opts as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *DeleteCLIDiscoverySourceInputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}
	if opts.DiscoverySourceName == "" {
		return false, errors.New("discovery source name is required")
	}
	return true, nil
}
