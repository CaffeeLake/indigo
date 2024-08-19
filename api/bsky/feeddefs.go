// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.feed.defs

import (
	"encoding/json"
	"fmt"

	comatprototypes "github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/lex/util"
)

// FeedDefs_BlockedAuthor is a "blockedAuthor" in the app.bsky.feed.defs schema.
type FeedDefs_BlockedAuthor struct {
	Did    string                 `json:"did" cborgen:"did"`
	Viewer *ActorDefs_ViewerState `json:"viewer,omitempty" cborgen:"viewer,omitempty"`
}

// FeedDefs_BlockedPost is a "blockedPost" in the app.bsky.feed.defs schema.
//
// RECORDTYPE: FeedDefs_BlockedPost
type FeedDefs_BlockedPost struct {
	LexiconTypeID string                  `json:"$type,const=app.bsky.feed.defs#blockedPost" cborgen:"$type,const=app.bsky.feed.defs#blockedPost"`
	Author        *FeedDefs_BlockedAuthor `json:"author" cborgen:"author"`
	Blocked       bool                    `json:"blocked" cborgen:"blocked"`
	Uri           string                  `json:"uri" cborgen:"uri"`
}

// FeedDefs_FeedViewPost is a "feedViewPost" in the app.bsky.feed.defs schema.
type FeedDefs_FeedViewPost struct {
	// feedContext: Context provided by feed generator that may be passed back alongside interactions.
	FeedContext *string                       `json:"feedContext,omitempty" cborgen:"feedContext,omitempty"`
	Post        *FeedDefs_PostView            `json:"post" cborgen:"post"`
	Reason      *FeedDefs_FeedViewPost_Reason `json:"reason,omitempty" cborgen:"reason,omitempty"`
	Reply       *FeedDefs_ReplyRef            `json:"reply,omitempty" cborgen:"reply,omitempty"`
}

type FeedDefs_FeedViewPost_Reason struct {
	FeedDefs_ReasonRepost *FeedDefs_ReasonRepost
}

func (t *FeedDefs_FeedViewPost_Reason) MarshalJSON() ([]byte, error) {
	if t.FeedDefs_ReasonRepost != nil {
		t.FeedDefs_ReasonRepost.LexiconTypeID = "app.bsky.feed.defs#reasonRepost"
		return json.Marshal(t.FeedDefs_ReasonRepost)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *FeedDefs_FeedViewPost_Reason) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.feed.defs#reasonRepost":
		t.FeedDefs_ReasonRepost = new(FeedDefs_ReasonRepost)
		return json.Unmarshal(b, t.FeedDefs_ReasonRepost)

	default:
		return nil
	}
}

// FeedDefs_GeneratorView is a "generatorView" in the app.bsky.feed.defs schema.
//
// RECORDTYPE: FeedDefs_GeneratorView
type FeedDefs_GeneratorView struct {
	LexiconTypeID       string                             `json:"$type,const=app.bsky.feed.defs#generatorView" cborgen:"$type,const=app.bsky.feed.defs#generatorView"`
	AcceptsInteractions *bool                              `json:"acceptsInteractions,omitempty" cborgen:"acceptsInteractions,omitempty"`
	Avatar              *string                            `json:"avatar,omitempty" cborgen:"avatar,omitempty"`
	Cid                 string                             `json:"cid" cborgen:"cid"`
	Creator             *ActorDefs_ProfileView             `json:"creator" cborgen:"creator"`
	Description         *string                            `json:"description,omitempty" cborgen:"description,omitempty"`
	DescriptionFacets   []*RichtextFacet                   `json:"descriptionFacets,omitempty" cborgen:"descriptionFacets,omitempty"`
	Did                 string                             `json:"did" cborgen:"did"`
	DisplayName         string                             `json:"displayName" cborgen:"displayName"`
	IndexedAt           string                             `json:"indexedAt" cborgen:"indexedAt"`
	Labels              []*comatprototypes.LabelDefs_Label `json:"labels,omitempty" cborgen:"labels,omitempty"`
	LikeCount           *int64                             `json:"likeCount,omitempty" cborgen:"likeCount,omitempty"`
	Uri                 string                             `json:"uri" cborgen:"uri"`
	Viewer              *FeedDefs_GeneratorViewerState     `json:"viewer,omitempty" cborgen:"viewer,omitempty"`
}

// FeedDefs_GeneratorViewerState is a "generatorViewerState" in the app.bsky.feed.defs schema.
type FeedDefs_GeneratorViewerState struct {
	Like *string `json:"like,omitempty" cborgen:"like,omitempty"`
}

// FeedDefs_Interaction is a "interaction" in the app.bsky.feed.defs schema.
type FeedDefs_Interaction struct {
	Event *string `json:"event,omitempty" cborgen:"event,omitempty"`
	// feedContext: Context on a feed item that was originally supplied by the feed generator on getFeedSkeleton.
	FeedContext *string `json:"feedContext,omitempty" cborgen:"feedContext,omitempty"`
	Item        *string `json:"item,omitempty" cborgen:"item,omitempty"`
}

// FeedDefs_NotFoundPost is a "notFoundPost" in the app.bsky.feed.defs schema.
//
// RECORDTYPE: FeedDefs_NotFoundPost
type FeedDefs_NotFoundPost struct {
	LexiconTypeID string `json:"$type,const=app.bsky.feed.defs#notFoundPost" cborgen:"$type,const=app.bsky.feed.defs#notFoundPost"`
	NotFound      bool   `json:"notFound" cborgen:"notFound"`
	Uri           string `json:"uri" cborgen:"uri"`
}

// FeedDefs_PostView is a "postView" in the app.bsky.feed.defs schema.
//
// RECORDTYPE: FeedDefs_PostView
type FeedDefs_PostView struct {
	LexiconTypeID string                             `json:"$type,const=app.bsky.feed.defs#postView" cborgen:"$type,const=app.bsky.feed.defs#postView"`
	Author        *ActorDefs_ProfileViewBasic        `json:"author" cborgen:"author"`
	Cid           string                             `json:"cid" cborgen:"cid"`
	Embed         *FeedDefs_PostView_Embed           `json:"embed,omitempty" cborgen:"embed,omitempty"`
	IndexedAt     string                             `json:"indexedAt" cborgen:"indexedAt"`
	Labels        []*comatprototypes.LabelDefs_Label `json:"labels,omitempty" cborgen:"labels,omitempty"`
	LikeCount     *int64                             `json:"likeCount,omitempty" cborgen:"likeCount,omitempty"`
	Record        *util.LexiconTypeDecoder           `json:"record" cborgen:"record"`
	ReplyCount    *int64                             `json:"replyCount,omitempty" cborgen:"replyCount,omitempty"`
	RepostCount   *int64                             `json:"repostCount,omitempty" cborgen:"repostCount,omitempty"`
	Threadgate    *FeedDefs_ThreadgateView           `json:"threadgate,omitempty" cborgen:"threadgate,omitempty"`
	Uri           string                             `json:"uri" cborgen:"uri"`
	Viewer        *FeedDefs_ViewerState              `json:"viewer,omitempty" cborgen:"viewer,omitempty"`
}

type FeedDefs_PostView_Embed struct {
	EmbedImages_View          *EmbedImages_View
	EmbedVideo_View           *EmbedVideo_View
	EmbedExternal_View        *EmbedExternal_View
	EmbedRecord_View          *EmbedRecord_View
	EmbedRecordWithMedia_View *EmbedRecordWithMedia_View
}

func (t *FeedDefs_PostView_Embed) MarshalJSON() ([]byte, error) {
	if t.EmbedImages_View != nil {
		t.EmbedImages_View.LexiconTypeID = "app.bsky.embed.images#view"
		return json.Marshal(t.EmbedImages_View)
	}
	if t.EmbedVideo_View != nil {
		t.EmbedVideo_View.LexiconTypeID = "app.bsky.embed.video#view"
		return json.Marshal(t.EmbedVideo_View)
	}
	if t.EmbedExternal_View != nil {
		t.EmbedExternal_View.LexiconTypeID = "app.bsky.embed.external#view"
		return json.Marshal(t.EmbedExternal_View)
	}
	if t.EmbedRecord_View != nil {
		t.EmbedRecord_View.LexiconTypeID = "app.bsky.embed.record#view"
		return json.Marshal(t.EmbedRecord_View)
	}
	if t.EmbedRecordWithMedia_View != nil {
		t.EmbedRecordWithMedia_View.LexiconTypeID = "app.bsky.embed.recordWithMedia#view"
		return json.Marshal(t.EmbedRecordWithMedia_View)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *FeedDefs_PostView_Embed) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.embed.images#view":
		t.EmbedImages_View = new(EmbedImages_View)
		return json.Unmarshal(b, t.EmbedImages_View)
	case "app.bsky.embed.video#view":
		t.EmbedVideo_View = new(EmbedVideo_View)
		return json.Unmarshal(b, t.EmbedVideo_View)
	case "app.bsky.embed.external#view":
		t.EmbedExternal_View = new(EmbedExternal_View)
		return json.Unmarshal(b, t.EmbedExternal_View)
	case "app.bsky.embed.record#view":
		t.EmbedRecord_View = new(EmbedRecord_View)
		return json.Unmarshal(b, t.EmbedRecord_View)
	case "app.bsky.embed.recordWithMedia#view":
		t.EmbedRecordWithMedia_View = new(EmbedRecordWithMedia_View)
		return json.Unmarshal(b, t.EmbedRecordWithMedia_View)

	default:
		return nil
	}
}

// FeedDefs_ReasonRepost is a "reasonRepost" in the app.bsky.feed.defs schema.
//
// RECORDTYPE: FeedDefs_ReasonRepost
type FeedDefs_ReasonRepost struct {
	LexiconTypeID string                      `json:"$type,const=app.bsky.feed.defs#reasonRepost" cborgen:"$type,const=app.bsky.feed.defs#reasonRepost"`
	By            *ActorDefs_ProfileViewBasic `json:"by" cborgen:"by"`
	IndexedAt     string                      `json:"indexedAt" cborgen:"indexedAt"`
}

// FeedDefs_ReplyRef is a "replyRef" in the app.bsky.feed.defs schema.
type FeedDefs_ReplyRef struct {
	// grandparentAuthor: When parent is a reply to another post, this is the author of that post.
	GrandparentAuthor *ActorDefs_ProfileViewBasic `json:"grandparentAuthor,omitempty" cborgen:"grandparentAuthor,omitempty"`
	Parent            *FeedDefs_ReplyRef_Parent   `json:"parent" cborgen:"parent"`
	Root              *FeedDefs_ReplyRef_Root     `json:"root" cborgen:"root"`
}

type FeedDefs_ReplyRef_Parent struct {
	FeedDefs_PostView     *FeedDefs_PostView
	FeedDefs_NotFoundPost *FeedDefs_NotFoundPost
	FeedDefs_BlockedPost  *FeedDefs_BlockedPost
}

func (t *FeedDefs_ReplyRef_Parent) MarshalJSON() ([]byte, error) {
	if t.FeedDefs_PostView != nil {
		t.FeedDefs_PostView.LexiconTypeID = "app.bsky.feed.defs#postView"
		return json.Marshal(t.FeedDefs_PostView)
	}
	if t.FeedDefs_NotFoundPost != nil {
		t.FeedDefs_NotFoundPost.LexiconTypeID = "app.bsky.feed.defs#notFoundPost"
		return json.Marshal(t.FeedDefs_NotFoundPost)
	}
	if t.FeedDefs_BlockedPost != nil {
		t.FeedDefs_BlockedPost.LexiconTypeID = "app.bsky.feed.defs#blockedPost"
		return json.Marshal(t.FeedDefs_BlockedPost)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *FeedDefs_ReplyRef_Parent) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.feed.defs#postView":
		t.FeedDefs_PostView = new(FeedDefs_PostView)
		return json.Unmarshal(b, t.FeedDefs_PostView)
	case "app.bsky.feed.defs#notFoundPost":
		t.FeedDefs_NotFoundPost = new(FeedDefs_NotFoundPost)
		return json.Unmarshal(b, t.FeedDefs_NotFoundPost)
	case "app.bsky.feed.defs#blockedPost":
		t.FeedDefs_BlockedPost = new(FeedDefs_BlockedPost)
		return json.Unmarshal(b, t.FeedDefs_BlockedPost)

	default:
		return nil
	}
}

type FeedDefs_ReplyRef_Root struct {
	FeedDefs_PostView     *FeedDefs_PostView
	FeedDefs_NotFoundPost *FeedDefs_NotFoundPost
	FeedDefs_BlockedPost  *FeedDefs_BlockedPost
}

func (t *FeedDefs_ReplyRef_Root) MarshalJSON() ([]byte, error) {
	if t.FeedDefs_PostView != nil {
		t.FeedDefs_PostView.LexiconTypeID = "app.bsky.feed.defs#postView"
		return json.Marshal(t.FeedDefs_PostView)
	}
	if t.FeedDefs_NotFoundPost != nil {
		t.FeedDefs_NotFoundPost.LexiconTypeID = "app.bsky.feed.defs#notFoundPost"
		return json.Marshal(t.FeedDefs_NotFoundPost)
	}
	if t.FeedDefs_BlockedPost != nil {
		t.FeedDefs_BlockedPost.LexiconTypeID = "app.bsky.feed.defs#blockedPost"
		return json.Marshal(t.FeedDefs_BlockedPost)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *FeedDefs_ReplyRef_Root) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.feed.defs#postView":
		t.FeedDefs_PostView = new(FeedDefs_PostView)
		return json.Unmarshal(b, t.FeedDefs_PostView)
	case "app.bsky.feed.defs#notFoundPost":
		t.FeedDefs_NotFoundPost = new(FeedDefs_NotFoundPost)
		return json.Unmarshal(b, t.FeedDefs_NotFoundPost)
	case "app.bsky.feed.defs#blockedPost":
		t.FeedDefs_BlockedPost = new(FeedDefs_BlockedPost)
		return json.Unmarshal(b, t.FeedDefs_BlockedPost)

	default:
		return nil
	}
}

// FeedDefs_SkeletonFeedPost is a "skeletonFeedPost" in the app.bsky.feed.defs schema.
type FeedDefs_SkeletonFeedPost struct {
	// feedContext: Context that will be passed through to client and may be passed to feed generator back alongside interactions.
	FeedContext *string                           `json:"feedContext,omitempty" cborgen:"feedContext,omitempty"`
	Post        string                            `json:"post" cborgen:"post"`
	Reason      *FeedDefs_SkeletonFeedPost_Reason `json:"reason,omitempty" cborgen:"reason,omitempty"`
}

type FeedDefs_SkeletonFeedPost_Reason struct {
	FeedDefs_SkeletonReasonRepost *FeedDefs_SkeletonReasonRepost
}

func (t *FeedDefs_SkeletonFeedPost_Reason) MarshalJSON() ([]byte, error) {
	if t.FeedDefs_SkeletonReasonRepost != nil {
		t.FeedDefs_SkeletonReasonRepost.LexiconTypeID = "app.bsky.feed.defs#skeletonReasonRepost"
		return json.Marshal(t.FeedDefs_SkeletonReasonRepost)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *FeedDefs_SkeletonFeedPost_Reason) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.feed.defs#skeletonReasonRepost":
		t.FeedDefs_SkeletonReasonRepost = new(FeedDefs_SkeletonReasonRepost)
		return json.Unmarshal(b, t.FeedDefs_SkeletonReasonRepost)

	default:
		return nil
	}
}

// FeedDefs_SkeletonReasonRepost is a "skeletonReasonRepost" in the app.bsky.feed.defs schema.
//
// RECORDTYPE: FeedDefs_SkeletonReasonRepost
type FeedDefs_SkeletonReasonRepost struct {
	LexiconTypeID string `json:"$type,const=app.bsky.feed.defs#skeletonReasonRepost" cborgen:"$type,const=app.bsky.feed.defs#skeletonReasonRepost"`
	Repost        string `json:"repost" cborgen:"repost"`
}

// FeedDefs_ThreadViewPost is a "threadViewPost" in the app.bsky.feed.defs schema.
//
// RECORDTYPE: FeedDefs_ThreadViewPost
type FeedDefs_ThreadViewPost struct {
	LexiconTypeID string                                  `json:"$type,const=app.bsky.feed.defs#threadViewPost" cborgen:"$type,const=app.bsky.feed.defs#threadViewPost"`
	Parent        *FeedDefs_ThreadViewPost_Parent         `json:"parent,omitempty" cborgen:"parent,omitempty"`
	Post          *FeedDefs_PostView                      `json:"post" cborgen:"post"`
	Replies       []*FeedDefs_ThreadViewPost_Replies_Elem `json:"replies,omitempty" cborgen:"replies,omitempty"`
}

type FeedDefs_ThreadViewPost_Parent struct {
	FeedDefs_ThreadViewPost *FeedDefs_ThreadViewPost
	FeedDefs_NotFoundPost   *FeedDefs_NotFoundPost
	FeedDefs_BlockedPost    *FeedDefs_BlockedPost
}

func (t *FeedDefs_ThreadViewPost_Parent) MarshalJSON() ([]byte, error) {
	if t.FeedDefs_ThreadViewPost != nil {
		t.FeedDefs_ThreadViewPost.LexiconTypeID = "app.bsky.feed.defs#threadViewPost"
		return json.Marshal(t.FeedDefs_ThreadViewPost)
	}
	if t.FeedDefs_NotFoundPost != nil {
		t.FeedDefs_NotFoundPost.LexiconTypeID = "app.bsky.feed.defs#notFoundPost"
		return json.Marshal(t.FeedDefs_NotFoundPost)
	}
	if t.FeedDefs_BlockedPost != nil {
		t.FeedDefs_BlockedPost.LexiconTypeID = "app.bsky.feed.defs#blockedPost"
		return json.Marshal(t.FeedDefs_BlockedPost)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *FeedDefs_ThreadViewPost_Parent) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.feed.defs#threadViewPost":
		t.FeedDefs_ThreadViewPost = new(FeedDefs_ThreadViewPost)
		return json.Unmarshal(b, t.FeedDefs_ThreadViewPost)
	case "app.bsky.feed.defs#notFoundPost":
		t.FeedDefs_NotFoundPost = new(FeedDefs_NotFoundPost)
		return json.Unmarshal(b, t.FeedDefs_NotFoundPost)
	case "app.bsky.feed.defs#blockedPost":
		t.FeedDefs_BlockedPost = new(FeedDefs_BlockedPost)
		return json.Unmarshal(b, t.FeedDefs_BlockedPost)

	default:
		return nil
	}
}

type FeedDefs_ThreadViewPost_Replies_Elem struct {
	FeedDefs_ThreadViewPost *FeedDefs_ThreadViewPost
	FeedDefs_NotFoundPost   *FeedDefs_NotFoundPost
	FeedDefs_BlockedPost    *FeedDefs_BlockedPost
}

func (t *FeedDefs_ThreadViewPost_Replies_Elem) MarshalJSON() ([]byte, error) {
	if t.FeedDefs_ThreadViewPost != nil {
		t.FeedDefs_ThreadViewPost.LexiconTypeID = "app.bsky.feed.defs#threadViewPost"
		return json.Marshal(t.FeedDefs_ThreadViewPost)
	}
	if t.FeedDefs_NotFoundPost != nil {
		t.FeedDefs_NotFoundPost.LexiconTypeID = "app.bsky.feed.defs#notFoundPost"
		return json.Marshal(t.FeedDefs_NotFoundPost)
	}
	if t.FeedDefs_BlockedPost != nil {
		t.FeedDefs_BlockedPost.LexiconTypeID = "app.bsky.feed.defs#blockedPost"
		return json.Marshal(t.FeedDefs_BlockedPost)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *FeedDefs_ThreadViewPost_Replies_Elem) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.feed.defs#threadViewPost":
		t.FeedDefs_ThreadViewPost = new(FeedDefs_ThreadViewPost)
		return json.Unmarshal(b, t.FeedDefs_ThreadViewPost)
	case "app.bsky.feed.defs#notFoundPost":
		t.FeedDefs_NotFoundPost = new(FeedDefs_NotFoundPost)
		return json.Unmarshal(b, t.FeedDefs_NotFoundPost)
	case "app.bsky.feed.defs#blockedPost":
		t.FeedDefs_BlockedPost = new(FeedDefs_BlockedPost)
		return json.Unmarshal(b, t.FeedDefs_BlockedPost)

	default:
		return nil
	}
}

// FeedDefs_ThreadgateView is a "threadgateView" in the app.bsky.feed.defs schema.
type FeedDefs_ThreadgateView struct {
	Cid    *string                    `json:"cid,omitempty" cborgen:"cid,omitempty"`
	Lists  []*GraphDefs_ListViewBasic `json:"lists,omitempty" cborgen:"lists,omitempty"`
	Record *util.LexiconTypeDecoder   `json:"record,omitempty" cborgen:"record,omitempty"`
	Uri    *string                    `json:"uri,omitempty" cborgen:"uri,omitempty"`
}

// FeedDefs_ViewerState is a "viewerState" in the app.bsky.feed.defs schema.
//
// Metadata about the requesting account's relationship with the subject content. Only has meaningful content for authed requests.
type FeedDefs_ViewerState struct {
	EmbeddingDisabled *bool   `json:"embeddingDisabled,omitempty" cborgen:"embeddingDisabled,omitempty"`
	Like              *string `json:"like,omitempty" cborgen:"like,omitempty"`
	ReplyDisabled     *bool   `json:"replyDisabled,omitempty" cborgen:"replyDisabled,omitempty"`
	Repost            *string `json:"repost,omitempty" cborgen:"repost,omitempty"`
	ThreadMuted       *bool   `json:"threadMuted,omitempty" cborgen:"threadMuted,omitempty"`
}
