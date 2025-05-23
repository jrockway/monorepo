// Package cmd implements the JSSO CLI.
package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	address    string
	timeout    time.Duration
	jsonOutput bool

	root, session, bearer string

	rootCmd = &cobra.Command{
		Use:   "jssoctl",
		Short: "jssoctl controls a jsso2 server",
		Long: `jssoctl connects to a jsso2 server and makes API requests on your behalf.

It can be used to administer JSSO, or interact with it as a normal user.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Stolen from https://github.com/carolynvs/stingoftheviper/blob/main/main.go.
			viper.SetEnvPrefix("jssoctl")
			viper.AutomaticEnv()
			// if err := viper.ReadInConfig(); err != nil {
			// 	if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// 		return err
			// 	}
			// }
			cmd.Flags().VisitAll(func(f *pflag.Flag) {
				// Apply the viper config value to the flag when the flag is not set and viper has a value
				if !f.Changed && viper.IsSet(f.Name) {
					val := viper.Get(f.Name)
					cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val)) //nolint:errcheck
				}
			})
			return nil
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func Execute() {
	ctx, c := context.WithTimeout(context.Background(), timeout)
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Fprintln(rootCmd.ErrOrStderr(), err)
		c()
		os.Exit(1)
	}
	c()
	os.Exit(0)
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&jsonOutput, "json-output", false, "set to output json instead of friendly text")
	rootCmd.PersistentFlags().StringVar(&address, "address", "localhost:4000", "address of the jsso grpc address")
	rootCmd.PersistentFlags().StringVar(&root, "root", "", "if set, authenticate with this root password")
	rootCmd.PersistentFlags().StringVar(&session, "session", "", "if set, authenticate with this base64-encoded session id")
	rootCmd.PersistentFlags().StringVar(&bearer, "bearer", "", "if set, authenticate with this bearer token")
	rootCmd.PersistentFlags().DurationVar(&timeout, "timeout", 5*time.Second, "time allowed for the command to run, including all network requests")
	rootCmd.AddCommand(usersCmd, devCmd)
}
