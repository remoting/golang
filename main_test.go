package main

import (
	"fmt"
	"os/exec"
	"testing"
)

func Test001(t *testing.T) {
	cmd := "echo `dmidecode -s baseboard-serial-number`-`dmidecode -s system-serial-number` | sed 's/[[:space:]|\\.|\\-]//g'"

	c := exec.Command("/bin/sh", "-c", cmd)

	output, err := c.CombinedOutput()

	fmt.Println(c.Stderr)
	fmt.Println(c.Stdout)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}
}
