package pgsql

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWhenSetupWithDefaultSetupAndTeardownReturnsNoError(t *testing.T) {
	con := TestPG{}
	err := con.Setup(nil)
	assert.NoError(t, err)

	time.Sleep(1 * time.Second)

	err = con.TearDown()
	assert.NoError(t, err)
}

func TestWhenImageNameIsProvidedAndDoesNotExistReturnsError(t *testing.T) {
	con := TestPG{}
	config := &Config{
		ImageName: "non-existent:latest",
	}
	err := con.Setup(config)
	assert.Error(t, err)
}

func TestWhenImageNameIsProvidedAndExistReturnsError(t *testing.T) {
	con := TestPG{}
	config := &Config{
		ImageName: "postgres:latest",
	}
	err := con.Setup(config)
	assert.NoError(t, err)

	time.Sleep(1 * time.Second)

	err = con.TearDown()
	assert.NoError(t, err)
}

func TestWhenExecutingQueryReturnsNoError(t *testing.T) {
	con := TestPG{}
	err := con.Setup(nil)
	assert.NoError(t, err)

	//Important! wait for database to setup
	time.Sleep(3 * time.Second)
	err = con.Execute("SELECT * FROM pg_catalog.pg_tables;")
	assert.NoError(t, err)

	err = con.TearDown()
	assert.NoError(t, err)
}

func TestWhenDatabaseHasPasswordItExecutesQuery(t *testing.T) {
	con := TestPG{}
	config := &Config{
		DBPassword: "pass",
	}
	err := con.Setup(config)
	assert.NoError(t, err)

	time.Sleep(3 * time.Second)
	err = con.Execute("SELECT * FROM pg_catalog.pg_tables;")
	assert.NoError(t, err)

	err = con.TearDown()
	assert.NoError(t, err)
}

func TestWhenPortIsProvidedItExecutesQuery(t *testing.T) {
	con := TestPG{}
	config := &Config{
		PORT: 1234,
	}
	err := con.Setup(config)
	assert.NoError(t, err)

	time.Sleep(3 * time.Second)
	err = con.Execute("SELECT * FROM pg_catalog.pg_tables;")
	assert.NoError(t, err)

	err = con.TearDown()
	assert.NoError(t, err)
}
