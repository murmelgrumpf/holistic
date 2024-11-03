package assert

import (
	"holistic/internals/common/logger"
	"strings"
)

func StringHasPrefix(object string, startsWith string, message string, context ...any) {
	if !strings.HasPrefix(object, startsWith) {
		logger.Panic(message, "startsWith", startsWith, "object", object, "context", context)
	}
}
