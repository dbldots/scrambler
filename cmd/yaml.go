// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	b64 "encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/richard-lyman/lithcrypt"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var yamlCmd = &cobra.Command{
	Use:   "yaml",
	Short: "Manage encrypted values in YAML files",
}

var yamlReadCmd = &cobra.Command{
	Use:   "read",
	Short: "Read encrypted YAML value",
	Long: `Convenience method to read an encrypted YAML value. For example:

scrambler yaml read config.yml database.password`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := checkSecret(); err != nil {
			return err
		}

		if len(args) < 2 {
			return errors.New(`You have to provide file and key`)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		buf, _ := ioutil.ReadFile(args[0])
		pass := []byte(viper.GetString("secret"))

		m := make(map[string]interface{})
		yaml.Unmarshal(buf, m)

		path := strings.Split(args[1], ".")
		find := search(m, path)

		value := cast.ToString(find)
		slice := []byte(value)

		if !strings.HasPrefix(value, "SCRAMBLED") {
			fmt.Println("Not an encrypted value")
			os.Exit(0)
		} else {
			match, _ := b64.StdEncoding.DecodeString(string(slice[10 : len(slice)-1]))
			decoded, _ := lithcrypt.Decrypt(pass, match)
			fmt.Println(string(decoded))
		}
	},
}

var yamlWriteCmd = &cobra.Command{
	Use:   "write",
	Short: "Write encrypted YAML value",
	Long: `Convenience method to write an encrypted YAML value. For example:

scrambler yaml write config.yml database.password P@ssw0rd`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := checkSecret(); err != nil {
			return err
		}

		if len(args) < 3 {
			return errors.New(`You have to provide file, key and value`)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		buf, _ := ioutil.ReadFile(args[0])
		pass := []byte(viper.GetString("secret"))

		m := make(map[string]interface{})
		yaml.Unmarshal(buf, m)

		path := strings.Split(args[1], ".")
		key := path[len(path)-1]

		var find interface{}
		if len(path) == 1 {
			find = m
		} else {
			parent := path[0 : len(path)-1]
			find = search(m, parent)
		}

		if find == nil {
			fmt.Println("Could not find parent key")
			os.Exit(0)
		}

		payload := []byte(args[2])

		fmt.Println("Using secret '" + viper.GetString("secret") + "'")

		encrypted, _ := lithcrypt.Encrypt(pass, payload)
		result := "SCRAMBLED(" + b64.StdEncoding.EncodeToString(encrypted) + ")"

		switch find.(type) {
		case map[interface{}]interface{}:
			find.(map[interface{}]interface{})[key] = result
		case map[string]interface{}:
			find.(map[string]interface{})[key] = result
		}

		buf2, _ := yaml.Marshal(m)
		ioutil.WriteFile(args[0], buf2, 0644)
	},
}

func init() {
	yamlCmd.AddCommand(yamlReadCmd)
	yamlCmd.AddCommand(yamlWriteCmd)
	RootCmd.AddCommand(yamlCmd)
}

// Copyright (c) 2014 Steve Francia
func search(source map[string]interface{}, path []string) interface{} {
	if len(path) == 0 {
		return source
	}

	// search for path prefixes, starting from the longest one
	for i := len(path); i > 0; i-- {
		prefixKey := strings.ToLower(strings.Join(path[0:i], "."))

		next, ok := source[prefixKey]
		if ok {
			// Fast path
			if i == len(path) {
				return next
			}

			// Nested case
			var val interface{}
			switch next.(type) {
			case map[interface{}]interface{}:
				val = search(cast.ToStringMap(next), path[i:])
			case map[string]interface{}:
				// Type assertion is safe here since it is only reached
				// if the type of `next` is the same as the type being asserted
				val = search(next.(map[string]interface{}), path[i:])
			default:
				// got a value but nested key expected, do nothing and look for next prefix
			}
			if val != nil {
				return val
			}
		}
	}

	return nil
}
