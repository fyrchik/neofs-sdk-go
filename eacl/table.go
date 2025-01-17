package eacl

import (
	"crypto/sha256"

	v2acl "github.com/nspcc-dev/neofs-api-go/v2/acl"
	cid "github.com/nspcc-dev/neofs-sdk-go/container/id"
	"github.com/nspcc-dev/neofs-sdk-go/session"
	"github.com/nspcc-dev/neofs-sdk-go/signature"
	"github.com/nspcc-dev/neofs-sdk-go/version"
)

// Table is a group of EACL records for single container.
//
// Table is compatible with v2 acl.EACLTable message.
type Table struct {
	version version.Version
	cid     *cid.ID
	token   *session.Token
	sig     *signature.Signature
	records []*Record
}

// CID returns identifier of the container that should use given access control rules.
func (t Table) CID() *cid.ID {
	return t.cid
}

// SetCID sets identifier of the container that should use given access control rules.
func (t *Table) SetCID(cid *cid.ID) {
	t.cid = cid
}

// Version returns version of eACL format.
func (t Table) Version() version.Version {
	return t.version
}

// SetVersion sets version of eACL format.
func (t *Table) SetVersion(version version.Version) {
	t.version = version
}

// Records returns list of extended ACL rules.
func (t Table) Records() []*Record {
	return t.records
}

// AddRecord adds single eACL rule.
func (t *Table) AddRecord(r *Record) {
	if r != nil {
		t.records = append(t.records, r)
	}
}

// SessionToken returns token of the session
// within which Table was set.
func (t Table) SessionToken() *session.Token {
	return t.token
}

// SetSessionToken sets token of the session
// within which Table was set.
func (t *Table) SetSessionToken(tok *session.Token) {
	t.token = tok
}

// Signature returns Table signature.
func (t Table) Signature() *signature.Signature {
	return t.sig
}

// SetSignature sets Table signature.
func (t *Table) SetSignature(sig *signature.Signature) {
	t.sig = sig
}

// ToV2 converts Table to v2 acl.EACLTable message.
//
// Nil Table converts to nil.
func (t *Table) ToV2() *v2acl.Table {
	if t == nil {
		return nil
	}

	v2 := new(v2acl.Table)

	if t.cid != nil {
		v2.SetContainerID(t.cid.ToV2())
	}

	if t.records != nil {
		records := make([]*v2acl.Record, 0, len(t.records))
		for _, record := range t.records {
			records = append(records, record.ToV2())
		}

		v2.SetRecords(records)
	}

	v2.SetVersion(t.version.ToV2())

	return v2
}

// NewTable creates, initializes and returns blank Table instance.
//
// Defaults:
//  - version: version.Current();
//  - container ID: nil;
//  - records: nil;
//  - session token: nil;
//  - signature: nil.
func NewTable() *Table {
	t := new(Table)
	t.SetVersion(*version.Current())

	return t
}

// CreateTable creates, initializes with parameters and returns Table instance.
func CreateTable(cid cid.ID) *Table {
	t := NewTable()
	t.SetCID(&cid)

	return t
}

// NewTableFromV2 converts v2 acl.EACLTable message to Table.
func NewTableFromV2(table *v2acl.Table) *Table {
	t := new(Table)

	if table == nil {
		return t
	}

	// set version
	if v := table.GetVersion(); v != nil {
		ver := version.Version{}
		ver.SetMajor(v.GetMajor())
		ver.SetMinor(v.GetMinor())

		t.SetVersion(ver)
	}

	// set container id
	if id := table.GetContainerID(); id != nil {
		if t.cid == nil {
			t.cid = new(cid.ID)
		}

		var h [sha256.Size]byte

		copy(h[:], id.GetValue())
		t.cid.SetSHA256(h)
	}

	// set eacl records
	v2records := table.GetRecords()
	t.records = make([]*Record, 0, len(v2records))

	for i := range v2records {
		t.records = append(t.records, NewRecordFromV2(v2records[i]))
	}

	return t
}

// Marshal marshals Table into a protobuf binary form.
func (t *Table) Marshal() ([]byte, error) {
	return t.ToV2().StableMarshal(nil)
}

// Unmarshal unmarshals protobuf binary representation of Table.
func (t *Table) Unmarshal(data []byte) error {
	fV2 := new(v2acl.Table)
	if err := fV2.Unmarshal(data); err != nil {
		return err
	}

	*t = *NewTableFromV2(fV2)

	return nil
}

// MarshalJSON encodes Table to protobuf JSON format.
func (t *Table) MarshalJSON() ([]byte, error) {
	return t.ToV2().MarshalJSON()
}

// UnmarshalJSON decodes Table from protobuf JSON format.
func (t *Table) UnmarshalJSON(data []byte) error {
	tV2 := new(v2acl.Table)
	if err := tV2.UnmarshalJSON(data); err != nil {
		return err
	}

	*t = *NewTableFromV2(tV2)

	return nil
}
