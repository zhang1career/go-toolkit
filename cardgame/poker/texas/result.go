package texas

import (
	"github.com/zhang1career/lib/cardgame"
	"github.com/zhang1career/lib/log"
)

type Result struct {
	IsWin   int
	Tip     string
}

func GetResult(communityCards []cardgame.Card, holeCards []cardgame.Card) []Result {
	game, err := New()
	if err != nil {
		log.Error(err.Error())
	}
	return game.isWin(communityCards, [][]cardgame.Card{holeCards})
}


func (this *Texas) isWin(communityCards []cardgame.Card, holeCards [][]cardgame.Card) []Result {
	ret := make([]Result, len(holeCards))
	maxIndex := 0; maxCate := 0; maxScore := 0.0
	for i, holeCard := range holeCards {
		originCards := append(communityCards, holeCard...)
		cate, score := this.Evaluate(originCards)
		ret[i].IsWin = 0
		ret[i].Tip = this.GetAllWinCates()[cate]
		if maxCate > cate {
			continue
		}
		if maxCate < cate {
			maxIndex = i; maxCate = cate; maxScore = score
			continue
		}
		if maxScore >= score {
			continue
		}
		maxIndex = i; maxScore = score
	}
	ret[maxIndex].IsWin = 1
	return ret
}

type Stat struct {
	Score       float64
	Count       int
	ScoreAve    float64
}

func GetStat(commonCards []cardgame.Card, ownCards []cardgame.Card) map[int]Stat {
	if len(commonCards) < 3 || len(commonCards) > 5 || len(ownCards) != 2 {
		return nil
	}
	
	game, err := New()
	if err != nil {
		log.Error(err.Error())
	}
	paramCards := game.AddCards(commonCards, ownCards)
	availableCards := game.Combinate(game.DelCards(game.GetAllCards(), paramCards), 7 - len(paramCards))
	
	stats := make(map[int]Stat)
	for _, availableCard := range availableCards {
		originCards := append(paramCards, availableCard...)
		cate, score := game.Evaluate(originCards)
		if stat, ok := stats[cate]; ok {
			stat.Score += score
			stat.Count++
			stats[cate] = stat
		} else {
			stats[cate] = Stat{Score: score, Count: 1}
		}
	}
	
	for cate, stat := range stats {
		stat.ScoreAve = stat.Score / float64(stat.Count)
		stats[cate] = stat
	}
	return stats
}