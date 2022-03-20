package instance

import (
	"github.com/stepanovdmitrii/go-esent/pkg/config"
)

// SysParameterCreator
type SysParameterCreator func() (*config.SystemParameter, error)

// SysParameterCreatorBool ...
func SysParameterCreatorBool(fn func(bool) (*config.SystemParameter, error), value bool) SysParameterCreator {
	return func() (*config.SystemParameter, error) {
		return fn(value)
	}
}

// SysParameterCreatoString ...
func SysParameterCreatoString(fn func(string) (*config.SystemParameter, error), value string) SysParameterCreator {
	return func() (*config.SystemParameter, error) {
		return fn(value)
	}
}

// SysParameterCreatoUInt32 ...
func SysParameterCreatoUInt32(fn func(uint32) (*config.SystemParameter, error), value uint32) SysParameterCreator {
	return func() (*config.SystemParameter, error) {
		return fn(value)
	}
}

// SetParameters Set many parameters to instance
func (i *Instance) SetParameters(parameters ...SysParameterCreator) error {
	for _, parameter := range parameters {
		value, err := parameter()
		if err != nil {
			return err
		}
		if err = i.SetSystemParameter(value); err != nil {
			return err
		}
	}
	return nil
}
