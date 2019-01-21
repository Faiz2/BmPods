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

type BmCommentStorage struct {
	db *BmMongodb.BmMongodb
}

func (s BmCommentStorage) NewCommentStorage(args []BmDaemons.BmDaemon) *BmCommentStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmCommentStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmCommentStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.Comment {
	in := BmModel.Comment{}
	var out []BmModel.Comment
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.Comment
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Comment)
	}
}

// GetOne model
func (s BmCommentStorage) GetOne(id string) (BmModel.Comment, error) {
	in := BmModel.Comment{ID: id}
	out := BmModel.Comment{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Comment for id %s not found", id)
	return BmModel.Comment{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmCommentStorage) Insert(c BmModel.Comment) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmCommentStorage) Delete(id string) error {
	in := BmModel.Comment{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Comment with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmCommentStorage) Update(c BmModel.Comment) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Comment with id does not exist")
	}

	return nil
}

func (s *BmCommentStorage) Count(c BmModel.Comment) int {
	r, _ := s.db.Count(&c)
	return r
}
