package main

import(
	"github.com/google/uuid"
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"fmt"
)

type Pessoa struct{

	ID string
	Nome string
}


func main(){

	p := Pessoa{
		ID: uuid.New().String(),
		Nome: "Joao",


		//Abrindo conexão com o banco
	}
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/aula" )
	if err != nil{
		log.Println(err)
	}

	log.Println("Conexão aberta com sql")


	defer db.Close()


	//Inserindo dados no banco
	 Inserir_dados, err := db.Prepare("insert into aula (id, nome) values(?,?)")
	 if err != nil{
	 	log.Println(err)
	 }

	 log.Println("Inserindo dados no banco....")

	 _, err = Inserir_dados.Exec(p.ID, p.Nome)


	 defer Inserir_dados.Close()


	 //Deletando usuarios pelo ID

	deletando_dados,_ := db.Prepare("delete from aula where id = ?")
	deletando_dados.Exec("6c4f6ee2-3fcd-4482-8f86-7e5f721fba41")
	log.Println("Apagando dado")


	// Consulta SQL para verificar se o usuário existe

	username := "joao"


	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM aula WHERE nome = ?", username).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count > 0 {
		fmt.Printf("O usuário '%s' existe no banco de dados.\n", username)
	} else {
		fmt.Printf("O usuário '%s' não existe no banco de dados.\n", username)
	}
}

