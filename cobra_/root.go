package cobra_

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "study-go",
	Short: "this is study-go",
	Long:  "this is study-go usage",
	//Run:   runRoot,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func runRoot(cmd *cobra.Command, args []string) {
	fmt.Printf("execute %s args:%v\n", cmd.Name(), args)
}

func DemoCobra() {

}
