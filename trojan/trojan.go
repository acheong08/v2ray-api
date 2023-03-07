package trojan

import (
	"fmt"
	"os"
	"os/exec"

	ps "github.com/mitchellh/go-ps"
)

type Trojan struct {
	pid int
}

func (trojan *Trojan) Start() error {
	// Check if ./v2ray exists
	_, err := exec.LookPath("./v2ray")
	if err != nil {
		return fmt.Errorf("v2ray binary not found: %w", err)
	}
	cmd := exec.Command("./v2ray", "run", "-config", "server.json")
	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start v2ray: %w", err)
	}
	// Get the PID
	trojan.pid = cmd.Process.Pid
	return nil
}

func (trojan *Trojan) Stop() error {
	cmd := exec.Command("kill", fmt.Sprint(trojan.pid))
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to stop v2ray: %w", err)
	}
	cmd.Wait()
	return nil
}

func (trojan *Trojan) Restart() error {
	if trojan.Status() == "stopped" {
		return trojan.Start()
	}
	if err := trojan.Stop(); err != nil {
		return fmt.Errorf("failed to stop v2ray: %w", err)
	}
	return trojan.Start()
}

func (trojan *Trojan) Status() string {
	// Check if PID is running
	process, err := ps.FindProcess(trojan.pid)
	if err != nil {
		return "stopped"
	}
	if process == nil {
		return "stopped"
	}
	return "running"
}

func (trojan *Trojan) Configure(config string) error {
	// Write config to server.json
	err := os.WriteFile("server.json", []byte(config), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to server.json: %w", err)
	}
	return nil
}

func (trojan *Trojan) GetConfig() (string, error) {
	// Read config from server.json
	config, err := os.ReadFile("server.json")
	if err != nil {
		return "", fmt.Errorf("failed to read server.json: %w", err)
	}
	return string(config), nil
}
