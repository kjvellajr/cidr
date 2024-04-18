/*
Copyright © 2024 Ken Vella kjvellajr@gmail.com
*/
package cmd

import (
	"fmt"
	"math"
	"net"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// maskCmd represents the mask command
var maskCmd = &cobra.Command{
	Use:   "mask",
	Short: "Calculate the network mask needed to fit a number of hosts",
	Long: `Calculate the network mask needed to fit a number of hosts. Accepts multiple host numbers.

Examples:
  $ cidr mask 10
  netmask for 10 hosts is /28

  $ cidr mask 200 390 12300
  netmask for 200 hosts is /24
  netmask for 390 hosts is /23
  netmask for 12300 hosts is /18

Usage:
  cidr mask <host_count> ...
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			os.Exit(1)
		}
		hosts := make([]int, 0, len(args))
		for _, host := range args {
			n, err := strconv.Atoi(host)
			if err != nil {
				fmt.Printf("unable to convert arg to int: %s %d\n", host, err)
				os.Exit(1)
			}
			hosts = append(hosts, n)
		}
		for _, v := range hosts {
			nm := maskForNHosts(v)
			s, err := nm.Size()
			if s == 0 {
				fmt.Printf("number too large: %d %v\n", err, v)
				os.Exit(1)
			}
			fmt.Printf("netmask for %d hosts is /%d\n", v, s)

		}
	},
}

func init() {
	rootCmd.AddCommand(maskCmd)
}

// maskForNHosts returns a netmask that would home a given number of hosts
// S = 32 - Ceil(log2(N+2))
func maskForNHosts(hosts int) net.IPMask {
	bits := 32 - math.Ceil(math.Log2(float64(hosts+2)))
	return net.CIDRMask(int(bits), 32)
}
