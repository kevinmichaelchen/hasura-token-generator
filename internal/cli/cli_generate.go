package cli

import (
	"github.com/kevinmichaelchen/tokesura/generate"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var generateTokenCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a Hasura JSON Web Token (JWT)",
	Long:  `Generates a Hasura JSON Web Token (JWT)`,
	Run:   fnGenerate,
}

func fnGenerate(cmd *cobra.Command, args []string) {
	out, err := generate.CreateToken(
		generate.WithSecret(secret),
		generate.WithSubject(subject),
		generate.WithAllowedRoles(allowedRoles),
		generate.WithDefaultRole(defaultRole),
		generate.WithUserID(userID),
	)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to generate token")
	}

	log.Info().
		Str("token", out.Token).
		Strs("allowed_roles", allowedRoles).
		Msg("Success")
}
