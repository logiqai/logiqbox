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

	"github.com/spf13/cobra"
)

var application, process, labels string

var tailExample = `
Tail all logs 
# logiqctl tail

Tail specific application logs
# logiqctl tail -p=logiq-flash-2 logiq-flash

Tail specific application and filter by its process id
# logiqctl logs -p=logiq-flash-2 logiq-flash

`

// tailCmd represents the tail command
var tailCmd = &cobra.Command{
	Use:     "tail",
	Aliases: []string{"t"},
	Example: tailExample,
	Short:   "Stream logs from logiq Observability stack",
	Long: `Tail command enables you to see logs from applications in realtime, See examples for options.

`,
	Run: func(cmd *cobra.Command, args []string) {
		var labelsArray []string
		var appName string
		var procId string
		app, err := services.RunSelectApplicationForNamespacePrompt(true)
		handleError(err)
		if app != nil {
			appName = app.Name
			process, err := services.RunSelectProcessesForNamespaceAndAppPrompt(app.Name, true)
			handleError(err)
			if process != nil {
				procId = process.ProcID
			} else {
				procId = "*"
			}
		} else {
			appName = "*"
		}
		fmt.Println("Crunching data for you...")
		services.Tail(appName, procId, labelsArray)

		return
	},
}

func init() {

	rootCmd.AddCommand(tailCmd)
	tailCmd.Flags().StringVarP(&process, "process", "p", "", `Filter logs by process id`)
	tailCmd.Flags().StringVarP(&labels, "labels", "l", "", `Filter logs by label`)

}
