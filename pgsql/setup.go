package pgsql

import (
	"fmt"
	"os/exec"
)

const (
	CONTAINER_NAME = "test-postgres"
	DEFAULT_IMAGE  = "postgres:latest"
	DEFAULT_USER   = "root"
	DEFAULT_DB     = "default"
	DEFAULT_PORT   = 5432
)

type Config struct {
	ImageName  string
	DBName     string
	DBUser     string
	DBPassword string
	PORT       int32
}

type TestPG struct {
	containerID string
}

func (t *TestPG) Setup(config *Config) error {
	config = setDefaultConfigValues(config)
	var passString string
	if config.DBPassword == "" {
		passString = "POSTGRES_HOST_AUTH_METHOD=trust"
	} else {
		passString = fmt.Sprintf("POSTGRES_PASSWORD=%s", config.DBPassword)
	}

	cmd := exec.Command("docker", "run",
		"-e", fmt.Sprintf("POSTGRES_DB=%s", config.DBName),
		"-e", fmt.Sprintf("POSTGRES_USER=%s", config.DBUser),
		"-e", passString,
		"-d", "-p", fmt.Sprintf("%d:5432", config.PORT),
		"--name", CONTAINER_NAME, config.ImageName)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	t.containerID = string(output)
	return nil
}

func setDefaultConfigValues(config *Config) *Config {
	if config == nil {
		config = &Config{}
	}

	if config.ImageName == "" {
		config.ImageName = DEFAULT_IMAGE
	}

	if config.DBUser == "" {
		config.DBUser = DEFAULT_USER
	}

	if config.DBName == "" {
		config.DBName = DEFAULT_DB
	}

	if config.PORT == 0 {
		config.PORT = DEFAULT_PORT
	}
	return config
}

func (t *TestPG) TearDown() error {
	err := t.stop()
	if err != nil {
		return err
	}

	err = t.remove()
	if err != nil {
		return err
	}
	return nil
}

func (t *TestPG) stop() error {
	cmd := exec.Command("docker", "stop", CONTAINER_NAME)

	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func (t *TestPG) remove() error {
	cmd := exec.Command("docker", "rm", CONTAINER_NAME)

	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func (t *TestPG) Execute(query string) error {
	cmd := exec.Command("docker", "exec", CONTAINER_NAME,
		"psql", "-p", "5432", "-U", DEFAULT_USER, "-d", DEFAULT_DB, "-c", query)

	output, err := cmd.CombinedOutput()
	println(string(output))
	if err != nil {
		return err
	}

	return nil
}
