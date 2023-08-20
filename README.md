# sloggcp

`sloggcp` provides a simple implementation of the `ReplaceAttr` 
function for `JSONHandler` from [slog](https://pkg.go.dev/log/slog@master).

This implementation adapts the default slog attributes to be compatible 
with [Google Cloud Platform's Structured Logging](https://cloud.google.com/logging/docs/structured-logging).

By using `sloggcp`, you can ensure correct log severity display in the GCP Logs Viewer 
and proper representation of other attributes.

It is applicable to any GCP service that delivers logs via [Logging agent](https://cloud.google.com/logging/docs/agent/logging/configuration#special-fields).

For example, CloudRun, etc.


## Usage

### Override default attributes

```go
package main

import (
	"github.com/vlad-tokarev/sloggcp"
	"log/slog"
	"os"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		ReplaceAttr: sloggcp.ReplaceAttr,
		AddSource:   true,
		Level:       slog.LevelDebug,
	}))
	slog.SetDefault(logger)
}

```

