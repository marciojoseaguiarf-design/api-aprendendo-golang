package handler

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Bateu aqui")   /* O print manda a mnsagem para o terminal */
	fmt.Fprint(w, "Resposta enviada para o navegador!")   /*Fprint manda a mensagem para o destino, exemplo o navegador ou o pstman*/
	defer r.Body.Close()

	// Ler a string enviada
	b, _ := io.ReadAll(r.Body)

	fmt.Println(string(b))

	query := strings.TrimSpace(string(b)) //Trimspace remove espaços em branco do inicio ao fim de uma string
	googleURL := "https://www.googleapis.com/books/v1/volumes?q=" + url.QueryEscape(query)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	gReq, _ := http.NewRequestWithContext(ctx,http.MethodGet, googleURL, nil) //o "_" depois da variavel gReq e para ignorar o erro
	resp, _ := http.DefaultClient.Do(gReq)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)






}