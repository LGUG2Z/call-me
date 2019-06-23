package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/LGUG2Z/call-me/client"
	"github.com/LGUG2Z/call-me/client/operations"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"gopkg.in/urfave/cli.v1"
)

var (
	Version string
	Commit  string
)

var ErrCommandRequiresAnArgument = fmt.Errorf("this command requires an argument")

func ErrEnvVarMustBeSet(envVar string) error { return fmt.Errorf("%s must be set", envVar) }

func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("call-me-client version %s (commit %s)\n", c.App.Version, Commit)
	}

	host := os.Getenv("HOST")
	apiKey := os.Getenv("API_KEY")

	app := cli.NewApp()

	app.Name = "call-me-client"
	app.Usage = "Command line client for call-me, a ci-agnostic test environment orchestrator"
	app.EnableBashCompletion = true
	app.Compiled = time.Now()
	app.Version = Version
	app.Authors = []cli.Author{{
		Name:  "J. Iqbal",
		Email: "jade@beamery.com",
	}}

	app.Before = func(c *cli.Context) error {
		if len(host) == 0 {
			return ErrEnvVarMustBeSet("HOST")
		}

		if len(apiKey) == 0 {
			return ErrEnvVarMustBeSet("API_KEY")
		}

		return nil
	}

	request := cli.Command{
		Name:  "request",
		Usage: "Request a slot for an environment",
		Action: cli.ActionFunc(func(c *cli.Context) error {
			if !c.Args().Present() {
				return ErrCommandRequiresAnArgument
			}

			environment := c.Args().First()

			callMe := client.NewHTTPClientWithConfig(strfmt.Default, &client.TransportConfig{Host: host})
			apiKeyHeaderAuth := httptransport.APIKeyAuth("X-API-KEY", "header", os.Getenv("API_KEY"))

			getParams := operations.NewGetMaybeParams()
			getParams.SetEnvironment(environment)

			var forbidden operations.GetMaybeForbidden

			start := time.Now()
			for {
				_, err := callMe.Operations.GetMaybe(getParams, apiKeyHeaderAuth)
				if err != nil {
					switch err.Error() {
					case forbidden.Error():
						fmt.Printf("waiting for free slot on environment: %s\n", environment)
						time.Sleep(time.Second * 5)
					default:
						return err
					}
				} else {
					break
				}
			}

			postParams := operations.NewPostMaybeParams()
			postParams.SetEnvironment(environment)

			if _, err := callMe.Operations.PostMaybe(postParams, apiKeyHeaderAuth); err != nil {
				return err
			}

			fmt.Printf("acquired slot for environment: %s\n", environment)
			fmt.Printf("waited %s\n", time.Since(start))
			return nil
		}),
	}

	release := cli.Command{
		Name:  "release",
		Usage: "Release a slot for an environment",
		Action: cli.ActionFunc(func(c *cli.Context) error {
			if !c.Args().Present() {
				return ErrCommandRequiresAnArgument
			}

			environment := c.Args().First()

			callMe := client.NewHTTPClientWithConfig(strfmt.Default, &client.TransportConfig{Host: host})
			apiKeyHeaderAuth := httptransport.APIKeyAuth("X-API-KEY", "header", os.Getenv("API_KEY"))

			deleteParams := operations.NewDeleteMaybeParams()
			deleteParams.SetEnvironment(environment)

			if _, err := callMe.Operations.DeleteMaybe(deleteParams, apiKeyHeaderAuth); err != nil {
				return err
			}

			fmt.Printf("slot has been freed for environment: %s", environment)
			return nil
		}),
	}

	app.Commands = []cli.Command{request, release}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
