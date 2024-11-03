package assert

import (
	"holistic/internals/common/logger"
)

func Nil(object any, message string, context ...any) {
	if object != nil {
		logger.Panic(message, "object", object, "context", context)
	}
}

func NotNil(object any, message string, context ...any) {
	if object == nil {
		logger.Panic(message, "object", object, "context", context)
	}
}
