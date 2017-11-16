package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"

	b64 "encoding/base64"

	"github.com/richard-lyman/lithcrypt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt a file",
	Long: `Decrypt a file with secured values. For example:

scrambler decrypt config.yml`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := checkSecret(); err != nil {
			return err
		}

		if len(args) == 0 {
			return errors.New(`You have to provide a file to decrypt`)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		buf, _ := ioutil.ReadFile(args[0])
		pass := []byte(viper.GetString("secret"))
		search := regexp.MustCompile(`SCRAMBLED\(.*\)`)

		fmt.Println("Using secret '" + viper.GetString("secret") + "'")

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
}
