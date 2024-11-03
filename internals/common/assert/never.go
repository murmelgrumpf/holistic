package assert

import (
	"holistic/internals/common/logger"
)

func Never(message string, context ...any) {
	logger.Panic(message, "context", context)
}
