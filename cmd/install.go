/*
Copyright © 2022 Tiger Wang <tiger@tensorsmart.com>
*/
package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tigerinus/good/common"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:     "install [package]",
	Aliases: []string{"hello", "add", "get"},
	Args:    cobra.MinimumNArgs(1),
	Short:   "Install a package named by its import path",
	Run: func(cmd *cobra.Command, args []string) {
		packageName := args[0]
		installPath, err := createPackageInstallPath(packageName)
		if err != nil {
			_logger.Error(err.Error())
			os.RemoveAll(installPath)
			os.Exit(1)
		}

		execCmd := common.GoInstallCmd(packageName, installPath)

		defer func() {
			if _, err := os.Stat(installPath); err != nil {
				return
			}

			_logger.Info("good: cleaning up mod cache at %s...\n", installPath)

			execCmd = common.GoCleanModCacheCmd(installPath)
			if err := execCmd.Run(); err != nil {
				_logger.Debug(err.Error())
			}
		}()

		_logger.Info("good: installing to %s...\n", installPath)
		if err := execCmd.Run(); err != nil {
			_logger.Error(err.Error())
			os.RemoveAll(installPath)
			os.Exit(1)
		}

		binPath := filepath.Join(installPath, "bin")
		items, err := ioutil.ReadDir(binPath)
		if err != nil {
			_logger.Error(err.Error())
			os.RemoveAll(installPath)
			os.Exit(1)
		}

		targetPath := viper.GetString(common.ConfigKeyLocalBinPath)
		for _, item := range items {
			if item.IsDir() {
				continue
			}

			targetFile := filepath.Join(targetPath, item.Name())

			if _, err := os.Stat(targetFile); err == nil {
				if err := os.Remove(targetFile); err != nil {
					_logger.Debug(err.Error())
				}
			}

			if err := os.Symlink(filepath.Join(binPath, item.Name()), targetFile); err != nil {
				_logger.Error(err.Error())
				os.RemoveAll(installPath)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createPackageInstallPath(packageName string) (string, error) {
	installPath := filepath.Join(viper.GetString(common.ConfigKeyInstallRootPath), packageName)

	_logger.Debug("good: creating install path %s...", installPath)
	if err := os.MkdirAll(installPath, 0o755); err != nil {
		return "", err
	}

	// create manifest file under installPath
	manifestFilepath := filepath.Join(installPath, common.ManifestFileName)
	_logger.Debug("good: creating manifest file %s...", manifestFilepath)
	if err := ioutil.WriteFile(manifestFilepath, []byte(packageName), 0o600); err != nil {
		return "", err
	}

	return installPath, nil
}
