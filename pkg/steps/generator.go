package steps

import (
        "fmt"
        "os"
        "strings"
        "unicode"

        "openapi-parsing/pkg/openapi"
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


func toAzureResourceName(name string) string {
        clean := name
        if strings.HasSuffix(clean, "ListResult") {
                clean = strings.TrimSuffix(clean, "ListResult")
        }
        var b strings.Builder
        b.WriteString("azure_")
        for i, r := range clean {
                if unicode.IsUpper(r) && i != 0 {
                        b.WriteByte('_')
                }
                b.WriteRune(unicode.ToLower(r))
        }
        return b.String()
}

type stepNode struct {
        ID       string     `yaml:"id"`
        Children []stepNode `yaml:"children,omitempty"`
}

type executor struct{
        Default any `yaml:"default"`
}

type yamlOutput struct {
        Kind     string    `yaml:"kind"`
        ID       string    `yaml:"id"`
        Executor executor  `yaml:"executor"`
        RootStep stepNode  `yaml:"rootStep"`
}

func stepTreeFromRoot(root Step) string {
        id := "UnknownCrawler"
        if root.Channel != nil {
                name := root.Channel.ID
                if strings.HasSuffix(name, "ListResult") {
                        name = strings.TrimSuffix(name, "ListResult")
                }
                id = name + "Crawler"
        }
        out := yamlOutput{
                Kind: "StepTree",
                ID:   id,
                Executor: executor{Default: map[string]any{}},
                RootStep: stepNodeFromStep(root),
        }
        data, err := yaml.Marshal(out)
        if err != nil {
                return ""
        }
        return string(data)
}

func stepNodeFromStep(s Step) stepNode {
        switch {
        case s.Channel != nil:
                var kids []stepNode
                for _, c := range s.Channel.Children {
                        kids = append(kids, stepNodeFromStep(c))
                }
                return stepNode{ID: s.Channel.ID, Children: kids}
        case s.Resource != nil:
                return stepNode{ID: s.Resource.ID}
        case s.Payload != nil:
                return stepNode{ID: s.Payload.ID}
        }
        return stepNode{}
}

func serializeStepRecursive(step Step) []string {
        schema := "# yaml-language-server: $schema=../../../../../crawler-generator/crawler.schema.json\n"
        switch {
        case step.Channel != nil:
                serialized, _ := yaml.Marshal(struct {
                        Kind         string       `yaml:"kind"`
                        ID           string       `yaml:"id"`
                        ClientMethod ClientMethod `yaml:"clientMethod"`
                }{
                        Kind:         "ChannelStep",
                        ID:           step.Channel.ID,
                        ClientMethod: step.Channel.ClientMethod,
                })
                var out []string
                out = append(out, schema+string(serialized))
                for _, child := range step.Channel.Children {
                        out = append(out, serializeStepRecursive(child)...)
                }
                return out
        case step.Resource != nil:
                serialized, _ := yaml.Marshal(struct {
                        Kind     string `yaml:"kind"`
                        ID       string `yaml:"id"`
                        Resource string `yaml:"resource"`
                }{
                        Kind:     "ResourceStep",
                        ID:       step.Resource.ID,
                        Resource: toAzureResourceName(step.Resource.Resource),
                })
                return []string{schema + string(serialized)}
        case step.Payload != nil:
                serialized, _ := yaml.Marshal(struct {
                        Kind string `yaml:"kind"`
                        ID   string `yaml:"id"`
                }{
                        Kind: "PayloadStep",
                        ID:   step.Payload.ID,
                })
                return []string{schema + string(serialized)}
        }
        return nil
}

func SerializeSteps(root Step) string {
        pieces := serializeStepRecursive(root)
        return strings.Join(pieces, "---\n")
}

func WriteStepTreeAndSteps(path string, root Step) error {
        schema := "# yaml-language-server: $schema=../../../../../crawler-generator/crawler.schema.json\n"
        tree := stepTreeFromRoot(root)
        var b strings.Builder
        b.WriteString(schema)
        b.WriteString(tree)
        b.WriteString("---\n")
        b.WriteString(SerializeSteps(root))
        return os.WriteFile(path, []byte(b.String()), 0644)
}

func OpResourceStep(op openapi.Op) Step {
        return Step{Resource: &ResourceStep{ID: op.ResponseType, Resource: op.ResponseType}}
}

func OpPayloadStep(op openapi.Op) Step {
        return Step{Payload: &PayloadStep{ID: op.ResponseType}}
}

func OpChannelStep(op openapi.Op, sdkVersion string) Step {
        var args []string
        for _, p := range op.Params {
                if p.Name == "api-version" {
                        continue
                }
                args = append(args, p.Name)
        }
        cm := ClientMethod{
                Package: fmt.Sprintf("github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/%s/arm%s/%s", op.ResourceProvider, op.ResourceProvider, sdkVersion),
                Client:  op.Client,
                Method:  op.Method,
                Args:    args,
        }
        return Step{Channel: &ChannelStep{ID: op.ResponseType, ClientMethod: cm}}
}

func GenerateSteps(ops []openapi.Op, sdkVersion string) Step {
        if len(ops) == 0 {
                return Step{}
        }
        rootChan := OpChannelStep(ops[0], sdkVersion).Channel
        rootChan.Children = append(rootChan.Children, OpResourceStep(ops[0]))
        rootChan.Children = append(rootChan.Children, OpPayloadStep(ops[0]))
        current := rootChan
        for _, op := range ops[1:] {
                nextChan := OpChannelStep(op, sdkVersion).Channel
                nextChan.Children = append(nextChan.Children, OpResourceStep(op))
                nextChan.Children = append(nextChan.Children, OpPayloadStep(op))
                current.Children = append(current.Children, Step{Channel: nextChan})
                current = current.Children[len(current.Children)-1].Channel
        }
        return Step{Channel: rootChan}
}
