package cmd

import (
	"fmt"
	"os/exec"
	"github.com/spf13/cobra"
)

var shell string

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "run a command",
	Long: `provide a command to run`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("exec called")
		out, err := exec.Command("sh", "-c", args[0]).Output()
		if err != nil {
			fmt.Printf("%s", err)
		}
		fmt.Printf("Command output: \n%s", out)
	},
}

func init() {

	RootCmd.AddCommand(execCmd)

	execCmd.Flags().StringVar(&shell, "shell", "", "specify a shell (supported types: bash)")

}
