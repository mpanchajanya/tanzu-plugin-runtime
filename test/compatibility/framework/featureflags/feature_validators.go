// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package featureflags

import (
	"fmt"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// Validate the setFeatureInputOptions as per runtime version i.e. check whether mandatory fields are set and throw error if missing
func (opts *SetFeatureInputOptions) Validate() (bool, error) {
	// Run Core Validators
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	switch opts.RuntimeVersion {
	case core.Version_latest, core.Version1_0, core.Version0_90, core.Version0_28:
		if opts.Plugin == "" {
			return false, fmt.Errorf("invalid 'plugin' for set context input options for the specified runtime version %v", opts.RuntimeVersion)
		}
		if opts.Key == "" {
			return false, fmt.Errorf("invalid 'key' for set context input options for the specified runtime version %v", opts.RuntimeVersion)
		}
		if opts.Value == "" {
			return false, fmt.Errorf("invalid 'value' for set context input options for the specified runtime version %v", opts.RuntimeVersion)
		}
		return true, nil
	default:
		return false, fmt.Errorf("SetFeature API is not supported for the specified runtime version %v", opts.RuntimeVersion)
	}
}

// Validate the opts as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *IsFeatureEnabledInputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	switch opts.RuntimeVersion {
	case core.Version_latest, core.Version1_0, core.Version0_90, core.Version0_28, core.Version0_25, core.Version0_11:
		if opts.Plugin == "" {
			return false, fmt.Errorf("invalid 'plugin' for set context input options for the specified runtime version %v", opts.RuntimeVersion)
		}
		if opts.Key == "" {
			return false, fmt.Errorf("invalid 'key' for set context input options for the specified runtime version %v", opts.RuntimeVersion)
		}
		return true, nil
	default:
		return false, fmt.Errorf("IsFeatureEnabled API is not supported for the specified runtime version %v", opts.RuntimeVersion)
	}
}

// Validate the opts as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *DeleteFeatureInputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}
	switch opts.RuntimeVersion {
	case core.Version_latest, core.Version1_0, core.Version0_90, core.Version0_28:
		if opts.Plugin == "" {
			return false, fmt.Errorf("invalid 'plugin' for set context input options for the specified runtime version %v", opts.RuntimeVersion)
		}
		if opts.Key == "" {
			return false, fmt.Errorf("invalid 'key' for set context input options for the specified runtime version %v", opts.RuntimeVersion)
		}
		return true, nil
	default:
		return false, fmt.Errorf("DeleteFeature API is not supported for the specified runtime version %v", opts.RuntimeVersion)
	}
}
