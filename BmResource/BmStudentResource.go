package BmResource

import (
	"errors"
	"github.com/alfredyang1986/BmPods/BmDataStorage"
	"github.com/alfredyang1986/BmPods/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

type BmStudentResource struct {
	BmStudentStorage  *BmDataStorage.BmStudentStorage
	BmKidStorage      *BmDataStorage.BmKidStorage
	BmTeacherStorage  *BmDataStorage.BmTeacherStorage
	BmGuardianStorage *BmDataStorage.BmGuardianStorage
	BmClassStorage *BmDataStorage.BmClassStorage
}

func (s BmStudentResource) NewStudentResource(args []BmDataStorage.BmStorage) BmStudentResource {
	var ss *BmDataStorage.BmStudentStorage
	var ks *BmDataStorage.BmKidStorage
	var gs *BmDataStorage.BmGuardianStorage
	var ts *BmDataStorage.BmTeacherStorage
	var cs *BmDataStorage.BmClassStorage
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
		} else if tp.Name() == "BmClassStorage" {
			cs = arg.(*BmDataStorage.BmClassStorage)
		}
	}
	return BmStudentResource{BmStudentStorage: ss, BmKidStorage: ks, BmGuardianStorage: gs, BmTeacherStorage: ts, BmClassStorage: cs}
}

// FindAll to satisfy api2go data source interface
func (s BmStudentResource) FindAll(r api2go.Request) (api2go.Responder, error) {

	//查詢class下的students
	classesID, ok := r.QueryParams["classesID"]
	if ok {
		modelID := classesID[0]
		filteredLeafs := []BmModel.Student{}
		model, err := s.BmClassStorage.GetOne(modelID)
		if err != nil {
			return &Response{}, err
		}
		for _, modelLeafID := range model.StudentsIDs {
			stud, err := s.BmStudentStorage.GetOne(modelLeafID)
			if err != nil {
				return &Response{}, err
			}

			//TODO:rebindmodel 抽離成 func
			stud.Guardians = []*BmModel.Guardian{}
			for _, gId := range stud.GuardiansIDs {
				g, err := s.BmGuardianStorage.GetOne(gId)
				if err != nil {
					return &Response{}, err
				}
				stud.Guardians = append(stud.Guardians, &g)
			}

			if stud.KidID != "" {
				k, err := s.BmKidStorage.GetOne(stud.KidID)
				if err != nil {
					return &Response{}, err
				}
				stud.Kid = &k
			}

			if stud.TeacherID != "" {
				k, err := s.BmTeacherStorage.GetOne(stud.TeacherID)
				if err != nil {
					return &Response{}, err
				}
				stud.Teacher = &k
			}

			filteredLeafs = append(filteredLeafs, stud)
		}

		return &Response{Res: filteredLeafs}, nil
	}

	var result []BmModel.Student
	studs := s.BmStudentStorage.GetAll(r, -1, -1)
	for _, stud := range studs {
		stud.Guardians = []*BmModel.Guardian{}
		for _, gId := range stud.GuardiansIDs {
			g, err := s.BmGuardianStorage.GetOne(gId)
			if err != nil {
				return &Response{}, err
			}
			stud.Guardians = append(stud.Guardians, &g)
		}

		if stud.KidID != "" {
			k, err := s.BmKidStorage.GetOne(stud.KidID)
			if err != nil {
				return &Response{}, err
			}
			stud.Kid = &k
		}

		if stud.TeacherID != "" {
			k, err := s.BmTeacherStorage.GetOne(stud.TeacherID)
			if err != nil {
				return &Response{}, err
			}
			stud.Teacher = &k
		}

		result = append(result, *stud)
	}
	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load users in chunks
func (s BmStudentResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []BmModel.Student
		number, size, offset, limit string
		startIndex, sizeInt, count, pages int
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

		startIndex = int(start)
		sizeInt = int(sizeI)
	} else {
		limitI, err := strconv.ParseUint(limit, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		offsetI, err := strconv.ParseUint(offset, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		startIndex = int(offsetI)
		sizeInt = int(limitI)
	}

	//查詢class下的students
	classesID, ok := r.QueryParams["classesID"]
	if ok {
		modelID := classesID[0]
		filteredLeafs := []BmModel.Student{}
		model, err := s.BmClassStorage.GetOne(modelID)
		if err != nil {
			return uint(0), &Response{}, err
		}
		for _, modelLeafID := range model.StudentsIDs {
			stud, err := s.BmStudentStorage.GetOne(modelLeafID)
			if err != nil {
				return uint(0), &Response{}, err
			}

			//TODO:rebindmodel 抽離成 func
			stud.Guardians = []*BmModel.Guardian{}
			for _, gId := range stud.GuardiansIDs {
				g, err := s.BmGuardianStorage.GetOne(gId)
				if err != nil {
					return uint(0), &Response{}, err
				}
				stud.Guardians = append(stud.Guardians, &g)
			}

			if stud.KidID != "" {
				k, err := s.BmKidStorage.GetOne(stud.KidID)
				if err != nil {
					return uint(0), &Response{}, err
				}
				stud.Kid = &k
			}

			if stud.TeacherID != "" {
				k, err := s.BmTeacherStorage.GetOne(stud.TeacherID)
				if err != nil {
					return uint(0), &Response{}, err
				}
				stud.Teacher = &k
			}

			filteredLeafs = append(filteredLeafs, stud)
		}

		count = len(filteredLeafs)
		pages = 1 + int(count / sizeInt)

		return uint(count), &Response{Res: filteredLeafs, QueryRes:"students", TotalPage:pages}, nil
	}

	for _, iter := range s.BmStudentStorage.GetAll(r, startIndex, sizeInt) {
		result = append(result, *iter)
	}

	in := BmModel.Student{}
	count = s.BmStudentStorage.Count(in)
	pages = 1 + int(count / sizeInt)

	return uint(count), &Response{Res: result, QueryRes:"students", TotalPage:pages}, nil
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
	model, ok := obj.(BmModel.Student)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	model.CreateTime = float64(time.Now().UnixNano() / 1e6)
	id := s.BmStudentStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
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
