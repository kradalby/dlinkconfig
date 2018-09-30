// +build: !windows

// Copyright © 2018 Kristoffer Dalby <kradalby@kradalby.no>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"github.com/kradalby/dlinkconfig/dlink"
	"github.com/spf13/cobra"
	"os"
	"syscall"
)

var (
	configFile string
	switchIP   string
	username   string
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure a D-Link DGS-3100 with a given configuration file",
	Long:  `This configure modus is intended for automatic provisioning from ISC DHCP`,

	Run: func(cmd *cobra.Command, args []string) {
		ret, _, err := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
		if err != 0 {
			panic(err)
		}
		switch ret {
		case 0:
			break
		default:
			os.Exit(0)
		}

		dlink.Configure(switchIP, username, configFile)
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	configureCmd.Flags().StringVarP(&configFile, "configFile", "c", "", "Switch configuration filename")
	configureCmd.Flags().StringVarP(&switchIP, "ip", "a", "", "Switch IP address")
	configureCmd.Flags().StringVarP(&username, "username", "u", "", "Switch username")
	configureCmd.MarkFlagRequired("configFile")
	configureCmd.MarkFlagRequired("ip")
	configureCmd.MarkFlagRequired("username")
}
