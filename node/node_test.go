package node

import (
	"context"
	"testing"

	libp2p "gx/ipfs/QmNh1kGFFdsPu79KNSaL4NUKUPb4Eiz4KHdMtFY6664RDp/go-libp2p"
	peerstore "gx/ipfs/QmXauCuJzmzapetmC6W4TuDJLL1yFFrVzSHoWv8YdbmnxH/go-libp2p-peerstore"

	"github.com/stretchr/testify/assert"
)

func TestNodeConstruct(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)

	nd, err := New(ctx)
	assert.NoError(err)
	assert.NotNil(nd.Host)

	nd.Stop()
}

func TestNodeNetworking(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)

	nd1, err := New(ctx, Libp2pOptions(libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0")))
	assert.NoError(err)
	nd2, err := New(ctx, Libp2pOptions(libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0")))
	assert.NoError(err)

	pinfo := peerstore.PeerInfo{
		ID:    nd2.Host.ID(),
		Addrs: nd2.Host.Addrs(),
	}

	err = nd1.Host.Connect(ctx, pinfo)
	assert.NoError(err)

	nd1.Stop()
	nd2.Stop()
}
