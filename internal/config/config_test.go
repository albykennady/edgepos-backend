package config

import "testing"

func TestLoadDefaultPort(t *testing.T) {

	t.Setenv("PORT", "")

	cfg := Load()

	if cfg.Port != "8080" {

		t.Errorf("expected default port 8080, got %s", cfg.Port)
	}

}

func TestLoadCustomPort(t *testing.T) {
	t.Setenv("PORT", "9000")
	cfg := Load()

	if cfg.Port != "9000" {
		t.Errorf("Custom Port expected %s got %s", "9000", cfg.Port)
	}
}
