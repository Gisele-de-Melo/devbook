package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

//Função foi utilizada para gerar o haskey e blockkey do arquivo .env
// func init() {
// 	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	fmt.Println(hashKey)
// 	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	fmt.Println(blockKey)
// }

func main() {
	utils.CarregarTemplates()
	config.Carregar()
	cookies.Configurar()
	r := router.Gerar()
	fmt.Println("Rodando WebApp!")
	log.Fatal(http.ListenAndServe(":3000", r))

}
