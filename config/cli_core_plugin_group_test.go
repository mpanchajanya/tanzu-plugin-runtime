// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

func TestSetCLICorePluginGroup(t *testing.T) {
	// Setup config test data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	currentTime := time.Now()

	tests := []struct {
		name   string
		input  *configtypes.CorePluginGroup
		total  int
		errStr string
	}{
		{
			name: "success add test",
			input: &configtypes.CorePluginGroup{
				Name:                 DefaultCorePluginGroupName,
				LastUpdatedTimestamp: &currentTime,
			},
			total: 1,
		},
		{
			name: "success add test",
			input: &configtypes.CorePluginGroup{
				Name:                 DefaultCorePluginGroupName,
				Version:              "v0.0.2",
				LastUpdatedTimestamp: &currentTime,
			},
			total: 1,
		},
		{
			name: "success add test",
			input: &configtypes.CorePluginGroup{
				Name:                 DefaultCorePluginGroupName,
				Version:              "",
				LastUpdatedTimestamp: &currentTime,
			},
			total: 1,
		},
		{
			name: "success add test",
			input: &configtypes.CorePluginGroup{
				Name:                 DefaultCorePluginGroupName,
				Version:              "v0.0.1",
				LastUpdatedTimestamp: &currentTime,
			},
			total: 1,
		},
		{
			name: "success update test",
			input: &configtypes.CorePluginGroup{
				Name:                 DefaultCorePluginGroupName,
				Version:              "v0.0.1",
				LastUpdatedTimestamp: &currentTime,
			},
			total: 1,
		},
	}
	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := SetCLICorePluginGroup(spec.input)
			if spec.errStr != "" {
				assert.Equal(t, spec.errStr, err.Error())
			} else {
				assert.NoError(t, err)
			}

			group, err := GetCLICorePluginGroup(spec.input.Name)
			assert.NoError(t, err)
			assert.Equal(t, spec.input.Name, group.Name)
			assert.Equal(t, spec.input.Version, group.Version)
			assert.Equal(t, spec.input.LastUpdatedTimestamp.UTC(), group.LastUpdatedTimestamp.UTC())

		})
	}
}
