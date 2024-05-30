package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strconv"
)

const enem = "https://www.gov.br/inep/pt-br/areas-de-atuacao/avaliacao-e-exames-educacionais/enem/provas-e-gabaritos"
const fcmscsp = "https://vestibular.brasilescola.uol.com.br/downloads/faculdade-ciencias-medicas-santa-casa-sao-paulo.htm"

func main() {
	resp, err := http.Get(enem)
	if err != nil {
		fmt.Printf("error to search page %d", err)
	}
	defer resp.Body.Close()

	tokenizer := html.NewTokenizer(resp.Body)
	for {
		tokenType := tokenizer.Next()
		token := tokenizer.Token()
		if tokenType == html.ErrorToken {
			if tokenizer.Err() == io.EOF {
				return
			}
			fmt.Println(tokenizer.Err())
		}
		if tokenType == html.StartTagToken {
			if token.Data == "a" { //data-id="2023"
				tt := tokenizer.Next()
				tokenAno := tokenizer.Token()
				for i := 2023; i >= 2012; i-- {
					if tt == html.TextToken && tokenAno.Data == strconv.Itoa(i) {
						fmt.Println(tokenAno.Data)
					}
				}
			}
		}
	}

}
