// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package fileexporter

import "encoding/json"
import "fmt"
import "reflect"
import "time"

// Config defines configuration for file exporter.
type Config struct {
	// Compression Codec used to export telemetry data. Supported compression
	// algorithms: `zstd`.
	Compression *ConfigCompression `json:"compression,omitempty" yaml:"compression,omitempty" mapstructure:"compression,omitempty"`

	// FlushInterval corresponds to the JSON schema field "flush_interval".
	FlushInterval *time.Duration `json:"flush_interval,omitempty" yaml:"flush_interval,omitempty" mapstructure:"flush_interval,omitempty"`

	// FormatType defines the data format of encoded telemetry data
	// Options:
	// - json [default]: OTLP json bytes.
	// - proto: OTLP binary protobuf bytes.
	Format ConfigFormat `json:"format,omitempty" yaml:"format,omitempty" mapstructure:"format,omitempty"`

	// Path of the file to write to. Path is relative to the current directory.
	Path string `json:"path" yaml:"path" mapstructure:"path"`

	// Rotation corresponds to the JSON schema field "rotation".
	Rotation *ConfigRotation `json:"rotation,omitempty" yaml:"rotation,omitempty" mapstructure:"rotation,omitempty"`
}

type ConfigCompression string

const ConfigCompressionBlank ConfigCompression = ""
const ConfigCompressionZstd ConfigCompression = "zstd"

type ConfigFormat string

const ConfigFormatJson ConfigFormat = "json"
const ConfigFormatProto ConfigFormat = "proto"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ConfigFormat) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ConfigFormat {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ConfigFormat, v)
	}
	*j = ConfigFormat(v)
	return nil
}

var enumValues_ConfigFormat = []interface{}{
	"json",
	"proto",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ConfigCompression) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ConfigCompression {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ConfigCompression, v)
	}
	*j = ConfigCompression(v)
	return nil
}

type ConfigRotation struct {
	// Localtime corresponds to the JSON schema field "localtime".
	Localtime *bool `json:"localtime,omitempty" yaml:"localtime,omitempty" mapstructure:"localtime,omitempty"`

	// MaxBackups corresponds to the JSON schema field "max_backups".
	MaxBackups *int `json:"max_backups,omitempty" yaml:"max_backups,omitempty" mapstructure:"max_backups,omitempty"`

	// MaxDays corresponds to the JSON schema field "max_days".
	MaxDays *int `json:"max_days,omitempty" yaml:"max_days,omitempty" mapstructure:"max_days,omitempty"`

	// MaxMegabytes corresponds to the JSON schema field "max_megabytes".
	MaxMegabytes *int `json:"max_megabytes,omitempty" yaml:"max_megabytes,omitempty" mapstructure:"max_megabytes,omitempty"`
}

var enumValues_ConfigCompression = []interface{}{
	"",
	"zstd",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Config) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["path"]; !ok || v == nil {
		return fmt.Errorf("field path in Config: required")
	}
	type Plain Config
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["format"]; !ok || v == nil {
		plain.Format = "json"
	}
	*j = Config(plain)
	return nil
}

func (cfg *Config)ValidateHelper() error {
	b, err := json.Marshal(cfg)
	if err != nil {
			return err
	}
	var config Config
	if err := json.Unmarshal(b, &config); err != nil {
			return err
	}
	return nil
}