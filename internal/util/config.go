package util

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

func ReadYamlFile[T any](path string) (*T, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to os.Open(): %w", err)
	}
	defer f.Close()

	var v T
	if err := yaml.NewDecoder(f).Decode(&v); err != nil {
		return nil, fmt.Errorf("failed to yaml.NewDecoder().Decode(): %w", err)
	}
	return &v, nil
}
