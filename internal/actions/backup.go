package actions

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/toshikidev/skm/internal/utils"
	"github.com/toshikidev/skm/pkg/lib"
	"gopkg.in/urfave/cli.v1"
)

func Backup(c *cli.Context) error {
	env := utils.MustGetEnvironment(c)
	if c.Bool("restic") {
		lib.MustHaveRestic(env)
		resticCfg := lib.MustLoadOrCreateResticSettings(env, c)
		lib.EnsureInitializedResticRepo(resticCfg, env)
		result := utils.Execute(env.StorePath, env.ResticPath, "backup", ".", "--password-file", resticCfg.PasswordFile, "--repo", resticCfg.Repository)
		if result {
			color.Green("%s Backup to %s complete", utils.CheckSymbol, resticCfg.Repository)
		}
		return nil
	}
	fileName := utils.GetBakFileName()
	dstFile := filepath.Join(os.Getenv("HOME"), fileName)

	result := utils.Execute(env.StorePath, "tar", "-czvf", dstFile, ".")
	if result {
		color.Green("%s All SSH keys backup to: %s", utils.CheckSymbol, dstFile)
	}

	return nil
}
