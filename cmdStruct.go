package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

func main() {
// http://askubuntu.com/questions/20034/differences-between-doublequotes-singlequotes-and-backticks-%C2%B4-%C2%B4-on-comm

// Single quotes ('') are used to preserve the literal value of each character enclosed within the quotes

// Using double quotes the literal value of all characters enclosed is preserved, except for the dollar sign, the backticks (backward single quotes, ``) and the backslash.

// When enclosed inside back-ticks, the shell interprets something to mean "the output of the command inside the back-ticks." This is referred to as command substitution, as the output of the command inside the back-ticks is substituted for the command itself

	cmd := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	var person struct {
		Name string
		Age  int
	}
	// http://golang.org/pkg/encoding/json/#NewDecoder
	if err := json.NewDecoder(stdout).Decode(&person); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	// at this point, no more commandline magic. We already have the data in our person struct 
	fmt.Printf("%s is %d years old\n", person.Name, person.Age)
}