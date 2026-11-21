package config

import (
	"os"
	"testing"
	"time"
)

func TestConfig_DefaultValues(t *testing.T) {
	// Xóa các env test nếu có
	os.Unsetenv("SERVER_PORT")
	os.Unsetenv("APP_ENV")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("kỳ vọng load config thành công: %v", err)
	}

	if cfg.Server.Port != 8080 {
		t.Errorf("kỳ vọng port mặc định là 8080 nhưng nhận %d", cfg.Server.Port)
	}

	if cfg.App.Environment != EnvDevelopment {
		t.Errorf("kỳ vọng môi trường mặc định là development nhưng nhận %s", cfg.App.Environment)
	}

	if cfg.Server.ShutdownTimeout != 15*time.Second {
		t.Errorf("kỳ vọng timeout là 15s nhưng nhận %v", cfg.Server.ShutdownTimeout)
	}
}

func TestConfig_ValidateInvalidPort(t *testing.T) {
	cfg := &Config{
		Server: ServerConfig{Port: 70000},
		App:    AppConfig{Environment: EnvDevelopment},
	}

	if err := cfg.Validate(); err == nil {
		t.Error("kỳ vọng lỗi khi port vượt quá 65535 nhưng không có")
	}
}

func TestConfig_ValidateInvalidEnv(t *testing.T) {
	cfg := &Config{
		Server: ServerConfig{Port: 8080},
		App:    AppConfig{Environment: "unknown_env"},
	}

	if err := cfg.Validate(); err == nil {
		t.Error("kỳ vọng lỗi khi môi trường không hợp lệ nhưng không có")
	}
}
