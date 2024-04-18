/*
Copyright © 2024 Ken Vella kjvellajr@gmail.com
*/
package cmd

import (
	"fmt"
	"net/netip"
	"os"

	"github.com/spf13/cobra"
)

// overlapCmd represents the overlap command
var overlapCmd = &cobra.Command{
	Use:   "overlap",
	Short: "Determine if two CIDRs overlap by sharing some IPs",
	Long: `Determine if two CIDRs overlap by sharing some IPs.

Examples:
  $ cidr overlap 10.10.10.124/28 10.10.10.127/29
  true

  $ cidr overlap 10.10.10.0/8 10.10.10.0/8
  false

Usage:
  cidr overlap <address>/<mask> <address>/<mask>
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			cmd.Help()
			os.Exit(1)
		}
		c1, err := netip.ParsePrefix(args[0])
		if err != nil {
			fmt.Printf("Invalid CIDR string: %v\n", err)
			os.Exit(1)
		}
		c2, err := netip.ParsePrefix(args[1])
		if err != nil {
			fmt.Printf("Invalid CIDR string: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(c1.Overlaps(c2))

		fmt.Println()
		calcCmd.Run(cmd, []string{args[0]})

		fmt.Println()
		calcCmd.Run(cmd, []string{args[1]})
	},
}

func init() {
	rootCmd.AddCommand(overlapCmd)
}
