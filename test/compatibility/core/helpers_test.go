// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructTestPluginCmd(t *testing.T) {
	tests := []struct {
		version   RuntimeVersion
		apis      []*API
		pluginCmd string
		err       string
	}{
		{
			version: Version_latest,
			apis: []*API{
				{
					Name:      SetContextAPI,
					Version:   Version_latest,
					Arguments: map[APIArgumentType]interface{}{},
					Output:    nil,
				},
			},
			pluginCmd: pluginLatest,
			err:       "",
		},
		{
			version: Version1_0,
			apis: []*API{
				{
					Name:      SetContextAPI,
					Version:   Version1_0,
					Arguments: map[APIArgumentType]interface{}{},
					Output:    nil,
				},
			},
			pluginCmd: pluginV10,
			err:       "",
		},
		{
			version: Version0_90,
			apis: []*API{
				{
					Name:      SetContextAPI,
					Version:   Version0_90,
					Arguments: map[APIArgumentType]interface{}{},
					Output:    nil,
				},
			},
			pluginCmd: pluginV090,
			err:       "",
		},
	}

	for _, tt := range tests {
		pCmd, err := ConstructTestPluginCmd(tt.version, tt.apis)
		if err != nil {
			assert.Equal(t, tt.err, err.Error())
		} else {
			assert.Contains(t, pCmd, tt.pluginCmd)
		}
	}
}
