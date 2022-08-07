/*
Copyright © 2022 Tiger Wang <tiger@tensorsmart.com>
*/
package cmd

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tigerinus/good/common"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:     "uninstall",
	Aliases: []string{"bye", "remove", "delete"},
	Short:   "Uninstall a package named by its import path",
	Run: func(cmd *cobra.Command, args []string) {
		packageName := args[0]

		installPath := filepath.Join(viper.GetString(common.ConfigKeyInstallRootPath), packageName)

		_logger.Info("good: uninstalling %s...", packageName)

		binPath := filepath.Join(installPath, "bin")
		items, err := ioutil.ReadDir(binPath)
		if err != nil {
			_logger.Error(err.Error())
			os.Exit(1)
		}

		targetPath := viper.GetString(common.ConfigKeyLocalBinPath)
		for _, item := range items {
			if item.IsDir() {
				continue
			}

			targetFile := filepath.Join(targetPath, item.Name())

			path, err := filepath.EvalSymlinks(targetFile)
			if err != nil {
				continue
			}

			if filepath.Join(binPath, item.Name()) == path {
				_logger.Debug("good: removing %s...", targetFile)
				os.Remove(targetFile)
			}
		}

		if err := filepath.WalkDir(installPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				_logger.Debug(err.Error())
				return filepath.SkipDir
			}

			if d.IsDir() {
				_logger.Debug("good: changing permission of %s to 0700...", path)
				if err := os.Chmod(path, 0o700); err != nil {
					_logger.Debug(err.Error())
					return nil // continue
				}
			} else {
				_logger.Debug("good: changing permission of %s to 0600...", path)
				if err := os.Chmod(path, 0o600); err != nil {
					_logger.Debug(err.Error())
					return nil // continue
				}
			}

			return nil // continue
		}); err != nil {
			_logger.Debug(err.Error())
		}

		_logger.Debug("good: removing entire %s...", packageName)
		if err := os.RemoveAll(installPath); err != nil {
			_logger.Error(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uninstallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uninstallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
