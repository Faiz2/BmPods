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

type BmPostStorage struct {
	db *BmMongodb.BmMongodb
}

func (s BmPostStorage) NewPostStorage(args []BmDaemons.BmDaemon) *BmPostStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmPostStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmPostStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.Post {
	in := BmModel.Post{}
	var out []BmModel.Post
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.Post
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Post)
	}
}

// GetOne model
func (s BmPostStorage) GetOne(id string) (BmModel.Post, error) {
	in := BmModel.Post{ID: id}
	out := BmModel.Post{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Post for id %s not found", id)
	return BmModel.Post{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmPostStorage) Insert(c BmModel.Post) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmPostStorage) Delete(id string) error {
	in := BmModel.Post{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Post with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmPostStorage) Update(c BmModel.Post) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Post with id does not exist")
	}

	return nil
}

func (s *BmPostStorage) Count(c BmModel.Post) int {
	r, _ := s.db.Count(&c)
	return r
}
