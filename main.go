package main

import (
	"Victor/gin-api-rest/database"
	"Victor/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()

}
