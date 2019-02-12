package ttree

import (
	"github.com/zhang1career/lib/datastruct/dim/bidim"
	"golang.org/x/sync/syncmap"
	"testing"
)

type dfsTest struct {
	lenX, lenY    int
	inbuf, outbuf [][]int
	start, stop   bidim.Coordinate
	dirs          []bidim.Coordinate
	arrivedNodes  syncmap.Map
}

func (s *dfsTest) index(node interface{}) bidim.Coordinate {
	return node.(bidim.Coordinate)
}

func (s *dfsTest) stepCount(b bidim.Coordinate) int {
	return s.outbuf[b.X][b.Y]
}

func (s *dfsTest) stepOn(b bidim.Coordinate, currentStep int) {
	s.outbuf[b.X][b.Y] = currentStep + 1
}

func (s *dfsTest) explore(node interface{}) (interface{}, error) {
	result := resultTest{}
	
	currentPos := s.index(node)
	currentStep := s.stepCount(currentPos)
	for _, dir := range s.dirs {
		nextPos := currentPos.Add(dir)
		if !s.available(&nextPos) {
			continue
		}
		s.stepOn(nextPos, currentStep)
		result.nodes = append(result.nodes, nextPos)
		result.value = s.outbuf
	}
	
	return &result, nil
}

func (s *dfsTest) parse(result interface{}) ([]interface{}, interface{}) {
	return result.(*resultTest).nodes, result.(*resultTest).value
}

func (s *dfsTest) available(b *bidim.Coordinate) bool {
	if b.X < 0 || b.X >= s.lenX || b.Y < 0 || b.Y >= s.lenY {
		return false
	}
	
	if s.isWall(b) || s.outbuf[b.X][b.Y]>0 {
		return false
	}
	return true
}

func (s *dfsTest) isWall(b *bidim.Coordinate) bool {
	if s.inbuf[b.X][b.Y] == 1 {
		return true
	}
	return false
}

func (s *dfsTest) root(node interface{}) bool {
	arrived, ok := s.arrivedNodes.Load(node)
	if !ok {
		return false
	}
	
	if arrived == true {
		return true
	}
	s.arrivedNodes.Store(node, true)
	return false
}

func (s *dfsTest) setStart(e bidim.Coordinate) {
	s.stepOn(e, 0)
}

func (s *dfsTest) setStop(e bidim.Coordinate) {
	s.stop = e
}

func (s *dfsTest) done(node interface{}) (done bool) {
	currentPos := s.index(node)
	if currentPos != s.stop {
		return false
	}
	return true
}

func TestDFS_Traverse(t *testing.T) {
	dt := dfsTest{
		lenX: 6,
		lenY: 5,
		inbuf: [][]int{
			{0, 1, 0, 0, 0},
			{0, 0, 0, 1, 0},
			{0, 1, 0, 1, 0},
			{1, 1, 1, 0, 0},
			{0, 1, 0, 0, 1},
			{0, 1, 0, 0, 0},
		},
		outbuf: [][]int{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		},
		start: bidim.Coordinate{0, 0},
		stop:  bidim.Coordinate{5, 4},
		dirs:  []bidim.Coordinate{{-1, 0}, {1, 0}, {0, -1}, {0, 1}},
		arrivedNodes: syncmap.Map{},
	}
	
	d := NewDFS(dt.explore, dt.parse, dt.root, dt.done)
	
	start := bidim.Coordinate{0, 0}
	dt.setStart(start)
	
	stop := bidim.Coordinate{5, 4}
	dt.setStop(stop)
	
	var input = []interface{}{start}
	var output []interface{}
	done, err := d.Traverse(input, &output)
	if err != nil {
		t.Error(err)
		return
	}
	if !done {
		t.Log("no solution")
		return
	}
	
	bidim.Printf("%3d\n", dt.outbuf)
	step := dt.outbuf[stop.X][stop.Y] - 1
	t.Logf("shortest step: %d\n", step)
}
