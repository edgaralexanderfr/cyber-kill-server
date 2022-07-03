package net

type MatchmakingFactory struct{}

func (matchmakingFactory *MatchmakingFactory) NewMatchmaking() MatchmakingInterface {
	return &Matchmaking{}
}
