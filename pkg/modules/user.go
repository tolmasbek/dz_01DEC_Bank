package modules

import (
	"bank-t/database"
	"database/sql"
	"fmt"
)

type User struct {
	Id        int64
	FirstName string
	LastName  string
	Age       int64
	Gender    string
	Login     string
	Password  string
	Role      string
	Remove    bool
}

func AddNewUser(db *sql.DB, user User) (ok bool, err error){
	_, err = db.Exec(database.AddNewUser, user.FirstName, user.LastName, user.Age, user.Gender, user.Login, user.Password, user.Role, user.Remove)
	if err != nil {
		fmt.Println("Can't insert to ATMs table new address, err is", err)
		return false, err
	}
	return true, nil
}

func ShowUser (database *sql.DB, login, password string){
	var UserA User
	_ = database.QueryRow(`Select *From users Where login=($1) and password=($2)`, login, password).Scan(
		&UserA.Id,
		&UserA.FirstName,
		&UserA.LastName,
		&UserA.Age,
		&UserA.Gender,
		&UserA.Login,
		&UserA.Password,
		&UserA.Role,
		&UserA.Remove)
	fmt.Printf("%s", UserA)
}

