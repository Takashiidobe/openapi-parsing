package steps

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ClientMethod struct {
	Package string   `yaml:"package"`
	Client  string   `yaml:"client"`
	Method  string   `yaml:"method"`
	Args    []string `yaml:"args"`
}

type ChannelStep struct {
	ID           string       `yaml:"id"`
	ClientMethod ClientMethod `yaml:"clientMethod"`
	Children     []Step       `yaml:"children,omitempty"`
}

type ResourceStep struct {
	ID       string `yaml:"id"`
	Resource string `yaml:"resource"`
}

type PayloadStep struct {
	ID          string   `yaml:"id"`
	ExcludeTags []string `yaml:"excludeTags,omitempty"`
}

type Step struct {
	Channel  *ChannelStep  `yaml:"channel,omitempty"`
	Resource *ResourceStep `yaml:"resource,omitempty"`
	Payload  *PayloadStep  `yaml:"payload,omitempty"`
}

type StepTree struct {
	Kind     string `yaml:"kind"`
	ID       string `yaml:"id"`
	RootStep Step   `yaml:"rootStep"`
}

func WriteStepTree(path string, tree StepTree) error {
	data, err := yaml.Marshal(tree)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func OpChannelStep(opPath, opID, sdkVersion string) Step {
	cm := ClientMethod{
		Package: opPath,
		Client:  opID,
		Method:  "",
	}
	ch := &ChannelStep{ID: opID, ClientMethod: cm}
	return Step{Channel: ch}
}
