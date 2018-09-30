package dlink

import (
	"fmt"
	// "os"
	"time"

	"github.com/sparrc/go-ping"
)

func ConfigLoop(host string, telnetPort int, user string, privileged bool) {
	destination := fmt.Sprintf("%s:%d", host, telnetPort)
	timeout := time.Second * 100000
	interval := time.Second

	pinger, err := ping.NewPinger(host)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("Host %s found, starting configuration\n", pkt.IPAddr)
		configure(destination, user)
	}

	pinger.Interval = interval
	pinger.Timeout = timeout
	pinger.SetPrivileged(privileged)

	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	pinger.Run()
}

func configure(destination string, user string) {
	t, err := NewTelnet(destination)
	t.CheckErr(err)

	// var data []byte
	t.Expect("UserName:")
	t.Sendln(user)
	t.Expect("DGS-3100# ")
	t.Sendln("config dhcp_auto enable")
	t.Expect("The configuration will take place on the next time the device will get DHCP address.")
	t.Expect("Success.")
	t.Expect("DGS-3100# ")
	t.Sendln("save")
	t.Expect("Overwrite file [startup-config] ?[Yes/press any key for no]....")
	t.Sendln("yes")
	t.Expect("Success.")
	t.Expect("DGS-3100# ")
	t.Sendln("reboot")
	t.Expect("This action may take a few minutes")
	// t.Expect(t, "Are you sure you want to proceed with system reboot now? (Y/N)[N] ")
	t.Sendln("Y")
	t.Expect("Shutting down ...")
	fmt.Println("Switch configuration done, rebooting...")
	fmt.Println("Please disconnect and move on to the next switch")
	time.Sleep(time.Second * 30)
	fmt.Println("Looking for new switch...")
	// data, err = t.ReadBytes('>')
	// checkErr(err)
	// os.Stdout.Write(data)
	// os.Stdout.WriteString("\n")
}
