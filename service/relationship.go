package service

import (
	"strconv"
	"tantan-simplify/model"
	"tantan-simplify/pkg/constant"
	"tantan-simplify/types"
)

type RelationshipHandler interface {
	ListRelationships() ([]types.Relationship, error)
	UpdateRelationship(toUserId int, state constant.State) (types.Relationship, error)
}

type RelationshipHandlerImpl struct {
	userId int
}

func GetRelationshipHandler(userId int) RelationshipHandler {
	return &RelationshipHandlerImpl{userId:userId}
}

func (rh *RelationshipHandlerImpl) ListRelationships() ([]types.Relationship, error) {
	relationships, err := model.ListRelationships(rh.userId)
	if err != nil {
		return nil, err
	}
	relationshipsResult := make([]types.Relationship, len(relationships))
	for k, relationship := range relationships {
		relationshipsResult[k].UserId = strconv.Itoa(relationship.ToUserId)
		relationshipsResult[k].State = relationship.State.String()
		relationshipsResult[k].Type = constant.RelationshipStr
	}
	return relationshipsResult, nil
}

func (rh *RelationshipHandlerImpl) UpdateRelationship(toUserId int, state constant.State) (types.Relationship, error) {
	var relationshipResult types.Relationship
	relationship, err := model.UpdateRelationship(rh.userId, toUserId, state)
	if err != nil{
		return relationshipResult, err
	}
	relationshipResult.UserId = strconv.Itoa(relationship.ToUserId)
	relationshipResult.State = relationship.State.String()
	relationshipResult.Type = constant.RelationshipStr
	return relationshipResult, nil
}