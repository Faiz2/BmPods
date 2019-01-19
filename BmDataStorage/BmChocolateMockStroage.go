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

// sorting
type byID []BmModel.Chocolate

func (c byID) Len() int {
	return len(c)
}

func (c byID) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c byID) Less(i, j int) bool {
	return c[i].GetID() < c[j].GetID()
}

// ChocolateStorage stores all of the tasty chocolate, needs to be injected into
// User and Chocolate Resource. In the real world, you would use a database for that.
type ChocolateStorage struct {
	chocolates map[string]*BmModel.Chocolate
	idCount    int

	db *BmMongodb.BmMongodb
}

func (s ChocolateStorage) NewChocolateStorage(args []BmDaemons.BmDaemon) *ChocolateStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &ChocolateStorage{make(map[string]*BmModel.Chocolate), 1, mdb}
}

// GetAll of the chocolate
func (s ChocolateStorage) GetAll(r api2go.Request) []BmModel.Chocolate {
	in := BmModel.Chocolate{}
	out := make([]BmModel.Chocolate, 10)
	err := s.db.FindMulti(r, &in, &out, -1, -1)
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

// GetOne tasty chocolate
func (s ChocolateStorage) GetOne(id string) (BmModel.Chocolate, error) {
	in := BmModel.Chocolate{ID: id}
	out := BmModel.Chocolate{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("User for id %s not found", id)
	return BmModel.Chocolate{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *ChocolateStorage) Insert(c BmModel.Chocolate) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *ChocolateStorage) Delete(id string) error {
	in := BmModel.Chocolate{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Chocolate with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing chocolate
func (s *ChocolateStorage) Update(c BmModel.Chocolate) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Chocolate with id does not exist")
	}

	return nil
}
