package instance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateInstance(t *testing.T) {
	instance, err := CreateInstance("testInstance")
	assert.Nil(t, err, "failed to create esent instance: error")
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	assert.NotNil(t, instance, "failed to create esent instance: instance is nil")

	err = instance.Init()
	assert.Nil(t, err, "failed to init esent instance")
	if err != nil {
		assert.FailNow(t, err.Error())
	}
}
