package postgres

import (
	"os/exec"
)

type TestPG struct {
	containerID string
}

func (t *TestPG) Setup() error {
	cmd := exec.Command("docker", "run", "-e", "POSTGRES_HOST_AUTH_METHOD=trust",
		"-e", "POSTGRES_DB=broker", "-e", "POSTGRES_USER=admin",
		"-d", "-p", "5432:5432", "--name", "test-postgres", "postgres:latest")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	t.containerID = string(output)
	return nil
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
	cmd := exec.Command("docker", "stop", "test-postgres")

	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func (t *TestPG) remove() error {
	cmd := exec.Command("docker", "rm", "test-postgres")

	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
} 
