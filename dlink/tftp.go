package dlink

import (
	"bufio"
	"log"
	"os"
	"path"
)

func Configure(destination string, user string, configFile string) {
	t, err := NewTelnet(destination)
	t.CheckErr(err)

	// var data []byte
	t.Expect("UserName:")
	t.Sendln(user)
	t.Expect("DGS-3100# ")
	err = enterConfig(t, configFile)
	if err != nil {
		log.Fatalln(err)
	}

	// data, err = t.ReadBytes('>')
	// checkErr(err)
	// os.Stdout.Write(data)
	// os.Stdout.WriteString("\n")
}

func enterConfig(t *Telnet, configFile string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	file, err := os.Open(path.Join(dir, "file.txt"))

	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t.Sendln(scanner.Text())
		t.Expect("DGS-3100# ")
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
