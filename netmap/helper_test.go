package netmap

import (
	"testing"

	subnetid "github.com/nspcc-dev/neofs-sdk-go/subnet/id"
	"github.com/stretchr/testify/require"
)

func newFilter(name string, k, v string, op Operation, fs ...*Filter) *Filter {
	f := NewFilter()
	f.SetName(name)
	f.SetKey(k)
	f.SetOperation(op)
	f.SetValue(v)
	f.SetInnerFilters(fs...)
	return f
}

func newSelector(name string, attr string, c Clause, count uint32, filter string) *Selector {
	s := NewSelector()
	s.SetName(name)
	s.SetAttribute(attr)
	s.SetCount(count)
	s.SetClause(c)
	s.SetFilter(filter)
	return s
}

func newPlacementPolicy(bf uint32, rs []*Replica, ss []*Selector, fs []*Filter) *PlacementPolicy {
	p := NewPlacementPolicy()
	p.SetContainerBackupFactor(bf)
	p.SetReplicas(rs...)
	p.SetSelectors(ss...)
	p.SetFilters(fs...)
	return p
}

func newReplica(c uint32, s string) *Replica {
	r := NewReplica()
	r.SetCount(c)
	r.SetSelector(s)
	return r
}

func newSubnetID(id uint32) *subnetid.ID {
	var s subnetid.ID
	s.SetNumber(id)
	return &s
}

func nodeInfoFromAttributes(props ...string) NodeInfo {
	attrs := make([]*NodeAttribute, len(props)/2)
	for i := range attrs {
		attrs[i] = NewNodeAttribute()
		attrs[i].SetKey(props[i*2])
		attrs[i].SetValue(props[i*2+1])
	}
	n := NewNodeInfo()
	n.SetAttributes(attrs...)
	return *n
}

func getTestNode(props ...string) *Node {
	m := make(map[string]string, len(props)/2)
	for i := 0; i < len(props); i += 2 {
		m[props[i]] = props[i+1]
	}
	return &Node{AttrMap: m}
}

type enumIface interface {
	FromString(string) bool
	String() string
}

type enumStringItem struct {
	val enumIface
	str string
}

func testEnumStrings(t *testing.T, e enumIface, items []enumStringItem) {
	for _, item := range items {
		require.Equal(t, item.str, item.val.String())

		s := item.val.String()

		require.True(t, e.FromString(s), s)

		require.EqualValues(t, item.val, e, item.val)
	}

	// incorrect strings
	for _, str := range []string{
		"some string",
		"undefined",
	} {
		require.False(t, e.FromString(str))
	}
}
