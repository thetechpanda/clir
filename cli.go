package clir

import (
	"fmt"
	"log"
	"os"

	"github.com/leaanthony/clir/internal/colour"
)

// Cli - The main application object
type Cli struct {
	version        string
	rootCommand    *Command
	defaultCommand *Command
	preRunCommand  func(*Cli) error
}

// Version - Get the Application version string
func (c *Cli) Version() string {
	return c.version
}

// Abort prints the given error and terminates the application
func (c *Cli) Abort(err error) {
	log.Fatal(err)
	os.Exit(1)
}

// AddCommand - Adds a command to the application
func (c *Cli) AddCommand(command *Command) {
	c.rootCommand.addCommand(command)
}

// PrintBanner prints the application banner!
func (c *Cli) PrintBanner() {
	fmt.Printf("%s %s - %s\n\n", colour.YellowString(c.rootCommand.name), colour.RedString(c.version), c.rootCommand.shortdescription)
}

// PrintHelp - Prints the application's help
func (c *Cli) PrintHelp() {
	c.rootCommand.PrintHelp()
}

// Run - Runs the application with the given arguments
func (c *Cli) Run(args ...string) error {
	if c.preRunCommand != nil {
		err := c.preRunCommand(c)
		if err != nil {
			return err
		}
	}
	if len(args) == 0 {
		args = os.Args[1:]
	}
	return c.rootCommand.run(args)
}

// DefaultCommand - Sets the given command as the command to run when
// no other commands given
func (c *Cli) DefaultCommand(defaultCommand *Command) *Cli {
	c.defaultCommand = defaultCommand
	return c
}

// NewSubCommand - Creates a new SubCommand for the application
func (c *Cli) NewSubCommand(name, description string) *Command {
	return c.rootCommand.NewSubCommand(name, description)
}

// PreRun - Calls the given function before running the specific command
func (c *Cli) PreRun(callback func(*Cli) error) {
	c.preRunCommand = callback
}

// BoolFlag - Adds a boolean flag to the root command
func (c *Cli) BoolFlag(name, description string, variable *bool) *Cli {
	c.rootCommand.BoolFlag(name, description, variable)
	return c
}

// StringFlag - Adds a string flag to the root command
func (c *Cli) StringFlag(name, description string, variable *string) *Cli {
	c.rootCommand.StringFlag(name, description, variable)
	return c
}

// IntFlag - Adds an int flag to the root command
func (c *Cli) IntFlag(name, description string, variable *int) *Cli {
	c.rootCommand.IntFlag(name, description, variable)
	return c
}

// Action - Define an action from this command
func (c *Cli) Action(callback Action) *Cli {
	c.rootCommand.Action(callback)
	return c
}

// LongDescription - Sets the long description for the command
func (c *Cli) LongDescription(longdescription string) *Cli {
	c.rootCommand.LongDescription(longdescription)
	return c
}
