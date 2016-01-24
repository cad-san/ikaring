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

type Stage struct {
	Name  string `json:"name"`
	Image string `json:"asset_path"`
}

type Regulation struct {
	Regular []Stage
	Gachi   []Stage
}

type Schedule struct {
	TimeBegin time.Time  `json:"datetime_begin"`
	TimeEnd   time.Time  `json:"datetime_end"`
	Stages    Regulation `json:"stages"`
	GachiRule string     `json:"gachi_rule"`
}

type Festival struct {
	TimeBegin time.Time `json:"datetime_begin"`
	TimeEnd   time.Time `json:"datetime_end"`
	TeamA     string    `json:"team_alpha_name"`
	TeamB     string    `json:"team_bravo_name"`
	Stages    []Stage   `json:"stages"`
}

type StageInfo struct {
	Festival     bool
	Schedules    *[]Schedule
	FesSchedules *[]Festival
}

type stageJSON struct {
	Festival  bool            `json:"festival"`
	Schedules json.RawMessage `json:"schedule"`
}

func decodeJSONSchedule(data []byte) (*StageInfo, error) {
	i := &StageInfo{}
	p := &stageJSON{}

	if err := json.Unmarshal(data, p); err != nil {
		return nil, err
	}

	if p.Festival {
		var fes []Festival
		if err := json.Unmarshal(p.Schedules, &fes); err == nil {
			i.FesSchedules = &fes
		}
	} else {
		var regular []Schedule
		if err := json.Unmarshal(p.Schedules, &regular); err == nil {
			i.Schedules = &regular
		}
	}

	return i, nil
}

func (s Schedule) String() string {
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

func (s Festival) String() string {
	timefmt := "01/02 15:04:05"
	str := fmt.Sprintf("%s - %s\n",
		s.TimeBegin.Format(timefmt), s.TimeEnd.Format(timefmt))

	str += fmt.Sprintf("フェスマッチ [%s vs %s]\n", s.TeamA, s.TeamB)
	for i, stage := range s.Stages {
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
