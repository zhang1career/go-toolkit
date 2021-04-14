package texas

import (
	"github.com/zhang1career/golab/cardgame"
	"github.com/zhang1career/golab/cardgame/poker"
)

const (
	HighCard            = 0     // 高牌
	OnePair             = 1     // 一对
	TwoPair             = 2     // 两对
	ThreeOfAKind        = 3     // 三条
	Straight            = 4     // 顺子
	Flush               = 5     // 同花
	FullHouse           = 6     // 葫芦
	FourOfAKind         = 7     // 四条
	StraightFlush       = 8     // 同花顺
	RoyalStraightFlush  = 9     // 皇家同花顺
)

type Texas struct {
	poker.Poker
}

func New() (*Texas, error) {
	p, err := poker.New()
	if err != nil {
		return nil, err
	}
	
	texas := Texas{
		Poker:      *p,
	}
	
	return &texas, nil
}

func (this *Texas) GetAllCards() []cardgame.Card {
	return []cardgame.Card {
		{1,1}, {1,2}, {1,3}, {1,4},
		{2,1}, {2,2}, {2,3}, {2,4},
		{3,1}, {3,2}, {3,3}, {3,4},
		{4,1}, {4,2}, {4,3}, {4,4},
		{5,1}, {5,2}, {5,3}, {5,4},
		{6,1}, {6,2}, {6,3}, {6,4},
		{7,1}, {7,2}, {7,3}, {7,4},
		{8,1}, {8,2}, {8,3}, {8,4},
		{9,1}, {9,2}, {9,3}, {9,4},
		{10,1}, {10,2}, {10,3}, {10,4},
		{11,1}, {11,2}, {11,3}, {11,4},
		{12,1}, {12,2}, {12,3}, {12,4},
		{13,1}, {13,2}, {13,3}, {13,4},
	}
}

func (this *Texas) GetAllWinCates() map[int]string {
	return map[int]string {
		HighCard           : "高牌",
		OnePair            : "一对",
		TwoPair            : "两对",
		ThreeOfAKind       : "三条",
		Straight           : "顺子",
		Flush              : "同花",
		FullHouse          : "葫芦",
		FourOfAKind        : "四条",
		StraightFlush      : "同花顺",
		RoyalStraightFlush : "皇家同花顺",
	}
}