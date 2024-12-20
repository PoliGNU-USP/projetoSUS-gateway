package models

// TODO: Tarefa do banco de dados dos usuários definir os parâmetros aqui
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
