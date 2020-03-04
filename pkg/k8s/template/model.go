package template

import "gopkg.in/yaml.v2"

type Source struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
}

func Build(data []byte) (*Source, error) {
	var t Source
	if err := yaml.Unmarshal(data, &t); err != nil {
		return nil, err
	}

	return &t, nil
}
