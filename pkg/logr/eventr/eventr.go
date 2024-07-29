package eventr

import (
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
)

const MaxLoggingLevel = 0

// A logr implementation generating K8s events
type eventr struct {
	recorder record.EventRecorder
	subject runtime.Object
	// keysAndValue []any
}

func NewEventr(recorder record.EventRecorder, subject runtime.Object) logr.Logger {
	return logr.Logger{}.WithSink(&eventr{recorder: recorder, subject: subject})
}

func (r *eventr) Init(info logr.RuntimeInfo) {}

func (r *eventr) Enabled(level int) bool {
	return level <= MaxLoggingLevel
}

func (r *eventr) Info(level int, msg string, keysAndValues ...any) {
	r.recorder.Event(r.subject, "Normal", "Info", msg)
}

func (r *eventr) Error(err error, msg string, keysAndValues ...any) {

}

func (r *eventr) WithValues(keysAndValues ...any) logr.LogSink {
	return r
}

func (r *eventr) WithName(name string) logr.LogSink {
	return r
}
