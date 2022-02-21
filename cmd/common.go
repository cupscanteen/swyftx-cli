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
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"regexp"
)

const SwyftxAPI = "api.swyftx.com.au"

// todo(dm): use requests validator instead of this.
func errCheck401(s string) bool {
	match, _ := regexp.MatchString("status: 401", s)
	if match {
		fmt.Println("401 Unauthorized error. You may need to refresh the Access Token")
		fmt.Println("To refresh run: 'swyftx-cli authenticate refresh'")
		fmt.Println("If you have not set the 'apikey' by running 'swyftx-cli authenticate --apikey <apikey>' you will need to do this before continuing.")
		return true
	}
	return false
}

func GenericPrinter(r interface{}, p bool) string {
	if p {
		return PrettyPrinter(r)

	}
	return Printer(r)
}
func PrettyPrinter(i interface{}) string {
	js, _ := json.MarshalIndent(i, "", "  ")
	return string(js)
}
func Printer(i interface{}) string {
	js, _ := json.Marshal(i)
	return string(js)
}

// AccessTokenGetter will attempt to fetch the Access Token from the
// configuration file. If we are testing it will get the valid mock access
// token.
func AccessTokenGetter() (string, error) {
	if os.Getenv("TESTING_ENABLED") != "" {
		return os.Getenv("FAKE_TOKEN"), nil
	}
	token := viper.GetString("token")
	if token == "" {
		err := errors.New("access token missing from configuration file")
		if err != nil {
			return "", err
		}
	}
	return token, nil
}
