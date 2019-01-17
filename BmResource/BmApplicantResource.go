package BmResource

import (
	"errors"
	"github.com/alfredyang1986/BmPods/BmDataStorage"
	"github.com/alfredyang1986/BmPods/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type BmApplicantResource struct {
	ApplicantStorage *BmDataStorage.BmApplicantStorage
}

func (c BmApplicantResource) NewApplicantResource(args []BmDataStorage.BmStorage) BmApplicantResource {
	var as *BmDataStorage.BmApplicantStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmApplicantStorage" {
			as = arg.(*BmDataStorage.BmApplicantStorage)
		}
	}
	return BmApplicantResource{ApplicantStorage: as}
}

// FindAll apeolates
func (c BmApplicantResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	//usersID, ok := r.QueryParams["usersID"]
	result := c.ApplicantStorage.GetAll(-1, -1)
	//if ok {
	//	// this means that we want to show all sweets of a user, this is the route
	//	// /v0/users/1/sweets
	//	userID := usersID[0]
	//	// filter out sweets with userID, in real world, you would just run a different database query
	//	filteredSweets := []BmModel.Applicant{}
	//	user, err := c.UserStorage.GetOne(userID)
	//	if err != nil {
	//		return &Response{}, err
	//	}
	//	for _, sweetID := range user.ApplicantsIDs {
	//		sweet, err := c.ApplicantStorage.GetOne(sweetID)
	//		if err != nil {
	//			return &Response{}, err
	//		}
	//		filteredSweets = append(filteredSweets, sweet)
	//	}
	//
	//	return &Response{Res: filteredSweets}, nil
	//}
	return &Response{Res: result}, nil
}

func (s BmApplicantResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	result := []BmModel.Applicant{}
	return 100, &Response{Res: result}, nil
}

// FindOne ape
func (c BmApplicantResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.ApplicantStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new ape
func (c BmApplicantResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	ape, ok := obj.(BmModel.Applicant)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.ApplicantStorage.Insert(ape)
	ape.ID = id
	return &Response{Res: ape, Code: http.StatusCreated}, nil
}

// Delete a ape :(
func (c BmApplicantResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.ApplicantStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a ape
func (c BmApplicantResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	ape, ok := obj.(BmModel.Applicant)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.ApplicantStorage.Update(ape)
	return &Response{Res: ape, Code: http.StatusNoContent}, err
}
