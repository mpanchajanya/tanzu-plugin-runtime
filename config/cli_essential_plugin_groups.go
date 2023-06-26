// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"fmt"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/nodeutils"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

// GetCLIEssentialPluginGroups retrieves cli discovery sources
func GetCLIEssentialPluginGroups() ([]*configtypes.EssentialPluginGroup, error) {
	// Retrieve client config node
	node, err := getClientConfigNode()
	if err != nil {
		return nil, err
	}

	return getCLIEssentialPluginGroups(node)
}

// GetCLIEssentialPluginGroup retrieves the CLIEssentialPluginGroup by name
func GetCLIEssentialPluginGroup(name string) (*configtypes.EssentialPluginGroup, error) {
	// Retrieve client config node
	node, err := getClientConfigNode()
	if err != nil {
		return nil, err
	}
	return getCLIEssentialPluginGroup(node, name)
}

// SetCLIEssentialPluginGroups Add/Update array of cli discovery sources to the yaml node
func SetCLIEssentialPluginGroups(discoverySources []*configtypes.EssentialPluginGroup) (err error) {
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}

	// Loop through each discovery source and add or update existing node
	for _, discoverySource := range discoverySources {
		persist, err := setCLIEssentialPluginGroup(node, discoverySource)
		if err != nil {
			return err
		}
		// Persist the config node to the file
		if persist {
			err = persistConfig(node)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// SetCLIEssentialPluginGroup add or update CLIEssentialPluginGroup
func SetCLIEssentialPluginGroup(c *configtypes.EssentialPluginGroup) error {
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}
	// Add or update the CLIEssentialPluginGroup
	persist, err := setCLIEssentialPluginGroup(node, c)
	if err != nil {
		return err
	}

	if persist {
		err = persistConfig(node)
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteCLIEssentialPluginGroup delete a CLIEssentialPluginGroup by name
func DeleteCLIEssentialPluginGroup(name string) error {
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()

	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}

	ctx, err := getCLIEssentialPluginGroup(node, name)
	if err != nil {
		return err
	}

	err = removeCLIEssentialPluginGroup(node, ctx.Name)
	if err != nil {
		return err
	}

	return persistConfig(node)
}

func getCLIEssentialPluginGroup(node *yaml.Node, name string) (*configtypes.EssentialPluginGroup, error) {
	// check if CLIEssentialPluginGroup name is empty
	if name == "" {
		return nil, errors.New("CLIEssentialPluginGroup name cannot be empty")
	}

	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}
	for _, ctx := range cfg.CoreCliOptions.EssentialPluginGroups {
		if ctx.Name == name {
			return ctx, nil
		}
	}
	return nil, fmt.Errorf("CLIEssentialPluginGroup %v not found", name)
}

func setCLIEssentialPluginGroup(node *yaml.Node, ctx *configtypes.EssentialPluginGroup) (persist bool, err error) {
	// Check if ctx.Name is empty
	if ctx.Name == "" {
		return false, errors.New("CLIEssentialPluginGroup name cannot be empty")
	}

	// Get Patch Strategies from config metadata
	patchStrategies, err := GetConfigMetadataPatchStrategy()
	if err != nil {
		patchStrategies = make(map[string]string)
	}

	var persistDiscoverySources bool

	// Convert CLIEssentialPluginGroup to node
	newCLIEssentialPluginGroupNode, err := convertObjectToNode(ctx)
	if err != nil {
		return persist, err
	}

	// Find the CLIEssentialPluginGroups node from the root node
	keys := []nodeutils.Key{
		{Name: KeyCLI, Type: yaml.MappingNode},
		{Name: KeyEssentials, Type: yaml.SequenceNode},
	}
	CLIEssentialPluginGroupsNode := nodeutils.FindNode(node.Content[0], nodeutils.WithForceCreate(), nodeutils.WithKeys(keys))
	if CLIEssentialPluginGroupsNode == nil {
		return persist, err
	}

	exists := false
	var result []*yaml.Node
	// Skip duplicate for CLIEssentialPluginGroup and server similar logic
	//nolint:dupl
	for _, CLIEssentialPluginGroupNode := range CLIEssentialPluginGroupsNode.Content {
		if index := nodeutils.GetNodeIndex(CLIEssentialPluginGroupNode.Content, "name"); index != -1 &&
			CLIEssentialPluginGroupNode.Content[index].Value == ctx.Name {
			exists = true
			// replace the nodes as per patch strategy
			_, err = nodeutils.DeleteNodes(newCLIEssentialPluginGroupNode.Content[0], CLIEssentialPluginGroupNode, nodeutils.WithPatchStrategyKey(KeyEssentials), nodeutils.WithPatchStrategies(patchStrategies))
			if err != nil {
				return false, err
			}
			persist, err = nodeutils.MergeNodes(newCLIEssentialPluginGroupNode.Content[0], CLIEssentialPluginGroupNode)
			if err != nil {
				return false, err
			}

			result = append(result, CLIEssentialPluginGroupNode)
			continue
		}
		result = append(result, CLIEssentialPluginGroupNode)
	}
	if !exists {
		result = append(result, newCLIEssentialPluginGroupNode.Content[0])
		persist = true
	}
	CLIEssentialPluginGroupsNode.Content = result
	return persistDiscoverySources || persist, err
}

func removeCLIEssentialPluginGroup(node *yaml.Node, name string) error {
	// check if CLIEssentialPluginGroup name is empty
	if name == "" {
		return errors.New("CLIEssentialPluginGroup name cannot be empty")
	}

	// Find the CLIEssentialPluginGroups node in the yaml node
	keys := []nodeutils.Key{
		{Name: KeyCLI},
		{Name: KeyEssentials},
	}
	CLIEssentialPluginGroupsNode := nodeutils.FindNode(node.Content[0], nodeutils.WithKeys(keys))
	if CLIEssentialPluginGroupsNode == nil {
		return nil
	}
	var CLIEssentialPluginGroups []*yaml.Node
	for _, CLIEssentialPluginGroupNode := range CLIEssentialPluginGroupsNode.Content {
		if index := nodeutils.GetNodeIndex(CLIEssentialPluginGroupNode.Content, "name"); index != -1 && CLIEssentialPluginGroupNode.Content[index].Value == name {
			continue
		}
		CLIEssentialPluginGroups = append(CLIEssentialPluginGroups, CLIEssentialPluginGroupNode)
	}
	CLIEssentialPluginGroupsNode.Content = CLIEssentialPluginGroups
	return nil
}

func getCLIEssentialPluginGroups(node *yaml.Node) ([]*configtypes.EssentialPluginGroup, error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}
	if cfg.CoreCliOptions != nil && cfg.CoreCliOptions.EssentialPluginGroups != nil {
		return cfg.CoreCliOptions.EssentialPluginGroups, nil
	}
	return nil, errors.New("cli discovery sources not found")
}
