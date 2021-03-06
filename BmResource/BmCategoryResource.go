package BmResource

import (
	"errors"
	"github.com/alfredyang1986/BmPods/BmDataStorage"
	"github.com/alfredyang1986/BmPods/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type BmCategoryResource struct {
	CategoryStorage *BmDataStorage.BmCategoryStorage
}

func (c BmCategoryResource) NewCategoryResource(args []BmDataStorage.BmStorage) BmCategoryResource {
	var as *BmDataStorage.BmCategoryStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmCategoryStorage" {
			as = arg.(*BmDataStorage.BmCategoryStorage)
		}
	}
	return BmCategoryResource{CategoryStorage: as}
}

// FindAll apeolates
func (c BmCategoryResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	result := c.CategoryStorage.GetAll(r,-1, -1)
	return &Response{Res: result}, nil
}

func (s BmCategoryResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	result := []BmModel.Category{}
	return 100, &Response{Res: result}, nil
}

// FindOne ape
func (c BmCategoryResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.CategoryStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new ape
func (c BmCategoryResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	ape, ok := obj.(BmModel.Category)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.CategoryStorage.Insert(ape)
	ape.ID = id
	return &Response{Res: ape, Code: http.StatusCreated}, nil
}

// Delete a ape :(
func (c BmCategoryResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.CategoryStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a ape
func (c BmCategoryResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	ape, ok := obj.(BmModel.Category)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.CategoryStorage.Update(ape)
	return &Response{Res: ape, Code: http.StatusNoContent}, err
}
