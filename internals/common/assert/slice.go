package assert

import (
	"holistic/internals/common/logger"
	"slices"
)

func SliceContains[T comparable](slice []T, object T, message string, context ...any) {
	if !slices.Contains(slice, object) {
		logger.Panic(message, "object", object, "slice", slice, "context", context)
	}
}

func SliceNotContains[T comparable](slice []T, object T, message string, context ...any) {
	if slices.Contains(slice, object) {
		logger.Panic(message, "object", object, "slice", slice, "context", context)
	}
}
