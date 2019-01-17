package BmResource

import (
	"github.com/manyminds/api2go"
	"github.com/alfredyang1986/BmPods/BmModel"
	"errors"
	"net/http"
	"github.com/alfredyang1986/BmPods/BmDataStorage"
	"reflect"
	"strconv"
)

type BmStudentResource struct {
	BmStudentStorage *BmDataStorage.BmStudentStorage
	BmKidStorage *BmDataStorage.BmKidStorage
	BmTeacherStorage *BmDataStorage.BmTeacherStorage
	BmGuardianStorage *BmDataStorage.BmGuardianStorage
}

func (s BmStudentResource) NewStudentResource(args []BmDataStorage.BmStorage) BmStudentResource {
	var ss *BmDataStorage.BmStudentStorage
	var ks *BmDataStorage.BmKidStorage
	var gs *BmDataStorage.BmGuardianStorage
	var ts *BmDataStorage.BmTeacherStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmStudentStorage" {
			ss = arg.(*BmDataStorage.BmStudentStorage)
		} else if tp.Name() == "BmKidStorage" {
			ks = arg.(*BmDataStorage.BmKidStorage)
		} else if tp.Name() == "BmGuardianStorage" {
			gs = arg.(*BmDataStorage.BmGuardianStorage)
		} else if tp.Name() == "BmTeacherStorage" {
			ts = arg.(*BmDataStorage.BmTeacherStorage)
		}
	}
	return BmStudentResource{BmStudentStorage: ss, BmKidStorage: ks, BmGuardianStorage: gs, BmTeacherStorage: ts}
}

// FindAll to satisfy api2go data source interface
func (s BmStudentResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []BmModel.Student
	studs := s.BmStudentStorage.GetAll(-1, -1)

	for _, user := range studs {
		user.Guardians = []*BmModel.Guardian{}
		for _, gId := range user.GuardiansIDs {
			g, err := s.BmGuardianStorage.GetOne(gId)
			if err != nil {
				return &Response{}, err
			}
			user.Guardians = append(user.Guardians, &g)
		}

		if user.KidID != "" {
			k, err := s.BmKidStorage.GetOne(user.KidID)
			if err != nil {
				return &Response{}, err
			}
			user.Kid = &k
		}

		if user.TeacherID != "" {
			k, err := s.BmTeacherStorage.GetOne(user.TeacherID)
			if err != nil {
				return &Response{}, err
			}
			user.Teacher = &k
		}

		result = append(result, *user)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load users in chunks
func (s BmStudentResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []BmModel.Student
		number, size, offset, limit string
	)

	numberQuery, ok := r.QueryParams["page[number]"]
	if ok {
		number = numberQuery[0]
	}
	sizeQuery, ok := r.QueryParams["page[size]"]
	if ok {
		size = sizeQuery[0]
	}
	offsetQuery, ok := r.QueryParams["page[offset]"]
	if ok {
		offset = offsetQuery[0]
	}
	limitQuery, ok := r.QueryParams["page[limit]"]
	if ok {
		limit = limitQuery[0]
	}

	if size != "" {
		sizeI, err := strconv.ParseInt(size, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		numberI, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		start := sizeI * (numberI - 1)
		for _, iter := range s.BmStudentStorage.GetAll(int(start), int(sizeI)) {
			result = append(result, *iter)
		}

	} else {
		limitI, err := strconv.ParseUint(limit, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		offsetI, err := strconv.ParseUint(offset, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		for _, iter := range s.BmStudentStorage.GetAll(int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := BmModel.Student{}
	count := s.BmStudentStorage.Count(in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the user with the given ID, otherwise an error
func (s BmStudentResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	user, err := s.BmStudentStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	user.Guardians = []*BmModel.Guardian{}
	for _, chocolateID := range user.GuardiansIDs {
		choc, err := s.BmGuardianStorage.GetOne(chocolateID)
		if err != nil {
			return &Response{}, err
		}
		user.Guardians = append(user.Guardians, &choc)
	}

	if user.KidID != "" {
		k, err := s.BmKidStorage.GetOne(user.KidID)
		if err != nil {
			return &Response{}, err
		}
		user.Kid = &k
	}

	if user.TeacherID != "" {
		k, err := s.BmTeacherStorage.GetOne(user.TeacherID)
		if err != nil {
			return &Response{}, err
		}
		user.Teacher = &k
	}

	return &Response{Res: user}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s BmStudentResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	user, ok := obj.(BmModel.Student)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.BmStudentStorage.Insert(user)
	user.ID = id

	return &Response{Res: user, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s BmStudentResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.BmStudentStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the user
func (s BmStudentResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	user, ok := obj.(BmModel.Student)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.BmStudentStorage.Update(user)
	return &Response{Res: user, Code: http.StatusNoContent}, err
}
