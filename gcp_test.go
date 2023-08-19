package sloggcp_test

import (
	sloggcp "github.com/vlad-tokarev/slog-GCP"
	"log/slog"
	"os"
	"testing"
	"time"
)

var (
	someTime   = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	someSource = slog.Source{
		Function: "test",
		File:     "test.go",
		Line:     1,
	}
)

// ExampleReplaceAttr shows how to replace default slog attributes with GCP compatible ones
// It writes to stderr, that is why output is empty
func ExampleReplaceAttr() {

	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		ReplaceAttr: sloggcp.ReplaceAttr,
		AddSource:   true,
		Level:       slog.LevelDebug,
	}))

	logger.Debug("test",
		slog.String("test", "test"),
	)

	// Output:

}

func TestReplaceAttr(t *testing.T) {

	type args struct {
		groups []string
		a      slog.Attr
	}
	tests := []struct {
		name string
		args args
		want slog.Attr
	}{
		{
			name: "TimeKey",
			args: args{
				groups: []string{},
				a:      slog.Time(slog.TimeKey, someTime),
			},
			want: slog.Time("time", someTime),
		},
		{
			name: "LevelKey",
			args: args{
				groups: []string{},
				a:      slog.Any(slog.LevelKey, slog.LevelDebug),
			},
			want: slog.String("severity", "DEBUG"),
		},
		{
			name: "SourceKey",
			args: args{
				groups: []string{},
				a:      slog.Any(slog.SourceKey, &someSource),
			},
			want: slog.Any("logging.googleapis.com/sourceLocation", &someSource),
		},
		{
			name: "MessageKey",
			args: args{
				groups: []string{},
				a:      slog.String(slog.MessageKey, "test"),
			},
			want: slog.String("message", "test"),
		},
		{
			name: "OtherKey",
			args: args{
				groups: []string{},
				a:      slog.String("test", "test"),
			},
			want: slog.String("test", "test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sloggcp.ReplaceAttr(tt.args.groups, tt.args.a); !got.Equal(tt.want) {
				t.Errorf("ReplaceAttr() = %v, want %v", got, tt.want)
			}
		})
	}
}
