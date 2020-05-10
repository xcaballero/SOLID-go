package main

import "fmt"

type MySQL struct {
}

func (db MySQL) QueryMySQL() []string {
	return []string{"alex", "john", "mike"}
}

type PostgreSQL struct {
}

func (db PostgreSQL) QueryPostgreSQL() map[string]string {
	return map[string]string{
		"a3f69c2b-d153-48fd-b10c-5b641657477b": "alex",
		"a4f69c2b-d153-48fd-b10c-5b641657477a": "john",
		"a5f69c2b-d153-48fd-b10c-5b641657477c": "mike",
	}
}

type UserRepository struct {
	db MySQL
	// db PostgreSQL
}

func (r UserRepository) GetUsers() []string {
	res := r.db.QueryMySQL()
	// res := r.db.QueryPostgreSQL()//
	// var users []string
	// for _, u := range res {
	// 	users = append(users, u)
	// }
	// return users
	return res
}

func main() {
	mysqlDB := MySQL{}
	// postgreSQLDB := PostgreSQL{}
	repo1 := UserRepository{db: mysqlDB}
	// repo2 := UserRepository{db: postgreSQLDB}
	fmt.Println(repo1.GetUsers())
	// fmt.Println(repo2.GetUsers())
}