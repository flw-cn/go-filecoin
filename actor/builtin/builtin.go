// Package builtin implements the predefined actors in Filecoin.
package builtin

import (
	"fmt"

	"github.com/filecoin-project/go-filecoin/actor/builtin/account"
	"github.com/filecoin-project/go-filecoin/actor/builtin/initactor"
	"github.com/filecoin-project/go-filecoin/actor/builtin/miner"
	"github.com/filecoin-project/go-filecoin/actor/builtin/paymentbroker"
	"github.com/filecoin-project/go-filecoin/actor/builtin/storagemarket"
	"github.com/filecoin-project/go-filecoin/exec"
	"github.com/filecoin-project/go-filecoin/types"
	cid "github.com/ipfs/go-cid"
)

// codeVersion identifies an ExecutableActor by its code and protocol version
type codeVersion struct {
	code            cid.Cid
	protocolVersion uint64
}

type Actors struct {
	actors map[codeVersion]exec.ExecutableActor
}

// GetActorCode returns executable code for an actor by code cid at a specific protocol version
func (ba Actors) GetActorCode(code cid.Cid, version uint64) (exec.ExecutableActor, error) {
	if !code.Defined() {
		return nil, fmt.Errorf("missing code")
	}
	actor, ok := ba.actors[codeVersion{code: code, protocolVersion: version}]
	if !ok {
		return nil, fmt.Errorf("unknown code: %s", code.String())
	}
	return actor, nil
}

type BuiltinActorsBuilder struct {
	actors map[codeVersion]exec.ExecutableActor
}

func NewActorsBuilder() *BuiltinActorsBuilder {
	return &BuiltinActorsBuilder{actors: map[codeVersion]exec.ExecutableActor{}}
}

func (bab *BuiltinActorsBuilder) AddAll(actors Actors) *BuiltinActorsBuilder {
	for cv, a := range actors.actors {
		bab.Add(cv.code, cv.protocolVersion, a)
	}
	return bab
}

func (bab *BuiltinActorsBuilder) Add(c cid.Cid, version uint64, actor exec.ExecutableActor) *BuiltinActorsBuilder {
	bab.actors[codeVersion{code: c, protocolVersion: version}] = actor
	return bab
}

func (bab *BuiltinActorsBuilder) Build() Actors {
	return Actors{actors: bab.actors}
}

// DefaultActors is list of all actors that ship with Filecoin.
// They are indexed by their CID.
var DefaultActors = NewActorsBuilder().
	Add(types.AccountActorCodeCid, 0, &account.Actor{}).
	Add(types.StorageMarketActorCodeCid, 0, &storagemarket.Actor{}).
	Add(types.PaymentBrokerActorCodeCid, 0, &paymentbroker.Actor{}).
	Add(types.MinerActorCodeCid, 0, &miner.Actor{}).
	Add(types.BootstrapMinerActorCodeCid, 0, &miner.Actor{Bootstrap: true}).
	Add(types.InitActorCodeCid, 0, &initactor.Actor{}).
	Build()
