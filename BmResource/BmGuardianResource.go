package BmResource

import (
	"errors"
	"github.com/alfredyang1986/BmPods/BmDataStorage"
	"github.com/alfredyang1986/BmPods/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type BmGuardianResource struct {
	BmGuardianStorage   *BmDataStorage.BmGuardianStorage
}

func (c BmGuardianResource) NewGuardianResource(args []BmDataStorage.BmStorage) BmGuardianResource {
	var cs *BmDataStorage.BmGuardianStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmGuardianStorage" {
			cs = arg.(*BmDataStorage.BmGuardianStorage)
		}
	}
	return BmGuardianResource{BmGuardianStorage: cs}
}

// FindAll guardians
func (c BmGuardianResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	//guardiansID, ok := r.QueryParams["guardiansID"]
	guardians := c.BmGuardianStorage.GetAll()
	//if ok {
	//	// this means that we want to show all guardians of a model, this is the route
	//	// /v0/models/1/guardians
	//	modelID := guardiansID[0]
	//	// filter out guardians with modelID, in real world, you would just run a different database query
	//	filteredLeafs := []BmModel.Guardian{}
	//	model, err := c.BmApplyStorage.GetOne(modelID)
	//	if err != nil {
	//		return &Response{}, err
	//	}
	//	for _, modelLeafID := range model.GuardiansIDs {
	//		sweet, err := c.BmGuardianStorage.GetOne(modelLeafID)
	//		if err != nil {
	//			return &Response{}, err
	//		}
	//		filteredLeafs = append(filteredLeafs, sweet)
	//	}
	//
	//	return &Response{Res: filteredLeafs}, nil
	//}
	return &Response{Res: guardians}, nil
}

// FindOne choc
func (c BmGuardianResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmGuardianStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmGuardianResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Guardian)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmGuardianStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmGuardianResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmGuardianStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmGuardianResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Guardian)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmGuardianStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
