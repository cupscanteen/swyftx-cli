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
	"github.com/carlmjohnson/requests"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

const SwyftxAPI = "api.swyftx.com.au"

var errorBody string

func StatusChecker(res *http.Response) error {
	err := requests.CheckStatus(200)(res)
	if requests.HasStatusErr(err, 400) {
		if copyErr := requests.ToString(&errorBody)(res); copyErr != nil {
			return copyErr
		}
	}
	if requests.HasStatusErr(err, 401) {
		return errors.New(`401 Unauthorized error. You may need to refresh the Access Token or authenticate using ` + appName)
	}
	return err
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
		log.Println("[!] using test TOKEN [!]")
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
