package gproject

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"runtime"

	"golang.org/x/oauth2/google"
)

func Default() string {
	explicit := FromEnv()
	if explicit != "" {
		return explicit
	}

	ctx := context.Background()
	credentials, err := google.FindDefaultCredentials(ctx)
	if err == nil && credentials.ProjectID != "" {
		return credentials.ProjectID
	}

	return FromCommand()
}

func FromEnv() string {
	keys := []string{"GOOGLE_CLOUD_PROJECT", "GCLOUD_PROJECT"}
	for _, key := range keys {
		if v := os.Getenv(key); v != "" {
			return v
		}
	}
	return ""
}

func FromCommand() string {
	var cmdName string
	switch runtime.GOOS {
	case "windows":
		cmdName = "gcloud.cmd"
	default:
		cmdName = "gcloud"
	}
	cmd := exec.Command(cmdName, "config", "config-helper", "--format", "json")
	output := bytes.NewBuffer(nil)
	cmd.Stdout = output

	if err := cmd.Run(); err != nil {
		return ""
	}

	data := struct {
		Configuration struct {
			Properties struct {
				Core struct {
					Project string
				}
			}
		}
	}{}

	if err := json.NewDecoder(output).Decode(&data); err != nil {
		return ""
	}
	return data.Configuration.Properties.Core.Project
}
