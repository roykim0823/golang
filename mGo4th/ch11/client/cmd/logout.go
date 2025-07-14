package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout for user",
	Long:  `Logging out a user from the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		endpoint := "/logout"
		user := User{Username: username, Password: password}

		// bytes.Buffer is both a Reader and a Writer
		buf := new(bytes.Buffer)
		err := user.ToJSON(buf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		req, err := http.NewRequest(http.MethodPost, SERVER+PORT+endpoint, buf)
		if err != nil {
			fmt.Println("GetAll – Error in req: ", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		c := &http.Client{
			Timeout: 15 * time.Second,
		}

		resp, err := c.Do(req)
		if err != nil {
			fmt.Println("Do:", err)
			return
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Println(resp)
			return
		} else {
			fmt.Println("User", user.Username, "logged out!")
		}
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
