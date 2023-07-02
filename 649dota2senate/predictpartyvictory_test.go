package _49dota2senate

import "testing"

func TestPredictPartyVictory(t *testing.T) {
      senate := "RDD"
	s :=  predictPartyVictory(senate)
	t.Log(s)
}
