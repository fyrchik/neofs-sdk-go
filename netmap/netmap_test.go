package netmap_test

import (
	"fmt"
	"math/rand"
	"testing"

	v2netmap "github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/nspcc-dev/neofs-sdk-go/netmap"
	netmaptest "github.com/nspcc-dev/neofs-sdk-go/netmap/test"
	"github.com/stretchr/testify/require"
)

func newNodeInfo(attribute, value string) (x netmap.NodeInfo) {
	key := make([]byte, 33)
	rand.Read(key)

	x.SetPublicKey(key)
	x.SetNetworkEndpoints("1", "2", "3")
	x.SetAttribute(attribute, value)
	return
}

func TestNetmapSorting(t *testing.T) {
	var nm netmap.NetMap
	nodes := []netmap.NodeInfo{
		newNodeInfo("Location", "Moscow"),
		newNodeInfo("Location", "SPB"),
		newNodeInfo("Location", "Moscow"),
		newNodeInfo("Location", "Other"),
		newNodeInfo("Location", "None"),
	}

	var s netmap.Sorter
	require.NoError(t, s.DecodeString("CASE Location WHEN Moscow 2 WHEN SPB 1 END"))

	v, err := nm.WeightedPlacementVectors([][]netmap.NodeInfo{nodes}, nil, s.WeightFunc())
	require.NoError(t, err)
	for i := range v[0] {
		fmt.Println(v[0][i].Attribute("Location"))
	}
	require.Equal(t, "Moscow", v[0][0].Attribute("Location"))
	require.Equal(t, "Moscow", v[0][1].Attribute("Location"))
	require.Equal(t, "SPB", v[0][2].Attribute("Location"))
}

func TestNetMapNodes(t *testing.T) {
	var nm netmap.NetMap

	require.Empty(t, nm.Nodes())

	nodes := []netmap.NodeInfo{netmaptest.NodeInfo(), netmaptest.NodeInfo()}

	nm.SetNodes(nodes)
	require.ElementsMatch(t, nodes, nm.Nodes())

	nodesV2 := make([]v2netmap.NodeInfo, len(nodes))
	for i := range nodes {
		nodes[i].WriteToV2(&nodesV2[i])
	}

	var m v2netmap.NetMap
	nm.WriteToV2(&m)

	require.ElementsMatch(t, nodesV2, m.Nodes())
}

func TestNetMap_SetEpoch(t *testing.T) {
	var nm netmap.NetMap

	require.Zero(t, nm.Epoch())

	const e = 158

	nm.SetEpoch(e)
	require.EqualValues(t, e, nm.Epoch())

	var m v2netmap.NetMap
	nm.WriteToV2(&m)

	require.EqualValues(t, e, m.Epoch())
}
