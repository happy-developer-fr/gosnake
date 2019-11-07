package screenclear

import (
	"fmt"
	"os"
	"os/exec"
)

func CleanScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}
