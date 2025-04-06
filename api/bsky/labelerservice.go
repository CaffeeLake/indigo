// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.labeler.service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	comatprototypes "github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/lex/util"
	cbg "github.com/whyrusleeping/cbor-gen"
)

func init() {
	util.RegisterType("app.bsky.labeler.service", &LabelerService{})
} //
// RECORDTYPE: LabelerService
type LabelerService struct {
	LexiconTypeID string                       `json:"$type,const=app.bsky.labeler.service" cborgen:"$type,const=app.bsky.labeler.service"`
	CreatedAt     string                       `json:"createdAt" cborgen:"createdAt"`
	Labels        *LabelerService_Labels       `json:"labels,omitempty" cborgen:"labels,omitempty"`
	Policies      *LabelerDefs_LabelerPolicies `json:"policies" cborgen:"policies"`
	// reasonTypes: The set of report reason 'codes' which are in-scope for this service to review and action. These usually align to policy categories. If not defined (distinct from empty array), all reason types are allowed.
	ReasonTypes []*string `json:"reasonTypes,omitempty" cborgen:"reasonTypes,omitempty"`
	// subjectCollections: Set of record types (collection NSIDs) which can be reported to this service. If not defined (distinct from empty array), default is any record type.
	SubjectCollections []string `json:"subjectCollections,omitempty" cborgen:"subjectCollections,omitempty"`
	// subjectTypes: The set of subject types (account, record, etc) this service accepts reports on.
	SubjectTypes []*string `json:"subjectTypes,omitempty" cborgen:"subjectTypes,omitempty"`
}

type LabelerService_Labels struct {
	LabelDefs_SelfLabels *comatprototypes.LabelDefs_SelfLabels
}

func (t *LabelerService_Labels) MarshalJSON() ([]byte, error) {
	if t.LabelDefs_SelfLabels != nil {
		t.LabelDefs_SelfLabels.LexiconTypeID = "com.atproto.label.defs#selfLabels"
		return json.Marshal(t.LabelDefs_SelfLabels)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *LabelerService_Labels) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "com.atproto.label.defs#selfLabels":
		t.LabelDefs_SelfLabels = new(comatprototypes.LabelDefs_SelfLabels)
		return json.Unmarshal(b, t.LabelDefs_SelfLabels)

	default:
		return nil
	}
}

func (t *LabelerService_Labels) MarshalCBOR(w io.Writer) error {

	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if t.LabelDefs_SelfLabels != nil {
		return t.LabelDefs_SelfLabels.MarshalCBOR(w)
	}
	return fmt.Errorf("cannot cbor marshal empty enum")
}
func (t *LabelerService_Labels) UnmarshalCBOR(r io.Reader) error {
	typ, b, err := util.CborTypeExtractReader(r)
	if err != nil {
		return err
	}

	switch typ {
	case "com.atproto.label.defs#selfLabels":
		t.LabelDefs_SelfLabels = new(comatprototypes.LabelDefs_SelfLabels)
		return t.LabelDefs_SelfLabels.UnmarshalCBOR(bytes.NewReader(b))

	default:
		return nil
	}
}
