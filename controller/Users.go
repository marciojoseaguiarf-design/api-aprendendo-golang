package controller

import (
	"api-aula-1/models"
	"api-aula-1/persistency"
	"api-aula-1/responses"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateUser")
	//ler body da requisição
	//fazer o unmarshal do body para um struct
	//fazer os tratamentos necessarios usando os metodos do user
	//Enviar para o repository
	//Preparar a resposta para o cliente

	//Lê o resquestbody e retorna um erro caso haja algum problema
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	//Descompacta (Unmarshal) o conteudo em JSON em uma Struct
	var newUser models.Users
	if err = json.Unmarshal(bodyRequest, &newUser); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	log.Println(newUser)
	//chama os metodos de preparação do user
	if err = newUser.Prepare("create"); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := persistency.Connect()
if err != nil {
	responses.Err(w, http.StatusInternalServerError, err)
	return
}
defer db.Close()

//Insere o usuário no banco de dados

	// Retorna o usuário criado
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(newUser); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
}

func FetchUser(w http.ResponseWriter, r *http.Request) {}

func UpdateUser(w http.ResponseWriter, r *http.Request) {}

func DeleteUser(w http.ResponseWriter, r *http.Request) {}
