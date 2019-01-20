package BmResource

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/alfredyang1986/BmPods/BmDataStorage"
	"github.com/alfredyang1986/BmPods/BmModel"
	"github.com/manyminds/api2go"
)

type BmRoomResource struct {
	BmRoomStorage *BmDataStorage.BmRoomStorage
	BmYardStorage *BmDataStorage.BmYardStorage
}

func (c BmRoomResource) NewRoomResource(args []BmDataStorage.BmStorage) BmRoomResource {
	var us *BmDataStorage.BmYardStorage
	var cs *BmDataStorage.BmRoomStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmYardStorage" {
			us = arg.(*BmDataStorage.BmYardStorage)
		} else if tp.Name() == "BmRoomStorage" {
			cs = arg.(*BmDataStorage.BmRoomStorage)
		}
	}
	return BmRoomResource{BmYardStorage: us, BmRoomStorage: cs}
}

// FindAll Rooms
func (c BmRoomResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	roomsID, ok := r.QueryParams["roomsID"]
	if ok {
		// this means that we want to show all rooms of a model, this is the route
		// /v0/models/1/rooms
		modelID := roomsID[0]
		// filter out rooms with modelID, in real world, you would just run a different database query
		filteredLeafs := []BmModel.Room{}
		model, err := c.BmYardStorage.GetOne(modelID)
		if err != nil {
			return &Response{}, err
		}
		for _, modelLeafID := range model.RoomsIDs {
			sweet, err := c.BmRoomStorage.GetOne(modelLeafID)
			if err != nil {
				return &Response{}, err
			}
			filteredLeafs = append(filteredLeafs, sweet)
		}

		return &Response{Res: filteredLeafs}, nil
	}
	rooms := c.BmRoomStorage.GetAll(r)
	return &Response{Res: rooms}, nil
}

// FindOne choc
func (c BmRoomResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmRoomStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmRoomResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Room)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmRoomStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmRoomResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmRoomStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmRoomResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Room)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmRoomStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
