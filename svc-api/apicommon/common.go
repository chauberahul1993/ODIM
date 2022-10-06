package apicommon

import (
	"github.com/ODIM-Project/ODIM/lib-utilities/common"
	"github.com/ODIM-Project/ODIM/lib-utilities/config"
	l "github.com/ODIM-Project/ODIM/lib-utilities/logs"
)

var (
	// ConfigFilePath holds the value of odim config file path
	ConfigFilePath string
)

func TrackConfigFileChanges() {
	eventChan := make(chan interface{})
	go common.TrackConfigFileChanges(ConfigFilePath, eventChan)
	for {
		l.Log.Info(<-eventChan) // new data arrives through eventChan channel
		l.Log.Logger.Level = config.Data.LogLevel
	}
}
