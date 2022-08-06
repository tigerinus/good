/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:     "install [package]",
	Aliases: []string{"hello", "add", "get"},
	Args:    cobra.MinimumNArgs(1),
	Short:   "[TODO] A brief description of your command",
	Long: `[TODO] A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		packageName := args[0]
		installPath, err := createPackageInstallPath(packageName)
		if err != nil {
			_logger.Error(err.Error())
			os.RemoveAll(installPath)
			os.Exit(1)
		}

		execCmd := exec.Command("go", "install", packageName)
		execCmd.Env = append(os.Environ(), fmt.Sprintf("GOPATH=%s", installPath))
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr

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

		targetPath := viper.GetString(configKeyLocalBinPath)
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
	installPath := filepath.Join(viper.GetString(configKeyInstallRootPath), packageName)

	if err := os.MkdirAll(installPath, 0o755); err != nil {
		return "", err
	}

	return installPath, nil
}