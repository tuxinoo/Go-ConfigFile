package utils

import "testing"

func TestConfigAuthUrl(t *testing.T) {
	config := LoadConfig("config.yaml.test")

	result := config.AuthURL()
	assertion := "https://backend.local.domain/auth/v1/"

	if result != assertion {
		t.Errorf("AuthURL %s should be equal to '%s'", result, assertion)
	}
}

func TestConfigBackendUrl(t *testing.T) {
	config := LoadConfig("config.yaml.test")

	result := config.BackendURL()
	assertion := "https://backend.local.domain/api/v1/"

	if result != assertion {
		t.Errorf("BackendURL %s should be equal to '%s'", result, assertion)
	}
}
