package main

import (
	"fmt"
	"math/rand"
)

const (
	win            = 100
	gamesPerSeries = 10
)

// 得分情况
type score struct {
	players  []int // 玩家目前得分
	current  int   // 当前玩家
	thisTurn int   // 本轮得分
}

// 行动
type action func(s *score) (turnIsOver bool)

func roll(s *score) bool {
	outcome := rand.Intn(6) + 1 //[1,6]
	fmt.Printf("%d roll! %d!", s.current, outcome)
	if outcome == 1 {
		return true
	}
	s.thisTurn += outcome
	fmt.Printf("get %d this turn!\n", s.thisTurn)
	return false
}

func stay(s *score) bool {
	s.players[s.current] += s.thisTurn
	fmt.Printf("%d stay. has %d now!", s.current, s.players[s.current])
	return true
}

// 根据当前分数情况制定策略
type strategy func(*score) action

func stayAtK(k int) strategy {
	return func(s *score) action {
		if s.thisTurn >= k {
			return stay
		}
		return roll
	}
}

// 玩游戏 return winner
func play(players []strategy) int {
	// 初始化 随机从某个人开始
	turnIsOver := false
	s := &score{make([]int, len(players)), rand.Intn(len(players)), 0}
	for s.players[s.current]+s.thisTurn < win {
		action := players[s.current](s)
		turnIsOver = action(s)
		if turnIsOver {
			fmt.Println("<turn is over>")
			s.current = (s.current + 1) % len(players)
			s.thisTurn = 0
		}
	}
	stay(s)
	return s.current
}

func main() {
	players := make([]strategy, 3)
	for i := range players {
		players[i] = stayAtK((i + 1) * 5)
	}
	fmt.Println(play(players), "win!")
}
