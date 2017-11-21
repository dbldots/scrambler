// Copyright © 2017 johannes-kostas goetzinger <dbldots@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var scrambledRegex = regexp.MustCompile(`SCRAMBLED:[^\n$]*`)
var scrambleRegex = regexp.MustCompile(`SCRAMBLE:[^\n$]*`)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "scrambler",
	Short: "Encrypt/decrypt sensible data",
	Long: `Scrambler aims to provide a simple way to let you secure sensible
information such as credentials you don't want to put into source control.`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(home)
	viper.SetConfigName(".scrambler")

	viper.SetEnvPrefix("scrambler")
	viper.AutomaticEnv()
	viper.ReadInConfig()
}

var secret []byte
var filler = []byte("123456789abcdefghijklmnopqrstuvz")

func checkSecret() error {
	if viper.GetString("secret") == "" {
		return errors.New(`Required "secret" have/has not been set`)
	}

	secret = []byte(viper.GetString("secret"))

	if len(secret) < 32 {
		rest := filler[:(32 - len(secret))]
		secret = append(secret, rest...)
	} else if len(secret) > 32 {
		secret = secret[:32]
	}

	return nil
}

func editFile(file string) {
	cmd := os.Getenv("EDITOR")
	if cmd == "" {
		cmd = "vi"
	}

	editor := exec.Command(cmd, file)
	editor.Stdin = os.Stdin
	editor.Stdout = os.Stdout
	editor.Stderr = os.Stderr
	editor.Run()
}
