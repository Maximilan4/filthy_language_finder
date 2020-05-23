package settings

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	def := getEnvironment("TEST", "default")
	if def != "default" {
		t.Errorf("Expected %s, got %s", "default", def)
	}
	os.Setenv("TEST", "test")
	def = getEnvironment("TEST", "default")
	if def != "test" {
		t.Errorf("Expected %s, got %s", "test", def)
	}
	os.Unsetenv("TEST")
}
