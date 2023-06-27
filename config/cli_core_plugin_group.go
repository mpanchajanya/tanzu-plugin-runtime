// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/nodeutils"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

const DefaultEssentialName = "vmware-cli/core"
const TanzuCLIEssentialPluginGroupName = "TANZU_CLI_ESSENTIAL_PLUGIN_GROUP_NAME"
const TanzuCLIEssentialPluginGroupVersion = "TANZU_CLI_ESSENTIAL_PLUGIN_GROUP_VERSION"

// GetCLIEssential retrieves cli discovery sources
func GetCLIEssential() (*configtypes.Essentials, error) {
	// Retrieve client config node
	node, err := getClientConfigNode()
	if err != nil {
		return nil, err
	}

	core, err := getCLIEssential(node)
	if err != nil {
		return nil, err
	}

	if core == nil {

		// Setup Default Essential Plugin groups
		defaultEssentialName := DefaultEssentialName
		EssentialName := os.Getenv(TanzuCLIEssentialPluginGroupName)
		if EssentialName != "" {
			defaultEssentialName = EssentialName
		}

		var defaultEssentialVersion string
		EssentialVersion := os.Getenv(TanzuCLIEssentialPluginGroupVersion)
		if EssentialVersion != "" {
			defaultEssentialVersion = EssentialVersion
		}

		defaultEssential := &configtypes.Essentials{
			Name:    defaultEssentialName,
			Version: defaultEssentialVersion,
		}
		return defaultEssential, nil

	}

	return core, nil

}

// SetCLIEssential add or update CLIEssential
func SetCLIEssential(c *configtypes.Essential) error {
	fmt.Println("SetCLIEssential", c.Name, c.Version, c.LastUpdatedTimestamp)
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}
	// Add or update the CLIEssential
	persist, err := setCLIEssential(node, c)
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

// DeleteCLIEssential delete a CLIEssential by name
func DeleteCLIEssential() error {
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()

	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}

	ctx, err := getCLIEssential(node)
	if err != nil {
		return err
	}

	err = removeCLIEssential(node, ctx.Name)
	if err != nil {
		return err
	}

	return persistConfig(node)
}

func getCLIEssential(node *yaml.Node) (*configtypes.Essential, error) {

	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}

	if cfg != nil && cfg.CoreCliOptions != nil {
		if cfg.CoreCliOptions.Essential != nil {
			return cfg.CoreCliOptions.Essential, nil
		}
	}

	return nil, fmt.Errorf("CLIEssential not found")
}

func setCLIEssential(node *yaml.Node, ctx *configtypes.Essential) (persist bool, err error) {
	// Check if ctx.Name is empty
	if ctx.Name == "" {
		return false, errors.New("CLIEssential name cannot be empty")
	}

	//// Get Patch Strategies from config metadata
	//patchStrategies, err := GetConfigMetadataPatchStrategy()
	//if err != nil {
	//	patchStrategies = make(map[string]string)
	//}

	// Convert CLIEssential to node
	newCLIEssentialNode, err := convertObjectToNode(ctx)
	if err != nil {
		return persist, err
	}

	// Find the CLIEssentials node from the root node
	keys := []nodeutils.Key{
		{Name: KeyCLI, Type: yaml.MappingNode},
		{Name: KeyEssential, Type: yaml.MappingNode},
	}

	CLIEssentialsNode := nodeutils.FindNode(node.Content[0], nodeutils.WithForceCreate(), nodeutils.WithKeys(keys))
	if CLIEssentialsNode == nil {
		return persist, err
	}

	// replace the nodes as per patch strategy
	//_, err = nodeutils.DeleteNodes(newCLIEssentialNode.Content[0], CLIEssentialsNode, nodeutils.WithPatchStrategyKey(KeyEssential), nodeutils.WithPatchStrategies(patchStrategies))
	//if err != nil {
	//	return false, err
	//}
	persist, err = nodeutils.MergeNodes(newCLIEssentialNode.Content[0], CLIEssentialsNode)
	if err != nil {
		return false, err
	}

	return persist, err
}

func removeCLIEssential(node *yaml.Node, name string) error {
	// check if CLIEssential name is empty
	if name == "" {
		return errors.New("CLIEssential name cannot be empty")
	}

	// Find the CLIEssentials node in the yaml node
	keys := []nodeutils.Key{
		{Name: KeyCLI},
		{Name: KeyEssential},
	}
	CLIEssentialsNode := nodeutils.FindNode(node.Content[0], nodeutils.WithKeys(keys))
	if CLIEssentialsNode == nil {
		return nil
	}
	var CLIEssentials []*yaml.Node
	for _, CLIEssentialNode := range CLIEssentialsNode.Content {
		if index := nodeutils.GetNodeIndex(CLIEssentialNode.Content, "name"); index != -1 && CLIEssentialNode.Content[index].Value == name {
			continue
		}
		CLIEssentials = append(CLIEssentials, CLIEssentialNode)
	}
	CLIEssentialsNode.Content = CLIEssentials
	return nil
}
