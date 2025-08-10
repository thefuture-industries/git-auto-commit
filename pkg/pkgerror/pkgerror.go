package pkgerror

import (
	"context"
	"errors"
	"net"
)

func CreateError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, context.DeadlineExceeded) {
		return Err_ContextDeadlineExceeded
	}

	if errors.Is(err, context.Canceled) {
		return Err_ContextCanceled
	}

	var netErr net.Error
	if errors.As(err, &netErr) {
		if netErr.Timeout() {
			return Err_Timeout
		}

		return Err_Network
	}

	// write to log and return log for user
	return err
}
