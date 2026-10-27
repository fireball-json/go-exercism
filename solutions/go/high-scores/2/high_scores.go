package highscores

import (
	"sort"
)

type HighScores struct {
	highest int
	latest  int
	top     []int
	scores  []int
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	raw := append([]int(nil), scores...)

	if len(raw) == 0 {
		return &HighScores{
			scores: raw,
			top:    []int{},
		}
	}

	latest := raw[len(raw)-1]
	// 复制一份用于排序，不破坏原分数顺序
	cp := append([]int(nil), raw...)
	sort.Ints(cp)

	// 最多取3个，不足就全部取
	take := len(cp)
	if take > 3 {
		take = 3
	}
	// 取出最大的take个（升序排序后末尾）
	topPart := cp[len(cp)-take:]
	// 反转，变成从高到低
	for i, j := 0, len(topPart)-1; i < j; i, j = i+1, j-1 {
		topPart[i], topPart[j] = topPart[j], topPart[i]
	}

	return &HighScores{
		scores:  raw,
		highest: topPart[0],
		latest:  latest,
		top:     topPart,
	}
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	return s.scores
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
	return s.latest
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
	return s.highest
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
	return s.top
}
