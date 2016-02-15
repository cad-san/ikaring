package ikaring

import (
	"encoding/json"
	"fmt"
	"time"
)

// Stage has stage infomation.
type Stage struct {
	Name  string `json:"name"`       // stage name
	Image string `json:"asset_path"` // url for stage image
}

// Regulation is buttle regulation for Regular Match & Gachi Match.
// It includes stage set.
type Regulation struct {
	Regular []Stage // stage set for Regular Match
	Gachi   []Stage // stage set for Gatch Match
}

// TimeSpan is a period for Buttle regulation
type TimeSpan struct {
	TimeBegin time.Time `json:"datetime_begin"` // Start Time
	TimeEnd   time.Time `json:"datetime_end"`   // End Time
}

// Schedule is a rule set for normal day.
// It has time span, stage set, and gachi rule.
type Schedule struct {
	TimeSpan
	Stages    Regulation `json:"stages"`     // Stage Set
	GachiRule string     `json:"gachi_rule"` // rule for Gatch Match
}

// Festival is a rule set in Festival
type Festival struct {
	TimeSpan
	TeamA  string  `json:"team_alpha_name"` // Name of Team A
	TeamB  string  `json:"team_bravo_name"` // Name of Team B
	Stages []Stage `json:"stages"`          // Stage Set
}

// StageInfo is all set of Stage Schedules.
type StageInfo struct {
	Festival     bool        // during the festival
	Schedules    *[]Schedule // Schedule info for normal day(nil in festival)
	FesSchedules *[]Festival // Schedule info for festival(nil in normal day)
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

	i.Festival = p.Festival

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

func (t TimeSpan) String() string {
	timefmt := "01/02 15:04:05"
	str := fmt.Sprintf("%s - %s",
		t.TimeBegin.Format(timefmt), t.TimeEnd.Format(timefmt))
	return str
}

func (s Schedule) String() string {
	str := s.TimeSpan.String()
	str += "\n"
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
	str := s.TimeSpan.String()
	str += "\n"
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
