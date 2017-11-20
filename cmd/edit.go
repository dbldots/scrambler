package cmd

import (
	b64 "encoding/base64"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"

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
		file, _ := ioutil.TempFile(os.TempDir(), "scrambler")
		defer os.Remove(file.Name())

		buf, _ := ioutil.ReadFile(args[0])

		result := scrambledRegex.ReplaceAllFunc(buf, func(s []byte) []byte {
			match, _ := b64.StdEncoding.DecodeString(string(s[10:len(s)]))
			decrypted, _ := decrypt(match)
			return append([]byte("SCRAMBLE:"), decrypted...)
		})

		ioutil.WriteFile(file.Name(), result, 0644)

		editor := exec.Command("vim", file.Name())
		editor.Stdin = os.Stdin
		editor.Stdout = os.Stdout
		editor.Stderr = os.Stderr
		editor.Run()

		buf, _ = ioutil.ReadFile(file.Name())

		result = scrambleRegex.ReplaceAllFunc(buf, func(s []byte) []byte {
			match := s[9:len(s)]
			encrypted, _ := encrypt(match)
			value := "SCRAMBLED:" + b64.StdEncoding.EncodeToString(encrypted)
			return []byte(value)
		})

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