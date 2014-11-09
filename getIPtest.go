package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	// "os"
	// "encoding/json"
	// "encoding/hex"
	// "log"
)

func main() {
	// var slaveIPAddress []byte
	// slaveIPAddress, _ := exec.Command("ifconfig", "|", "grep", "inet").CombinedOutput()
	// slaveIPAddress, _ := exec.Command("ifconfig", "|", "grep", "-Eo", "'inet (addr:)?([0-9]*\\.){3}[0-9]*'", "|", "grep", "-Eo", "'([0-9]*\\.){3}[0-9]*'", "|", "grep", "-v", "'127.0.0.1'").CombinedOutput() 
	// fmt.Println(slaveIPAddress)
	// fmt.Println(hex.Dump(slaveIPAddress))


	cmd := exec.Command("ifconfig")

	IPAddressBytes, _ := cmd.Output() // display operating system name...why do we need the -a?
	IPAddress := string(IPAddressBytes)
	inetAddressRegexpPattern := "inet (addr:)?([0-9]*\\.){3}[0-9]*"
	re := regexp.MustCompile(inetAddressRegexpPattern)
	IPAddress = re.FindAllString(IPAddress, -1)[1]
	IPAddress = strings.Split(IPAddress, " ")[1]
	fmt.Println(IPAddress)

	// stdout, err := cmd.StdoutPipe() // returns a pipe

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := cmd.Start(); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(cmd.Stdout)
	// http://stackoverflow.com/questions/8875038/redirect-stdout-pipe-of-child-process-in-golang
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// cmd.Run()
}
	


