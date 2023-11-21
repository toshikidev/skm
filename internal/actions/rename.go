package actions

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/toshikidev/skm/internal/utils"
	"gopkg.in/urfave/cli.v1"
)

func Rename(c *cli.Context) error {
	env := utils.MustGetEnvironment(c)
	var alias, newAlias string

	if c.NArg() == 2 {
		alias = c.Args().Get(0)
		newAlias = c.Args().Get(1)

		err := os.Rename(filepath.Join(env.StorePath, alias), filepath.Join(env.StorePath, newAlias))

		if err == nil {
			color.Green("%s SSH key [%s] renamed to [%s]", utils.CheckSymbol, alias, newAlias)
		} else {
			color.Red("%s Failed to rename SSH key!", utils.CrossSymbol)
		}
	} else {
		color.Red("%s Please input current alias name and new alias name", utils.CrossSymbol)
	}

	return nil
}
