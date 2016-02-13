package ikaring

import (
	"testing"
)

func TestDecodeFriendList(t *testing.T) {
	jsonStr := `
    [
        {
            "hashed_id":"hash_1",
            "online":true,
            "mode":"regular",
            "intention":{
                "id":null,
                "image":null
            },
            "mii_name":"player1",
            "mii_url":"https://example.com/player1.png"
        }
    ]
    `

	l, err := decodeJSONFriendList([]byte(jsonStr))

	if err != nil {
		t.Errorf("decodeJSONFriendList() should not be error with %v", err)
	}
	if len(l) != 1 {
		t.Errorf("list should be 1 length but %d", len(l))
	}
	expect := Friend{
		Player: Player{
			ID:     "hash_1",
			Name:   "player1",
			MiiURL: "https://example.com/player1.png",
		},
		Online: true,
		Mode:   "regular",
	}
	actual := l[0]
	if actual != expect {
		t.Errorf("decode failure\nexpect:%v\nactual:%v\n", expect, actual)
	}
}

func TestFriendStringer(t *testing.T) {
	var m PlayMode = "playing"
	if m.String() != "Splatoon プレイ中" {
		t.Errorf("convert playing-mode failure:%v", m.String())
	}
}
