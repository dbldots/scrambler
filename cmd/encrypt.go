package cmd

import (
	"errors"
	"fmt"

	b64 "encoding/base64"

	"github.com/richard-lyman/lithcrypt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		pass := []byte(viper.GetString("secret"))

		fmt.Println("Using secret '" + viper.GetString("secret") + "'")

		encrypted, _ := lithcrypt.Encrypt(pass, payload)
		result := "SCRAMBLED(" + b64.StdEncoding.EncodeToString(encrypted) + ")"

		fmt.Println("Encrypted Result:")
		fmt.Println("---------------------------------------")
		fmt.Println(result)
	},
}

func init() {
	RootCmd.AddCommand(encryptCmd)
}
