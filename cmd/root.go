/*
Copyright © 2022 Tiger Wang <tiger@tensorsmart.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tigerinus/good/common"
)

// var cfgFile string
var _logger *common.Logger

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "good",
	Short: "Install and uninstall a Go package to an isolated path",
	Long: `A tool for installing a Go package to an isolated path, to keep the global GOPATH
clean. Because of isolated path, uninstalling is also possible.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	_logger = common.NewLogger()

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// configFilepath := filepath.Join(xdg.ConfigHome, "good.toml")
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is "+configFilepath+")")
	rootCmd.PersistentFlags().BoolVarP(&_logger.DebugMode, "debug", "d", false, "debug mode")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle") // commented by tiger
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// if cfgFile != "" {
	// 	// Use config file from the flag.
	// 	viper.SetConfigFile(cfgFile)
	// } else {
	// 	viper.AddConfigPath(xdg.ConfigHome)
	// 	viper.SetConfigType("toml")
	// 	viper.SetConfigName("good")
	// }

	viper.SetDefault(common.ConfigKeyInstallRootPath, filepath.Join(xdg.DataHome, "good"))
	viper.SetDefault(common.ConfigKeyLocalBinPath, filepath.Join(xdg.Home, ".local", "bin"))

	viper.SetEnvPrefix("good")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	// if err := viper.ReadInConfig(); err == nil {
	// 	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	// } else {
	// 	if err := viper.SafeWriteConfig(); err != nil {
	// 		fmt.Fprintln(os.Stderr, "Error writing config file:", err)
	// 	}
	// }
}
