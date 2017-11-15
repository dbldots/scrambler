package cmd

import (
	"fmt"
	"io/ioutil"
	"regexp"

	b64 "encoding/base64"

	"github.com/richard-lyman/lithcrypt"
	"github.com/spf13/cobra"
)

var file string

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt a file",
	Long: `Decrypt a file with secured values. For example:

scrambler decrypt --secret P@assw0rd --file config.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		buf, _ := ioutil.ReadFile(file)
		pass := []byte(secret)
		search := regexp.MustCompile(`SCRAMBLED\(.*\)`)

		result := search.ReplaceAllFunc(buf, func(s []byte) []byte {
			match, _ := b64.StdEncoding.DecodeString(string(s[10 : len(s)-1]))
			decoded, _ := lithcrypt.Decrypt(pass, match)
			return decoded
		})

		fmt.Print(string(result))
	},
}

func init() {
	RootCmd.AddCommand(decryptCmd)
	decryptCmd.Flags().StringVarP(&file, "file", "f", "", "File to decrypt")
	decryptCmd.MarkFlagRequired("file")
}
