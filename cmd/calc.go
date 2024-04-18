/*
Copyright © 2024 Ken Vella kjvellajr@gmail.com
*/
package cmd

import (
	"encoding/binary"
	"fmt"
	"math"
	"net"
	"net/netip"
	"os"

	"github.com/spf13/cobra"
)

// calcCmd represents the calc command
var calcCmd = &cobra.Command{
	Use:   "calc",
	Short: "Show network metadata from CIDR",
	Long: `Show network metadata from CIDR.

Examples:
  $ cidr calc 10.10.10.124/28
  Network:      10.10.10.112/28 (Class A)
  Netmask:      255.255.255.240
  First:        10.10.10.112
  Last:         10.10.10.127
  Total Hosts:  16

Usage:
  cidr calc <address>/<mask>
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Help()
			os.Exit(1)
		}
		p, err := netip.ParsePrefix(args[0])
		if err != nil {
			fmt.Printf("Invalid CIDR string: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Network:\t%s (%s)\n", p.Masked(), getIPClass(p.Masked().Addr()))

		ip, ok := netip.AddrFromSlice(net.CIDRMask(p.Bits(), 32))
		if ok {
			fmt.Printf("Netmask:\t%s\n", ip)
		}

		bcastIP := bcast(p)
		fmt.Printf("First:\t\t%s\n", p.Masked().Addr())
		fmt.Printf("Last:\t\t%s\n", bcastIP)
		fmt.Printf("Total Hosts:\t%d\n", ipsForCIDR(p))
	},
}

func init() {
	rootCmd.AddCommand(calcCmd)
}

// bcast returns the broadcast IP for a Prefix
// https://play.golang.org/p/Igo6Ct3gx_
func bcast(p netip.Prefix) netip.Addr {
	ab, err := p.Masked().Addr().MarshalBinary()
	if err != nil {
		return netip.Addr{}
	}
	bcastIP := make(net.IP, len(ab))
	a := net.IP(ab).To4()
	nm := net.IP(net.CIDRMask(p.Bits(), 32))
	binary.BigEndian.PutUint32(bcastIP, binary.BigEndian.Uint32(a)|^binary.BigEndian.Uint32(nm))
	addr, ok := netip.AddrFromSlice(bcastIP)
	if !ok {
		return netip.Addr{}
	}
	return addr
}

func getIPClass(ip netip.Addr) string {
	// Get first byte of 4 bytes representation of IP in string format
	fb := fmt.Sprintf("%08b", ip.As4()[0])

	var c string
	if fb[:1] == "0" {
		c = "Class A"
	} else if fb[:2] == "10" {
		c = "Class B"
	} else if fb[:3] == "110" {
		c = "Class C"
	} else if fb[:4] == "1110" {
		c = "Class D"
	} else if fb[:5] == "11110" {
		c = "Class E"
	}

	return c
}

// Returns the number of usaable IPs for a given Prefix
func ipsForCIDR(p netip.Prefix) int {
	return int(math.Pow(2, float64(32-p.Bits())))
}

// ipsForMask returns the number IPs contained in a netmask
func ipsForMask(m net.IPMask) int {
	s, _ := m.Size()
	return int(math.Pow(2, float64(32-s)))
}
