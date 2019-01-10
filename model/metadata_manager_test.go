package model

import (
	"tantan-simplify/config"
	"tantan-simplify/pkg/constant"
	"testing"
)

var user0 Users
var user1 Users

func TestDB(t *testing.T)  {
	Setup()
	t.Run("TestAddUser", TestAddUser)
	t.Run("TestListAllUsers", TestListAllUsers)
	t.Run("TestUpdateRelationship", TestUpdateRelationshipMatch)
	t.Run("TestListRelationships", TestListRelationships)
	TearDown()
}

func Setup()  {
	postgreSQLConfig := config.PostgreSQL{
		User:"xzy",
		Password:"",
		Addr:"127.0.0.1:5432",
		DBName:"tantan",
	}
	MustSetDB(&postgreSQLConfig)
}

func TestAddUser(t *testing.T) {
	_, err := AddUser("user0")
	_, err = AddUser("user1")
	if err != nil{
		t.Fatalf("Test add user failed, error: %s", err)
	}
}

func TestListAllUsers(t *testing.T) {
	users, err := ListAllUsers()
	if err != nil{
		t.Fatalf("Test list all users failed, error: %s", err)
	}
	user0 = users[0]
	user1 = users[1]
}

func TestUpdateRelationshipMatch(t *testing.T) {
	_, err := UpdateRelationship(user1.Id, user0.Id, constant.Liked)
	_, err = UpdateRelationship(user0.Id, user1.Id, constant.Liked)
	if err != nil {
		t.Fatalf("Test update relationship failed, error: %s", err)
	}
}

func TestListRelationships(t *testing.T) {
	_, err := ListRelationships(user0.Id)
	if err != nil{
		t.Fatalf("Test list relationships failed, error: %s", err)
	}
}

func TearDown(){
	var relationship Relationships
	_, err := db.QueryOne(&relationship, `select id, from_user_id, to_user_id, state from relationships where from_user_id=? and to_user_id=?`, user0.Id, user1.Id)
	if err == nil{
		err = db.Delete(&relationship)
	}
	_, err = db.QueryOne(&relationship, `select id, from_user_id, to_user_id, state from relationships where from_user_id=? and to_user_id=?`, user1.Id, user0.Id)
	if err == nil{
		err = db.Delete(&relationship)
	}
	err = db.Delete(&user0)
	err = db.Delete(&user1)
}