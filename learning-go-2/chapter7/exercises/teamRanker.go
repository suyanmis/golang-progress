package main

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

// ex 1
type Team struct {
	Name    string
	Players []string
}
type League struct {
	Teams []Team
	Wins  map[string]int
}
type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	io.WriteString(w, strings.Join(r.Ranking(), "\n"))
}

func NewLeague(t []Team) *League {
	wins := make(map[string]int)
	for _, team := range t {
		wins[team.Name] = 0
	}
	return &League{
		Teams: t,
		Wins:  wins,
	}
}
func (l *League) MatchResult(t1name string, t1score int, t2name string, t2score int) {
	var winner string
	if t1score > t2score {
		winner = t1name
	} else if t2score < t1score {
		winner = t2name
	} else {
		return
	}
	if _, ok := l.Wins[t2name]; !ok {
		l.Wins[winner] = 1
	} else {
		l.Wins[winner]++
	}
}

func (l League) Ranking() []string {
	type kv struct {
		Key   string
		Value int
	}
	var sortedList []kv
	for team, score := range l.Wins {
		sortedList = append(sortedList, kv{team, score})
	}
	sort.Slice(sortedList, func(i, j int) bool {
		return sortedList[i].Value > sortedList[j].Value
	})
	var teamNames []string
	for _, kv := range sortedList {
		teamNames = append(teamNames, fmt.Sprintf("%s %d", kv.Key, kv.Value))
	}
	return teamNames
}

func main() {
	t1 := Team{Name: "Team Crazy", Players: []string{"Crazy1", "Crazy2", "Crazy3", "Crazy4", "Crazy5"}}
	t2 := Team{Name: "Team Banger", Players: []string{"Banger1", "Banger2", "Banger3", "Banger4", "Banger5"}}
	t3 := Team{Name: "Team Dominants", Players: []string{"Dominants1", "Dominants2", "Dominants3", "Dominants4", "Dominants5"}}
	teams := []Team{t1, t2, t3}
	league := NewLeague(teams)
	league.MatchResult("Team Crazy", 55, "Team Banger", 49)
	league.MatchResult("Team Dominants", 55, "Team Banger", 49)
	league.MatchResult("Team Dominants", 55, "Team Crazy", 49)
	league.MatchResult("Team Dominants", 55, "Team Banger", 49)
	league.MatchResult("Team Banger", 55, "Team Crazy", 49)
	league.MatchResult("Team Crazy", 55, "Team Dominants", 49)
	fmt.Printf("League Ranking:\n%s\n", strings.Join(league.Ranking(), "\n"))
}
