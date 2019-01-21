package BmResource

import (
	"errors"
	"net/http"
	"reflect"
	"strconv"

	"github.com/alfredyang1986/BmPods/BmDataStorage"
	"github.com/alfredyang1986/BmPods/BmModel"
	"github.com/manyminds/api2go"
)

type BmPostResource struct {
	BmPostStorage    *BmDataStorage.BmPostStorage
	BmCommentStorage *BmDataStorage.BmCommentStorage
}

func (s BmPostResource) NewPostResource(args []BmDataStorage.BmStorage) BmPostResource {
	var bs *BmDataStorage.BmPostStorage
	var is *BmDataStorage.BmCommentStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmPostStorage" {
			bs = arg.(*BmDataStorage.BmPostStorage)
		} else if tp.Name() == "BmCommentStorage" {
			is = arg.(*BmDataStorage.BmCommentStorage)
		} else {
		}
	}
	return BmPostResource{BmPostStorage: bs, BmCommentStorage: is}
}

// FindAll to satisfy api2go data source interface
func (s BmPostResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []BmModel.Post
	models := s.BmPostStorage.GetAll(r, -1, -1)

	for _, model := range models {
		// get all sweets for the model
		model.Comments = []*BmModel.Comment{}
		for _, kID := range model.CommentsIDs {
			choc, err := s.BmCommentStorage.GetOne(kID)
			if err != nil {
				return &Response{}, err
			}
			model.Comments = append(model.Comments, &choc)
		}

		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s BmPostResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []BmModel.Post
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
		for _, iter := range s.BmPostStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.BmPostStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := BmModel.Post{}
	count := s.BmPostStorage.Count(in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s BmPostResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.BmPostStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	model.Comments = []*BmModel.Comment{}
	for _, kID := range model.CommentsIDs {
		choc, err := s.BmCommentStorage.GetOne(kID)
		if err != nil {
			return &Response{}, err
		}
		model.Comments = append(model.Comments, &choc)
	}

	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s BmPostResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(BmModel.Post)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.BmPostStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s BmPostResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.BmPostStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s BmPostResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(BmModel.Post)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.BmPostStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
