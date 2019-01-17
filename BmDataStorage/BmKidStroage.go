package BmDataStorage

import (
	"fmt"
	"github.com/alfredyang1986/BmPods/BmModel"
	"github.com/alfredyang1986/BmPods/BmDaemons"
	"github.com/alfredyang1986/BmPods/BmDaemons/BmMongodb"
	"github.com/manyminds/api2go"
	"errors"
	"net/http"
)

// BmKidStorage stores all of the tasty modelleaf, needs to be injected into
// User and Kid Resource. In the real world, you would use a database for that.
type BmKidStorage struct {
	kids map[string]*BmModel.Kid
	idCount    int

	db *BmMongodb.BmMongodb
}

var BmKidStorageName = "BmKidStorage"

func (s BmKidStorage) NewKidStorage(args []BmDaemons.BmDaemon) *BmKidStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmKidStorage{make(map[string]*BmModel.Kid), 1, mdb}
}

// GetAll of the modelleaf
func (s BmKidStorage) GetAll() []BmModel.Kid {
	in := BmModel.Kid{}
	out := make([]BmModel.Kid, 10)
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
func (s BmKidStorage) GetOne(id string) (BmModel.Kid, error) {
	in := BmModel.Kid{ ID:id }
	out := BmModel.Kid{ ID:id }
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("User for id %s not found", id)
	return BmModel.Kid{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *BmKidStorage) Insert(c BmModel.Kid) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmKidStorage) Delete(id string) error {
	in := BmModel.Kid{ ID:id }
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Kid with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *BmKidStorage) Update(c BmModel.Kid) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Kid with id does not exist")
	}

	return nil
}
