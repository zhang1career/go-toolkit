package sel

import (
	"fmt"
	"github.com/zhang1career/lib/ast"
)

type Sel struct {
}

func New() ast.Calculable {
	return &Sel{}
}

func (this *Sel) Calc(params []ast.Valuable) interface{} {
	url := fmt.Sprintf("/%s/%s?%s", params[1].Evaluate(), params[0].Evaluate(), params[2].Evaluate())
	return url
	//resp, err := http.Get(url)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//return string(body)
}