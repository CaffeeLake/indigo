// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.server.updateEmail

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// ServerUpdateEmail_Input is the input argument to a com.atproto.server.updateEmail call.
type ServerUpdateEmail_Input struct {
	Email string `json:"email" cborgen:"email"`
	// token: Requires a token from com.atproto.sever.requestEmailUpdate if the account's email has been confirmed.
	Token *string `json:"token,omitempty" cborgen:"token,omitempty"`
}

// ServerUpdateEmail calls the XRPC method "com.atproto.server.updateEmail".
func ServerUpdateEmail(ctx context.Context, c *xrpc.Client, input *ServerUpdateEmail_Input) error {
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "com.atproto.server.updateEmail", nil, input, nil); err != nil {
		return err
	}

	return nil
}