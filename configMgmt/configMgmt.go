package configMgmt

import (
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

func SetVal(key, val string) error {
	if res := GetVal(key); res != "Value not found" {
		if val == "true" || val == "false" {
			parseBool, err := strconv.ParseBool(val)
			if err != nil {
				return err
			}
			viper.Set(key, parseBool)
		} else {
			viper.Set(key, val)
		}
		err := viper.WriteConfig()
		if err != nil {
			return err
		}
	}
	return nil
}
