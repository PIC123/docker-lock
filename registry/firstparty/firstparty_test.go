package firstparty

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/michaelperel/docker-lock/registry"
)

// TestDefaultWrapper ensures that the default wrapper
// is one that can handle images without a prefix.
func TestDefaultWrapper(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	client := &registry.HTTPClient{}

	wrapper, err := DefaultWrapper(client, "")
	if err != nil {
		t.Fatal("could not get default wrapper")
	}

	if _, ok := wrapper.(*DockerWrapper); !ok {
		t.Fatal("expected DockerWrapper")
	}
}
