package ikaring

import (
	"fmt"
	"testing"
)

func TestDecodeRanking(t *testing.T) {
	jsonStr := `
    {
        "regular": [
            {
                "hashed_id":"hash_1",
                "rank":["1"],
                "score":["1","0","0"],
                "mii_url":"https://example.com/player1.png",
                "mii_name":"player1",
                "weapon":"c9f7cead9ee5ada35437d7c2ea8ddae6ca1dacfc6c9b01d5939cfc0ff59fe0ea-dummy.png",
                "head":"head1.png", "clothes":"clothes1.png", "shoes":"shoes1.png",
                "weapon2x":"weapon1.png", "head2x":"head1.png", "clothes2x":"clothes1.png","shoes2x":"shoes1.png"
            },
            {
                "hashed_id":"hash_2",
                "rank":["2"],
                "score":["2","0","0"],
                "mii_url":"https://example.com/player2.png",
                "mii_name":"player2",
                "weapon":"c9f7cead9ee5ada35437d7c2ea8ddae6ca1dacfc6c9b01d5939cfc0ff59fe0ea-dummy.png",
                "head":"head2.png", "clothes":"clothes2.png", "shoes":"shoes2.png",
                "weapon2x":"weapon2.png", "head2x":"head2.png", "clothes2x":"clothes2.png", "shoes2x":"shoes2.png"
            } ],
        "gachi":[
            {
                "hashed_id":"hash_1",
                "rank":["1"],
                "score":["3","2","1"],
                "mii_url":"https://example.com/player1.png",
                "mii_name":"player1",
                "weapon":"c9f7cead9ee5ada35437d7c2ea8ddae6ca1dacfc6c9b01d5939cfc0ff59fe0ea-dummy.png",
                "head":"head1.png", "clothes":"clothes1.png", "shoes":"shoes1.png",
                "weapon2x":"weapon1.png", "head2x":"head1.png", "clothes2x":"clothes1.png","shoes2x":"shoes1.png"
            },
            {
                "hashed_id":"hash_2",
                "rank":["2"],
                "score":["2","1","0"],
                "mii_url":"https://example.com/player2.png",
                "mii_name":"player2",
                "weapon":"c9f7cead9ee5ada35437d7c2ea8ddae6ca1dacfc6c9b01d5939cfc0ff59fe0ea-dummy.png",
                "head":"head2.png", "clothes":"clothes2.png", "shoes":"shoes2.png",
                "weapon2x":"weapon2.png", "head2x":"head2.png", "clothes2x":"clothes2.png", "shoes2x":"shoes2.png"
            }]
    }
    `

	r, err := decodeJSONRanking([]byte(jsonStr))

	if err != nil {
		t.Errorf("decodeJSONRanking() should not be error with %v", err)
	}

	if len(r.Regular) != 2 {
		t.Errorf("r.Regular should be 2 length but %d", len(r.Gachi))
	}

	if len(r.Gachi) != 2 {
		t.Errorf("r.Gachi should be 2 length but %d", len(r.Gachi))
	}

	expect := []Player{}
	for i := range r.Regular {
		p := Player{}
		p.Name = fmt.Sprintf("player%d", i+1)
		p.ID = fmt.Sprintf("hash_%d", i+1)
		p.MiiURL = fmt.Sprintf("https://example.com/player%d.png", i+1)
		p.Weapon = "わかばシューター"
		expect = append(expect, p)
	}

	for i, p := range r.Regular {
		if p.Rank != i+1 {
			t.Errorf("player[%d].Score parse failed %d", i, p.Rank)
		}
		if p.Score != (i+1)*100 {
			t.Errorf("player[%d].Score parse failed %d", i, p.Score)
		}
		if p.Player != expect[i] {
			t.Errorf("player[%d] parse faild expect %v but %v", i, expect[i], p.Player)
		}
	}
	for i, p := range r.Gachi {
		if p.Rank != i+1 {
			t.Errorf("player rank shoud be sorted but player[%d].Rank is %d", i, p.Rank)
		}
		if p.Player != expect[i] {
			t.Errorf("player[%d] parse faild expect %v but %v", i, expect[i], p.Player)
		}
	}
}
