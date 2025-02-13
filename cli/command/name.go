package command

import (
	"errors"
	"strings"

	"github.com/taubyte/dreamland/cli/flags"
	"github.com/urfave/cli/v2"
)


func NameWithDefault(c *cli.Command, def string) {
	flag := flags.Name
	flag.DefaultText = def
	flag.Value = def

	attachName(c, &flag)
}

func Name(c *cli.Command) {
	attachName(c, &flags.Name)
}

func attachName(c *cli.Command, flag cli.Flag) {
	c.Flags = append(c.Flags, flag)

	// If nothing is passed in as c.ArgsUsage, then Name is set to just name.
	if len(c.ArgsUsage) == 0 {
		// Same result as NameWithDefault function
		c.ArgsUsage = "name"
	} else {
		// Attaches name string to name in the Name object
		// Is the inclusion of the "name," string here for sorting purposes in the getName function?
		c.ArgsUsage = "name," + c.ArgsUsage
	}

	action := c.Action

	c.Action = func(ctx *cli.Context) error {
		name, err := getName(ctx)
		if err != nil {
			return err
		}
		ctx.Set("name", name)
		return action(ctx)
	}

}

// when name is args0 or flag -n this method will get
// or return an error
func getName(c *cli.Context) (name string, err error) {
	name = c.Args().First()
	if len(name) == 0 {
		name = c.String("name")
		if len(name) == 0 {
			err = errors.New("Please provide a name")
			return
		}
	} else {
		if strings.HasPrefix(c.Args().Get(1), "-") {
			err = errors.New("Parse arguments failed: write [arguments] after -flags")
			return
		}
	}

	return
}
