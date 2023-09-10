package steam

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/Philipp15b/go-steam/v3/netutil"
)

func TestConnectToContext(t *testing.T) {

	c := NewClient()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for event := range c.Events() {
			switch e := event.(type) {

			case nil:
				return

			default:
				t.Logf("\t\tReceived event of type %T", e)
			}
		}
	}()

	// Times out on dial.
	badServer := "162.254.198.104:27018"

	// Should work.
	okServer := "155.133.246.52:27018"

	// Fail 5 times with bad server to test for deadlocks.
	for i := 0; i < 5; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)

		t.Log("Testing [ConnectToContext] on bad server", badServer)
		err := c.ConnectToContext(
			ctx,
			netutil.ParsePortAddr(badServer),
		)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				t.Logf("\t[ConnectToContext] timed out as expected: %v", err)
			} else {
				t.Fatalf("\t[ConnectToContext] failed: %v", err)
			}
		}
		cancel()
	}

	// Connect to OK server.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	t.Log("Testing [ConnectToContext] on OK server", okServer)
	err := c.ConnectToContext(
		ctx,
		netutil.ParsePortAddr(okServer),
	)
	if err != nil {
		t.Fatalf("\t[ConnectToContext] failed (is the server still OK?): %v", err)
	}

	t.Logf("\t[ConnectToContext] connected without issues")

	time.Sleep(time.Second)
	c.Emit(nil) // Stop the event loop.

	wg.Wait()
}
