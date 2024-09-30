package pkg

import (
	"fmt"
	"os"
	"os/exec"
)

func RunCommand(name string, arg ...string) {
	//fmt.Println("Running command:", name, arg)
	//fmt.Println()
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(cmd.Stderr)
	}
	//fmt.Println("\nCommand executed successfully:", name, arg)
}
