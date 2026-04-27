package handler

import (
 "fmt"
 "net/http"
 

)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Bateu aqui")   /* O print manda a mnsagem para o terminal */
	fmt.Fprint(w, "Resposta enviada para o navegador!")   /*Fprint manda a mensagem para o destino, exemplo o navegador ou o pstman*/
}