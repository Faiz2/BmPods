package BmDataStorage

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/manyminds/api2go"
	"github.com/alfredyang1986/BmPods/BmModel"
	"github.com/alfredyang1986/BmPods/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmPods/BmDaemons"
)

// BmApplyStorage stores all applys
type BmApplyStorage struct {
	applys   map[string]*BmModel.Apply
	idCount int

	db *BmMongodb.BmMongodb
}

var ApplyStorageName = "BmApplyStorage"

func (s BmApplyStorage) NewApplyStorage(args []BmDaemons.BmDaemon) *BmApplyStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmApplyStorage{make(map[string]*BmModel.Apply), 1, mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmApplyStorage) GetAll(skip int, take int) []*BmModel.Apply {
	in := BmModel.Apply{}
	var out []BmModel.Apply
	err := s.db.FindMulti(&in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.Apply
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Apply)
	}
}

// GetOne model
func (s BmApplyStorage) GetOne(id string) (BmModel.Apply, error) {
	in := BmModel.Apply{ ID:id }
	out := BmModel.Apply{ ID:id }
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Apply for id %s not found", id)
	return BmModel.Apply{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmApplyStorage) Insert(c BmModel.Apply) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmApplyStorage) Delete(id string) error {
	in := BmModel.Apply{ ID:id }
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Apply with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmApplyStorage) Update(c BmModel.Apply) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Apply with id does not exist")
	}

	return nil
}

func (s *BmApplyStorage) Count(c BmModel.Apply) int {
	r, _ := s.db.Count(&c)
	return r
}