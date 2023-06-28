// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCLICorePluginGroup(t *testing.T) {
	testCases := []struct {
		name                       string
		envVariables               map[string]string
		expectedPluginGroupName    string
		expectedPluginGroupVersion string
		expectedError              error
	}{
		{
			name:                       "Default values",
			envVariables:               map[string]string{},
			expectedPluginGroupName:    DefaultCorePluginGroupName,
			expectedPluginGroupVersion: "",
			expectedError:              nil,
		},
		{
			name: "Custom values",
			envVariables: map[string]string{
				TanzuCLICorePluginGroupName:    "custom-group",
				TanzuCLICorePluginGroupVersion: "1.0.0",
			},
			expectedPluginGroupName:    "custom-group",
			expectedPluginGroupVersion: "1.0.0",
			expectedError:              nil,
		},
		{
			name: "Missing core plugin group name",
			envVariables: map[string]string{
				TanzuCLICorePluginGroupVersion: "1.0.0",
			},
			expectedPluginGroupName:    DefaultCorePluginGroupName,
			expectedPluginGroupVersion: "1.0.0",
			expectedError:              nil,
		},
		{
			name: "Missing core plugin group version",
			envVariables: map[string]string{
				TanzuCLICorePluginGroupName: "custom-group",
			},
			expectedPluginGroupName:    "custom-group",
			expectedPluginGroupVersion: "",
			expectedError:              nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set environment variables
			for key, value := range tc.envVariables {
				err := os.Setenv(key, value)
				assert.Nil(t, err)
			}

			// Call the function
			corePluginGroup, err := GetCLICorePluginGroup()

			// Assert the values and errors
			if corePluginGroup.Name != tc.expectedPluginGroupName {
				t.Errorf("Expected core plugin group name: %s, got: %s", tc.expectedPluginGroupName, corePluginGroup.Name)
			}

			if corePluginGroup.Version != tc.expectedPluginGroupVersion {
				t.Errorf("Expected core plugin group version: %s, got: %s", tc.expectedPluginGroupVersion, corePluginGroup.Version)
			}

			if err != tc.expectedError {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}

			// Reset environment variables
			for key := range tc.envVariables {
				err := os.Unsetenv(key)
				assert.Nil(t, err)
			}
		})
	}
}
