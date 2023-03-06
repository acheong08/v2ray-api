package trojan

import (
	"fmt"
	"os/exec"
)

type Trojan struct {
	pid int
}

func (trojan *Trojan) Start() {
	cmd := exec.Command("v2ray", "run", "-c", "server.json")
	cmd.Start()
	// Get the PID
	trojan.pid = cmd.Process.Pid
}

func (trojan *Trojan) Stop() {
	cmd := exec.Command("kill", "-9", fmt.Sprint(trojan.pid))
	cmd.Start()
}
