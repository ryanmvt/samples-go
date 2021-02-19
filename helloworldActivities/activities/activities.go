package activities

import (
	"context"

	"go.temporal.io/sdk/activity"
)

// HelloActivity disconnected activity run test
func HelloActivity(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "name", name)
	return "Hello " + name + "!", nil
}
