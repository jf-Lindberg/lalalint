package configMgmt

import (
	"github.com/jf-Lindberg/lalalint/helper"
	"github.com/spf13/viper"
	"strconv"
)

func GetConfig() map[string]interface{} {
	return viper.AllSettings()
}

func GetVal(key string) interface{} {
	if res := viper.Get(key); res != nil {
		return viper.Get(key)
	}
	return "Error not found"
}

func SetVal(key, val string) {
	if res := GetVal(key); res != "Value not found" {
		if val == "true" || val == "false" {
			bool, err := strconv.ParseBool(val)
			if err != nil {
				helper.LogFatal(err)
			}
			viper.Set(key, bool)
		} else {
			viper.Set(key, val)
		}
		viper.WriteConfig()
		return
	}

}
