package BmDataStorage

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/alfredyang1986/BmPods/BmDaemons"
	"github.com/alfredyang1986/BmPods/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmPods/BmModel"
	"github.com/manyminds/api2go"
)

// BmReservableItemStorage stores all reservableItems
type BmReservableItemStorage struct {
	reservableItems map[string]*BmModel.ReservableItem
	idCount         int

	db *BmMongodb.BmMongodb
}

func (s BmReservableItemStorage) NewReservableItemStorage(args []BmDaemons.BmDaemon) *BmReservableItemStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmReservableItemStorage{make(map[string]*BmModel.ReservableItem), 1, mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmReservableItemStorage) GetAll(skip int, take int) []*BmModel.ReservableItem {
	in := BmModel.ReservableItem{}
	var out []BmModel.ReservableItem
	err := s.db.FindMulti(&in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.ReservableItem
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.ReservableItem)
	}
}

// GetOne model
func (s BmReservableItemStorage) GetOne(id string) (BmModel.ReservableItem, error) {
	in := BmModel.ReservableItem{ID: id}
	out := BmModel.ReservableItem{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("ReservableItem for id %s not found", id)
	return BmModel.ReservableItem{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmReservableItemStorage) Insert(c BmModel.ReservableItem) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmReservableItemStorage) Delete(id string) error {
	in := BmModel.ReservableItem{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ReservableItem with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmReservableItemStorage) Update(c BmModel.ReservableItem) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ReservableItem with id does not exist")
	}

	return nil
}

func (s *BmReservableItemStorage) Count(c BmModel.ReservableItem) int {
	r, _ := s.db.Count(&c)
	return r
}
