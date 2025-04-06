// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package chat

// schema: chat.bsky.convo.removeReaction

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// ConvoRemoveReaction_Input is the input argument to a chat.bsky.convo.removeReaction call.
type ConvoRemoveReaction_Input struct {
	ConvoId   string `json:"convoId" cborgen:"convoId"`
	MessageId string `json:"messageId" cborgen:"messageId"`
	Value     string `json:"value" cborgen:"value"`
}

// ConvoRemoveReaction_Output is the output of a chat.bsky.convo.removeReaction call.
type ConvoRemoveReaction_Output struct {
	Message *ConvoDefs_MessageView `json:"message" cborgen:"message"`
}

// ConvoRemoveReaction calls the XRPC method "chat.bsky.convo.removeReaction".
func ConvoRemoveReaction(ctx context.Context, c *xrpc.Client, input *ConvoRemoveReaction_Input) (*ConvoRemoveReaction_Output, error) {
	var out ConvoRemoveReaction_Output
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "chat.bsky.convo.removeReaction", nil, input, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
