/*
Copyright © 2022 Cupscanteen Industries

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
	"encoding/json"
	"fmt"
	"regexp"
)

const SwyftxAPI = "api.swyftx.com.au"

func prettyPrint(i interface{}) string {
	js, _ := json.MarshalIndent(i, "", "\t")
	return string(js)
}

func printer(i interface{}) string {
	js, _ := json.Marshal(i)
	return string(js)
}

func errCheck401(s string) bool {
	match, _ := regexp.MatchString("status: 401", s)
	if match {
		fmt.Println("401 Unauthorized error. You may need to refresh the Access Token")
		fmt.Println("To refresh run: 'excellerate authenticate refresh'")
		fmt.Println("If you have not set the 'apikey' by running 'excellerate authenticate --apikey <apikey>' you will need to do this before continuing.")
		return true
	}
	return false
}

func assetPrinter(result AssetHistoryAll, prettify bool) string {
	if prettify {
		return prettyPrint(result)
	}
	return printer(result)
}
