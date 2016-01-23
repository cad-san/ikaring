package ikaring

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ikaClient struct {
	http.Client
}

type stage struct {
	Name  string `json:"name"`
	Image string `json:"asset_path"`
}

type regulation struct {
	Regular []stage
	Gachi   []stage
}

type schedule struct {
	TimeBegin time.Time  `json:"datetime_begin"`
	TimeEnd   time.Time  `json:"datetime_end"`
	Stages    regulation `json:"stages"`
	GachiRule string     `json:"gachi_rule"`
}

type stageInfo struct {
	Festival  bool       `json:"festival"`
	Schedules []schedule `json:"schedule"`
}

func decodeJSONSchedule(data []byte) (*stageInfo, error) {
	info := &stageInfo{}
	if err := json.Unmarshal(data, info); err != nil {
		return nil, err
	}
	return info, nil
}

func (s schedule) String() string {
	timefmt := "01/02 15:04:05"
	str := fmt.Sprintf("%s - %s\n",
		s.TimeBegin.Format(timefmt), s.TimeEnd.Format(timefmt))

	str += "レギュラーマッチ\n"
	for i, stage := range s.Stages.Regular {
		if i == 0 {
			str += "\t"
		} else {
			str += ", "
		}
		str += stage.Name
	}
	str += "\n"

	str += fmt.Sprintf("ガチマッチ (%s)\n", s.GachiRule)
	for i, stage := range s.Stages.Gachi {
		if i == 0 {
			str += "\t"
		} else {
			str += ", "
		}
		str += stage.Name
	}
	str += "\n"
	return str
}
