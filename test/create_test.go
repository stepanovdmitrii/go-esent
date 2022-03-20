package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stepanovdmitrii/go-esent/pkg/config"
	"github.com/stepanovdmitrii/go-esent/pkg/instance"
	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	if e := os.RemoveAll(testDataDir); e != nil {
		assert.FailNowf(t, "drop testdata dir error", "%s", e.Error())
	}

	defer func() {
		if e := os.RemoveAll(testDataDir); e != nil {
			assert.Fail(t, "failed to clean testdata directory", e.Error())
		}
	}()

	inst, err := instance.CreateInstance("createInstance_test")
	if err != nil {
		assert.FailNow(t, "failed to create esent instance", err.Error())
	}
	if inst == nil {
		assert.FailNow(t, "instance must be not nil")
	}

	defer func() {
		if e := inst.Term(); e != nil {
			assert.Fail(t, "failed to terminate instance", e.Error())
		}
	}()

	if e := os.MkdirAll(testDataDir, 0755); e != nil {
		assert.FailNow(t, "create testdata dir error", e.Error())
	}

	systemPath := filepath.Join(testDataDir, "system")
	logsPath := filepath.Join(testDataDir, "logs")
	tempPath := filepath.Join(testDataDir, "temp")
	baseName := "goo"

	err = inst.SetParameters(
		instance.SysParameterCreatorBool(config.SysParameterCreatePathIfNotExist, true),
		instance.SysParameterCreatorBool(config.SysParameterCircularLog, true),

		instance.SysParameterCreatoString(config.SysParameterSystemPath, systemPath),
		instance.SysParameterCreatoString(config.SysParameterBaseName, baseName),
		instance.SysParameterCreatoString(config.SysParameterLogFilePath, logsPath),
		instance.SysParameterCreatoString(config.SysParameterTempPath, tempPath),

		instance.SysParameterCreatoUInt32(config.SysParameterDbExtensionSize, 512),
	)
	if err != nil {
		assert.FailNow(t, "failed to configure esent instance", err.Error())
	}

	err = inst.Init()
	if err != nil {
		assert.FailNow(t, "failed to init esent instance", err.Error())
	}

	assert.DirExists(t, systemPath, "system directory doesn't exist")
	assert.DirExists(t, logsPath, "logs directory doesn't exist")
	assert.NoDirExists(t, tempPath, "temp path was created")

	assert.FileExists(t, logsPath+"\\"+baseName+".log", "log file doesn't exist")
	assert.FileExists(t, logsPath+"\\"+baseName+"res00001.jrs", "journal file 1 doesn't exist")
	assert.FileExists(t, logsPath+"\\"+baseName+"res00002.jrs", "journal file 2 doesn't exist")
	assert.FileExists(t, logsPath+"\\"+baseName+"tmp.log", "temp log file doesn't exist")
	assert.FileExists(t, systemPath+"\\"+baseName+".chk", "checkpoint file doesn't exist")
}
