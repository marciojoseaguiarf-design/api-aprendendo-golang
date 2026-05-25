package controller

import (
	"context"
	"encoding/json" 
	"io"
	"log" 
	"net/http"
	"net/url"
	"strings"
	"time"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	log.Println("Requisição recebida em /HandleSearch")

	// 1. Validar se o método é POST (já que espera dados no Body)
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "Método não permitido. Use POST"})
		return
	}

	// 2. Ler o corpo da requisição enviado pelo cliente
	b, err := io.ReadAll(r.Body)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "Erro ao ler o corpo da requisição"})
		return
	}
	defer r.Body.Close() // Fechar o corpo após a leitura bem-sucedida

	query := strings.TrimSpace(string(b))
	if query == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "O termo de pesquisa não pode estar vazio"})
		return
	}

	log.Printf("Termo pesquisado: %s\n", query)

	// 3. Montar o URL da Google Books API de forma segura
	googleURL := "https://www.googleapis.com/books/v1/volumes?q=" + url.QueryEscape(query)

	// 4. Criar o Contexto com Timeout de 10 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 5. Criar a requisição HTTP externa
	gReq, err := http.NewRequestWithContext(ctx, http.MethodGet, googleURL, nil)
	if err != nil {
		log.Println("Erro ao criar requisição para a Google:", err)
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno do servidor"})
		return
	}

	// 6. Executar a requisição externa
	resp, err := http.DefaultClient.Do(gReq)
	if err != nil {
		log.Println("Erro ao contactar a API da Google:", err)
		writeJSON(w, http.StatusBadGateway, map[string]string{"error": "Não foi possível contactar o serviço externo"})
		return
	}
	defer resp.Body.Close() // Obrigatório para requisições de cliente HTTP

	// 7. Ler a resposta devolvida pela Google
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro ao ler resposta externa"})
		return
	}

	// 8. Enviar a resposta final de volta ao navegador/Postman (Apenas UMA escrita autorizada)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(resp.StatusCode) // Repassa o status retornado pela Google (geralmente 200 OK)
	_, _ = w.Write(body)
}

// Função auxiliar corrigida (letra maiúscula em Header() e tratamento do JSON)
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("Erro ao codificar JSON de resposta:", err)
	}
}

