package porter

import (
	"testing"

	"github.com/deislabs/porter/pkg/exec"

	"github.com/deislabs/porter/pkg/config"
	"github.com/deislabs/porter/pkg/mixin"
)

type TestPorter struct {
	*Porter
	TestConfig *config.TestConfig
}

// NewTestPorter initializes a porter test client, with the output buffered, and an in-memory file system.
func NewTestPorter(t *testing.T) *TestPorter {
	tc := config.NewTestConfig(t)
	p := New()
	p.Config = tc.Config
	p.MixinProvider = &TestMixinProvider{}
	return &TestPorter{
		Porter:     p,
		TestConfig: tc,
	}
}

// TODO: use this later to not actually execute a mixin during a unit test
type TestMixinProvider struct {
}

func (p *TestMixinProvider) GetMixins() ([]mixin.Metadata, error) {
	mixins := []mixin.Metadata{
		{Name: "exec"},
	}
	return mixins, nil
}

func (p *TestMixinProvider) GetMixinSchema(m mixin.Metadata) (string, error) {
	t := exec.NewSchemaBox()
	return t.FindString("exec.json")
}
