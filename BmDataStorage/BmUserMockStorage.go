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

// UserStorage stores all users
type UserStorage struct {
	users   map[string]*BmModel.User
	idCount int

	db *BmMongodb.BmMongodb
}

func (s UserStorage) NewUserStorage(args []BmDaemons.BmDaemon) *UserStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UserStorage{make(map[string]*BmModel.User), 1, mdb}
}

// GetAll returns the user map (because we need the ID as key too)
func (s UserStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.User {
	in := BmModel.User{}
	var out []BmModel.User
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.User
		//tmp := make(map[string]*BmModel.User)
		for _, iter := range out {
			s.db.ResetIdWithId_(&iter)
			tmp = append(tmp, &iter)
			//tmp[iter.ID] = &iter
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.User)
	}
}

// GetOne user
func (s UserStorage) GetOne(id string) (BmModel.User, error) {
	in := BmModel.User{ID: id}
	out := BmModel.User{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("User for id %s not found", id)
	return BmModel.User{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a user
func (s *UserStorage) Insert(c BmModel.User) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UserStorage) Delete(id string) error {
	in := BmModel.User{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("User with id %s does not exist", id)
	}

	return nil
}

// Update a user
func (s *UserStorage) Update(c BmModel.User) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("User with id does not exist")
	}

	return nil
}

func (s *UserStorage) Count(c BmModel.User) int {
	r, _ := s.db.Count(&c)
	return r
}
