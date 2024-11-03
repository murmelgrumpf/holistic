package assert

import (
	"log"
	"slices"
)

func SliceContains[T comparable](slice []T, object T, message string, context ...any) {
	if !slices.Contains(slice, object) {
		log.Panic(message, " | ", object, " > ", slice, " | ", context)
	}
}

func SliceNotContains[T comparable](slice []T, object T, message string, context ...any) {
	if slices.Contains(slice, object) {
		log.Panic(message, " | ", object, " > ", slice, " | ", context)
	}
}
