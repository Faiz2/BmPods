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
func (s UserStorage) GetAll() map[string]*BmModel.User {
	return s.users
}

// GetOne user
func (s UserStorage) GetOne(id string) (BmModel.User, error) {
	user, ok := s.users[id]
	if ok {
		return *user, nil
	}
	errMessage := fmt.Sprintf("User for id %s not found", id)
	return BmModel.User{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a user
func (s *UserStorage) Insert(c BmModel.User) string {
	// mock data in memory
	//id := fmt.Sprintf("%d", s.idCount)
	//c.ID = id
	//s.users[id] = &c
	//s.idCount++

	// real data in the mongodb
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UserStorage) Delete(id string) error {
	_, exists := s.users[id]
	if !exists {
		return fmt.Errorf("User with id %s does not exist", id)
	}
	delete(s.users, id)

	return nil
}

// Update a user
func (s *UserStorage) Update(c BmModel.User) error {
	_, exists := s.users[c.ID]
	if !exists {
		return fmt.Errorf("User with id %s does not exist", c.ID)
	}
	s.users[c.ID] = &c

	return nil
}