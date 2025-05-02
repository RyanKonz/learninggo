package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type Team struct {
	name    string
	players []string
}

type League struct {
	Teams map[string]Team
	Wins  map[string]int
}

type Ranker interface {
	Ranking() []string
}

func main() {
	league := League{
		Teams: map[string]Team{
			"usa": {
				name:    "usa",
				players: []string{"p1", "p2", "p3"},
			},
			"italy": {
				name:    "italy",
				players: []string{"p1", "p2", "p3"},
			},
			"new zealand": {
				name:    "new zealand",
				players: []string{"p1", "p2", "p3"},
			},
		},
		Wins: map[string]int{},
	}

	league.MatchResult("usa", 30, "italy", 10)
	league.MatchResult("italy", 30, "new zealand", 10)
	league.MatchResult("new zealand", 30, "usa", 10)
	league.MatchResult("new zealand", 30, "italy", 10)
	league.MatchResult("usa", 30, "new zealand", 10)
	league.MatchResult("usa", 30, "italy", 10)
	league.MatchResult("usa", 30, "new zealand", 10)
	RankPrinter(&league, os.Stdout)
}

func (l *League) MatchResult(team1 string, score1 int, team2 string, score2 int) {
	if _, ok := l.Teams[team1]; !ok {
		return
	}

	if _, ok := l.Teams[team2]; !ok {
		return
	}

	switch {
	case score1 > score2:
		l.Wins[team1]++
		fmt.Println(team1, "wins!")
	case score1 < score2:
		l.Wins[team2]++
		fmt.Println(team2, "wins!")
	default:
		fmt.Println("tie game - no winners")
	}
}

func (l *League) Ranking() []string {
	names := make([]string, 0, len(l.Teams))
	for k := range l.Teams {
		names = append(names, k)
	}

	sort.Slice(names, func(i, j int) bool {
		return l.Wins[names[i]] > l.Wins[names[j]]
	})

	return names
}

func RankPrinter(r Ranker, w io.Writer) {
	results := r.Ranking()

	//var s string
	//for _, v := range results {
	//	s += v + "\n"
	//}
	//w.Write([]byte(s))

	for _, v := range results {
		io.WriteString(w, v)
		w.Write([]byte("\n"))
	}
}
