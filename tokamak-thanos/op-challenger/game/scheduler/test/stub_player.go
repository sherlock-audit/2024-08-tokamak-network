package test

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/tokamak-network/tokamak-thanos/op-challenger/game/types"
)

type StubGamePlayer struct {
	Addr          common.Address
	ProgressCount int
	StatusValue   types.GameStatus
	Dir           string
	PrestateErr   error
}

func (g *StubGamePlayer) ValidatePrestate(_ context.Context) error {
	return g.PrestateErr
}

func (g *StubGamePlayer) ProgressGame(_ context.Context) types.GameStatus {
	g.ProgressCount++
	return g.StatusValue
}

func (g *StubGamePlayer) Status() types.GameStatus {
	return g.StatusValue
}
