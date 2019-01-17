package BmResource

import (
	"github.com/alfredyang1986/BmPods/BmDataStorage"
	"github.com/manyminds/api2go"
	"errors"
	"net/http"
	"github.com/alfredyang1986/BmPods/BmModel"
	"reflect"
)

type BmKidResource struct {
	BmKidStorage *BmDataStorage.BmKidStorage
	//TODO:Replace Apply
	BmModelStorageExample *BmDataStorage.BmModelStorageExample
}

func (c BmKidResource) NewKidResource(args []BmDataStorage.BmStorage) BmKidResource {
	var us *BmDataStorage.BmModelStorageExample
	var cs *BmDataStorage.BmKidStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == BmDataStorage.ModelStorageName {
			us = arg.(*BmDataStorage.BmModelStorageExample)
		} else if tp.Name() == BmDataStorage.ModelLeafStorageName {
			cs = arg.(*BmDataStorage.BmKidStorage)
		}
	}
	return BmKidResource { BmModelStorageExample:us, BmKidStorage:cs }
}

// FindAll kids
func (c BmKidResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	kidsID, ok := r.QueryParams["kidsID"]
	kids := c.BmKidStorage.GetAll()
	if ok {
		// this means that we want to show all kids of a model, this is the route
		// /v0/models/1/kids
		modelID := kidsID[0]
		// filter out kids with modelID, in real world, you would just run a different database query
		filteredLeafs := []BmModel.Kid{}
		model, err := c.BmModelStorageExample.GetOne(modelID)
		if err != nil {
			return &Response{}, err
		}
		for _, modelLeafID := range model.ModelLeafsIDs {
			sweet, err := c.BmKidStorage.GetOne(modelLeafID)
			if err != nil {
				return &Response{}, err
			}
			filteredLeafs = append(filteredLeafs, sweet)
		}

		return &Response{Res: filteredLeafs}, nil
	}
	return &Response{Res: kids}, nil
}

// FindOne choc
func (c BmKidResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmKidStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmKidResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Kid)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmKidStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmKidResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmKidStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmKidResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Kid)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmKidStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}