package instance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateInstance(t *testing.T) {
	instance, err := CreateInstance("test113213")
	assert.Nil(t, err, "failed to create esent instance: error")
	assert.NotNil(t, instance, "failed to create esent instance: instance is nil")
	err = instance.Init()
	assert.Nil(t, err, "failed to init esent instance")
}
