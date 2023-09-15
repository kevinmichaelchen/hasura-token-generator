package cli

import (
	"github.com/kevinmichaelchen/tokesura/verify"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verifies a Hasura JSON Web Token (JWT)",
	Long:  `Verifies a Hasura JSON Web Token (JWT)`,
	Run:   fnVerify,
}

func fnVerify(cmd *cobra.Command, args []string) {
	err := verify.Verify(secret, token)
	if err != nil {
		log.Err(err).Msg("token verification failed")
		return
	}

	log.Info().Msg("Token has been validated and verified")
}
