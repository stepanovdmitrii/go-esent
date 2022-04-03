package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stepanovdmitrii/go-esent/pkg/config"
	"github.com/stepanovdmitrii/go-esent/pkg/esent"
	"github.com/stretchr/testify/assert"
)

const (
	testDataDir      string = "testdata"
	testDatabaseName string = "testDatabase.edb"
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

	inst, err := esent.CreateInstance("createInstance_test")
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
	databaseDirPath := filepath.Join(testDataDir, "database")
	databaseFilePath := filepath.Join(databaseDirPath, testDatabaseName)
	baseName := "goo"

	err = inst.SetParameters(
		esent.SysParameterCreatorBool(config.SysParameterCreatePathIfNotExist, true),
		esent.SysParameterCreatorBool(config.SysParameterCircularLog, true),

		esent.SysParameterCreatoString(config.SysParameterSystemPath, systemPath),
		esent.SysParameterCreatoString(config.SysParameterBaseName, baseName),
		esent.SysParameterCreatoString(config.SysParameterLogFilePath, logsPath),
		esent.SysParameterCreatoString(config.SysParameterTempPath, tempPath),

		esent.SysParameterCreatoUInt32(config.SysParameterDbExtensionSize, 512),
	)
	if err != nil {
		assert.FailNow(t, "failed to configure esent instance", err.Error())
	}

	err = inst.Init()
	if err != nil {
		assert.FailNow(t, "failed to init esent instance", err.Error())
	}

	var session *esent.Session
	session, err = inst.BeginSession()
	if err != nil {
		assert.FailNow(t, "failed to begin session", err.Error())
	}
	defer func() {
		if e := session.EndSession(); e != nil {
			assert.Fail(t, "failed to end session", e.Error())
		}
	}()

	var database *esent.Database
	database, err = session.CreateDatabase(databaseFilePath, esent.DbDefault)
	if err != nil {
		assert.FailNow(t, "failed to create database", err.Error())
	}
	defer func() {
		if e := database.CloseDatabase(); e != nil {
			assert.Fail(t, "failed to close database", e.Error())
		}
	}()

	var table *esent.Table
	table, err = database.CreateTable("test", esent.DefaultTablePages, esent.DefaultTableDensity)
	if err != nil {
		assert.FailNow(t, "failed to create table", err.Error())
	}
	defer func() {
		if e := table.CloseTable(); e != nil {
			assert.Fail(t, "failed to close table", e.Error())
		}
	}()

	assert.DirExists(t, systemPath, "system directory doesn't exist")
	assert.DirExists(t, logsPath, "logs directory doesn't exist")
	assert.DirExists(t, databaseDirPath, "database directory doesn't exist")
	assert.NoDirExists(t, tempPath, "temp path was created")

	assert.FileExists(t, logsPath+"\\"+baseName+".log", "log file doesn't exist")
	assert.FileExists(t, logsPath+"\\"+baseName+"res00001.jrs", "journal file 1 doesn't exist")
	assert.FileExists(t, logsPath+"\\"+baseName+"res00002.jrs", "journal file 2 doesn't exist")
	assert.FileExists(t, logsPath+"\\"+baseName+"tmp.log", "temp log file doesn't exist")
	assert.FileExists(t, systemPath+"\\"+baseName+".chk", "checkpoint file doesn't exist")
	assert.FileExists(t, databaseFilePath, "database file doesn't exist")
}
