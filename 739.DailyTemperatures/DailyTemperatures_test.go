package _39_DailyTemperatures

import "testing"

func TestDailyTemperatures(t *testing.T) {
	var data []int
	data = append(data,73,74,75,71,69,72,76,73)
	res := dailyTemperatures(data)
    t.Log(res)
}
