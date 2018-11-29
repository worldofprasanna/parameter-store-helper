package main

import (
  "fmt"
  "github.com/jessevdk/go-flags"
  "log"
  "os"
  "github.com/worldofprasanna/parameter-store-helper/cmd"
)

type ApplicationOption struct{
  Profile string `short:"p" long:"profile" description:"AWS Profile to use. If not specified will use the default profile" required:"false"`
}

var option ApplicationOption
var parser = flags.NewParser(&option, flags.Default)

func main() {
  _, err := parser.AddCommand("search",
    "search using path",
    "search the ssm path and retrieve the value",
    &cmd.SearchCmd{})

  if err != nil {
    log.Fatal(err)
  }

  extraArgs, err := parser.Parse()
  if err != nil {
    if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
      os.Exit(0)
    } else {
      fmt.Println("Unused Args: ", extraArgs)
      panic(err)
    }
  }
}

