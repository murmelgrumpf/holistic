package assert

import "log"

func Never(message string, context ...any) {
	log.Panic(message, " | ", context)
}
