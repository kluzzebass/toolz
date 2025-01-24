package toolz

import (
	"context"
	"time"
)

func Sleep(ctx context.Context, delay time.Duration) {
	select {
	case <-ctx.Done():
	case <-time.After(delay):
	}
}
