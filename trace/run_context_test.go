package trace

import (
	"testing"
	"context"
	"fmt"
)

func TestNewChildRunContext(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	fmt.Println(ctx.Value(runContextKey).(RunContext))
}
