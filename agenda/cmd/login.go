package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"agenda/entity/User"
)
var login = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login command is called.")
		username,_ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		information := &User.User{username, password, "", nil, nil}
		User.LogIn(information)
	},
}

func init() {
	RootCmd.AddCommand(login)
	login.Flags().StringP("username", "u", "", "Username")
	login.Flags().StringP("password", "p", "", "User password")
}