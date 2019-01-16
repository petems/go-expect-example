package main

import (
	"bytes"
	"fmt"
	"time"

	expect "github.com/Netflix/go-expect"
	"github.com/hinshun/vt10x"
	"github.com/howeyc/gopass"
)

func main() {
	stdoutBuf := new(bytes.Buffer)
	c, state, err := vt10x.NewVT10XConsole(
		expect.WithStdout(stdoutBuf),
	)
	defer c.Close()

	donec := make(chan struct{})
	go func() {
		defer close(donec)
		c.ExpectString("Enter passphrase (empty for no passphrase): ")
		time.Sleep(time.Second)
		c.Send("h")
		c.Send("u")
		c.Send("n")
		c.Send("t")
		c.Send("e")
		c.Send("r")
		c.Send("2")
		c.SendLine("")
	}()

	passwordInput, err := gopass.GetPasswdPrompt("Enter passphrase (empty for no passphrase): ", true, c.Tty(), c.Tty())

	<-donec

	if err != nil {
		fmt.Printf("\nError: %s\n", err)
	}

	if passwordInput != nil {
		fmt.Printf("\nResponse: %s\n", passwordInput)
	}

	fmt.Printf("\nState: %s\n", expect.StripTrailingEmptyLines(state.String()))
	fmt.Printf("\nStdout: %s\n", stdoutBuf)
}
