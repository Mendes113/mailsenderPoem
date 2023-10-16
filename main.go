package main

import (
	
	"fmt"
	
	"math/rand"
	"net/http"
	"net/smtp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	 
)





func main() {
	randomPoem, err := setupPoem()
	if err != nil {
		fmt.Println("Erro ao configurar o poema:", err)
		return
	}

	err = setupEmail(randomPoem)
	if err != nil {
		fmt.Println("Erro ao configurar o email:", err)
		return
	}


	
}

func sendMail(from, password, host string, port int, to []string, message []byte) error {
	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(fmt.Sprintf("%s:%d", host, port), auth, from, to, message)
	if err != nil {
		return err
	}

	return nil
}


func scrapePoems(url string) ([]string, error) {
	var poems []string

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Erro ao acessar a página. Código de status: %d", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}
	doc.Find(".")
	// Extrair os poemas
	doc.Find(".frase").Each(func(index int, item *goquery.Selection) {
		poemText := item.Text()
		poemText = strings.TrimSpace(poemText)
		poems = append(poems, poemText)
	})

	return poems, nil
}



func setupPoem() (string, error) {
	url := "https://www.pensador.com/poemas_de_amor/"

	poems, err := scrapePoems(url)
	if err != nil {
		return "", err
	}

	if len(poems) == 0 {
		return "", fmt.Errorf("Nenhum poema encontrado")
	}

	

	rand.Seed(time.Now().UnixNano())

	randomPoemIndex := rand.Intn(len(poems)) // Gera um índice aleatório para escolher um poema
	randomPoem := poems[randomPoemIndex]
	fmt.Println(randomPoem)
	return randomPoem, nil
}






func setupEmail(body string) error {
	from := ""
	password := "" // Sua senha do Elastic Email
	to := ""
	subject := "Poema do dia"
	

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := sendMail(from, password, "smtp.elasticemail.com", 2525, []string{to}, []byte(msg))
	if err != nil {
		fmt.Println("Erro ao enviar o email:", err)
	} else {
		fmt.Println("Email enviado com sucesso!")


	}

	return nil
}
