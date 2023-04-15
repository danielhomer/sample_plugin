package sample_plugin

import (
	"github.com/roadrunner-server/errors"
	"go.uber.org/zap"
	"time"
)

const name = "sample_plugin"

type Configurer interface {
	// UnmarshalKey takes a single key and unmarshal it into a Struct.
	//
	// func (h *HttpService) Init(cp config.Configurator) error {
	//     h.config := &HttpConfig{}
	//     if err := configProvider.UnmarshalKey("http", h.config); err != nil {
	//         return err
	//     }
	// }
	UnmarshalKey(name string, out interface{}) error

	// Unmarshal the config into a Struct. Make sure that the tags
	// on the fields of the structure are properly set.
	Unmarshal(out interface{}) error

	// Get used to get config section
	Get(name string) interface{}

	// Overwrite used to overwrite particular values in the unmarshalled config
	Overwrite(values map[string]interface{}) error

	// Has checks if config section exists.
	Has(name string) bool

	// GracefulTimeout represents timeout for all servers registered in the endure
	GracefulTimeout() time.Duration

	// RRVersion returns running RR version
	RRVersion() string
}

type Logger interface {
	NamedLogger(name string) *zap.Logger
}

type Plugin struct {
	log *zap.Logger
	cfg *Config
}

func (p *Plugin) Init(cfg Configurer, logger Logger) error {
	const op = errors.Op("my_plugin_init")

	if !cfg.Has(name) {
		return errors.E(errors.Disabled)
	}

	err := cfg.UnmarshalKey(name, &p.cfg)
	if err != nil {
		return errors.E(op, err)
	}

	p.cfg.InitDefaults()

	p.log = logger.NamedLogger(name)

	return nil
}

func (p *Plugin) Serve() chan error {
	errCh := make(chan error, 1)

	p.log.Info("HELLO FROM SERVE!")
	p.log.Info(p.cfg.keySample)

	return errCh
}

func (p *Plugin) Stop() error {
	return nil
}

func (p *Plugin) Name() string {
	return name
}
