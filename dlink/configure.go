package dlink

import (
	"bufio"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/sparrc/go-ping"
	"log"
	"os"
	"path"
	"time"
)

type ConfigurationFunc func(string, string, string)

func RunConfigurationPingLoop(host string, telnetPort int, user string, privileged bool, configFile string, confFunc ConfigurationFunc) {
	destination := fmt.Sprintf("%s:%d", host, telnetPort)
	timeout := time.Second * 100000
	interval := time.Second
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()

	pinger, err := ping.NewPinger(host)
	if err != nil {
		log.Fatalf(color.RedString("ERROR: %s\n"), err.Error())
	}

	pinger.OnRecv = func(pkt *ping.Packet) {
		s.Stop()
		log.Printf(color.BlueString("Host %s found, starting configuration\n"), pkt.IPAddr)
		confFunc(destination, user, configFile)
	}

	pinger.OnFinish = func(stats *ping.Statistics) {
		s.Start()
	}

	pinger.Interval = interval
	pinger.Timeout = timeout
	pinger.SetPrivileged(privileged)

	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	pinger.Run()
}

func ConfigureFromFile(destination string, user string, configFile string) {
	t, err := NewTelnet(destination)
	if err != nil {
		log.Fatalln(err)
	}

	// var data []byte
	Login(t, user)
	err = EnterConfigFile(t, user, configFile)
	WriteConfig(t)
	Reboot(t)
	log.Println(color.RedString("Please disconnect and move on to the next switch"))
	time.Sleep(time.Second * 20)
	log.Println(color.MagentaString("Looking for new switch..."))

	// data, err = t.ReadBytes('>')
	// checkErr(err)
	// os.Stdout.Write(data)
	// os.Stdout.WriteString("\n")
}

func Login(t *Telnet, user string) {
	t.Expect("UserName:")
	t.Sendln(user)
	t.Expect("DGS-3100# ")
	log.Println(color.GreenString("Login complete"))
}

func EnterConfigFile(t *Telnet, user string, configFile string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	file, err := os.Open(path.Join(dir, configFile))

	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		command := scanner.Text()
		log.Println(color.CyanString("Executing: "), color.MagentaString(command))
		err := t.SendlnWithError(command)
		if err != nil {
			log.Printf(color.RedString("Error: %s\n"), err)
			log.Println(color.BlueString("Reconnecting..."))

			err := t.Reconnect()
			if err != nil {
				return err
			}

			Login(t, user)
			continue
		}

		err = t.ExpectWithError("Success")
		if err != nil {
			log.Printf(color.RedString("Error: %s\n"), err)
			log.Println(color.BlueString("Reconnecting..."))

			err := t.Reconnect()
			if err != nil {
				return err
			}

			Login(t, user)
		}
		err = t.ExpectWithError("DGS-3100# ")
		if err != nil {
			log.Printf(color.RedString("Error: %s\n"), err)
			log.Println(color.BlueString("Reconnecting..."))

			err := t.Reconnect()
			if err != nil {
				return err
			}

			Login(t, user)
		}

	}

	if err := scanner.Err(); err != nil {
		return err
	}

	log.Println("Configuration entered successfully")
	time.Sleep(time.Second * 5)
	return nil
}

func WriteConfig(t *Telnet) {
	t.Sendln("save")
	t.Expect("Overwrite file [startup-config] ?[Yes/press any key for no]....")
	t.Sendln("yes")
	t.Expect("Success.")
	t.Expect("DGS-3100# ")
	log.Println(color.GreenString("Configuration saved"))
}

func Reboot(t *Telnet) {
	t.Sendln("reboot")
	t.Expect("This action may take a few minutes")
	// t.Expect(t, "Are you sure you want to proceed with system reboot now? (Y/N)[N] ")
	t.Sendln("Y")
	t.Expect("Shutting down ...")
	log.Println(color.YellowString("Switch rebooting"))
}
