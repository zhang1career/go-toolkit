package sel

import (
	"fmt"
	"github.com/zhang1career/lib/ast"
	"io/ioutil"
	"log"
	"net/http"
)

type Sel struct {
	target      ast.Valuable
	source      ast.Valuable
	cond        ast.Valuable
}

func New() ast.Calculable {
	return &Sel{}
}

func (this *Sel) Calc(map[string]ast.Valuable) interface{} {
	url := fmt.Sprintf("/%s/%s?%s", this.source, this.target, this.cond)
	
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	
	return string(body)
}