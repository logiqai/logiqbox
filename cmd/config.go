/*
Copyright © 2020 Logiq.ai <cli@logiq.ai>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"

	"github.com/logiqai/logiqctl/services"

	"github.com/logiqai/logiqctl/utils"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config SUBCOMMAND",
	Short: "Modify or View logiqctl configuration options",
	Long: `
# View current context
	logiqctl config view

# Set default cluster
	logiqctl set-cluster END-POINT

# Set default context
	logiqctl set-context namespace

# Set default context with the help of an interactive prompt
	logiqctl set-context

`,
}

func NewSetClusterCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "set-cluster",
		Example: "logiqctl set-cluster END-POINT",
		Short:   "Sets a logiq cluster entry point",
		Long: `
Sets the cluster, a valid logiq cluster end point is required for all the operations
		`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				fmt.Println("Incorrect Usage")
				fmt.Println(cmd.Example)
				return
			}
			viper.Set(utils.KeyCluster, args[0])
			viper.Set(utils.KeyPort, utils.DefaultPort)
			err := viper.WriteConfig()
			if err != nil {
				fmt.Print(err)
				return
			}
			printCluster()
		},
	}
	return cmd
}

func printCluster() {
	cluster := viper.GetString(utils.KeyCluster)
	if cluster != "" {
		fmt.Printf("Cluster Endpoint set to: %s\n", cluster)
	} else {
		fmt.Println("Default Cluster not set")
	}
}

func printNamespace() {
	ns := viper.GetString(utils.KeyNamespace)
	if ns != "" {
		fmt.Printf("Default Context set to: %s\n", ns)
	} else {
		fmt.Println("Default Context not set")
	}
}

func NewViewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "view",
		Short: "View current defaults",
		Run: func(cmd *cobra.Command, args []string) {
			printCluster()
			printNamespace()
		},
	}
}

func NewSetContextCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "set-context",
		Example: "set-context NAMESPACE",
		Short:   "Sets a default namespace in logiqctl",
		Long: `
This will the default context for all the operations.
		`,
		PreRun: preRun,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 1 {
				fmt.Printf("Incorrect usage")
				fmt.Println(cmd.Example)
				return
			}
			if len(args) == 1 {
				setContext(args[0])
				printNamespace()
				return
			}
			selectedNs, err := services.RunSelectNamespacePrompt()
			if err != nil {
				fmt.Printf("Incorrect usage")
				fmt.Println(cmd.Example)
				return
			}
			setContext(selectedNs)
			printNamespace()
		},
	}
	return cmd
}

func setContext(arg string) {
	viper.Set(utils.KeyNamespace, arg)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Print(err)
		return
	}
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(NewSetClusterCommand())
	configCmd.AddCommand(NewSetContextCommand())
	configCmd.AddCommand(NewViewCommand())

}