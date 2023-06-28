// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"os"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

const DefaultCorePluginGroupName = "vmware-cli/core"
const TanzuCLICorePluginGroupName = "TANZU_CLI_CORE_PLUGIN_GROUP_NAME"
const TanzuCLICorePluginGroupVersion = "TANZU_CLI_CORE_PLUGIN_GROUP_VERSION"

// GetCLICorePluginGroup retrieves cli discovery sources
func GetCLICorePluginGroup() (*configtypes.CorePluginGroup, error) {
	// Retrieve client config node
	node, err := getClientConfigNode()
	if err != nil {
		return nil, err
	}

	essentials, err := getCLIEssentials(node)
	if err != nil {
		essentials = &configtypes.Essentials{
			LastUpdatedTimestamp: nil,
		}
	}

	// Setup Default Essential Plugin groups
	defaultCorePluginGroupName := DefaultCorePluginGroupName
	corePluginGroupName := os.Getenv(TanzuCLICorePluginGroupName)
	if corePluginGroupName != "" {
		defaultCorePluginGroupName = corePluginGroupName
	}

	var defaultCorePluginGroupVersion string
	corePluginGroupVersion := os.Getenv(TanzuCLICorePluginGroupVersion)
	if corePluginGroupVersion != "" {
		defaultCorePluginGroupVersion = corePluginGroupVersion
	}

	defaultCorePluginGroup := &configtypes.CorePluginGroup{
		Name:                 defaultCorePluginGroupName,
		Version:              defaultCorePluginGroupVersion,
		LastUpdatedTimestamp: essentials.LastUpdatedTimestamp,
	}
	return defaultCorePluginGroup, nil
}
