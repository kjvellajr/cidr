# cidr

Simple command line utility for calculating networks based on IP addresses and mask bits.

Built with [cobra-cli](https://github.com/spf13/cobra-cli) and based on [chanux/ipeter](https://codeberg.org/chanux/ipeter/src/commit/71afd0f992821776a8d7adea3547357383735b1b).

## install

```bash
go install github.com/kjvellajr/cidr
```

## usage

```bash
$ cidr calc 10.10.10.124/28
Network:        10.10.10.112/28 (Class A)
Netmask:        255.255.255.240
First:          10.10.10.112
Last:           10.10.10.127
Total Hosts:    16

$ cidr contains 10.10.10.0/8 10.20.10.10 11.20.10.10
10.20.10.10	: true
11.20.10.10	: false

$ cidr mask 200 390 12300
netmask for 200 hosts is /24
netmask for 390 hosts is /23
netmask for 12300 hosts is /18

$ cidr overlap 10.10.10.124/28 10.10.10.127/29
true
```
