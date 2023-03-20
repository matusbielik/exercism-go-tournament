package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type (
	matchResult string
	position    uint8
)

const (
	MATCH_WIN  matchResult = "win"
	MATCH_LOSS matchResult = "loss"
	MATCH_DRAW matchResult = "draw"

	POSITION_FIRST  position = 1
	POSITION_SECOND position = 2

	TEAM_COL_WIDTH = 31
)

func NewMatchResultByPosition(matchResultStr string, position position) (matchResult, error) {
	if matchResultStr == "draw" {
		return MATCH_DRAW, nil
	}
	switch position {
	case POSITION_FIRST:
		switch matchResultStr {
		case "win":
			return MATCH_WIN, nil
		case "loss":
			return MATCH_LOSS, nil
		default:
			return "", fmt.Errorf("invalid match result")
		}
	case POSITION_SECOND:
		switch matchResultStr {
		case "win":
			return MATCH_LOSS, nil
		case "loss":
			return MATCH_WIN, nil
		default:
			return "", fmt.Errorf("invalid match result")
		}
	default:
		return "", fmt.Errorf("invalid position")
	}
}

type teamSummary struct {
	MatchesPlayed uint32
	Wins          uint32
	Draws         uint32
	Losses        uint32
	Points        uint32
	TeamName      string
}

func (ts teamSummary) Print(writer io.Writer) {
	fmt.Fprintf(
		writer,
		"%s|  %d |  %d |  %d |  %d |  %d\n",
		prepareTeamString(ts.TeamName, TEAM_COL_WIDTH), ts.MatchesPlayed, ts.Wins, ts.Draws, ts.Losses, ts.Points,
	)
}

func NewTeamSummary(teamName string, mRes matchResult) *teamSummary {
	switch mRes {
	case MATCH_WIN:
		return &teamSummary{TeamName: teamName, MatchesPlayed: 1, Wins: 1, Points: 3}
	case MATCH_DRAW:
		return &teamSummary{TeamName: teamName, MatchesPlayed: 1, Draws: 1, Points: 1}
	case MATCH_LOSS:
		return &teamSummary{TeamName: teamName, MatchesPlayed: 1, Losses: 1}
	default:
		panic("invalid match result")
	}
}

func (ts *teamSummary) Update(matchResult matchResult, position position) {
	ts.MatchesPlayed += 1

	switch matchResult {
	case MATCH_WIN:
		ts.Wins += 1
		ts.Points += 3
	case MATCH_DRAW:
		ts.Draws += 1
		ts.Points += 1
	case MATCH_LOSS:
		ts.Losses += 1
	}
}

type teamSummaryMap map[string]*teamSummary

func (tsMap teamSummaryMap) Update(teamName string, matchResultStr string, position position) error {
	matchResult, err := NewMatchResultByPosition(matchResultStr, position)
	if err != nil {
		return err
	}
	if team, ok := tsMap[teamName]; ok {
		team.Update(matchResult, position)
		return nil
	}
	tsMap[teamName] = NewTeamSummary(teamName, matchResult)

	return nil
}

func (tsMap teamSummaryMap) Print(writer io.Writer) {
	fmt.Fprintf(
		writer,
		"%s| %s |  %s |  %s |  %s |  %s\n",
		prepareTeamString("Team", TEAM_COL_WIDTH), "MP", "W", "D", "L", "P",
	)
	for _, ts := range tsMap.toList() {
		ts.Print(writer)
	}
}

func (tsMap teamSummaryMap) toList() []*teamSummary {
	summaryList := make([]*teamSummary, 0, len(tsMap))

	for _, summaryRow := range tsMap {
		summaryList = append(summaryList, summaryRow)
	}
	sort.Slice(summaryList, func(i, j int) bool {
		if summaryList[i].Points == summaryList[j].Points {
			return summaryList[i].TeamName < summaryList[j].TeamName
		} else {
			return summaryList[i].Points > summaryList[j].Points
		}
	})

	return summaryList
}

func Tally(reader io.Reader, writer io.Writer) error {
	summaryMap := teamSummaryMap{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, ";")
		if len(parts) != 3 {
			return fmt.Errorf("invalid line: %s", line)
		}
		team1Name, team2Name, matchResultStr := parts[0], parts[1], parts[2]

		if err := summaryMap.Update(team1Name, matchResultStr, POSITION_FIRST); err != nil {
			return err
		}
		if err := summaryMap.Update(team2Name, matchResultStr, POSITION_SECOND); err != nil {
			return err
		}
	}

	summaryMap.Print(writer)

	return nil
}

func prepareTeamString(teamName string, teamColWidth int) string {
	teamColPaddingLen := teamColWidth - len(teamName)
	return fmt.Sprintf("%s%s", teamName, strings.Repeat(" ", teamColPaddingLen))
}
