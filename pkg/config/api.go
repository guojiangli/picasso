package config

import (
	"strings"
	"time"

	"picasso/pkg/utils/kcast"
	"picasso/pkg/utils/kmap"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// ErrInvalidKey ...
var ErrInvalidKey = errors.New("invalid key")

// Get returns the value associated with the key
func (c *Config) Get(key string) interface{} {
	key = strings.ToLower(key)
	return c.find(key)
}

// Get returns the value associated with the key
func Get(key string) interface{} {
	key = strings.ToLower(key)
	return defaultConfig.find(key)
}

// GetString returns the value associated with the key as a string with default defaultConfig.
func GetString(key string) string {
	key = strings.ToLower(key)
	return defaultConfig.GetString(key)
}

// GetString returns the value associated with the key as a string.
func (c *Config) GetString(key string) string {
	key = strings.ToLower(key)
	return kcast.ToString(c.Get(key))
}

// GetBool returns the value associated with the key as a boolean with default defaultConfig.
func GetBool(key string) bool {
	key = strings.ToLower(key)
	return defaultConfig.GetBool(key)
}

// GetBool returns the value associated with the key as a boolean.
func (c *Config) GetBool(key string) bool {
	key = strings.ToLower(key)
	return kcast.ToBool(c.Get(key))
}

// GetInt returns the value associated with the key as an integer with default defaultConfig.
func GetInt(key string) int {
	key = strings.ToLower(key)
	return defaultConfig.GetInt(key)
}

// GetInt returns the value associated with the key as an integer.
func (c *Config) GetInt(key string) int {
	key = strings.ToLower(key)
	return kcast.ToInt(c.Get(key))
}

// GetInt64 returns the value associated with the key as an integer with default defaultConfig.
func GetInt64(key string) int64 {
	key = strings.ToLower(key)
	return defaultConfig.GetInt64(key)
}

// GetInt64 returns the value associated with the key as an integer.
func (c *Config) GetInt64(key string) int64 {
	key = strings.ToLower(key)
	return kcast.ToInt64(c.Get(key))
}

// GetFloat64 returns the value associated with the key as a float64 with default defaultConfig.
func GetFloat64(key string) float64 {
	key = strings.ToLower(key)
	return defaultConfig.GetFloat64(key)
}

// GetFloat64 returns the value associated with the key as a float64.
func (c *Config) GetFloat64(key string) float64 {
	key = strings.ToLower(key)
	return kcast.ToFloat64(c.Get(key))
}

// GetTime returns the value associated with the key as time with default defaultConfig.
func GetTime(key string) time.Time {
	key = strings.ToLower(key)
	return defaultConfig.GetTime(key)
}

// GetTime returns the value associated with the key as time.
func (c *Config) GetTime(key string) time.Time {
	key = strings.ToLower(key)
	return kcast.ToTime(c.Get(key))
}

// GetDuration returns the value associated with the key as a duration with default defaultConfig.
func GetDuration(key string) time.Duration {
	key = strings.ToLower(key)
	return defaultConfig.GetDuration(key)
}

// GetDuration returns the value associated with the key as a duration.
func (c *Config) GetDuration(key string) time.Duration {
	key = strings.ToLower(key)
	return kcast.ToDuration(c.Get(key))
}

// GetStringSlice returns the value associated with the key as a slice of strings with default defaultConfig.
func GetStringSlice(key string) []string {
	key = strings.ToLower(key)
	return defaultConfig.GetStringSlice(key)
}

// GetStringSlice returns the value associated with the key as a slice of strings.
func (c *Config) GetStringSlice(key string) []string {
	key = strings.ToLower(key)
	return kcast.ToStringSlice(c.Get(key))
}

// GetSlice returns the value associated with the key as a slice of strings with default defaultConfig.
func GetSlice(key string) []interface{} {
	key = strings.ToLower(key)
	return defaultConfig.GetSlice(key)
}

// GetSlice returns the value associated with the key as a slice of strings.
func (c *Config) GetSlice(key string) []interface{} {
	key = strings.ToLower(key)
	return kcast.ToSlice(c.Get(key))
}

// GetStringMap returns the value associated with the key as a map of interfaces with default defaultConfig.
func GetStringMap(key string) map[string]interface{} {
	key = strings.ToLower(key)
	return defaultConfig.GetStringMap(key)
}

// GetStringMap returns the value associated with the key as a map of interfaces.
func (c *Config) GetStringMap(key string) map[string]interface{} {
	key = strings.ToLower(key)
	return kcast.ToStringMap(c.Get(key))
}

// GetStringMapString returns the value associated with the key as a map of strings with default defaultConfig.
func GetStringMapString(key string) map[string]string {
	key = strings.ToLower(key)
	return defaultConfig.GetStringMapString(key)
}

// GetStringMapString returns the value associated with the key as a map of strings.
func (c *Config) GetStringMapString(key string) map[string]string {
	key = strings.ToLower(key)
	return kcast.ToStringMapString(c.Get(key))
}

// GetSliceStringMap returns the value associated with the slice of maps.
func (c *Config) GetSliceStringMap(key string) []map[string]interface{} {
	key = strings.ToLower(key)
	return kcast.ToSliceStringMap(c.Get(key))
}

// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings with default defaultConfig.
func GetStringMapStringSlice(key string) map[string][]string {
	key = strings.ToLower(key)
	return defaultConfig.GetStringMapStringSlice(key)
}

// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
func (c *Config) GetStringMapStringSlice(key string) map[string][]string {
	key = strings.ToLower(key)
	return kcast.ToStringMapStringSlice(c.Get(key))
}

func (c *Config) find(key string) interface{} {
	dd, ok := c.keyMap.Load(key)
	if ok {
		return dd
	}

	paths := strings.Split(key, c.keyDelim)
	c.mu.RLock()
	defer c.mu.RUnlock()
	m := kmap.DeepSearchInMap(c.override, paths[:len(paths)-1]...)
	dd = m[paths[len(paths)-1]]
	c.keyMap.Store(key, dd)
	return dd
}

// UnmarshalWithExpect unmarshal key, returns expect if failed
func UnmarshalWithExpect(key string, expect interface{}) interface{} {
	return defaultConfig.UnmarshalWithExpect(key, expect)
}

// UnmarshalWithExpect unmarshal key, returns expect if failed
func (c *Config) UnmarshalWithExpect(key string, expect interface{}) interface{} {
	err := c.UnmarshalKey(key, expect)
	if err != nil {
		return expect
	}
	return expect
}

// UnmarshalKey takes a single key and unmarshal it into a Struct with default defaultdefaultConfig.
func UnmarshalKey(key string, rawVal interface{}) error {
	return defaultConfig.UnmarshalKey(key, rawVal)
}

// UnmarshalKey takes a single key and unmarshal it into a Struct.
func (c *Config) UnmarshalKey(key string, rawVal interface{}) error {
	config := mapstructure.DecoderConfig{
		DecodeHook: mapstructure.StringToTimeDurationHookFunc(),
		Result:     rawVal,
		TagName:    defaultGetOptionTag,
	}
	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		return err
	}
	if key == "" {
		c.mu.RLock()
		defer c.mu.RUnlock()
		return decoder.Decode(c.override)
	}

	value := c.Get(key)
	if value == nil {
		return nil
	}
	return decoder.Decode(value)
}
