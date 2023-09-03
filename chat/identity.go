// This file was auto-generated by Fern from our API Definition.

package chat

import (
	uuid "github.com/gofrs/uuid/v5"
	identity "github.com/rivet-gg/rivet-go/identity"
)

type GetDirectThreadResponse struct {
	ThreadId *uuid.UUID       `json:"thread_id,omitempty"`
	Identity *identity.Handle `json:"identity,omitempty"`
}
