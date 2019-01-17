package BmDataStorage

import (
	"github.com/alfredyang1986/BmPods/BmModel"
	"github.com/alfredyang1986/BmPods/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmPods/BmDaemons"
	"fmt"
	"github.com/manyminds/api2go"
	"errors"
	"net/http"
)

// ApplyeeStorage stores all users
type BmApplyeeStorage struct {
	db *BmMongodb.BmMongodb
}

func (s BmApplyeeStorage) NewApplyeeStorage(args []BmDaemons.BmDaemon) *BmApplyeeStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmApplyeeStorage{mdb}
}

// GetAll returns the user map (because we need the ID as key too)
func (s BmApplyeeStorage) GetAll(skip int, take int) []*BmModel.Applyee {
	in := BmModel.Applyee{}
	var out []BmModel.Applyee
	err := s.db.FindMulti(&in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.Applyee
		//tmp := make(map[string]*BmModel.Applyee)
		for _, iter := range out {
			s.db.ResetIdWithId_(&iter)
			tmp = append(tmp, &iter)
			//tmp[iter.ID] = &iter
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Applyee)
	}
}

// GetOne user
func (s BmApplyeeStorage) GetOne(id string) (BmModel.Applyee, error) {
	in := BmModel.Applyee{ ID:id }
	out := BmModel.Applyee{ ID:id }
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Applyee for id %s not found", id)
	return BmModel.Applyee{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a user
func (s *BmApplyeeStorage) Insert(c BmModel.Applyee) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmApplyeeStorage) Delete(id string) error {
	in := BmModel.Applyee{ ID:id }
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Applyee with id %s does not exist", id)
	}

	return nil
}

// Update a user
func (s *BmApplyeeStorage) Update(c BmModel.Applyee) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Applyee with id does not exist")
	}

	return nil
}

func (s *BmApplyeeStorage) Count(c BmModel.Applyee) int {
	r, _ := s.db.Count(&c)
	return r
}