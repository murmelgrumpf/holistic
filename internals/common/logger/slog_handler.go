package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"holistic/internals/common/logging"
	"log/slog"
	"sync"
)

const (
	timeFormat = "[15:04:05]"

	debugPrefix = logging.Purple + logging.Bold + "Debug: " + logging.Reset
	errorPrefix = logging.Red + logging.Bold + "Error: " + logging.Reset
	warnPrefix  = logging.Yellow + logging.Bold + "Warn: " + logging.Reset
	infoPrefix  = logging.Cyan + logging.Bold + "Info: " + logging.Reset
)

var handler = slog.New(newHandler())

var Error = handler.Error
var Info = handler.Info
var Debug = handler.Debug
var Warn = handler.Warn

func Panic(msg string, args ...any) {
	Error(msg, args...)
	panic(msg)
}

type slogHandler struct {
	handler slog.Handler
	buffer  *bytes.Buffer
	mutex   *sync.Mutex
}

func (handler *slogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return handler.handler.Enabled(ctx, level)
}

func (handler *slogHandler) WithAttrs(attributes []slog.Attr) slog.Handler {
	return &slogHandler{handler: handler.handler.WithAttrs(attributes), buffer: handler.buffer, mutex: handler.mutex}
}

func (handler *slogHandler) WithGroup(name string) slog.Handler {
	return &slogHandler{handler: handler.handler.WithGroup(name), buffer: handler.buffer, mutex: handler.mutex}
}

func (handler *slogHandler) Handle(ctx context.Context, r slog.Record) error {

	var prefix string
	switch r.Level {
	case slog.LevelDebug:
		prefix = debugPrefix
	case slog.LevelInfo:
		prefix = infoPrefix
	case slog.LevelWarn:
		prefix = warnPrefix
	case slog.LevelError:
		prefix = errorPrefix
	}

	handler.mutex.Lock()
	defer func() {
		handler.buffer.Reset()
		handler.mutex.Unlock()
	}()

	if r.NumAttrs() == 0 {
		fmt.Println(
			logging.LightGray + r.Time.Format(timeFormat) +
				prefix +
				logging.Reset + r.Message,
		)
		return nil
	}

	if err := handler.handler.Handle(ctx, r); err != nil {
		logging.Delimiter()
		fmt.Printf(logging.Red, "Error when calling slog handle", logging.Reset)
		logging.Delimiter()
	}

	var pretty bytes.Buffer
	if err := json.Indent(&pretty, handler.buffer.Bytes(), "", "\t"); err != nil {
		logging.Delimiter()
		fmt.Printf(logging.Red, "Error when trying to Indent slog attributes", logging.Reset)
		logging.Delimiter()
	}

	fmt.Print(
		logging.LightGray + r.Time.Format(timeFormat) +
			prefix +
			logging.White + r.Message + " " +
			logging.LightGreen + string(pretty.Bytes()) + logging.Reset,
	)

	return nil
}

func suppressDefaults(
	next func([]string, slog.Attr) slog.Attr,
) func([]string, slog.Attr) slog.Attr {
	return func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey ||
			a.Key == slog.LevelKey ||
			a.Key == slog.MessageKey {
			return slog.Attr{}
		}
		if next == nil {
			return a
		}
		return next(groups, a)
	}
}

func newHandler() *slogHandler {
	opts := &slog.HandlerOptions{}
	buffer := &bytes.Buffer{}
	return &slogHandler{
		buffer: buffer,
		handler: slog.NewJSONHandler(buffer, &slog.HandlerOptions{
			Level:       opts.Level,
			AddSource:   opts.AddSource,
			ReplaceAttr: suppressDefaults(opts.ReplaceAttr),
		}),
		mutex: &sync.Mutex{},
	}
}
