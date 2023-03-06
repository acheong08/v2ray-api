package trojan

import (
	v2ray "github.com/v2fly/v2ray-core/v5"
)

type Trojan struct {
	instance *v2ray.Instance
}

func (trojan *Trojan) Create(raw_config string) error {
	config, err := v2ray.LoadConfig("json", raw_config)
	if err != nil {
		return err
	}
	trojan.instance, err = v2ray.New(config)
	if err != nil {
		return err
	}
	return nil
}

func (trojan *Trojan) CreateAndRun(raw_config string) error {
	config, err := v2ray.LoadConfig("json", raw_config)
	if err != nil {
		return err
	}
	trojan.instance, err = v2ray.New(config)
	if err != nil {
		return err
	}
	err = trojan.Start()
	return err
}

func (trojan *Trojan) Start() error {
	err := trojan.instance.Start()
	if err != nil {
		return err
	}
	return nil
}

func (trojan *Trojan) Stop() error {
	err := trojan.instance.Close()
	if err != nil {
		return err
	}
	return nil
}

func (trojan *Trojan) RestartWithNewConfig(raw_config string) error {
	err := trojan.Stop()
	if err != nil {
		return err
	}
	err = trojan.Create(raw_config)
	if err != nil {
		return err
	}
	err = trojan.Start()
	if err != nil {
		return err
	}
	return nil
}

func (trojan *Trojan) Status() string {
	if trojan.instance == nil {
		return "nil"
	}
	return "exists"
}

func New() *Trojan {
	return &Trojan{}
}
