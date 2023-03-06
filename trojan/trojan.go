package trojan

import (
	"v2ray.com/core/common"
	"v2ray.com/core/infra/conf/serial"

	v2ray "v2ray.com/core"
)

func init() {
	// Register JSON loader
	common.Must(v2ray.RegisterConfigLoader(&v2ray.ConfigFormat{
		Name:      "JSON",
		Extension: []string{"json"},
		Loader:    serial.LoadJSONConfig,
	}))
}

type Trojan struct {
	instance *v2ray.Instance
}

func (trojan *Trojan) Create(raw_config string) error {
	// Create io.Reader from raw_config
	config, err := v2ray.LoadConfig("json", raw_config, nil)
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
	err := trojan.Create(raw_config)
	if err != nil {
		return err
	}
	err = trojan.Start()
	if err != nil {
		return err
	}
	return nil
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
