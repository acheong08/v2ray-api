package trojan

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/v2fly/v2ray-core/v5/common"

	v2ray "github.com/v2fly/v2ray-core/v5"
)

func init() {
	// Register JSON loader
	common.Must(v2ray.RegisterConfigLoader(&v2ray.ConfigFormat{
		Name:      []string{"json"},
		Extension: []string{"json"},
		Loader: func(input interface{}) (*v2ray.Config, error) {
			config := new(v2ray.Config)
			switch v := input.(type) {
			case string:
				err := json.Unmarshal([]byte(v), config)
				if err != nil {
					return nil, err
				}
				return config, nil
			case []byte:
				err := json.Unmarshal(v, config)
				if err != nil {
					return nil, err
				}
			case io.Reader:
				err := json.NewDecoder(v).Decode(config)
				if err != nil {
					return nil, err
				}
			}
			return config, nil
		},
	}))
}

type Trojan struct {
	instance *v2ray.Instance
}

func (trojan *Trojan) Create(raw_config string) error {
	// Create io.Reader from raw_config
	config, err := v2ray.LoadConfig("json", strings.NewReader(raw_config))
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
