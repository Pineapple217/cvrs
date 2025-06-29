package config

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"strings"

	_ "github.com/joho/godotenv/autoload"

	"github.com/go-viper/mapstructure/v2"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	Workforce Workforce `yaml:"workforce"`
	Database  Database  `yaml:"database"`
}

func (c *Config) SetDefault() {
	c.Workforce.SetDefault()
	c.Database.SetDefault()
}

func Load() (Config, error) {
	slog.Info("Loading configs")
	k := koanf.New(".")

	var conf Config
	conf.SetDefault()

	err := k.Load(file.Provider("./config.yaml"), yaml.Parser())
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return Config{}, fmt.Errorf("failed to load config file: %w", err)
		}
		slog.Warn("Config file not found, falling back to environment variables")
	}

	unmarshalCfg := koanf.UnmarshalConf{
		Tag:       "yaml",
		FlatPaths: false,
		DecoderConfig: &mapstructure.DecoderConfig{
			DecodeHook: mapstructure.ComposeDecodeHookFunc(
				mapstructure.StringToTimeDurationHookFunc()),
			Metadata:         nil,
			Result:           &conf,
			WeaklyTypedInput: true,
			ErrorUnused:      true,
			TagName:          "yaml",
		},
	}

	err = k.UnmarshalWithConf("", &conf, unmarshalCfg)
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	err = k.Load(env.ProviderWithValue("", ".", func(s string, v string) (string, any) {
		key := strings.ReplaceAll(strings.ToLower(s), "_", ".")
		// Check to exist if we have a configuration option already and see if it's a slice
		// If there is a comma in the value, split the value into a slice by the comma.
		if strings.Contains(v, ",") {
			return key, strings.Split(v, ",")
		}

		// Otherwise return the new key with the unaltered value
		return key, v
	}), nil)
	if err != nil {
		return Config{}, err
	}

	keys := make(map[string]any, len(k.Keys()))
	for _, key := range k.Keys() {
		keys[strings.ToLower(key)] = k.Get(key)
	}
	k.Delete("")
	err = k.Load(confmap.Provider(keys, "."), nil)
	if err != nil {
		return Config{}, err
	}

	unmarshalCfg.DecoderConfig.ErrorUnused = false
	unmarshalCfg.DecoderConfig.ZeroFields = true // Empty default slices/maps if a value is configured
	err = k.UnmarshalWithConf("", &conf, unmarshalCfg)
	if err != nil {
		return Config{}, err
	}

	return conf, nil
}
