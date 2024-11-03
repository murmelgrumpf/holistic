package assert

import "log"

func Nil(object any, message string, context ...any) {
	if object != nil {
		log.Panic(message, " | ", object, " | ", context)
	}
}

func NotNil(object any, message string, context ...any) {
	if object == nil {
		log.Panic(message, " | ", object, " | ", context)
	}
}
