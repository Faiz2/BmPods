package BmResource

import (
	"github.com/alfredyang1986/BmPods/BmDataStorage"
	"github.com/manyminds/api2go"
	"errors"
	"net/http"
	"strconv"
	"github.com/alfredyang1986/BmPods/BmModel"
	"reflect"
)

type BmUserResource struct {
	ChocStorage *BmDataStorage.ChocolateStorage
	UserStorage *BmDataStorage.UserStorage
}

func (s BmUserResource) NewUserResource(args []BmDataStorage.BmStorage) BmUserResource {
	var us *BmDataStorage.UserStorage
	var cs *BmDataStorage.ChocolateStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UserStorage" {
			us = arg.(*BmDataStorage.UserStorage)
		} else if tp.Name() == "ChocolateStorage" {
			cs = arg.(*BmDataStorage.ChocolateStorage)
		}
	}
	return BmUserResource{ UserStorage:us, ChocStorage:cs }
}

// FindAll to satisfy api2go data source interface
func (s BmUserResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []BmModel.User
	users := s.UserStorage.GetAll(-1, -1)

	for _, user := range users {
		// get all sweets for the user
		user.Chocolates = []*BmModel.Chocolate{}
		for _, chocolateID := range user.ChocolatesIDs {
			choc, err := s.ChocStorage.GetOne(chocolateID)
			if err != nil {
				return &Response{}, err
			}
			user.Chocolates = append(user.Chocolates, &choc)
		}
		result = append(result, *user)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load users in chunks
func (s BmUserResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []BmModel.User
		number, size, offset, limit string
	)

	numberQuery, ok := r.QueryParams["page[number]"]
	if ok {
		number = numberQuery[0]
	}
	sizeQuery, ok := r.QueryParams["page[size]"]
	if ok {
		size = sizeQuery[0]
	}
	offsetQuery, ok := r.QueryParams["page[offset]"]
	if ok {
		offset = offsetQuery[0]
	}
	limitQuery, ok := r.QueryParams["page[limit]"]
	if ok {
		limit = limitQuery[0]
	}

	if size != "" {
		sizeI, err := strconv.ParseInt(size, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		numberI, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		start := sizeI * (numberI - 1)
		for _, iter := range s.UserStorage.GetAll(int(start), int(sizeI)) {
			result = append(result, *iter)
		}

	} else {
		limitI, err := strconv.ParseUint(limit, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		offsetI, err := strconv.ParseUint(offset, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		for _, iter := range s.UserStorage.GetAll(int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := BmModel.User{}
	count := s.UserStorage.Count(in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the user with the given ID, otherwise an error
func (s BmUserResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	user, err := s.UserStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	user.Chocolates = []*BmModel.Chocolate{}
	for _, chocolateID := range user.ChocolatesIDs {
		choc, err := s.ChocStorage.GetOne(chocolateID)
		if err != nil {
			return &Response{}, err
		}
		user.Chocolates = append(user.Chocolates, &choc)
	}
	return &Response{Res: user}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s BmUserResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	user, ok := obj.(BmModel.User)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.UserStorage.Insert(user)
	user.ID = id

	return &Response{Res: user, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s BmUserResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.UserStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the user
func (s BmUserResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	user, ok := obj.(BmModel.User)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.UserStorage.Update(user)
	return &Response{Res: user, Code: http.StatusNoContent}, err
}