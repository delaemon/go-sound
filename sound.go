package sound

import "os/exec"

func Play(filename string) error{
	return exec.Command("open", filename).Run()
}
