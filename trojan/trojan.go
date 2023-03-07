package trojan

import (
	"fmt"
	"os"
	"os/exec"

	ps "github.com/mitchellh/go-ps"
)

type Trojan struct {
	process *os.Process
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
	trojan.process = cmd.Process
	return nil
}

func (trojan *Trojan) Stop() error {
	if trojan.Status() == "stopped" {
		return nil
	}
	// Kill the process
	err := trojan.process.Kill()
	if err != nil {
		return fmt.Errorf("failed to kill v2ray process: %w", err)
	}
	// Wait for the process to stop
	_, err = trojan.process.Wait()
	if err != nil {
		return fmt.Errorf("failed to wait for v2ray process: %w", err)
	}
	// Release the process
	err = trojan.process.Release()
	if err != nil {
		return fmt.Errorf("failed to release v2ray process: %w", err)
	}
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
	// Check if process is running
	process, err := ps.FindProcess(trojan.process.Pid)
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
