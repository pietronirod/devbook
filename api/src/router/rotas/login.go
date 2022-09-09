package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasLogin = Rota{
	Uri:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}
