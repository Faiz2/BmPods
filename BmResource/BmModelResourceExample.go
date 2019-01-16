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

type BmModelResourceExample struct {
	BmModelLeafStorageExample *BmDataStorage.BmModelLeafStorageExample
	BmModelStorageExample *BmDataStorage.BmModelStorageExample
}

func (s BmModelResourceExample) NewModelResource(args []BmDataStorage.BmStorage) BmModelResourceExample {
	var us *BmDataStorage.BmModelStorageExample
	var cs *BmDataStorage.BmModelLeafStorageExample
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == BmDataStorage.ModelStorageName {
			us = arg.(*BmDataStorage.BmModelStorageExample)
		} else if tp.Name() == BmDataStorage.ModelLeafStorageName {
			cs = arg.(*BmDataStorage.BmModelLeafStorageExample)
		}
	}
	return BmModelResourceExample{ BmModelStorageExample:us, BmModelLeafStorageExample:cs }
}

// FindAll to satisfy api2go data source interface
func (s BmModelResourceExample) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []BmModel.ModelExample
	models := s.BmModelStorageExample.GetAll(-1, -1)

	for _, model := range models {
		// get all sweets for the model
		model.ModelLeafs = []*BmModel.ModelLeafExample{}
		for _, modelLeafID := range model.ModelLeafsIDs {
			choc, err := s.BmModelLeafStorageExample.GetOne(modelLeafID)
			if err != nil {
				return &Response{}, err
			}
			model.ModelLeafs = append(model.ModelLeafs, &choc)
		}
		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s BmModelResourceExample) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []BmModel.ModelExample
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
		for _, iter := range s.BmModelStorageExample.GetAll(int(start), int(sizeI)) {
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

		for _, iter := range s.BmModelStorageExample.GetAll(int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := BmModel.ModelExample{}
	count := s.BmModelStorageExample.Count(in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s BmModelResourceExample) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.BmModelStorageExample.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	model.ModelLeafs = []*BmModel.ModelLeafExample{}
	for _, modelLeafID := range model.ModelLeafsIDs {
		choc, err := s.BmModelLeafStorageExample.GetOne(modelLeafID)
		if err != nil {
			return &Response{}, err
		}
		model.ModelLeafs = append(model.ModelLeafs, &choc)
	}
	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s BmModelResourceExample) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(BmModel.ModelExample)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.BmModelStorageExample.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s BmModelResourceExample) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.BmModelStorageExample.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s BmModelResourceExample) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(BmModel.ModelExample)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.BmModelStorageExample.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}