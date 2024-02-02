package main

import "testing"

func TestLeague_MatchResult(t *testing.T) {
	type fields struct {
		Name  string
		Teams map[string]Team
		Wins  map[string]int
	}
	type args struct {
		team1  string
		score1 int
		team2  string
		score2 int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &League{
				Name:  tt.fields.Name,
				Teams: tt.fields.Teams,
				Wins:  tt.fields.Wins,
			}
			l.MatchResult(tt.args.team1, tt.args.score1, tt.args.team2, tt.args.score2)
		})
	}
}
