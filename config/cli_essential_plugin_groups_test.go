// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

func TestSetCLIEssentialPluginGroups(t *testing.T) {
	// Setup config test data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	currentTime := time.Now()

	tests := []struct {
		name   string
		input  []*configtypes.EssentialPluginGroup
		total  int
		errStr string
	}{
		{
			name: "success add test",
			input: []*configtypes.EssentialPluginGroup{
				{
					Name:                 "test",
					Version:              "v0.0.1",
					LastUpdatedTimestamp: &currentTime,
				},
			},
			total: 1,
		},
		{
			name: "success add test",
			input: []*configtypes.EssentialPluginGroup{
				{
					Name:                 "test2",
					Version:              "v0.0.1",
					LastUpdatedTimestamp: &currentTime,
				},
			},
			total: 2,
		},
		{
			name: "success update test",
			input: []*configtypes.EssentialPluginGroup{
				{
					Name:                 "test",
					Version:              "v0.0.1",
					LastUpdatedTimestamp: &currentTime,
				},
			},
			total: 2,
		},
		{
			name: "should not persist same test",
			input: []*configtypes.EssentialPluginGroup{
				{
					Name:                 "test",
					Version:              "v0.0.1",
					LastUpdatedTimestamp: &currentTime,
				},
			},
			total: 2,
		},
		{
			name: "failed discovery source with empty name",
			input: []*configtypes.EssentialPluginGroup{
				{
					Name:                 "",
					Version:              "v0.0.1",
					LastUpdatedTimestamp: &currentTime,
				},
			},
			errStr: "discovery source name cannot be empty",
			total:  4,
		},
	}
	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := SetCLIEssentialPluginGroups(spec.input)
			if spec.errStr != "" {
				assert.Equal(t, spec.errStr, err.Error())
			} else {
				assert.NoError(t, err)
			}

			sources, err := GetCLIEssentialPluginGroups()
			assert.NoError(t, err)
			assert.Equal(t, spec.total, len(sources))
		})
	}
}

func TestDeleteCLIEssentialPluginGroup(t *testing.T) {
	// Setup config test data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})
	currentTime := time.Now()

	defer func() {
		cleanUp()
	}()

	tests := []struct {
		name   string
		src    []*configtypes.EssentialPluginGroup
		input  string
		count  int
		errStr string
	}{
		{
			name: "should return err on deleting non existing source",
			src: []*configtypes.EssentialPluginGroup{
				{
					Name:                 "test",
					Version:              "v0.0.1",
					LastUpdatedTimestamp: &currentTime,
				},
			},
			input:  "test-notfound",
			count:  1,
			errStr: "cli discovery source not found",
		},
		{
			name: "should delete existing test source",
			src: []*configtypes.EssentialPluginGroup{
				{
					Name:                 "test",
					Version:              "v0.0.1",
					LastUpdatedTimestamp: &currentTime,
				},
			},
			input: "test",
			count: 0,
		},
		{
			name: "should delete test2 source",
			src: []*configtypes.EssentialPluginGroup{
				{
					Name:                 "test",
					Version:              "v0.0.1",
					LastUpdatedTimestamp: &currentTime,
				},
				{
					Name:                 "test2",
					Version:              "v0.0.1",
					LastUpdatedTimestamp: &currentTime,
				},
			},
			count: 1,
			input: "test2",
		},
	}
	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := SetCLIEssentialPluginGroups(spec.src)
			assert.NoError(t, err)
			err = DeleteCLIEssentialPluginGroup(spec.input)
			if spec.errStr != "" {
				assert.Equal(t, err.Error(), spec.errStr)
			} else {
				assert.NoError(t, err)
			}
			sources, err := GetCLIEssentialPluginGroups()
			assert.NoError(t, err)
			assert.Equal(t, spec.count, len(sources))
		})
	}
}

func TestIntegrationSetGetDeleteCLIEssentialPluginGroup(t *testing.T) {
	// Setup config test data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})
	currentTime := time.Now()

	defer func() {
		cleanUp()
	}()

	sources := []*configtypes.EssentialPluginGroup{
		{
			Name:                 "test",
			Version:              "v0.0.1",
			LastUpdatedTimestamp: &currentTime,
		},
	}

	// Get from the empty config
	ds, err := GetCLIEssentialPluginGroup("test")
	assert.Equal(t, "cli discovery source not found", err.Error())
	assert.Nil(t, ds)

	// Add source to empty config
	err = SetCLIEssentialPluginGroups(sources)
	assert.NoError(t, err)

	ds, err = GetCLIEssentialPluginGroup("test")
	assert.Nil(t, err)
	assert.Equal(t, sources[0], *ds)

	// Delete existing source
	err = DeleteCLIEssentialPluginGroup("test")
	assert.NoError(t, err)

	ds, err = GetCLIEssentialPluginGroup("test")
	assert.Equal(t, "cli discovery source not found", err.Error())
	assert.Nil(t, ds)

}
