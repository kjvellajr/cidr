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

// containsCmd represents the contains command
var containsCmd = &cobra.Command{
	Use:   "contains",
	Short: "Check if a list of IPs are contained within a given CIDR",
	Long: `Check if a list of IPs are contained within a given CIDR.

Examples:
  $ cidr contains 10.10.10.0/8 10.20.10.10
  10.20.10.10   : true

  $ cidr contains 10.10.10.0/8 10.20.10.10 11.20.10.10
  10.20.10.10   : true
  11.20.10.10   : false

Usage:
  cidr contains <address>/<mask> <ips> ...
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			os.Exit(1)
		}
		c, err := netip.ParsePrefix(args[0])
		if err != nil {
			fmt.Printf("Invalid CIDR string: %v\n", err)
			os.Exit(1)
		}
		var an []netip.Addr
		for _, s := range args[1:] {
			a, err := netip.ParseAddr(s)
			if err != nil {
				fmt.Printf("Invalid IP string: %v\n", err)
				os.Exit(1)
			}

			an = append(an, a)
		}
		rs := cidrContains(c, an)
		for i, r := range rs {
			fmt.Printf("%s\t: %t\n", an[i], r)
		}
	},
}

func init() {
	rootCmd.AddCommand(containsCmd)
}

func cidrContains(c netip.Prefix, na []netip.Addr) []bool {
	var r []bool
	for _, a := range na {
		r = append(r, c.Contains(a))
	}
	return r
}
