package ___

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func output(reader io.ReadCloser) error {
	buf := make([]byte, 1024)
	for {
		num, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if num > 0 {
			fmt.Printf("%s", string(buf[:num]))
		}
	}
	//return nil
}

func Exec(arg ...string) error {
	cmd := exec.Command("go", arg...)

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		log.Printf("Error executing command: %s......\n", err.Error())
		return err
	}

	go output(stdout)
	go output(stderr)

	if err := cmd.Wait(); err != nil {
		log.Printf("Error waiting for command execution: %s......\n", err.Error())
		return err
	}

	log.Printf("done")

	return nil
}
