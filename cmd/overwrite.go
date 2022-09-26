/*
Copyright © 2022 Filip Lindberg fili21@student.bth.se
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// overwriteCmd represents the overwrite command
var overwriteCmd = &cobra.Command{
	Use:   "overwrite",
	Short: "Not yet implemented",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("overwrite called")
	},
}

func init() {
	rootCmd.AddCommand(overwriteCmd)
}
