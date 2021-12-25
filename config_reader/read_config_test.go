package config

import (
	"os"
	"testing"
)

func TestProcessNoConfigFilePresent(t *testing.T) {

	_, err := ReadConfig()
	if err == nil {
		t.Errorf("ReadConfig method without any valid config file should fail.")
	} else {
		if err.Error() != "Fatal error reading config file: Config File \"config\" Not Found in \"[/etc/webcam-snapshot-worker]\"" {
			t.Errorf("Default config should be in /etc/webcam-snapshot-worker/config.toml, not in other place, error was '%s'.", err.Error())
		}
	}
}

func TestProcessServerNoDataInConfig(t *testing.T) {
	os.Setenv("WEBCAM_SNAPSHOT_WORKER_SERVICE_CONFIG_FILE_LOCATION", "./config_files_test/server_no_data/")
	_, err := ReadConfig()
	if err == nil {
		t.Errorf("ReadConfig method without server data config should fail.")
	} else {
		if err.Error() != "Fatal error config: no server host was found." {
			t.Errorf("Error should be \"Fatal error config: no server host was found.\" but error was '%s'.", err.Error())
		}
	}
}

func TestProcessServerOnlyHostInConfig(t *testing.T) {
	os.Setenv("WEBCAM_SNAPSHOT_WORKER_SERVICE_CONFIG_FILE_LOCATION", "./config_files_test/server_only_host/")
	_, err := ReadConfig()
	if err == nil {
		t.Errorf("ReadConfig method without server port should fail.")
	} else {
		if err.Error() != "Fatal error config: no server port was found." {
			t.Errorf("Error should be \"Fatal error config: no server port was found.\" but error was '%s'.", err.Error())
		}
	}
}

func TestProcessServerNoUserPasswordInConfig(t *testing.T) {
	os.Setenv("WEBCAM_SNAPSHOT_WORKER_SERVICE_CONFIG_FILE_LOCATION", "./config_files_test/server_only_host_port/")
	_, err := ReadConfig()
	if err == nil {
		t.Errorf("ReadConfig method without user password should fail.")
	} else {
		if err.Error() != "Fatal error config: no server user was found." {
			t.Errorf("Error should be \"Fatal error config: no server user was found.\" but error was '%s'.", err.Error())
		}
	}
}

func TestOKConfig(t *testing.T) {
	os.Setenv("WEBCAM_SNAPSHOT_WORKER_SERVICE_CONFIG_FILE_LOCATION", "./config_files_test/ok/")
	config, err := ReadConfig()
	if err != nil {
		t.Errorf("ReadConfig with ok config shouln't return errors. Returned: %s.", err.Error())
	}
	if config.Server.Host != "localhost" {
		t.Errorf("Server Host should be localhost. Returned: %s.", config.Server.Host)
	}
	if config.Server.Port != 5672 {
		t.Errorf("Server Port should be 5672. Returned: %d.", config.Server.Port)
	}
	if config.Server.User != "guest" {
		t.Errorf("Server Host should be guest. Returned: %s.", config.Server.User)
	}
	if config.Server.Password != "pass" {
		t.Errorf("Server Password should be pass. Returned: %s.", config.Server.Password)
	}

	if config.Incoming.Name != "incoming" {
		t.Errorf("Incoming name should be incoming. Returned: %s.", config.Incoming.Name)
	}

	if config.Outgoing.Name != "outgoing" {
		t.Errorf("Outgoing name should be outgoing. Returned: %s.", config.Outgoing.Name)
	}
}
