package ikaring

import (
	"testing"
)

func TestDecodeRegularSchedule(t *testing.T) {
	jsonStr := `
    {
        "festival":false,
        "schedule": [
            {
                "datetime_begin":"2016-01-01T11:00:00.000+09:00",
                "datetime_end":"2016-01-01T15:00:00.000+09:00",
                "stages":
                {
                    "regular":[
                        {"name":"ハコフグ倉庫","asset_path":"1.png"},
                        {"name":"シオノメ油田","asset_path":"2.png"}
                    ],
                    "gachi":[
                        {"name":"デカライン高架下","asset_path":"3.png"},
                        {"name":"アロワナモール","asset_path":"4.png"}
                    ]
                },
                "gachi_rule":"ガチエリア"
            },
            {
                "datetime_begin":"2016-01-01T15:00:00.000+09:00",
                "datetime_end":"2016-01-01T19:00:00.000+09:00",
                "stages":{
                    "regular":[
                        {"name":"Ｂバスパーク","asset_path":"5.png"},
                        {"name":"ホッケふ頭","asset_path":"6.png"}
                    ],
                    "gachi":[
                        {"name":"ネギトロ炭鉱","asset_path":"7.png"},
                        {"name":"モズク農園","asset_path":"8.png"}
                    ]
                },
                "gachi_rule":"ガチヤグラ"
            }]
    }
    `
	info, err := decodeJSONSchedule([]byte(jsonStr))

	if err != nil {
		t.Errorf("decode JSON failed with : %v", err)
	}

	if info.Festival {
		t.Errorf("info.festival should be false")
	}

	if info.FesSchedules != nil {
		t.Errorf("info.FesSchedules should be nil")
	}

	if info.Schedules == nil {
		t.Errorf("info.Schedules should not be nil")
	}
}
