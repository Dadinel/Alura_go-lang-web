package controllers

import (
	"log"
	"models"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	// produtos := []Produto{
	// 	{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quatidade: 5},
	// 	{"Tênis", "Confortável", 89, 3},
	// 	{"Fone", "Muito", 54, 5},
	// }
	todosOsProdutos := models.BuscaTodosOsProdutos()

	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversão do preço: ", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversão da quantidade: ", err)
		}

		models.CriarNovoProduto(nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}
	goToIndex(w, r)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := getIdFromUrl(r)
	models.DeletaProduto(idDoProduto)
	goToIndex(w, r)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := getIdFromUrl(r)
	produto := models.BuscaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertidoParaInt, err := strconv.Atoi(id)

		if err != nil {
			log.Println("Erro na conversão do id: ", err)
		}

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversão do preço: ", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversão da quantidade: ", err)
		}

		models.AtualizaProduto(idConvertidoParaInt, nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}
	goToIndex(w, r)
}

func goToIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", 301)
}

func getIdFromUrl(r *http.Request) string {
	return r.URL.Query().Get("id")
}
