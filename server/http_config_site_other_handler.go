package server

import (
	"github.com/evcc-io/evcc/core/keys"
	"github.com/evcc-io/evcc/server/db/settings"
	
)

func setExperimental(pub publisher) func(bool) error {
	return func(b bool) error {
		settings.SetBool(keys.Experimental, b)
		pub(keys.Experimental, b)
		return nil
	}
}

func getExperimental() bool {
	b, _ := settings.Bool(keys.Experimental)
	return b
}

