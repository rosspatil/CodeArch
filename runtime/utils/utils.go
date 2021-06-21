package utils

import (
	"os"
	"regexp"
	"strings"

	"github.com/rosspatil/codearch/runtime/models"
)

const (
	envRegexTmp = "^\\${.+}$"
)

var (
	EnvRegex = regexp.MustCompile(envRegexTmp)
)

func ResolveValue(key string, c *models.Controller) interface{} {
	if value, found := ResolveEnvironmentVariable(key); found {
		return value
	}
	return c.Path(key).Data()
}

func ResolveValues(keys []string, c *models.Controller) []interface{} {
	values := make([]interface{}, len(keys))
	for i, key := range keys {
		values[i] = ResolveValue(key, c)
	}
	return values
}

func ResolveEnvironmentVariable(str string) (string, bool) {
	if EnvRegex.MatchString(str) {
		str = extractEnvVariable(str)
		return os.Getenv(str), true
	}
	return str, false
}

func extractEnvVariable(str string) string {
	str = strings.TrimPrefix(str, "${")
	str = strings.TrimSuffix(str, "}")
	return str
}
