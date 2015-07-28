package main

import (
  "github.com/jordan0day/envver/auth"
  "github.com/jordan0day/envver/environment"
  "fmt"
  "os"
  "github.com/codegangsta/cli"
)

type Arguments struct {
  ClientId        string
  ClientSecret    string
  TokenUrl        string
  ManagerUrl      string
  ProductName     string
  EnvironmentName string
}

func main() {
  app := cli.NewApp()
  app.Name = "Envver"
  app.Usage = "Retrieve product environment settings from OpenAperture and echoes them to the console ."
  app.Version = "0.0.2"
  app.Flags = []cli.Flag{
    cli.StringFlag{
      Name:   "id, i",
      Usage:  "The OAuth Client ID used to talk to the OpenAperture manager.",
      EnvVar: "OA_CLIENT_ID",
    },
    cli.StringFlag{
      Name:   "secret, s",
      Usage:  "The OAuth Client Secret used to talk to the OpenAperture manager.",
      EnvVar: "OA_CLIENT_SECRET",
    },
    cli.StringFlag{
      Name:   "auth_url, a",
      Usage:  "OAuth token URL used to talk to retreive an OAuth token for the OpenAperture manager.",
      EnvVar: "OA_AUTH_TOKEN_URL",
    },
    cli.StringFlag{
      Name:   "url, u",
      Usage:  "Base URL of the OpenAperture Manager service.",
      EnvVar: "OA_URL",
    },
    cli.StringFlag{
      Name:   "product, p",
      Usage:  "The OpenAperture Product Name of the product for which you're retrieving environment settings.",
      EnvVar: "OA_PRODUCT_NAME",
    },
    cli.StringFlag{
      Name:   "environment, e",
      Usage:  "The OpenAperture Product Environment name for which you're retrieving environment settings.",
      EnvVar: "OA_PRODUCT_ENVIRONMENT_NAME",
    },
  }
  app.Action = func(c *cli.Context) {
    args := getArguments(c)
    if validateArguments(args) != true {
      cli.ShowAppHelp(c)
      return
    }

    fetchVariables(args)
  }

  app.RunAndExitOnError()//os.Args)
}

func getArguments(c *cli.Context) (*Arguments) {
  var args = Arguments{
    ClientId: c.GlobalString("id"),
    ClientSecret: c.GlobalString("secret"),
    TokenUrl: c.GlobalString("auth_url"),
    ManagerUrl: c.GlobalString("url"),
    ProductName: c.String("product"),
    EnvironmentName: c.String("environment"),
  }
  return &args
}

func validateArguments(a *Arguments) (bool) {
  if a == nil {
    return false
  }

  if len(a.ClientId) <= 0 {
    fmt.Println("Please specify the OAuth Client ID, either via the --id flag or by setting the 'OA_CLIENT_ID' environmental variable.")
    return false
  }

  if len(a.ClientSecret) <= 0 {
    fmt.Println("Please specify the OAuth Client Secret, either via the --secret flag or by setting the 'OA_CLIENT_SECRET' environmental variable.")
    return false
  }

  if len(a.TokenUrl) <= 0 {
    fmt.Println("Please specify the OAUth Access Token URL, either via the --auth_url flag or by setting the 'OA_AUTH_TOKEN_URL' environmental variable.")
    return false
  }

  if len(a.ManagerUrl) <= 0 {
    fmt.Println("Please specify the URL of the OpenAperture Manager service, either via the --url flag or by setting the 'OA_URL' environment variable.")
  }

  if len(a.ProductName) <= 0 {
    fmt.Println("Please specify the Product Name, either via the --product flag or by setting the 'OA_PRODUCT_NAME' environmental variable.")
    return false
  }

  return true
}

func fetchVariables(args *Arguments) {
  var variables []environment.EnvironmentVariable

  credentials := auth.ClientCredentials{
    ClientId: args.ClientId,
    ClientSecret: args.ClientSecret,
  }

  authResponse, err := auth.GetAuthToken(credentials, args.TokenUrl)

  if err != nil {
    fmt.Printf("Couldn't retrieve auth token: %v\n", err)
    os.Exit(1)
  }

  if len(args.EnvironmentName) > 0 {
    variables, err = environment.GetEnvironmentVariables(args.ProductName, args.EnvironmentName, authResponse.AccessToken, args.ManagerUrl)
  } else {
    variables, err = environment.GetProductEnvironmentVariables(args.ProductName, authResponse.AccessToken, args.ManagerUrl)
  }

  if err != nil {
    fmt.Printf("Error retrieving variables: %v\n", err)
    os.Exit(1)
  }

  for _, v := range variables {
    fmt.Printf("%s=%s\n", v.Name, v.Value)
  }
}