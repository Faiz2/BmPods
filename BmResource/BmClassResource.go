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

type BmClassResource struct {
	BmClassStorage          *BmDataStorage.BmClassStorage
	BmStudentStorage        *BmDataStorage.BmStudentStorage
	BmTeacherStorage        *BmDataStorage.BmTeacherStorage
	BmUnitStorage           *BmDataStorage.BmUnitStorage
	BmYardStorage           *BmDataStorage.BmYardStorage
	BmSessioninfoStorage    *BmDataStorage.BmSessioninfoStorage
	BmReservableitemStorage *BmDataStorage.BmReservableitemStorage
}

func (s BmClassResource) NewClassResource(args []BmDataStorage.BmStorage) BmClassResource {
	var us *BmDataStorage.BmClassStorage
	var ys *BmDataStorage.BmYardStorage
	var ss *BmDataStorage.BmSessioninfoStorage
	var cs *BmDataStorage.BmStudentStorage
	var ts *BmDataStorage.BmTeacherStorage
	var ns *BmDataStorage.BmUnitStorage
	var rs *BmDataStorage.BmReservableitemStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmClassStorage" {
			us = arg.(*BmDataStorage.BmClassStorage)
		} else if tp.Name() == "BmStudentStorage" {
			cs = arg.(*BmDataStorage.BmStudentStorage)
		} else if tp.Name() == "BmTeacherStorage" {
			ts = arg.(*BmDataStorage.BmTeacherStorage)
		} else if tp.Name() == "BmUnitStorage" {
			ns = arg.(*BmDataStorage.BmUnitStorage)
		} else if tp.Name() == "BmYardStorage" {
			ys = arg.(*BmDataStorage.BmYardStorage)
		} else if tp.Name() == "BmSessioninfoStorage" {
			ss = arg.(*BmDataStorage.BmSessioninfoStorage)
		} else if tp.Name() == "BmReservableitemStorage" {
			rs = arg.(*BmDataStorage.BmReservableitemStorage)
		}
	}
	return BmClassResource{BmClassStorage: us, BmYardStorage: ys, BmSessioninfoStorage: ss, BmStudentStorage: cs, BmTeacherStorage: ts, BmUnitStorage: ns, BmReservableitemStorage: rs}
}

// FindAll to satisfy api2go data source interface
func (s BmClassResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []BmModel.Class
	models := s.BmClassStorage.GetAll(r, -1, -1)

	for _, model := range models {
		// get all sweets for the model
		model.Students = []*BmModel.Student{}
		for _, tmpID := range model.StudentsIDs {
			choc, err := s.BmStudentStorage.GetOne(tmpID)
			if err != nil {
				return &Response{}, err
			}
			model.Students = append(model.Students, &choc)
		}
		model.Teachers = []*BmModel.Teacher{}
		for _, tmpID := range model.TeachersIDs {
			choc, err := s.BmTeacherStorage.GetOne(tmpID)
			if err != nil {
				return &Response{}, err
			}
			model.Teachers = append(model.Teachers, &choc)
		}
		model.Units = []*BmModel.Unit{}
		for _, tmpID := range model.UnitsIDs {
			choc, err := s.BmUnitStorage.GetOne(tmpID)
			if err != nil {
				return &Response{}, err
			}
			model.Units = append(model.Units, &choc)
		}

		if model.YardID != "" {
			yard, err := s.BmYardStorage.GetOne(model.YardID)
			if err != nil {
				return &Response{}, err
			}
			model.Yard = yard
		}
		if model.SessioninfoID != "" {
			item, err := s.BmSessioninfoStorage.GetOne(model.SessioninfoID)
			if err != nil {
				return &Response{}, err
			}
			model.Sessioninfo = item
		}

		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s BmClassResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []BmModel.Class
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
		for _, model := range s.BmClassStorage.GetAll(r, int(start), int(sizeI)) {

			model.Students = []*BmModel.Student{}
			for _, tmpID := range model.StudentsIDs {
				choc, err := s.BmStudentStorage.GetOne(tmpID)
				if err != nil {
					return 0, &Response{}, err
				}
				model.Students = append(model.Students, &choc)
			}
			model.Teachers = []*BmModel.Teacher{}
			for _, tmpID := range model.TeachersIDs {
				choc, err := s.BmTeacherStorage.GetOne(tmpID)
				if err != nil {
					return 0, &Response{}, err
				}
				model.Teachers = append(model.Teachers, &choc)
			}
			model.Units = []*BmModel.Unit{}
			for _, tmpID := range model.UnitsIDs {
				choc, err := s.BmUnitStorage.GetOne(tmpID)
				if err != nil {
					return 0, &Response{}, err
				}
				model.Units = append(model.Units, &choc)
			}

			if model.YardID != "" {
				yard, err := s.BmYardStorage.GetOne(model.YardID)
				if err != nil {
					return 0, &Response{}, err
				}
				model.Yard = yard
			}
			if model.SessioninfoID != "" {
				item, err := s.BmSessioninfoStorage.GetOne(model.SessioninfoID)
				if err != nil {
					return 0, &Response{}, err
				}
				model.Sessioninfo = item
			}

			result = append(result, *model)
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

		for _, model := range s.BmClassStorage.GetAll(r, int(offsetI), int(limitI)) {

			model.Students = []*BmModel.Student{}
			for _, tmpID := range model.StudentsIDs {
				choc, err := s.BmStudentStorage.GetOne(tmpID)
				if err != nil {
					return 0, &Response{}, err
				}
				model.Students = append(model.Students, &choc)
			}
			model.Teachers = []*BmModel.Teacher{}
			for _, tmpID := range model.TeachersIDs {
				choc, err := s.BmTeacherStorage.GetOne(tmpID)
				if err != nil {
					return 0, &Response{}, err
				}
				model.Teachers = append(model.Teachers, &choc)
			}
			model.Units = []*BmModel.Unit{}
			for _, tmpID := range model.UnitsIDs {
				choc, err := s.BmUnitStorage.GetOne(tmpID)
				if err != nil {
					return 0, &Response{}, err
				}
				model.Units = append(model.Units, &choc)
			}

			if model.YardID != "" {
				yard, err := s.BmYardStorage.GetOne(model.YardID)
				if err != nil {
					return 0, &Response{}, err
				}
				model.Yard = yard
			}
			if model.SessioninfoID != "" {
				item, err := s.BmSessioninfoStorage.GetOne(model.SessioninfoID)
				if err != nil {
					return 0, &Response{}, err
				}
				model.Sessioninfo = item
			}

			result = append(result, *model)
		}
	}

	in := BmModel.Class{}
	count := s.BmClassStorage.Count(in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s BmClassResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.BmClassStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	model.Students = []*BmModel.Student{}
	for _, tmpID := range model.StudentsIDs {
		choc, err := s.BmStudentStorage.GetOne(tmpID)
		if err != nil {
			return &Response{}, err
		}
		model.Students = append(model.Students, &choc)
	}
	model.Teachers = []*BmModel.Teacher{}
	for _, tmpID := range model.TeachersIDs {
		choc, err := s.BmTeacherStorage.GetOne(tmpID)
		if err != nil {
			return &Response{}, err
		}
		model.Teachers = append(model.Teachers, &choc)
	}
	model.Units = []*BmModel.Unit{}
	for _, tmpID := range model.UnitsIDs {
		choc, err := s.BmUnitStorage.GetOne(tmpID)
		if err != nil {
			return &Response{}, err
		}
		model.Units = append(model.Units, &choc)
	}

	if model.YardID != "" {
		yard, err := s.BmYardStorage.GetOne(model.YardID)
		if err != nil {
			return &Response{}, err
		}
		model.Yard = yard
	}
	if model.SessioninfoID != "" {
		item, err := s.BmSessioninfoStorage.GetOne(model.SessioninfoID)
		if err != nil {
			return &Response{}, err
		}
		model.Sessioninfo = item
	}

	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s BmClassResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(BmModel.Class)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	model.CreateTime = float64(time.Now().UnixNano() / 1e6)
	id := s.BmClassStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s BmClassResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.BmClassStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s BmClassResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(BmModel.Class)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.BmClassStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
