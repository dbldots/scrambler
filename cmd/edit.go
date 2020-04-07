package cmd

import (
	b64 "encoding/base64"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a secured file",
	Long: `Edit a file that contains/should contain encrypted values. For example:

scrambler edit config.yml`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := checkSecret(); err != nil {
			return err
		}

		if len(args) == 0 {
			return errors.New(`You have to provide a file to edit`)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		extension := filepath.Ext(args[0])
		file, _ := ioutil.TempFile(os.TempDir(), "scrambler*" + extension)
		defer os.Remove(file.Name())

		buf, _ := ioutil.ReadFile(args[0])
		var result []byte

		if len(buf) > 10 && string(buf[0:10]) == `:SCRAMBLED` {
			content, _ := b64.StdEncoding.DecodeString(string(buf[11:]))
			decrypted, _ := decrypt(content)
			result = append([]byte(":SCRAMBLE\n"), decrypted...)
		} else {
			result = scrambledRegex.ReplaceAllFunc(buf, func(s []byte) []byte {
				match, _ := b64.StdEncoding.DecodeString(string(s[10:len(s)]))
				decrypted, _ := decrypt(match)
				return append([]byte("SCRAMBLE:"), decrypted...)
			})
		}

		ioutil.WriteFile(file.Name(), result, 0644)
		editFile(file.Name())
		buf, _ = ioutil.ReadFile(file.Name())

		if len(buf) > 9 && string(buf[0:9]) == `:SCRAMBLE` {
			content := buf[10:]
			encrypted, _ := encrypt(content)
			value := ":SCRAMBLED\n" + b64.StdEncoding.EncodeToString(encrypted)
			result = []byte(value)
		} else {
			result = scrambleRegex.ReplaceAllFunc(buf, func(s []byte) []byte {
				match := s[9:len(s)]
				encrypted, _ := encrypt(match)
				value := "SCRAMBLED:" + b64.StdEncoding.EncodeToString(encrypted)
				return []byte(value)
			})
		}

		ioutil.WriteFile(args[0], result, 0644)
	},
}

func init() {
	RootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
