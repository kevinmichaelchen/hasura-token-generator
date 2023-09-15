package cli

import (
	"fmt"
	"github.com/kevinmichaelchen/tokesura/secrets"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var generateAdminSecretCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates an alphanumeric secret for Hasura",
	Long:  `Generates an alphanumeric secret for Hasura`,
	Run:   fnGenerateAlphanumericSecret,
}

var secretCmd = &cobra.Command{
	Use:   "secret",
	Short: "Generates secrets for Hasura",
	Long:  `Generates secrets for Hasura`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO should be at least 32 chars for HS256
		// https://hasura.io/docs/latest/auth/authentication/jwt/#jwt-json-key
	},
}

func fnGenerateAlphanumericSecret(cmd *cobra.Command, args []string) {
	s, err := secrets.GenerateRandomString(secretLen)
	if err != nil {
		log.Err(err).Msg("Failed to generate admin secret")
		return
	}

	fmt.Println(s)
}
