package BmResource

import (
	"github.com/alfredyang1986/BmPods/BmDataStorage"
	"github.com/manyminds/api2go"
	"errors"
	"net/http"
	"github.com/alfredyang1986/BmPods/BmModel"
	"reflect"
)

type BmChocolateResource struct {
	ChocStorage *BmDataStorage.ChocolateStorage
	UserStorage *BmDataStorage.UserStorage
}

func (c BmChocolateResource) NewChocolateResource(args ...interface{}) BmChocolateResource {
	var us *BmDataStorage.UserStorage
	var cs *BmDataStorage.ChocolateStorage
	sds := args[0].([]BmDataStorage.BmStorage)
	for _, arg := range sds {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UserStorage" {
			us = arg.(*BmDataStorage.UserStorage)
		} else if tp.Name() == "ChocolateStorage" {
			cs = arg.(*BmDataStorage.ChocolateStorage)
		}
	}
	return BmChocolateResource { UserStorage:us, ChocStorage:cs }



	return BmChocolateResource{}
}

// FindAll chocolates
func (c BmChocolateResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	usersID, ok := r.QueryParams["usersID"]
	sweets := c.ChocStorage.GetAll()
	if ok {
		// this means that we want to show all sweets of a user, this is the route
		// /v0/users/1/sweets
		userID := usersID[0]
		// filter out sweets with userID, in real world, you would just run a different database query
		filteredSweets := []BmModel.Chocolate{}
		user, err := c.UserStorage.GetOne(userID)
		if err != nil {
			return &Response{}, err
		}
		for _, sweetID := range user.ChocolatesIDs {
			sweet, err := c.ChocStorage.GetOne(sweetID)
			if err != nil {
				return &Response{}, err
			}
			filteredSweets = append(filteredSweets, sweet)
		}

		return &Response{Res: filteredSweets}, nil
	}
	return &Response{Res: sweets}, nil
}

// FindOne choc
func (c BmChocolateResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.ChocStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmChocolateResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Chocolate)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.ChocStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmChocolateResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.ChocStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmChocolateResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Chocolate)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.ChocStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}