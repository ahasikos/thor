// Copyright (c) 2018 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package comm

import (
	"github.com/ethereum/go-ethereum/p2p/discv5"
	"github.com/stretchr/testify/assert"
	"github.com/vechain/thor/v2/chain"
	"github.com/vechain/thor/v2/genesis"
	"github.com/vechain/thor/v2/muxdb"
	"github.com/vechain/thor/v2/state"
	"github.com/vechain/thor/v2/txpool"
	"testing"
	"time"
)

func newCommunicator() *Communicator {
	db := muxdb.NewMem()
	stater := state.NewStater(db)
	gen := genesis.NewDevnet()
	blk, _, _, _ := gen.Build(stater)

	repo, _ := chain.NewRepository(db, blk)
	txPool := txpool.New(repo, stater, txpool.Options{
		Limit:           10000,
		LimitPerAccount: 16,
		MaxLifetime:     10 * time.Minute,
	})

	return New(repo, txPool)
}

func TestCommunicatorProtocols(t *testing.T) {
	comm := newCommunicator()
	protocols := comm.Protocols()

	assert.Equal(t, "thor", protocols[0].Name)
	assert.Equal(t, uint(1), protocols[0].Version)
	assert.Equal(t, uint64(8), protocols[0].Length)
}

func TestCommunicatorDiscTopic(t *testing.T) {
	comm := newCommunicator()
	discTopic := comm.DiscTopic()

	assert.Equal(t, discv5.Topic("thor1@a4988aba7aea69f6"), discTopic)
}

func TestCommunicatorPeerCount(t *testing.T) {
	comm := newCommunicator()

	assert.Equal(t, 0, comm.PeerCount())
}

func TestCommunicatorPeerStats(t *testing.T) {
	comm := newCommunicator()

	assert.Empty(t, comm.PeersStats())
}
