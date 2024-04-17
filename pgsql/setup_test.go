package pgsql

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWhenSettingUpNewContainerWithDefaultValueReturnNoError(t *testing.T) {
	con := TestPG{}
	err := con.Setup()
	assert.NoError(t, err)
}

func TestWhenTearingDownContainerReturnNoError(t *testing.T) {
	con := TestPG{}
	err := con.Setup()
	assert.NoError(t, err)

	time.Sleep(10*time.Second)

	err = con.TearDown()
	assert.NoError(t, err)
}
