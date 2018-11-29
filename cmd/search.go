package cmd

import "fmt"

type SearchCmd struct {
  BasePath string `short:"b" long:"basepath" description:"Base path to search recursively" required:"true"`
}

func (cmd *SearchCmd) Execute(args []string) error {
  cmd.search(cmd.BasePath)
  return nil
}

func (cmd *SearchCmd) search(path string) {
  fmt.Println("Search will start from the path "+ path)
}
