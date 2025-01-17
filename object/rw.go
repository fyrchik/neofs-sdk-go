package object

import (
	"github.com/nspcc-dev/neofs-api-go/v2/object"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-sdk-go/checksum"
	cid "github.com/nspcc-dev/neofs-sdk-go/container/id"
	"github.com/nspcc-dev/neofs-sdk-go/owner"
	"github.com/nspcc-dev/neofs-sdk-go/session"
	"github.com/nspcc-dev/neofs-sdk-go/signature"
	"github.com/nspcc-dev/neofs-sdk-go/version"
)

// wrapper over v2 Object that provides
// public getter and private setters.
type rwObject object.Object

// ToV2 converts Object to v2 Object message.
func (o *rwObject) ToV2() *object.Object {
	return (*object.Object)(o)
}

func (o *rwObject) setHeaderField(setter func(*object.Header)) {
	obj := (*object.Object)(o)
	h := obj.GetHeader()

	if h == nil {
		h = new(object.Header)
		obj.SetHeader(h)
	}

	setter(h)
}

func (o *rwObject) setSplitFields(setter func(*object.SplitHeader)) {
	o.setHeaderField(func(h *object.Header) {
		split := h.GetSplit()
		if split == nil {
			split = new(object.SplitHeader)
			h.SetSplit(split)
		}

		setter(split)
	})
}

// ID returns object identifier.
func (o *rwObject) ID() *ID {
	return NewIDFromV2(
		(*object.Object)(o).
			GetObjectID(),
	)
}

func (o *rwObject) setID(v *ID) {
	(*object.Object)(o).
		SetObjectID(v.ToV2())
}

// Signature returns signature of the object identifier.
func (o *rwObject) Signature() *signature.Signature {
	return signature.NewFromV2(
		(*object.Object)(o).GetSignature())
}

func (o *rwObject) setSignature(v *signature.Signature) {
	(*object.Object)(o).SetSignature(v.ToV2())
}

// Payload returns payload bytes.
func (o *rwObject) Payload() []byte {
	return (*object.Object)(o).GetPayload()
}

func (o *rwObject) setPayload(v []byte) {
	(*object.Object)(o).SetPayload(v)
}

// Version returns version of the object.
func (o *rwObject) Version() *version.Version {
	return version.NewFromV2(
		(*object.Object)(o).
			GetHeader().
			GetVersion(),
	)
}

func (o *rwObject) setVersion(v *version.Version) {
	o.setHeaderField(func(h *object.Header) {
		h.SetVersion(v.ToV2())
	})
}

// PayloadSize returns payload length of the object.
func (o *rwObject) PayloadSize() uint64 {
	return (*object.Object)(o).
		GetHeader().
		GetPayloadLength()
}

func (o *rwObject) setPayloadSize(v uint64) {
	o.setHeaderField(func(h *object.Header) {
		h.SetPayloadLength(v)
	})
}

// ContainerID returns identifier of the related container.
func (o *rwObject) ContainerID() *cid.ID {
	return cid.NewFromV2(
		(*object.Object)(o).
			GetHeader().
			GetContainerID(),
	)
}

func (o *rwObject) setContainerID(v *cid.ID) {
	o.setHeaderField(func(h *object.Header) {
		h.SetContainerID(v.ToV2())
	})
}

// OwnerID returns identifier of the object owner.
func (o *rwObject) OwnerID() *owner.ID {
	return owner.NewIDFromV2(
		(*object.Object)(o).
			GetHeader().
			GetOwnerID(),
	)
}

func (o *rwObject) setOwnerID(v *owner.ID) {
	o.setHeaderField(func(h *object.Header) {
		h.SetOwnerID(v.ToV2())
	})
}

// CreationEpoch returns epoch number in which object was created.
func (o *rwObject) CreationEpoch() uint64 {
	return (*object.Object)(o).
		GetHeader().
		GetCreationEpoch()
}

func (o *rwObject) setCreationEpoch(v uint64) {
	o.setHeaderField(func(h *object.Header) {
		h.SetCreationEpoch(v)
	})
}

// PayloadChecksum returns checksum of the object payload.
func (o *rwObject) PayloadChecksum() *checksum.Checksum {
	return checksum.NewFromV2(
		(*object.Object)(o).
			GetHeader().
			GetPayloadHash(),
	)
}

func (o *rwObject) setPayloadChecksum(v *checksum.Checksum) {
	o.setHeaderField(func(h *object.Header) {
		h.SetPayloadHash(v.ToV2())
	})
}

// PayloadHomomorphicHash returns homomorphic hash of the object payload.
func (o *rwObject) PayloadHomomorphicHash() *checksum.Checksum {
	return checksum.NewFromV2(
		(*object.Object)(o).
			GetHeader().
			GetHomomorphicHash(),
	)
}

func (o *rwObject) setPayloadHomomorphicHash(v *checksum.Checksum) {
	o.setHeaderField(func(h *object.Header) {
		h.SetHomomorphicHash(v.ToV2())
	})
}

// Attributes returns object attributes.
func (o *rwObject) Attributes() []*Attribute {
	attrs := (*object.Object)(o).
		GetHeader().
		GetAttributes()

	res := make([]*Attribute, 0, len(attrs))

	for i := range attrs {
		res = append(res, NewAttributeFromV2(attrs[i]))
	}

	return res
}

func (o *rwObject) setAttributes(v ...*Attribute) {
	attrs := make([]*object.Attribute, 0, len(v))

	for i := range v {
		attrs = append(attrs, v[i].ToV2())
	}

	o.setHeaderField(func(h *object.Header) {
		h.SetAttributes(attrs)
	})
}

// PreviousID returns identifier of the previous sibling object.
func (o *rwObject) PreviousID() *ID {
	return NewIDFromV2(
		(*object.Object)(o).
			GetHeader().
			GetSplit().
			GetPrevious(),
	)
}

func (o *rwObject) setPreviousID(v *ID) {
	o.setSplitFields(func(split *object.SplitHeader) {
		split.SetPrevious(v.ToV2())
	})
}

// Children return list of the identifiers of the child objects.
func (o *rwObject) Children() []*ID {
	ids := (*object.Object)(o).
		GetHeader().
		GetSplit().
		GetChildren()

	res := make([]*ID, 0, len(ids))

	for i := range ids {
		res = append(res, NewIDFromV2(ids[i]))
	}

	return res
}

func (o *rwObject) setChildren(v ...*ID) {
	ids := make([]*refs.ObjectID, 0, len(v))

	for i := range v {
		ids = append(ids, v[i].ToV2())
	}

	o.setSplitFields(func(split *object.SplitHeader) {
		split.SetChildren(ids)
	})
}

// SplitID return split identity of split object. If object is not split
// returns nil.
func (o *rwObject) SplitID() *SplitID {
	return NewSplitIDFromV2(
		(*object.Object)(o).
			GetHeader().
			GetSplit().
			GetSplitID(),
	)
}

func (o *rwObject) setSplitID(id *SplitID) {
	o.setSplitFields(func(split *object.SplitHeader) {
		split.SetSplitID(id.ToV2())
	})
}

// ParentID returns identifier of the parent object.
func (o *rwObject) ParentID() *ID {
	return NewIDFromV2(
		(*object.Object)(o).
			GetHeader().
			GetSplit().
			GetParent(),
	)
}

func (o *rwObject) setParentID(v *ID) {
	o.setSplitFields(func(split *object.SplitHeader) {
		split.SetParent(v.ToV2())
	})
}

// Parent returns parent object w/o payload.
func (o *rwObject) Parent() *Object {
	h := (*object.Object)(o).
		GetHeader().
		GetSplit()

	parSig := h.GetParentSignature()
	parHdr := h.GetParentHeader()

	if parSig == nil && parHdr == nil {
		return nil
	}

	oV2 := new(object.Object)
	oV2.SetObjectID(h.GetParent())
	oV2.SetSignature(parSig)
	oV2.SetHeader(parHdr)

	return NewFromV2(oV2)
}

func (o *rwObject) setParent(v *Object) {
	o.setSplitFields(func(split *object.SplitHeader) {
		split.SetParent((*object.Object)(v.rwObject).GetObjectID())
		split.SetParentSignature((*object.Object)(v.rwObject).GetSignature())
		split.SetParentHeader((*object.Object)(v.rwObject).GetHeader())
	})
}

func (o *rwObject) initRelations() {
	o.setHeaderField(func(h *object.Header) {
		h.SetSplit(new(object.SplitHeader))
	})
}

func (o *rwObject) resetRelations() {
	o.setHeaderField(func(h *object.Header) {
		h.SetSplit(nil)
	})
}

// SessionToken returns token of the session
// within which object was created.
func (o *rwObject) SessionToken() *session.Token {
	return session.NewTokenFromV2(
		(*object.Object)(o).
			GetHeader().
			GetSessionToken(),
	)
}

func (o *rwObject) setSessionToken(v *session.Token) {
	o.setHeaderField(func(h *object.Header) {
		h.SetSessionToken(v.ToV2())
	})
}

// Type returns type of the object.
func (o *rwObject) Type() Type {
	return TypeFromV2(
		(*object.Object)(o).
			GetHeader().
			GetObjectType(),
	)
}

func (o *rwObject) setType(t Type) {
	o.setHeaderField(func(h *object.Header) {
		h.SetObjectType(t.ToV2())
	})
}

func (o *rwObject) cutPayload() *rwObject {
	ov2 := new(object.Object)
	*ov2 = *(*object.Object)(o)
	ov2.SetPayload(nil)

	return (*rwObject)(ov2)
}

func (o *rwObject) HasParent() bool {
	return (*object.Object)(o).
		GetHeader().
		GetSplit() != nil
}

// Marshal marshals object into a protobuf binary form.
func (o *rwObject) Marshal() ([]byte, error) {
	return (*object.Object)(o).StableMarshal(nil)
}

// Unmarshal unmarshals protobuf binary representation of object.
func (o *rwObject) Unmarshal(data []byte) error {
	return (*object.Object)(o).Unmarshal(data)
}

// MarshalJSON encodes object to protobuf JSON format.
func (o *rwObject) MarshalJSON() ([]byte, error) {
	return (*object.Object)(o).MarshalJSON()
}

// UnmarshalJSON decodes object from protobuf JSON format.
func (o *rwObject) UnmarshalJSON(data []byte) error {
	return (*object.Object)(o).UnmarshalJSON(data)
}
