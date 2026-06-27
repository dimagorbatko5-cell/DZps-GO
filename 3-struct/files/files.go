package files

import (
	"os"
	"strings"
)

func ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func WriteFile(content []byte, name string) error {
	return os.WriteFile(name, content, 0644)
}

func IsJSON(name string) bool {
	return strings.HasSuffix(name, ".json")
}
