package BmResource

import (
	"github.com/alfredyang1986/BmPods/BmDataStorage"
	"github.com/manyminds/api2go"
	"errors"
	"net/http"
	"github.com/alfredyang1986/BmPods/BmModel"
	"reflect"
)

type BmModelLeafResourceExample struct {
	BmModelLeafStorageExample *BmDataStorage.BmModelLeafStorageExample
	BmModelStorageExample *BmDataStorage.BmModelStorageExample
}

func (c BmModelLeafResourceExample) NewModelLeafResource(args []BmDataStorage.BmStorage) BmModelLeafResourceExample {
	var us *BmDataStorage.BmModelStorageExample
	var cs *BmDataStorage.BmModelLeafStorageExample
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmModelStorageExample" {
			us = arg.(*BmDataStorage.BmModelStorageExample)
		} else if tp.Name() == "BmModelLeafStorageExample" {
			cs = arg.(*BmDataStorage.BmModelLeafStorageExample)
		}
	}
	return BmModelLeafResourceExample { BmModelStorageExample:us, BmModelLeafStorageExample:cs }
}

// FindAll modelleafs
func (c BmModelLeafResourceExample) FindAll(r api2go.Request) (api2go.Responder, error) {
	modelsID, ok := r.QueryParams["modelsID"]
	modelleafs := c.BmModelLeafStorageExample.GetAll()
	if ok {
		// this means that we want to show all modelleafs of a model, this is the route
		// /v0/models/1/modelleafs
		modelID := modelsID[0]
		// filter out modelleafs with modelID, in real world, you would just run a different database query
		filteredLeafs := []BmModel.ModelLeafExample{}
		model, err := c.BmModelStorageExample.GetOne(modelID)
		if err != nil {
			return &Response{}, err
		}
		for _, modelLeafID := range model.ModelLeafsIDs {
			sweet, err := c.BmModelLeafStorageExample.GetOne(modelLeafID)
			if err != nil {
				return &Response{}, err
			}
			filteredLeafs = append(filteredLeafs, sweet)
		}

		return &Response{Res: filteredLeafs}, nil
	}
	return &Response{Res: modelleafs}, nil
}

// FindOne choc
func (c BmModelLeafResourceExample) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmModelLeafStorageExample.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmModelLeafResourceExample) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.ModelLeafExample)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmModelLeafStorageExample.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmModelLeafResourceExample) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmModelLeafStorageExample.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmModelLeafResourceExample) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.ModelLeafExample)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmModelLeafStorageExample.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}