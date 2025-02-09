// This file was auto-generated by Fern from our API Definition.

package group

import (
	uuid "github.com/gofrs/uuid/v5"
)

type ConsumeInviteResponse struct {
	GroupId *uuid.UUID `json:"group_id,omitempty"`
}

type CreateInviteRequest struct {
	// How long until the group invite expires (in milliseconds).
	Ttl *float64 `json:"ttl,omitempty"`
	// How many times the group invite can be used.
	UseCount *float64 `json:"use_count,omitempty"`
}

type CreateInviteResponse struct {
	// The code that will be passed to `rivet.api.group#ConsumeInvite` to join a group.
	Code string `json:"code"`
}

type GetInviteResponse struct {
	Group *Handle `json:"group,omitempty"`
}
