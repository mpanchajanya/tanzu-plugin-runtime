package config

import (
	"fmt"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/nodeutils"
	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	"gopkg.in/yaml.v3"
)

// GetCLIEssentials retrieves cli discovery sources
func GetCLIEssentials() (*configtypes.Essentials, error) {
	// Retrieve client config node
	node, err := getClientConfigNode()
	if err != nil {
		return nil, err
	}

	return getCLIEssentials(node)
}

// SetCLIEssentials add or update CLICorePluginGroup
func SetCLIEssentials(c *configtypes.Essentials) error {
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}
	// Add or update the CLICorePluginGroup
	persist, err := setCLIEssentials(node, c)
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

func getCLIEssentials(node *yaml.Node) (*configtypes.Essentials, error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}

	if cfg != nil && cfg.CoreCliOptions != nil && cfg.CoreCliOptions.Essentials != nil {
		return cfg.CoreCliOptions.Essentials, nil
	}

	return nil, fmt.Errorf("essentials not found")
}

func setCLIEssentials(node *yaml.Node, e *configtypes.Essentials) (persist bool, err error) {
	// Get Patch Strategies from config metadata
	patchStrategies, err := GetConfigMetadataPatchStrategy()
	if err != nil {
		patchStrategies = make(map[string]string)
	}

	// Convert CLICorePluginGroup to node
	newCLICorePluginGroupNode, err := convertObjectToNode(e)
	if err != nil {
		return persist, err
	}

	// Find the CLICorePluginGroups node from the root node
	keys := []nodeutils.Key{
		{Name: KeyCLI, Type: yaml.MappingNode},
		{Name: KeyEssentials, Type: yaml.MappingNode},
	}

	CLICorePluginGroupsNode := nodeutils.FindNode(node.Content[0], nodeutils.WithForceCreate(), nodeutils.WithKeys(keys))
	if CLICorePluginGroupsNode == nil {
		return persist, err
	}

	// replace the nodes as per patch strategy
	_, err = nodeutils.DeleteNodes(newCLICorePluginGroupNode.Content[0], CLICorePluginGroupsNode, nodeutils.WithPatchStrategyKey(KeyEssentials), nodeutils.WithPatchStrategies(patchStrategies))
	if err != nil {
		return false, err
	}

	persist, err = nodeutils.MergeNodes(newCLICorePluginGroupNode.Content[0], CLICorePluginGroupsNode)
	if err != nil {
		return false, err
	}

	return persist, err
}
