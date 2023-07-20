// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

func TestDeleteContextAdditionalMetadata(t *testing.T) {
	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{cfgNextGen: ``, cfg: ``, cfgMetadata: ``})

	defer func() {
		cleanUp()
	}()

	var testcases = []struct {
		name       string
		ctx        *configtypes.Context
		ctxOptions ContextOptions
		out        map[string]interface{}
		errStr     string
	}{

		{
			name: "should add additional metadata",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"issuer":  "vmw1",
					"issuer2": "vmw1",
					"issuer3": "vmw1",
				},
			},
			ctxOptions: WithAdditionalMetadataKey("issuer4"),
			out: map[string]interface{}{
				"issuer":  "vmw1",
				"issuer2": "vmw1",
				"issuer3": "vmw1",
			},
		},

		{
			name: "should update additional metadata",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"issuer": "vmw2",
				},
			},
			ctxOptions: WithAdditionalMetadataKey("issuer"),
			out: map[string]interface{}{
				"issuer2": "vmw1",
				"issuer3": "vmw1",
			},
		},
		{
			name: "should delete all additional metadata",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"issuer":  "vmw2",
					"issuer2": "vmw2",
					"issuer3": "vmw2",
					"issuer4": "vmw2",
				},
			},
			ctxOptions: WithDeleteAllAdditionalMetadata(),
			out:        nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			// perform test
			err := SetContext(tc.ctx, false)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}

			err = DeleteContextAdditionalMetadata(tc.ctx.Name, tc.ctxOptions)
			assert.NoError(t, err)

			ctx, err := GetContext(tc.ctx.Name)
			assert.NoError(t, err)

			assert.Equal(t, tc.out, ctx.AdditionalMetadata)
		})
	}
}

func TestContextAdditionalMetadataStringToString(t *testing.T) {
	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{cfgNextGen: ``, cfg: ``, cfgMetadata: ``})

	defer func() {
		cleanUp()
	}()

	var testcases = []struct {
		name   string
		ctx    *configtypes.Context
		out    map[string]interface{}
		errStr string
	}{

		{
			name: "should add additional metadata tags:tag1",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"issuer": "vmw1",
				},
			},
		},
		{
			name: "should update additional metadata tags:tag2",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"issuer": "vmw2",
				},
			},
		},

		{
			name: "should clear additional metadata string:string",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"issuer": "",
				},
			},
		},

		{
			name: "should update additional metadata tags:tag1",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"issuer": "vmw1",
				},
			},
		},
		{
			name: "should update additional metadata tags:tag2",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"issuer": "vmw2",
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			// perform test
			err := SetContext(tc.ctx, false)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}
			ctx, err := GetContext(tc.ctx.Name)
			assert.NoError(t, err)

			if tc.out != nil {
				assert.Equal(t, tc.out, ctx.AdditionalMetadata)
			} else {
				assert.Equal(t, tc.ctx.AdditionalMetadata, ctx.AdditionalMetadata)
			}
		})
	}
}

func TestContextAdditionalMetadataStringToInt(t *testing.T) {
	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{cfgNextGen: ``, cfg: ``, cfgMetadata: ``})

	defer func() {
		cleanUp()
	}()

	var testcases = []struct {
		name   string
		ctx    *configtypes.Context
		out    map[string]interface{}
		errStr string
	}{

		// Additional metadata of map[string]int

		{
			name: "should add additional metadata tags:0",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"contextId": 0,
				},
			},
			out: map[string]interface{}{
				"contextId": 0,
			},
		},
		{
			name: "should update additional metadata tags:1",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"contextId": 1,
				},
			},
			out: map[string]interface{}{
				"contextId": 1,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			// perform test
			err := SetContext(tc.ctx, false)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}
			ctx, err := GetContext(tc.ctx.Name)
			assert.NoError(t, err)

			if tc.out != nil {
				assert.Equal(t, tc.out, ctx.AdditionalMetadata)
			} else {
				assert.Equal(t, tc.ctx.AdditionalMetadata, ctx.AdditionalMetadata)
			}
		})
	}
}

func TestContextAdditionalMetadataStringToStringArray(t *testing.T) {
	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{cfgNextGen: ``, cfg: ``, cfgMetadata: ``})

	defer func() {
		cleanUp()
	}()

	var testcases = []struct {
		name   string
		ctx    *configtypes.Context
		out    map[string]interface{}
		errStr string
	}{

		{
			name: "should add additional metadata \"tags\": []string{},",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"issuers": []string{},
				},
			},
			out: map[string]interface{}{
				"issuers": []interface{}{},
			},
		},
		{
			name: "should update additional metadata \"tags\": []string{\"tag1\"},",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"issuers": []interface{}{"vmw1"},
				},
			},
		},
		{
			name: "should update additional metadata \"tags\": []string{\"tag2\"},",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"issuers": []interface{}{"vmw1", "vmw2"},
				},
			},
		},
		{
			name: "should update additional metadata \"tags\": []string{\"tag1\", \"tag2\"},",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"issuers": []interface{}{"vmw1", "vmw2"},
				},
			},
		},
		{
			name: "should update additional metadata \"tags\": []string{\"tag3\", \"tag4\"},\n",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"issuers": []interface{}{"vmw1", "vmw2", "vmw3", "vmw4"},
				},
			},
		},
		{
			name: "should update additional metadata \"tags\": []string{\"tag3\", \"tag4\", \"tag1\"},",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"issuers": []interface{}{"vmw1", "vmw2", "vmw3", "vmw4"},
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			// perform test
			err := SetContext(tc.ctx, false)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}
			ctx, err := GetContext(tc.ctx.Name)
			assert.NoError(t, err)

			if tc.out != nil {
				assert.Equal(t, tc.out, ctx.AdditionalMetadata)
			} else {
				assert.Equal(t, tc.ctx.AdditionalMetadata, ctx.AdditionalMetadata)
			}
		})
	}
}

func TestContextAdditionalMetadataStringToMap(t *testing.T) {
	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{cfgNextGen: ``, cfg: ``, cfgMetadata: ``})

	defer func() {
		cleanUp()
	}()

	var testcases = []struct {
		name   string
		ctx    *configtypes.Context
		out    map[string]interface{}
		errStr string
	}{

		{
			name: "should add additional metadata \"auth\": map[string]string{\"a1\": \"x\",\n},",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"auth": map[string]string{
						"a1": "x",
					},
				},
			},
			out: map[string]interface{}{
				"auth": map[string]interface{}{
					"a1": "x",
				},
			},
		},
		{
			name: "should update additional metadata \"auth\": map[string]string{\n\"a2\": \"y\",\n},",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"auth": map[string]string{
						"a2": "y",
					},
				},
			},
			out: map[string]interface{}{
				"auth": map[string]interface{}{
					"a1": "x",
					"a2": "y",
				},
			},
		},
		{
			name: "should update additional metadata \"auth\": map[string]string{\n\"a1\": \"y\",\n},",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"auth": map[string]string{
						"a1": "y",
					},
				},
			},
			out: map[string]interface{}{
				"auth": map[string]interface{}{
					"a1": "y",
					"a2": "y",
				},
			},
		},
		{
			name: "should update additional metadata \"auth\": map[string]string{\n\"a1\": \"z\",\n\"a3\": \"z\",\n},",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"auth": map[string]string{
						"a1": "z",
						"a3": "z",
					},
				},
			},
			out: map[string]interface{}{
				"auth": map[string]interface{}{
					"a1": "z",
					"a3": "z",
					"a2": "y",
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			// perform test
			err := SetContext(tc.ctx, false)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}
			ctx, err := GetContext(tc.ctx.Name)
			assert.NoError(t, err)
			if tc.out != nil {
				assert.Equal(t, tc.out, ctx.AdditionalMetadata)
			} else {
				assert.Equal(t, tc.ctx.AdditionalMetadata, ctx.AdditionalMetadata)
			}
		})
	}
}

func TestContextAdditionalMetadataStringToStruct(t *testing.T) {
	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{cfgNextGen: ``, cfg: ``, cfgMetadata: ``})

	defer func() {
		cleanUp()
	}()

	var testcases = []struct {
		name   string
		ctx    *configtypes.Context
		out    map[string]interface{}
		errStr string
	}{

		{
			name: "should add additional metadata \"globalAuth\": configtypes.GlobalServerAuth{\n\t\t\t\t\t\tAccessToken: \"token1\",\n\t\t\t\t\t},",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"globalAuth": configtypes.GlobalServerAuth{
						AccessToken: "token1",
					},
				},
			},
			out: map[string]interface{}{
				"globalAuth": map[string]interface{}{
					"accessToken": "token1",
				},
			},
		},
		{
			name: "should update additional metadata \"globalAuth\": configtypes.GlobalServerAuth{\n\t\t\t\t\t\tAccessToken: \"token1\",\n\t\t\t\t\t\tIDToken:     \"id-1\",\n\t\t\t\t\t},",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"globalAuth": configtypes.GlobalServerAuth{
						AccessToken: "token1",
						IDToken:     "id-1",
					},
				},
			},
			out: map[string]interface{}{
				"globalAuth": map[string]interface{}{
					"accessToken": "token1",
					"IDToken":     "id-1",
				},
			},
		},
		{
			name: "should update additional metadata \"globalAuth\": configtypes.GlobalServerAuth{\n\t\t\t\t\t\tAccessToken: \"token1\",\n\t\t\t\t\t\tIDToken:     \"id-1\",\n\t\t\t\t\t\tPermissions: []string{\"p1\", \"p2\"},\n\t\t\t\t\t},",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"globalAuth": configtypes.GlobalServerAuth{
						AccessToken: "token1",
						IDToken:     "id-1",
						Permissions: []string{"p1", "p2"},
					},
				},
			},
			out: map[string]interface{}{
				"globalAuth": map[string]interface{}{
					"accessToken": "token1",
					"IDToken":     "id-1",
					"permissions": []interface{}{"p1", "p2"},
				},
			},
		},

		{
			name: "should update additional metadata \"globalAuth\": configtypes.GlobalServerAuth{\n\t\t\t\t\t\tAccessToken: \"token1\",\n\t\t\t\t\t\tIDToken:     \"id-1\",\n\t\t\t\t\t\tPermissions: []string{\"p1\", \"p2\"},\n\t\t\t\t\t},\n\t\t\t\t\t\"globalAuth2\": configtypes.GlobalServerAuth{\n\t\t\t\t\t\tAccessToken: \"token1\",\n\t\t\t\t\t\tIDToken:     \"id-1\",\n\t\t\t\t\t\tPermissions: []string{\"p1\", \"p2\"},\n\t\t\t\t\t},",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"globalAuth": configtypes.GlobalServerAuth{
						AccessToken: "token1",
						IDToken:     "id-1",
						Permissions: []string{"p1", "p2"},
					},
					"globalAuth2": configtypes.GlobalServerAuth{
						AccessToken: "token1",
						IDToken:     "id-1",
						Permissions: []string{"p1", "p2"},
					},
				},
			},
			out: map[string]interface{}{
				"globalAuth": map[string]interface{}{
					"accessToken": "token1",
					"IDToken":     "id-1",
					"permissions": []interface{}{"p1", "p2"},
				},
				"globalAuth2": map[string]interface{}{
					"accessToken": "token1",
					"IDToken":     "id-1",
					"permissions": []interface{}{"p1", "p2"},
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			// perform test
			err := SetContext(tc.ctx, false)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}
			ctx, err := GetContext(tc.ctx.Name)
			assert.NoError(t, err)

			if tc.out != nil {
				assert.Equal(t, tc.out, ctx.AdditionalMetadata)
			} else {
				assert.Equal(t, tc.ctx.AdditionalMetadata, ctx.AdditionalMetadata)
			}
		})
	}
}

func TestContextAdditionalMetadataStringToStructArray(t *testing.T) {
	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{cfgNextGen: ``, cfg: ``, cfgMetadata: ``})

	defer func() {
		cleanUp()
	}()
	var testcases = []struct {
		name   string
		ctx    *configtypes.Context
		out    map[string]interface{}
		errStr string
	}{

		{
			name: "should add additional metadata \"globalAuth\": configtypes.GlobalServerAuth{\n\t\t\t\t\t\tAccessToken: \"token1\",\n\t\t\t\t\t},",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"globalAuth": []*configtypes.GlobalServerAuth{
						{
							AccessToken: "token1",
						},

						{
							AccessToken: "token2",
						},
					},
				},
			},
			out: map[string]interface{}{
				"globalAuth": []interface{}{
					map[string]interface{}{
						"accessToken": "token1",
					},
					map[string]interface{}{
						"accessToken": "token2",
					},
				},
			},
		},

		{
			name: "should update additional metadata \"globalAuth\": configtypes.GlobalServerAuth{\n\t\t\t\t\t\tAccessToken: \"token1\",\n\t\t\t\t\t\tIDToken:     \"id-1\",\n\t\t\t\t\t},",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"globalAuth": []*configtypes.GlobalServerAuth{
						{
							AccessToken: "token1",
							IDToken:     "id-1",
						},
					},
				},
			},
			out: map[string]interface{}{
				"globalAuth": []interface{}{
					map[string]interface{}{
						"accessToken": "token1",
						"IDToken":     "id-1",
					},
					map[string]interface{}{
						"accessToken": "token2",
					},
				},
			},
		},

		{
			name: "should update additional metadata \"globalAuth\": configtypes.GlobalServerAuth{\n\t\t\t\t\t\tAccessToken: \"token1\",\n\t\t\t\t\t\tIDToken:     \"id-1\",\n\t\t\t\t\t\tPermissions: []string{\"p1\", \"p2\"},\n\t\t\t\t\t},\n\t\t\t\t\t\"globalAuth2\": configtypes.GlobalServerAuth{\n\t\t\t\t\t\tAccessToken: \"token1\",\n\t\t\t\t\t\tIDToken:     \"id-1\",\n\t\t\t\t\t\tPermissions: []string{\"p1\", \"p2\"},\n\t\t\t\t\t},",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				AdditionalMetadata: map[string]interface{}{
					"globalAuth": []*configtypes.GlobalServerAuth{
						{
							Permissions: []string{"p1"},
						},
					},
					"globalAuth2": []*configtypes.GlobalServerAuth{
						{
							AccessToken: "token2",
							Permissions: []string{"p2"},
						},
					},
				},
			},
			out: map[string]interface{}{
				"globalAuth": []interface{}{
					map[string]interface{}{
						"accessToken": "token1",
						"IDToken":     "id-1",
						"permissions": []interface{}{"p1"},
					},
					map[string]interface{}{
						"accessToken": "token2",
					},
				},
				"globalAuth2": []interface{}{
					map[string]interface{}{
						"accessToken": "token2",
						"permissions": []interface{}{"p2"},
					},
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			// perform test
			err := SetContext(tc.ctx, false)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}
			ctx, err := GetContext(tc.ctx.Name)
			assert.NoError(t, err)
			if tc.out != nil {
				assert.Equal(t, tc.out, ctx.AdditionalMetadata)
			} else {
				assert.Equal(t, tc.ctx.AdditionalMetadata, ctx.AdditionalMetadata)
			}
		})
	}
}
