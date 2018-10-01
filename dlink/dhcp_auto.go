package dlink

import (
	"log"
)

func ConfigureDHCPAuto(destination string, user string, _ string) {

	t, err := NewTelnet(destination)
	if err != nil {
		log.Fatalln(err)
	}

	Login(t, user)
	t.Sendln("config dhcp_auto enable")
	t.Expect("DGS-3100# ")
	WriteConfig(t)
	Reboot(t)
}

// func configure(destination string, user string) {
// 	t, err := NewTelnet(destination)
// 	t.CheckErr(err)
//
// 	// var data []byte
// 	t.Expect("UserName:")
// 	t.Sendln(user)
// 	t.Expect("DGS-3100# ")
// 	t.Sendln("config dhcp_auto enable")
// 	t.Expect("The configuration will take place on the next time the device will get DHCP address.")
// 	t.Expect("Success.")
// 	t.Expect("DGS-3100# ")
// 	t.Sendln("save")
// 	t.Expect("Overwrite file [startup-config] ?[Yes/press any key for no]....")
// 	t.Sendln("yes")
// 	t.Expect("Success.")
// 	t.Expect("DGS-3100# ")
// 	t.Sendln("reboot")
// 	t.Expect("This action may take a few minutes")
// 	// t.Expect(t, "Are you sure you want to proceed with system reboot now? (Y/N)[N] ")
// 	t.Sendln("Y")
// 	t.Expect("Shutting down ...")
// 	fmt.Println("Switch configuration done, rebooting...")
// 	fmt.Println("Please disconnect and move on to the next switch")
// 	time.Sleep(time.Second * 30)
// 	fmt.Println("Looking for new switch...")
// 	// data, err = t.ReadBytes('>')
// 	// checkErr(err)
// 	// os.Stdout.Write(data)
// 	// os.Stdout.WriteString("\n")
// }
