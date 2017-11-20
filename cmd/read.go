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

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a file",
	Long: `Read a file with secured values. For example:

scrambler read config.yml`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := checkSecret(); err != nil {
			return err
		}

		if len(args) == 0 {
			return errors.New(`You have to provide a file to read`)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		buf, _ := ioutil.ReadFile(args[0])
		pass := []byte(viper.GetString("secret"))
		search := regexp.MustCompile(`SCRAMBLED:[^\n ]*`)

		result := search.ReplaceAllFunc(buf, func(s []byte) []byte {
			match, _ := b64.StdEncoding.DecodeString(string(s[10:len(s)]))
			decoded, _ := lithcrypt.Decrypt(pass, match)
			return decoded
		})

		fmt.Print(string(result))
	},
}

func init() {
	RootCmd.AddCommand(readCmd)
}
