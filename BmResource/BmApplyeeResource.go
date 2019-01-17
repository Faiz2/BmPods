package BmResource

import (
	"net/http"
	"github.com/manyminds/api2go"
	"github.com/alfredyang1986/BmPods/BmModel"
	"errors"
	"github.com/alfredyang1986/BmPods/BmDataStorage"
	"reflect"
)

type BmApplyeeResource struct {
	ApplyeeStorage *BmDataStorage.BmApplyeeStorage
}

func (c BmApplyeeResource) NewApplyeeResource(args []BmDataStorage.BmStorage) BmApplyeeResource {
	var as *BmDataStorage.BmApplyeeStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmApplyeeStorage" {
			as = arg.(*BmDataStorage.BmApplyeeStorage)
		}
	}
	return BmApplyeeResource{ ApplyeeStorage: as }
}

// FindAll apeolates
func (c BmApplyeeResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	//usersID, ok := r.QueryParams["usersID"]
	sweets := c.ApplyeeStorage.GetAll(-1, -1)
	//if ok {
	//	// this means that we want to show all sweets of a user, this is the route
	//	// /v0/users/1/sweets
	//	userID := usersID[0]
	//	// filter out sweets with userID, in real world, you would just run a different database query
	//	filteredSweets := []BmModel.Applyee{}
	//	user, err := c.UserStorage.GetOne(userID)
	//	if err != nil {
	//		return &Response{}, err
	//	}
	//	for _, sweetID := range user.ApplyeesIDs {
	//		sweet, err := c.ApplyeeStorage.GetOne(sweetID)
	//		if err != nil {
	//			return &Response{}, err
	//		}
	//		filteredSweets = append(filteredSweets, sweet)
	//	}
	//
	//	return &Response{Res: filteredSweets}, nil
	//}
	return &Response{Res: sweets}, nil
}

func (s BmApplyeeResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	result := []BmModel.Applyee{}
	return 100, &Response{Res: result}, nil
}

// FindOne ape
func (c BmApplyeeResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.ApplyeeStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new ape
func (c BmApplyeeResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	ape, ok := obj.(BmModel.Applyee)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.ApplyeeStorage.Insert(ape)
	ape.ID = id
	return &Response{Res: ape, Code: http.StatusCreated}, nil
}

// Delete a ape :(
func (c BmApplyeeResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.ApplyeeStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a ape
func (c BmApplyeeResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	ape, ok := obj.(BmModel.Applyee)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.ApplyeeStorage.Update(ape)
	return &Response{Res: ape, Code: http.StatusNoContent}, err
}