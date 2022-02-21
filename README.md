# swyftx-cli

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/cupscanteen/swyftx-cli)

> Use [Swyftx] API from the command-line.

[swyftx]: https://swyftx.com.au

## Installation

This works best when run as an executable. All major operating systems have release candidates,
download and run it locally to get started.

Alternatively, you can clone this repository and install it locally.

```shell
git clone https://github.com/cupscanteen/swyftx-cli
cd swyftx-cli
# install a local copy in your GOBIN
go install
```

Requires Go 1.16 and above.

## Usage/Examples

After installing or downloading `swyftx-cli` run the following command in your terminal to get
started.

```shell
swyftx-cli 

# Output (this may be out of sync with latest outputs)
swyftx-cli is a command line interface (CLI) for the Swyftx API.

This CLI uses the Swyftx API for all queries using a mixture of authenticated and unauthenticated endpoints. Subcommands which
require authentication will fail if the 'authenticate' subcommand has not been initialized or if the Access Token has expired.

Usage:
  swyftx-cli [command]

Available Commands:
  authenticate Registers a Swyftx API Key for use throughout the application.
  completion   Generate the autocompletion script for the specified shell
  help         Help about any command
  markets      The subcommands of the markets Swyftx endpoints
  portfolio    Portfolio endpoints for Swyftx

Flags:
  -c, --config string   config file (default is $HOME/.swyftx-cli.yaml)
      --debug           Debug verbose output
  -h, --help            help for swyftx-cli

Additional help topics:
  swyftx-cli orders       Display commands that retrieve Orders information

Use "swyftx-cli [command] --help" for more information about a command.
```

### Authentication

Many endpoints do not require authentication with [Swyftx] and can be used without further
configuration. To access restricted content the `swyftx-cli` must have access to a valid account API
key. To create an API follow the [Swyftx] documentation on how to get one, [here][swyftx-api-docs]

Once you have a valid API key, you can authenticate with [Swyftx] using the `swyfxt-cli`.

```shell
# authenticate with Swyftx
swyftx-cli authenticate --apikey <api-key-here>
```

This will write the key to a file called `.swyftx-cli.yaml` in your `$HOME` or `%userprofile%`
directory depending on your system. Importantly, this command will fetch a valid Access Token from
Swyftx and write it to the same file. Both the API Key and Access Token are required to retrieve
data from protected endpoints.

Alternatively, you can supply a config file with it and a valid Access Token by passing it with each
invocation of the CLI using `swyftx-cli -c <path-to-file>`.

#### Config File Template

```yaml
apikey: 9jVg117muQOb3rdM...truncated
token: eyJhbGciOiJSUzI1Ni..truncated
```

## Running Tests

This application uses the [requests] package to make HTTP calls. A benefit of this is its 
ability to record and replay HTTP sessions. Swyftx-cli uses this extensively to test against.

However, there are some setup costs before running tests. The application needs three environment 
variables, two for running existing tests and another for generating new tests.

**Existing tests**

Set an environment variable called `FAKE_TOKEN` with the bearer token found in any of the 
`portfolio` testdata files found at:

`swyftx-cli/cmd/testdata/portfolio/*.req.txt`. Copy the `<TOKEN>` from `Authorization: Bearer 
<TOKEN>` line in any of those files. Then set a variable named `TESTING_ENABLED` to `1`. 

Run the tests and they should pass. The application must have those environment variables unset 
to function correctly otherwise you may experience subtle issues when attempting to run the 
application when you've finished testing.

**New tests**

To create new HTTP session records, set `SWYFTX_TESTDATA_RECORD=1` and `FAKE_TOKEN=`. You will 
need to have a valid API key and Access Token authenticated via `swyftx-cli authenticate 
--apikey <key>` so that the token is set correctly in the config file.

Once set, running tests will create new HTTP session records inside the `testdata` folder. Run 
the tests as many times a necessary to write your tests but note that data continually changes 
and can break your test output at anytime. I usually run it once to save the output then stop 
recording and write tests from the new data. 

**Unauthenticated endpoints**

If you are writing tests for endpoints which **do not** require authentication then you do not need 
to set `FAKE_TOKEN` as it is not sent with the request.


## Swyftx Documentation

[Documentation](https://docs.swyftx.com.au) for swyftx API endpoints match as closely as possible to
the `swyftx-cli` subcommands.

[swyftx-api-docs]: https://help.swyftx.com.au/en/articles/3825168-how-to-create-an-api-key
[requests]: https://github.com/carlmjohnson/requests
