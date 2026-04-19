package options

import (
	"fmt"
	"os"
)

type Options struct {
	MachineID     string
	MachineFolder string

	Region      string
	DiskImage   string
	DiskSize    string
	MachineType string
	Token       string
	GitlabToken string
}

func FromEnv(skipMachine bool) (*Options, error) {
	retOptions := &Options{}

	var err error
	if !skipMachine {
		retOptions.MachineID, err = fromEnvOrError("MACHINE_ID")
		if err != nil {
			return nil, err
		}
		// prefix with devpod-
		retOptions.MachineID = "devpod-" + retOptions.MachineID
		
		retOptions.MachineFolder, err = fromEnvOrError("MACHINE_FOLDER")
		if err != nil {
			return nil, err
		}
	}

	retOptions.Token, err = fromEnvOrError("TOKEN")
	if err != nil {
		return nil, err
	}
	retOptions.DiskSize, err = fromEnvOrError("DISK_SIZE")
	if err != nil {
		return nil, err
	}
	retOptions.DiskImage, err = fromEnvOrError("DISK_IMAGE")
	if err != nil {
		return nil, err
	}
	retOptions.MachineType, err = fromEnvOrError("MACHINE_TYPE")
	if err != nil {
		return nil, err
	}
	retOptions.Region, err = fromEnvOrError("REGION")
	if err != nil {
		return nil, err
	}

	// Optional — if set, the token is written to /etc/environment on the
	// droplet at first boot so the DevPod agent (and anything on the machine)
	// can resolve ${localEnv:GITLAB_TOKEN} in devcontainer.json. End users
	// launching workspaces never see or set this.
	retOptions.GitlabToken = os.Getenv("GITLAB_TOKEN")

	return retOptions, nil
}

func fromEnvOrError(name string) (string, error) {
	val := os.Getenv(name)
	if val == "" {
		return "", fmt.Errorf("couldn't find option %s in environment, please make sure %s is defined", name, name)
	}

	return val, nil
}
