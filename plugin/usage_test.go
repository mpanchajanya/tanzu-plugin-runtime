// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

func TestUsageFunc(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Error(err)
	}
	c := make(chan []byte)
	go readOutput(t, r, c)

	// Set up for our test
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()
	os.Stdout = w
	os.Stderr = w

	cmd := &cobra.Command{
		Use:   "Sub1",
		Short: "Sub1 description",
	}
	err = UsageFunc(cmd)
	assert.Nil(t, err)

	err = w.Close()
	assert.Nil(t, err)

	got := <-c
	assert.Contains(t, string(got), "Usage:")
}

func SampleTestCommand(t *testing.T, target types.Target) *cobra.Command {
	var pluginsCmd = &cobra.Command{
		Use:   "plugin",
		Short: "Plugin tests",
	}

	var fetchCmd = &cobra.Command{
		Use:   "fetch",
		Short: "Fetch the plugin tests",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("fetch")
			return nil
		},
	}

	var pushCmd = &cobra.Command{
		Use:   "push",
		Short: "Push the plugin tests",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("push")
			return nil
		},
	}

	var descriptor = PluginDescriptor{
		Name:        "test",
		Target:      target,
		Aliases:     []string{"t"},
		Description: "Test the CLI",
		Group:       AdminCmdGroup,
		Version:     "v1.1.0",
		BuildSHA:    "1234567",
	}

	var local string

	fetchCmd.Flags().StringVarP(&local, "local", "l", "", "path to local repository")
	_ = fetchCmd.MarkFlagRequired("local")

	p, err := NewPlugin(&descriptor)
	assert.Nil(t, err)

	p.AddCommands(
		fetchCmd,
		pushCmd,
		pluginsCmd,
	)

	return p.Cmd
}

func TestGlobalTestPluginCommandHelpText(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Error(err)
	}
	c := make(chan []byte)
	go readOutput(t, r, c)

	// Set up for our test
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()
	os.Stdout = w
	os.Stderr = w

	cmd := SampleTestCommand(t, types.TargetGlobal)
	cmd.SetArgs([]string{"-h"})

	err = cmd.Execute()
	assert.Nil(t, err)

	err = w.Close()
	assert.Nil(t, err)

	got := string(<-c)

	expected := `Test the CLI

Usage:
  tanzu test [command]

Aliases:
  test, t

Available Commands:
  fetch         Fetch the plugin tests
  push          Push the plugin tests

Flags:
  -h, --help   help for test

Additional help topics:
  test plugin        Plugin tests

Use "tanzu test [command] --help" for more information about a command.
`
	assert.Equal(t, expected, got)
}

func TestKubernetesTestPluginCommandHelpText(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Error(err)
	}
	c := make(chan []byte)
	go readOutput(t, r, c)

	// Set up for our test
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()
	os.Stdout = w
	os.Stderr = w

	cmd := SampleTestCommand(t, types.TargetK8s)
	cmd.SetArgs([]string{"-h"})

	err = cmd.Execute()
	assert.Nil(t, err)

	err = w.Close()
	assert.Nil(t, err)

	got := string(<-c)

	expected := `Test the CLI

Usage:
  tanzu test [command]
  tanzu kubernetes test [command]

Aliases:
  test, t

Available Commands:
  fetch         Fetch the plugin tests
  push          Push the plugin tests

Flags:
  -h, --help   help for test

Additional help topics:
  test plugin        Plugin tests

Use "tanzu test [command] --help" for more information about a command.
Use "tanzu kubernetes test [command] --help" for more information about a command.
`
	assert.Equal(t, expected, got)
}

func TestMissionControlTestPluginCommandHelpText(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Error(err)
	}
	c := make(chan []byte)
	go readOutput(t, r, c)

	// Set up for our test
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()
	os.Stdout = w
	os.Stderr = w

	// Prepare the root command with MissionControl target
	rootCmd := SampleTestCommand(t, types.TargetTMC)

	// Set the arguments as if the user typed them in the command line
	rootCmd.SetArgs([]string{"--help"})

	// Execute the command which will trigger the help
	err = rootCmd.Execute()
	assert.Nil(t, err)

	err = w.Close()
	assert.Nil(t, err)

	got := string(<-c)

	expected := `Test the CLI

Usage:
  tanzu mission-control test [command]

Aliases:
  test, t

Available Commands:
  fetch         Fetch the plugin tests
  push          Push the plugin tests

Flags:
  -h, --help   help for test

Additional help topics:
  test plugin        Plugin tests

Use "tanzu mission-control test [command] --help" for more information about a command.
`
	assert.Equal(t, expected, got)
}
