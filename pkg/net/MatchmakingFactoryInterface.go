package net

type MatchmakingFactoryInterface interface {
	NewMatchmaking() MatchmakingInterface
}
