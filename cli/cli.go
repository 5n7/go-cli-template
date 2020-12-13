package cli

import (
	"fmt"
	"os"

	hello "github.com/skmatz/go-cli-template"
)

// CLI represents a CLI.
type CLI struct{}

func New() *CLI {
	return &CLI{}
}

// Options represents options for the CIL.
type Options struct {
	Name string
}

// Run runs the CLI.
func (c *CLI) Run(opt Options) error {
	if opt.Name == "" {
		return fmt.Errorf("name is empty")
	}

	hello.Hello(os.Stdout, opt.Name)
	return nil
}
