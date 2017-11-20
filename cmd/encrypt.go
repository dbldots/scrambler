package cmd

import (
	"errors"
	"fmt"

	b64 "encoding/base64"

	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt a value",
	Long: `Use 'encrypt' to secure a value. For example:

scrambler encrypt "sensitive information"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := checkSecret(); err != nil {
			return err
		}

		if len(args) == 0 {
			return errors.New(`You have to provide a value to encrypt`)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		payload := []byte(args[0])

		encrypted, _ := encrypt(payload)
		result := "SCRAMBLED:" + b64.StdEncoding.EncodeToString(encrypted)
		//result := "SCRAMBLED:" + fmt.Sprintf("%0x", encrypted)

		fmt.Println("Encrypted Result:")
		fmt.Println("---------------------------------------")
		fmt.Println(result)
	},
}

func init() {
	RootCmd.AddCommand(encryptCmd)
}
