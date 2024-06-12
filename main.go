package main

import (
	"github.com/marcelosilvadev/gopportunities.git/config"
	"github.com/marcelosilvadev/gopportunities.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger := config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errorf("Erro ao conectar ao banco de dados: %v", err)
		return
	}

	//Inicia Router
	router.Initialize()
}
