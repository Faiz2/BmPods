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

// BmModelStorageExample stores all models
type BmModelStorageExample struct {
	models   map[string]*BmModel.ModelExample
	idCount int

	db *BmMongodb.BmMongodb
}

var ModelStorageName = "BmModelStorageExample"

func (s BmModelStorageExample) NewModelStorage(args []BmDaemons.BmDaemon) *BmModelStorageExample {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmModelStorageExample{make(map[string]*BmModel.ModelExample), 1, mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmModelStorageExample) GetAll(skip int, take int) []*BmModel.ModelExample {
	in := BmModel.ModelExample{}
	var out []BmModel.ModelExample
	err := s.db.FindMulti(&in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.ModelExample
		//tmp := make(map[string]*BmModel.ModelExample)
		for _, iter := range out {
			s.db.ResetIdWithId_(&iter)
			tmp = append(tmp, &iter)
			//tmp[iter.ID] = &iter
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.ModelExample)
	}
}

// GetOne model
func (s BmModelStorageExample) GetOne(id string) (BmModel.ModelExample, error) {
	in := BmModel.ModelExample{ ID:id }
	out := BmModel.ModelExample{ ID:id }
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("ModelExample for id %s not found", id)
	return BmModel.ModelExample{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmModelStorageExample) Insert(c BmModel.ModelExample) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmModelStorageExample) Delete(id string) error {
	in := BmModel.ModelExample{ ID:id }
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ModelExample with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmModelStorageExample) Update(c BmModel.ModelExample) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ModelExample with id does not exist")
	}

	return nil
}

func (s *BmModelStorageExample) Count(c BmModel.ModelExample) int {
	r, _ := s.db.Count(&c)
	return r
}