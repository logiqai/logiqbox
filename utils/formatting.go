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

package utils

import (
	"encoding/json"
	"fmt"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/dustin/go-humanize"
)

var FlagOut string
var FlagTimeFormat string

func GetTimeAsString(s int64) string {
	t := time.Unix(s, 0)
	switch FlagTimeFormat {
	case "epoch":
		return fmt.Sprintf("%d", s)
	case "RFC3339":
		return t.Format(time.RFC3339)
	default:
		return humanize.Time(t)
	}
}

func PrintResponse(data interface{}) bool {
	if FlagOut == "json" {
		b, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
		return true
	} else if FlagOut == "yaml" {
		b, err := yaml.Marshal(data)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
		return true
	}
	return false
}