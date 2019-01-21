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

type BmCommentResource struct {
	BmCommentStorage *BmDataStorage.BmCommentStorage
	BmPostStorage    *BmDataStorage.BmPostStorage
}

func (s BmCommentResource) NewCommentResource(args []BmDataStorage.BmStorage) BmCommentResource {
	var bs *BmDataStorage.BmCommentStorage
	var cs *BmDataStorage.BmPostStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmCommentStorage" {
			bs = arg.(*BmDataStorage.BmCommentStorage)
		} else if tp.Name() == "BmPostStorage" {
			cs = arg.(*BmDataStorage.BmPostStorage)
		} else {
		}
	}
	return BmCommentResource{BmCommentStorage: bs, BmPostStorage: cs}
}

func (c BmCommentResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	postsID, ok := r.QueryParams["postsID"]
	posts := c.BmCommentStorage.GetAll(r, -1, -1)
	if ok {
		// this means that we want to show all kids of a model, this is the route
		// /v0/models/1/kids
		modelID := postsID[0]
		// filter out kids with modelID, in real world, you would just run a different database query
		filteredLeafs := []BmModel.Comment{}
		model, err := c.BmPostStorage.GetOne(modelID)
		if err != nil {
			return &Response{}, err
		}
		for _, modelLeafID := range model.CommentsIDs {
			sweet, err := c.BmCommentStorage.GetOne(modelLeafID)
			if err != nil {
				return &Response{}, err
			}
			filteredLeafs = append(filteredLeafs, sweet)
		}

		return &Response{Res: filteredLeafs}, nil
	}
	return &Response{Res: posts}, nil
}

// FindAll to satisfy api2go data source interface
// func (s BmCommentResource) FindAll(r api2go.Request) (api2go.Responder, error) {
// 	var result []BmModel.Comment
// 	models := s.BmCommentStorage.GetAll(r, -1, -1)

// 	for _, model := range models {

// 		if model.PostID != "" {
// 			post, err := s.BmPostStorage.GetOne(model.PostID)
// 			if err != nil {
// 				return &Response{}, err
// 			}
// 			model.Post = &post
// 		}

// 		result = append(result, *model)
// 	}

// 	return &Response{Res: result}, nil
// }

// PaginatedFindAll can be used to load models in chunks
func (s BmCommentResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []BmModel.Comment
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
		for _, iter := range s.BmCommentStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.BmCommentStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := BmModel.Comment{}
	count := s.BmCommentStorage.Count(in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s BmCommentResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.BmCommentStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s BmCommentResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(BmModel.Comment)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.BmCommentStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s BmCommentResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.BmCommentStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s BmCommentResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(BmModel.Comment)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.BmCommentStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
