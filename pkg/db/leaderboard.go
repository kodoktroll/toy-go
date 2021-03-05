package db

type Leaderboard struct {
	Count int
	Users []*User
}

func (db *Database) GetLeaderboard() (*Leaderboard, error) {
	scores := db.Client.ZRangeWithScores(Ctx, "leaderboard", 0, -1)
	if scores == nil {
		return nil, ErrNil
	}
	count := len(scores.Val())
	users := make([]*User, count)
	for i, member := range scores.Val() {
		users[i] = &User{
			Username: member.Member.(string),
			Points:   int(member.Score),
			Rank:     i,
		}
	}
	return &Leaderboard{
		Count: count,
		Users: users,
	}, nil
}
