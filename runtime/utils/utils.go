package utils

import "github.com/Jeffail/gabs/v2"

func ResolveValue(key string, c *gabs.Container) interface{} {
	return c.Path(key).Data()
}

func ResolveValues(keys []string, c *gabs.Container) []interface{} {
	values := make([]interface{}, len(keys))
	for i, key := range keys {
		values[i] = c.Path(key).Data()
	}
	return values
}
