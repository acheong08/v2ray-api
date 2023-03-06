package trojan

import (
	"fmt"
	"os/exec"
)

type Trojan struct {
	pid int
}

func (trojan *Trojan) Start() error {
	cmd := exec.Command("./v2ray", "run", "-c", "server.json")
	err := cmd.Start()
	if err != nil {
		return err
	}
	// Get the PID
	trojan.pid = cmd.Process.Pid
	return nil
}

func (trojan *Trojan) Stop() error {
	cmd := exec.Command("kill", "-9", fmt.Sprint(trojan.pid))
	err := cmd.Start()
	return err
}

func (trojan *Trojan) Restart() error {
	if trojan.Status() == "stopped" {
		trojan.Start()
		return nil
	}
	err := trojan.Stop()
	if err != nil {
		return err
	}
	trojan.Start()
	return err
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
