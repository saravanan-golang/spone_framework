package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zhufuyi/sponge/pkg/gofile"

	"user_module/configs"
)

func TestInit(t *testing.T) {
	configFile := configs.Path("user_service.yml")
	err := Init(configFile)
	if gofile.IsExists(configFile) {
		assert.NoError(t, err)
	} else {
		assert.Error(t, err)
	}

	c := Get()
	assert.NotNil(t, c)

	str := Show()
	assert.NotEmpty(t, str)
	t.Log(str)

	// set nil
	Set(nil)
	defer func() {
		recover()
	}()
	Get()
}

func TestInitNacos(t *testing.T) {
	configFile := configs.Path("user_service_cc.yml")
	_, err := NewCenter(configFile)
	if gofile.IsExists(configFile) {
		assert.NoError(t, err)
	} else {
		assert.Error(t, err)
	}
}
