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
	"context"
	"errors"
	"fmt"
	"github.com/carlmjohnson/requests"
	"github.com/spf13/cobra"
	"net/http"
)

var (
	revokeApikey      bool
	revokeRefId       string
	revokeAccessToken bool
	getKeys           bool
	keyPretty         bool
)

func init() {
	authenticateCmd.AddCommand(revokeCmd)
	// API key revocation flags
	revokeCmd.Flags().BoolVarP(&revokeApikey, "apikey", "", false, "revoke current API key")
	revokeCmd.Flags().StringVarP(&revokeRefId, "ref-id", "", "", "API key reference ID")
	// Get the keys and reference id's
	revokeCmd.Flags().BoolVarP(&getKeys, "get-keys", "", false, "list the all API keys for the authorized account (must use an API key with 'app.api.read' permissions)")
	// Access Token
	revokeCmd.Flags().BoolVarP(&revokeAccessToken, "token", "", false, "revoke the current Access Token")

	revokeCmd.Flags().BoolVarP(&keyPretty, "pretty", "", false, "Pretty print the response")
}

var revokeCmd = &cobra.Command{
	Use:   "revoke",
	Short: `Revoke an accounts key or token. Must have already authenticated the account with ` + appName,
	Long: `This subcommand can revoke a valid API key or Access Token.

To revoke the API key, the key used to do the revoking must have 'app.api.revoke' permissions. You must pass the 'keyRef' with the
revocation command. The 'keyRef' can be gathered by getting the list of keys from the account first. 
To do that in ` + appName + ` run '` + appName + ` authenticate revoke --get-keys' and note the 'keyRef' that matches the key label you wish to revoke.

The ` + appName + ` must have already authenticated successfully with Swyftx by using the 'authenticate' subcommand.`,
	RunE: revokeIt,
}

func revokeIt(cmd *cobra.Command, args []string) error {

	token, err := AccessTokenGetter()
	cobra.CheckErr(err)

	if getKeys {
		result, err := getApiKeys(&client, token)
		cobra.CheckErr(err)
		stdout := GenericPrinter(result, keyPretty)
		fmt.Println(stdout)
		return nil
	}

	if revokeApikey {
		if revokeRefId == "" {
			cobra.CheckErr(errors.New(`must provide reference ID that matches the current API key in use. get the ID by running '` + appName + ` authenticate revoke --get-keys'`))
		}
		err := revokeApiKeyRequest(&client, token)
		cobra.CheckErr(err)
	}

	if revokeAccessToken {
		err := logoutAccessToken(&client, token)
		cobra.CheckErr(err)
	}
	err = cmd.Help()
	if err != nil {
		return err
	}
	return nil
}

func logoutAccessToken(c *http.Client, token string) error {
	url := "/auth/logout/"
	err := requests.
		URL(url).
		Method(http.MethodPost).
		Client(c).
		Host(SwyftxAPI).
		ContentType("application/json").
		Header("Authorization", fmt.Sprintf("Bearer %s", token)).
		AddValidator(StatusChecker).
		Fetch(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Access Token revoked. Re-authentication is required before attempting to access any restricted endpoints.")
	return nil
}

// revokeApiKeyRequest
func revokeApiKeyRequest(c *http.Client, token string) error {
	url := "/user/apiKeys/revoke/"
	body := ApiKeyRevocation{
		KeyRef: revokeRefId,
		KeyId:  "",
	}
	err := requests.
		URL(url).
		Method(http.MethodPost).
		Client(c).
		Host(SwyftxAPI).
		ContentType("application/json").
		Header("Authorization", fmt.Sprintf("Bearer %s", token)).
		BodyJSON(&body).
		AddValidator(StatusChecker).
		Fetch(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("API key revoked. Re-authentication with a valid API key is required before attempting to access any restricted endpoints.")
	return nil
}

type ApiKeyRevocation struct {
	KeyId  string `json:"id"`
	KeyRef string `json:"keyRef"`
}

func getApiKeys(c *http.Client, token string) (*[]ApiKeys, error) {
	url := "/user/apiKeys/"
	var response *[]ApiKeys
	err := requests.
		URL(url).
		Client(c).
		Host(SwyftxAPI).
		ContentType("application/json").
		Header("Authorization", fmt.Sprintf("Bearer %s", token)).
		//BodyJSON(&apiKey).
		ToJSON(&response).
		AddValidator(StatusChecker).
		Fetch(context.Background())
	if err != nil {
		return nil, err
	}

	return response, err
}

type ApiKeys struct {
	KeyRef  string `json:"keyRef,omitempty"`
	Id      string `json:"id"`
	Label   string `json:"label,omitempty"`
	Scope   string `json:"scope,omitempty"`
	Created int    `json:"created,omitempty"`
}
