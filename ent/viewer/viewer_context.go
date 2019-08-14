package viewer

import (
	"context"
	"fmt"
	"net/http"
)

type contextKey struct {
	name string
}

var viewerCtxKey = &contextKey{"viewer"}

func ForContext(ctx context.Context) (ViewerContext, error) {
	val, ok := ctx.Value(viewerCtxKey).(ViewerContext)
	if !ok {
		return nil, fmt.Errorf("invalid viewer passed to context")
	}
	return val, nil
}

func NewRequestWithContext(r *http.Request, viewer ViewerContext) *http.Request {
	ctx := context.WithValue(r.Context(), viewerCtxKey, viewer)
	return r.WithContext(ctx)
}

// ViewerContext interface is to be implemented by clients to indicate
// who's trying to view the ent
type ViewerContext interface {

	// returns the logged in user. should return nil if logged out
	// Can't have a generic User object to return because Go is hard

	// TODO rename to LoggedInEntity() that returns Entity instead of GetUser() that returns interface{}
	// would need to reorganize a few things
	GetUser() interface{}

	// returns the ID of the logged in viewer
	// LoggedInViewerID()?
	GetViewerID() string

	// Boolean indicating that the viewer is logged in or not
	// TODO kill this from interface. imply it from viewerID
	HasIdentity() bool // this is implied from above

	// Returns a boolean indicating if this user can see everything
	// Needed for places without actual privacy checks or for admin tools
	// or other such things
	// rename to CanSeeAllContent()? OverridePrivacyChecks()?
	// Or just have a second interface that's not necessary to mandate and do a typecheck on that
	// ViewerContextWhichOverridesPrivacy
	IsOmniscient() bool
}

// LoggedOutViewerContext struct is the default ViewerContext provided by ent framework
type LoggedOutViewerContext struct{}

// GetUser returns the Logged in User. For the LoggedOutViewerContext, this is nil
// TODO make the pointer implement the interface instead
func (viewer LoggedOutViewerContext) GetUser() interface{} {
	return nil
}

// GetViewerID returns the ID of the logged in viewer
func (viewer LoggedOutViewerContext) GetViewerID() string {
	return ""
}

// HasIdentity returns a boolean indicating that there's no identity associated with this viewer
func (viewer LoggedOutViewerContext) HasIdentity() bool {
	return false
}

// IsOmniscient returns a boolean indicating that the LoggedOutViewerContext cannot see most things by default
func (viewer LoggedOutViewerContext) IsOmniscient() bool {
	return false
}

// LoggedOutViewer returns an instance of LoggedOutViewerContext as the default Viewer in the ent framework
func LoggedOutViewer() ViewerContext {
	return LoggedOutViewerContext{}
}
