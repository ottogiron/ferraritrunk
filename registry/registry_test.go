package registry

import (
	"testing"

	"github.com/ottogiron/ferraritrunk/config"
)

func TestRegisterBackendFactory(t *testing.T) {

	cs := &config.ConfigurationSchema{
		Name: "test",
	}

	err := RegisterBackendFactory(nil, cs)

	if err != nil {
		t.Errorf("The first factory should be registered correctly for %s", cs.Name)
	}

	err = RegisterBackendFactory(nil, cs)

	if err == nil {
		t.Errorf("The registration should fail for %s", cs.Name)
	}
}

func TestGetBackendSchemas(t *testing.T) {
	RegisterBackendFactory(nil, &config.ConfigurationSchema{
		Name: "test",
	})

	schemas := BackendSchemas()

	slen := len(schemas)
	if slen != 1 {
		t.Errorf("Expected schemas size %d was", slen)
	}
}

func TestGetBackendrSchema(t *testing.T) {
	RegisterBackendFactory(nil, &config.ConfigurationSchema{Name: "test"})

	schema, _ := BackendSchema("test")

	if schema.Name != "test" {
		t.Errorf("expected schema name to be 'test' was %s", schema.Name)
	}
}
