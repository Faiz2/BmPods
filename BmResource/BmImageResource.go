package BmResource

import (
	"errors"
	"github.com/alfredyang1986/BmPods/BmDataStorage"
	"github.com/alfredyang1986/BmPods/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type BmImageResource struct {
	BmImageStorage   *BmDataStorage.BmImageStorage
	BmSessioninfoStorage *BmDataStorage.BmSessioninfoStorage
}

func (c BmImageResource) NewImageResource(args []BmDataStorage.BmStorage) BmImageResource {
	var us *BmDataStorage.BmSessioninfoStorage
	var cs *BmDataStorage.BmImageStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmSessioninfoStorage" {
			us = arg.(*BmDataStorage.BmSessioninfoStorage)
		} else if tp.Name() == "BmImageStorage" {
			cs = arg.(*BmDataStorage.BmImageStorage)
		}
	}
	return BmImageResource{BmSessioninfoStorage: us, BmImageStorage: cs}
}

// FindAll images
func (c BmImageResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	imagesID, ok := r.QueryParams["imagesID"]
	images := c.BmImageStorage.GetAll()
	if ok {
		// this means that we want to show all images of a model, this is the route
		// /v0/models/1/images
		modelID := imagesID[0]
		// filter out images with modelID, in real world, you would just run a different database query
		filteredLeafs := []BmModel.Image{}
		model, err := c.BmSessioninfoStorage.GetOne(modelID)
		if err != nil {
			return &Response{}, err
		}
		for _, modelLeafID := range model.ImagesIDs {
			sweet, err := c.BmImageStorage.GetOne(modelLeafID)
			if err != nil {
				return &Response{}, err
			}
			filteredLeafs = append(filteredLeafs, sweet)
		}

		return &Response{Res: filteredLeafs}, nil
	}
	return &Response{Res: images}, nil
}

// FindOne choc
func (c BmImageResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmImageStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmImageResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Image)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmImageStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmImageResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmImageStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmImageResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Image)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmImageStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
