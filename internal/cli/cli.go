package cli

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"os"
)

var (
	ldFlags      LDFlags
	secret       string
	subject      string
	allowedRoles []string
	defaultRole  string
	userID       string
	token        string
	secretLen    int
)

// LDFlags contain fields that get linked and compiled into the final binary
// program at build time.
type LDFlags struct {
	Version string
	Commit  string
	Date    string
}

func init() {
	rootCmd.AddCommand(generateTokenCmd)
	rootCmd.AddCommand(verifyCmd)
	rootCmd.AddCommand(secretCmd)

	secretCmd.AddCommand(generateAdminSecretCmd)
	secretCmd.PersistentFlags().IntVarP(&secretLen, "length", "l", 32, "secret length")

	rootCmd.PersistentFlags().StringVarP(&secret, "secret", "", "", "secret string")
	rootCmd.MarkFlagRequired("secret")

	verifyCmd.Flags().StringVarP(&token, "token", "t", "", "Token to verify")
	verifyCmd.MarkFlagRequired("token")

	generateTokenCmd.Flags().StringVarP(&subject, "subject", "s", "", "Token subject (\"sub\" from RFC 7519)")
	generateTokenCmd.MarkFlagRequired("subject")

	// StringSliceVarP is comma-delimited, whereas StringArrayVarP can have
	// multiple occurrences of the same flag
	generateTokenCmd.Flags().StringSliceVarP(&allowedRoles, "allowedRoles", "a", []string{}, "Allowed Hasura roles")
	generateTokenCmd.MarkFlagRequired("allowedRoles")

	generateTokenCmd.Flags().StringVarP(&defaultRole, "defaultRole", "d", "", "Default Hasura role")
	generateTokenCmd.MarkFlagRequired("defaultRole")

	generateTokenCmd.Flags().StringVarP(&userID, "userID", "", "", "Hasura User ID")
}

var rootCmd = &cobra.Command{
	Use:   "tokesura",
	Short: "Utility for working with Hasura JSON Web Tokens (JWTs)",
	Long:  `Utility for working with Hasura JSON Web Tokens (JWTs)`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Main(ldf LDFlags) {
	ldFlags = ldf
	if err := rootCmd.Execute(); err != nil {
		log.Error("execution failed", "err", err)
		os.Exit(1)
	}
}
