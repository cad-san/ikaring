package ikaring

import (
	"encoding/json"
)

type Intention struct {
	ID       *string
	ImageURL *string `json:"image"`
}

type Friend struct {
	Player
	Online    bool
	Mode      PlayMode
	Intention Intention
}

// PlayStatus is user
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
	}
	return table[string(m)]
}
