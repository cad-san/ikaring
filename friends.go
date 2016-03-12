package ikaring

import (
	"encoding/json"
)

// Intention is stasus of friend's recruitment for playing
type Intention struct {
	ID       *string
	ImageURL *string `json:"image"`
}

// Friend has infomation in FriendList
type Friend struct {
	Player
	Online    bool
	Mode      PlayMode
	Intention Intention
}

// PlayMode indicates mode that friend plays with splatoon.
type PlayMode string

func decodeJSONFriendList(data []byte) ([]Friend, error) {
	var f []Friend
	err := json.Unmarshal(data, &f)
	return f, err
}

func (m PlayMode) String() string {
	var table = map[string]string{
		"playing": "Splatoon プレイ中",
		"regular": "レギュラーマッチでイカしてるぜ！",
		"gachi":   "ガチマッチでウデだめししてるぜ！",
		"tag":     "タッグマッチでガチってるぜ！",
		"private": "プライベートマッチで自由に楽しんでるぜ！",
		"fes":     "フェスマッチでお祭りさわぎ！",
		"online":  "オンライン",
		"offline": "オフライン",
	}
	return table[string(m)]
}
