package public

import "context"

// UserReader is a narrow contract for other modules.
// It avoids leaking users' domain model or persistence concerns.
type UserReader interface {
	// Exists returns true if a user exists.
	Exists(ctx context.Context, userID string) (bool, error)
}
