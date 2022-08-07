/*
Copyright © 2022 Tiger Wang <tiger@tensorsmart.com>
*/
package cmd

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tigerinus/good/common"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all installed packages",
	Run: func(cmd *cobra.Command, args []string) {
		installRootPath := filepath.Join(viper.GetString(common.ConfigKeyInstallRootPath))

		if err := filepath.WalkDir(installRootPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				_logger.Debug(err.Error())
				return filepath.SkipDir
			}

			if !d.IsDir() && filepath.Base(path) == common.ManifestFileName {
				content, err := ioutil.ReadFile(path)
				if err != nil {
					_logger.Debug(err.Error())
					return filepath.SkipDir
				}

				fmt.Println(string(content))
				return nil // continue
			}

			return nil // continue
		}); err != nil {
			_logger.Debug(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
