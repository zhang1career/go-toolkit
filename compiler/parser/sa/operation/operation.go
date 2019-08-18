package operation

import (
	"github.com/zhang1career/lib/datastruct/stack"
	"github.com/zhang1career/lib/compiler"
)

var defaultSyntaxMap = map[string]string {
}

type Sa struct {
	syntaxMap map[string]string
	wordStack *stack.Stack
}

func New() *Sa {
	return &Sa{
		syntaxMap: defaultSyntaxMap,
		wordStack: stack.New(),
	}
}



func (this *Sa) Analyze(words []string) compiler.Dim {
	ret := make(compiler.Dim)
	
	for _, word := range words {
		
		syntax, ok := this.syntaxMap[word]
		if !ok {
			this.wordStack.Push(word)
			continue
		}
		
		s := syntax.New()
		k := s.GetValue()
		
		if s.PreCount() {
			v := this.wordStack.Pop()
		}
		
		if s.PostCount() {
			v := this.wordStack.Pop()
		}
		
	}
	
	return ret
}