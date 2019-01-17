package BmDataStorage

import (
	"errors"
	"fmt"
	"github.com/alfredyang1986/BmPods/BmDaemons"
	"github.com/alfredyang1986/BmPods/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmPods/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
)

// BmModelLeafStorageExample stores all of the tasty modelleaf, needs to be injected into
// User and ModelLeafExample Resource. In the real world, you would use a database for that.
type BmModelLeafStorageExample struct {
	modelleafs map[string]*BmModel.ModelLeafExample
	idCount    int

	db *BmMongodb.BmMongodb
}

func (s BmModelLeafStorageExample) NewModelLeafStorage(args []BmDaemons.BmDaemon) *BmModelLeafStorageExample {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmModelLeafStorageExample{make(map[string]*BmModel.ModelLeafExample), 1, mdb}
}

// GetAll of the modelleaf
func (s BmModelLeafStorageExample) GetAll() []BmModel.ModelLeafExample {
	in := BmModel.ModelLeafExample{}
	out := make([]BmModel.ModelLeafExample, 10)
	err := s.db.FindMulti(&in, &out, -1, -1)
	if err == nil {
		//tmp := make([]*BmModel.User, 10)
		for _, iter := range out {
			s.db.ResetIdWithId_(&iter)
			//tmp = append(tmp, &iter)
		}
		return out
	} else {
		return nil
	}
}

// GetOne tasty modelleaf
func (s BmModelLeafStorageExample) GetOne(id string) (BmModel.ModelLeafExample, error) {
	in := BmModel.ModelLeafExample{ID: id}
	out := BmModel.ModelLeafExample{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("User for id %s not found", id)
	return BmModel.ModelLeafExample{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *BmModelLeafStorageExample) Insert(c BmModel.ModelLeafExample) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmModelLeafStorageExample) Delete(id string) error {
	in := BmModel.ModelLeafExample{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ModelLeafExample with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *BmModelLeafStorageExample) Update(c BmModel.ModelLeafExample) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ModelLeafExample with id does not exist")
	}

	return nil
}
