package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//StringConexaoBanco string de conexao com o banco de dados
	StringConexaoBanco = ""

	//Porta que a API vai estar rodando
	Porta = 0

	// SecretKey é a chaven que vai ser usada para assinar o token
	SecretKey []byte
)

// Carregar vai inicializar as variaveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_BANCO"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
