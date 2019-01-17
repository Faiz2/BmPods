package BmResource

import (
	"errors"
	"github.com/alfredyang1986/BmPods/BmDataStorage"
	"github.com/alfredyang1986/BmPods/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
	"strconv"
)

type BmReservableItemResource struct {
	BmReservableItemStorage *BmDataStorage.BmReservableItemStorage
	BmSessioninfoStorage    *BmDataStorage.BmSessioninfoStorage
}

func (s BmReservableItemResource) NewReservableItemResource(args []BmDataStorage.BmStorage) BmReservableItemResource {
	var us *BmDataStorage.BmReservableItemStorage
	var ts *BmDataStorage.BmSessioninfoStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmReservableItemStorage" {
			us = arg.(*BmDataStorage.BmReservableItemStorage)
		} else if tp.Name() == "BmSessioninfoStorage" {
			ts = arg.(*BmDataStorage.BmSessioninfoStorage)
		}
	}
	return BmReservableItemResource{BmReservableItemStorage: us, BmSessioninfoStorage: ts}
}

// FindAll to satisfy api2go data source interface
func (s BmReservableItemResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []BmModel.ReservableItem
	models := s.BmReservableItemStorage.GetAll(-1, -1)

	for _, model := range models {

		if model.SessioninfoID != "" {
			sessioninfo, err := s.BmSessioninfoStorage.GetOne(model.SessioninfoID)
			if err != nil {
				return &Response{}, err
			}
			model.Sessioninfo = sessioninfo
		}

		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s BmReservableItemResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []BmModel.ReservableItem
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
		for _, iter := range s.BmReservableItemStorage.GetAll(int(start), int(sizeI)) {
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

		for _, iter := range s.BmReservableItemStorage.GetAll(int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := BmModel.ReservableItem{}
	count := s.BmReservableItemStorage.Count(in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s BmReservableItemResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.BmReservableItemStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s BmReservableItemResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(BmModel.ReservableItem)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.BmReservableItemStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s BmReservableItemResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.BmReservableItemStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s BmReservableItemResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(BmModel.ReservableItem)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.BmReservableItemStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
