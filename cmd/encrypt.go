package cmd

import (
	"fmt"

	b64 "encoding/base64"

	"github.com/richard-lyman/lithcrypt"
	"github.com/spf13/cobra"
)

var value string

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt a value",
	Long: `Use 'encrypt' to secure a value. For example:

scrambler encrypt --secret P@assw0rd --value your-secret`,
	Run: func(cmd *cobra.Command, args []string) {
		payload := []byte(value)
		pass := []byte(secret)
		encrypted, _ := lithcrypt.Encrypt(pass, payload)
		result := "SCRAMBLED(" + b64.StdEncoding.EncodeToString(encrypted) + ")"
		fmt.Println(result)
	},
}

func init() {
	RootCmd.AddCommand(encryptCmd)
	encryptCmd.Flags().StringVarP(&value, "value", "v", "", "The value to encrypt")
	encryptCmd.MarkFlagRequired("value")
}
