package netmap

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

func TestSorter(t *testing.T) {
	var s Sorter
	require.NoError(t, s.DecodeString("CASE Location WHEN Moscow 1 END"))
	spew.Dump(s)
}
