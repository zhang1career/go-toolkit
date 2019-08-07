package ins

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/zhang1career/lib/ast"
	"log"
	"net/http"
)

type Ins struct {
}

func New() ast.Calculable {
	return &Ins{}
}

func (this *Ins) Calc() interface{} {
	url := fmt.Sprintf("/%s/%s?%s", this.source, this.cond)
	
	message := map[string]interface{}{
		"hello": "world",
		"life":  42,
		"embedded": map[string]string{
			"yes": "of course!",
		},
	}
	
	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}
	
	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}
	
	var result map[string]interface{}
	
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalln(err)
	}
	
	return result["data"]
}