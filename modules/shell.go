package modules

/* Import modules */
import (
	"os/exec"
)

func Shell(cmd string) (string, error) {
	out, err := exec.Command("/bin/bash", "-c", cmd).Output()
	if err != nil {
		return "Unable to execute", err
	}
	var strOut string = string(out)
	return strOut, nil
}
