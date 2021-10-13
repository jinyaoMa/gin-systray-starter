package tray

import (
	"fmt"
	"this/tray/locales"
)

const defaultMessage = "???"

var locale = "en"

func setLocale(name string) (ok bool) {
	switch name {
	case "en", "zh":
		if locale != name {
			locale = name
			return true
		}
	}
	return false
}

func getMessage(key string) string {
	switch locale {
	case "en":
		return locales.En[key]
	case "zh":
		return locales.Zh[key]
	}
	return defaultMessage
}

func getMessageWithParams(key string, params ...interface{}) string {
	switch locale {
	case "en":
		return fmt.Sprintf(locales.En[key], params...)
	case "zh":
		return fmt.Sprintf(locales.Zh[key], params...)
	}
	return defaultMessage
}
