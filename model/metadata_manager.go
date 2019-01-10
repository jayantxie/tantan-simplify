package model

import (
	"tantan-simplify/errors"
	"tantan-simplify/pkg/constant"
)

type Users struct {
	Id   int    `sql:"id"`
	Name string `sql:"name"`
}

type Relationships struct {
	Id         int            `sql:"id"`
	FromUserId int            `sql:"from_user_id"`
	ToUserId   int            `sql:"to_user_id"`
	State      constant.State `sql:"state"` // 1: liked    2: disliked    3: matched
}

func AddUser(name string) (Users, error) {
	user := Users{Name: name}
	err := db.Insert(&user)
	return user, err
}

func ListAllUsers() (users []Users, err error) {
	_, err = db.Query(&users, `select id, name from users`)
	return users, err
}

func ListRelationships(fromUserId int) (relationships []Relationships, err error) {
	_, err = db.Query(&relationships, `select from_user_id, to_user_id, state from relationships where from_user_id=?`, fromUserId)
	return relationships, err
}

// row lock
func UpdateRelationship(fromUserId, toUserId int, state constant.State) (relationshipResult Relationships, err error) {
	// check id illegal
	var user Users
	_, err = db.QueryOne(&user, `select id, name from users where id=?`, fromUserId)
	if err != nil {
		return relationshipResult, errors.ErrUserNotFound
	}
	_, err = db.QueryOne(&user, `select id, name from users where id=?`, toUserId)
	if err != nil {
		return relationshipResult, errors.ErrUserNotFound
	}

	// update relationship
	var relationship Relationships
	var reversedRelationship Relationships
	reversedLiked := false
	tx, err := db.Begin()
	if err != nil {
		return relationshipResult, err
	}
	_, err = tx.QueryOne(&reversedRelationship, `select id, from_user_id, to_user_id, state from relationships where from_user_id=? and to_user_id=? for update`, toUserId, fromUserId)
	if err == nil && reversedRelationship.State == constant.Liked {
		reversedLiked = true
	}
	_, err = db.QueryOne(&relationship, `select id, from_user_id, to_user_id, state from relationships where from_user_id=? and to_user_id=?`, fromUserId, toUserId)
	if err != nil {
		if reversedLiked && state == constant.Liked {
			reversedRelationship.State = constant.Matched
			err = tx.Update(&reversedRelationship)
			if err != nil {
				err1 := tx.Rollback()
				if err1 != nil {
					return relationshipResult, err1
				} else {
					return relationshipResult, err
				}
			}
			relationship = Relationships{
				FromUserId: fromUserId,
				ToUserId:   toUserId,
				State:      constant.Matched,
			}
		} else {
			relationship = Relationships{
				FromUserId: fromUserId,
				ToUserId:   toUserId,
				State:      state,
			}
		}
		err = tx.Insert(&relationship)
		if err != nil {
			err1 := tx.Rollback()
			if err1 != nil {
				return relationshipResult, err1
			} else {
				return relationshipResult, err
			}
		}
		err = tx.Commit()
		return relationship, err
	}
	err = errors.ErrOperated
	return relationshipResult, err
}
