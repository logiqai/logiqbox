package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/logiqai/logiqctl/utils"

	"github.com/manifoldco/promptui"

	"github.com/logiqai/logiqctl/api/v1/query"
)

var templates = &promptui.SelectTemplates{
	Label:    "{{ . }}?",
	Active:   "\U000000BB {{ .Name | green }} ({{ .Details | red }})",
	Inactive: "  {{ .Name | cyan }} ({{ .Details | red }})",
	Selected: "Application {{ .Name | red }} selected",
}

const (
	FMT = "%-24s | %-16s | %-16s | %-16s | %-16s | %s\n"
)

func printSyslogHeader() {
	fmt.Printf(FMT, "Timestamp", "Namespace", "Application", "Process/Pod", "Facility", "Log message")
}

func printSyslogMessage(logMap map[string]interface{}, output string) {
	if output == OUTPUT_COLUMNS {
		fmt.Printf(FMT,
			logMap["timestamp"],
			logMap["namespace"],
			logMap["app_name"],
			logMap["proc_id"],
			logMap["facility_string"],
			logMap["message"],
		)
	} else if output == OUTPUT_JSON {
		v, err := json.Marshal(logMap)
		if err == nil {
			fmt.Printf("%s\n", v)
		} else {
			fmt.Printf("Error marshalling JSON %v", logMap)
			os.Exit(-1)
		}
	} else {
		fmt.Printf("%s %s %s %s %s %s\n",
			logMap["timestamp"],
			logMap["namespace"],
			logMap["app_name"],
			logMap["proc_id"],
			logMap["facility_string"],
			logMap["message"],
		)
		if utils.GetLineBreak() {
			fmt.Println()
		}
	}
}

func printSyslogMessageForType(log *query.SysLogMessage, output string) {
	if output == OUTPUT_COLUMNS {
		fmt.Printf("%s | %s | %s | %s | %s\n\n",
			log.Timestamp,
			log.SeverityString,
			log.FacilityString,
			log.ProcID,
			log.Message,
		)
	} else if output == OUTPUT_RAW {
		fmt.Printf("%s %s %s %s %s %s %s\n",
			log.Timestamp,
			log.SeverityString,
			log.FacilityString,
			log.Namespace,
			log.AppName,
			log.ProcID,
			log.Message,
		)
		if utils.GetLineBreak() {
			fmt.Println()
		}
	} else if output == OUTPUT_JSON {
		v, err := json.Marshal(log)
		if err == nil {
			fmt.Printf("%s\n", v)
		} else {
			fmt.Printf("Error marshalling JSON %v", *log)
			os.Exit(-1)
		}
	}
}
