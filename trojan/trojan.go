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

func (trojan *Trojan) Restart() {
	trojan.Stop()
	trojan.Start()
}

func (trojan *Trojan) Status() string {
	// Check if PID is running
	cmd := exec.Command("ps", "-p", fmt.Sprint(trojan.pid))
	err := cmd.Run()
	if err != nil {
		return "stopped"
	} else {
		return "running"
	}
}
