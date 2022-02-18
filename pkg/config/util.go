package config

import (
	"fmt"

	"picasso/pkg/utils/kcast"
)

func deepSearch(m map[string]interface{}, path []string) map[string]interface{} {
	for _, k := range path {
		m2, ok := m[k]
		if !ok {
			m3 := make(map[string]interface{})
			m[k] = m3
			m = m3
			continue
		}
		m3, ok := m2.(map[string]interface{})
		if !ok {
			m3 = make(map[string]interface{})
			m[k] = m3
		}
		m = m3
	}
	return m
}

func lookup(prefix string, target map[string]interface{}, data map[string]interface{}, sep string) {
	for k, v := range target {
		pp := fmt.Sprintf("%s%s%s", prefix, sep, k)
		if prefix == "" {
			pp = fmt.Sprintf("%s", k)
		}
		if dd, err := kcast.ToStringMapE(v); err == nil {
			lookup(pp, dd, data, sep)
		} else {
			data[pp] = v
		}
	}
}

func (c *Config) traverse(sep string) map[string]interface{} {
	data := make(map[string]interface{})
	lookup("", c.override, data, sep)
	return data
}
