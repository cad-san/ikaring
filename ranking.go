package ikaring

import (
	"encoding/json"
	"strconv"
	"strings"
)

// Player has player info.
type Player struct {
	ID     string `json:"hashed_id"`
	Name   string `json:"mii_name"`
	MiiURL string `json:"mii_url"`
	Weapon string
}

// PlayerScore has player info for ranking
type PlayerScore struct {
	Player
	Rank  int
	Score int
}

// RankingInfo has ranking for Regular Match and Gachi Match
type RankingInfo struct {
	Regular []PlayerScore
	Gachi   []PlayerScore
}

func decodeJSONRanking(data []byte) (*RankingInfo, error) {
	r := &RankingInfo{}
	err := json.Unmarshal(data, r)
	return r, err
}

// UnmarshalJSON parse JSON for PlayerScore for SplatNet Ranking.
func (p *PlayerScore) UnmarshalJSON(b []byte) error {
	var raw = struct {
		Player
		Rank  []string
		Score []string
	}{}
	err := json.Unmarshal(b, &raw)
	if err != nil {
		return err
	}
	p.Player = raw.Player
	p.Rank, err = strconv.Atoi(strings.Join(raw.Rank, ""))
	if err != nil {
		p.Rank = 0
	}
	p.Score, err = strconv.Atoi(strings.Join(raw.Score, ""))
	if err != nil {
		p.Score = 0
	}
	return nil
}
