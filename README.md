# dlinkconfig

Tool for provisioning D-Link DGS-3100 over Telnet.

## Install

Get precompiled binaries from [builds.kradalby.no](https://builds.kradalby.no) or install with Go:

```
go get -u github.com/kradalby/dlinkconfig
```

## Usage

```
$ dlinkconfig
Tool for provisioning D-Link DGS-3100 over Telnet

Usage:
  dlinkconfig [command]

Available Commands:
  configure   Configure a D-Link DGS-3100 with a given config
  dhcpAuto    Activate DHCP on a D-Link DGS-3100
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.dlinkconfig.yaml)
  -h, --help            help for dlinkconfig
  -t, --toggle          Help message for toggle

Use "dlinkconfig [command] --help" for more information about a command.

```
