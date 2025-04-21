package users

import (
	"context"

	"github.com/Pineapple217/cvrs/pkg/database"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

func GetCmd() *cobra.Command {
	userCmd := &cobra.Command{
		Use:   "user",
		Short: "Manage users",
	}
	var username, password string
	var isAdmin bool

	addUserCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new user",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := database.NewDatabase("file:./data/database.db?_fk=1&_journal_mode=WAL")
			if err != nil {
				return err
			}
			bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
			if err != nil {
				return err
			}
			_, err = db.Client.User.
				Create().
				SetUsername(username).
				SetPassword(bytes).
				SetIsAdmin(isAdmin).
				Save(context.Background())
			return err
		},
	}
	addUserCmd.Flags().StringVarP(&username, "username", "u", "", "Username for the new user")
	addUserCmd.Flags().StringVarP(&password, "password", "p", "", "Password for the new user")
	addUserCmd.Flags().BoolVar(&isAdmin, "admin", false, "Set the user as admin")
	addUserCmd.SilenceUsage = true

	userCmd.AddCommand(addUserCmd)
	return userCmd
}
