package BmResource

import (
	"github.com/manyminds/api2go"
	"github.com/alfredyang1986/BmPods/BmModel"
	"errors"
	"net/http"
	"github.com/alfredyang1986/BmPods/BmDataStorage"
	"reflect"
	"strconv"
)

type BmTeacherResource struct {
	BmTeacherStorage *BmDataStorage.BmTeacherStorage
}

func (s BmTeacherResource) NewTeacherResource(args []BmDataStorage.BmStorage) BmTeacherResource {
	var ss *BmDataStorage.BmTeacherStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmTeacherStorage" {
			ss = arg.(*BmDataStorage.BmTeacherStorage)
		}
	}
	return BmTeacherResource{BmTeacherStorage: ss}
}

// FindAll to satisfy api2go data source interface
func (s BmTeacherResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	result := s.BmTeacherStorage.GetAll(-1, -1)
	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load users in chunks
func (s BmTeacherResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []BmModel.Teacher
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
		for _, iter := range s.BmTeacherStorage.GetAll(int(start), int(sizeI)) {
			result = append(result, iter)
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

		for _, iter := range s.BmTeacherStorage.GetAll(int(offsetI), int(limitI)) {
			result = append(result, iter)
		}
	}

	in := BmModel.Teacher{}
	count := s.BmTeacherStorage.Count(in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the user with the given ID, otherwise an error
func (s BmTeacherResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	user, err := s.BmTeacherStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	return &Response{Res: user}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s BmTeacherResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	user, ok := obj.(BmModel.Teacher)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.BmTeacherStorage.Insert(user)
	user.ID = id

	return &Response{Res: user, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s BmTeacherResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.BmTeacherStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the user
func (s BmTeacherResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	user, ok := obj.(BmModel.Teacher)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.BmTeacherStorage.Update(user)
	return &Response{Res: user, Code: http.StatusNoContent}, err
}
