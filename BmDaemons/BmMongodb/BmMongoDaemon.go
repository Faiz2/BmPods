package BmMongodb

import (
	"gopkg.in/mgo.v2"
	"errors"
	"reflect"
	"github.com/alfredyang1986/BmPods/BmModel"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmmate"
)

const (
	BMJson    string = "json"
	BMJsonAPI string = "jsonapi"
	BMMongo   string = "bson"
)

type BmMongodb struct {
	Host string
	Port string
	Database string
}

type NoPtr struct {
}

func (m BmMongodb) NewMongoDBDaemon(args map[string]string) *BmMongodb {
	return &BmMongodb{
		Host: args["host"],
		Port: args["port"],
		Database: args["database"] }
}

func (m *BmMongodb) InsertBmObject(ptr BmModel.BmModelBase) (string, error) {
	session, err := mgo.Dial(m.Host + ":" + m.Port)
	if err != nil {
		return "", errors.New("dial db error")
	}
	defer session.Close()

	oid := m.generateModelId_(ptr)
	v := reflect.ValueOf(ptr).Elem()
	cn := v.Type().Name()
	c := session.DB(m.Database).C(cn)

	rst, err := Struct2map(v)
	rst["_id"] = oid
	err = c.Insert(rst)
	if err == nil {
		return m.resetIdWithId_(ptr), nil
	}
	return "", err
}

func (m *BmMongodb) ExistsBmObject(ptr BmModel.BmModelBase, out BmModel.BmModelBase) (bool, error) {
	session, err := mgo.Dial(m.Host + ":" + m.Port)
	if err != nil {
		return false, errors.New("dial db error")
	}
	defer session.Close()

	oid := m.resetId_WithId(ptr)
	v := reflect.ValueOf(ptr).Elem()
	cn := v.Type().Name()
	c := session.DB(m.Database).C(cn)

	rst, err := Struct2map(v)
	rst["_id"] = oid
	err = c.Find(rst).One(out)
	if err != nil {
		fmt.Println(err)
	}

	m.resetIdWithId_(out)
	return true, nil
}

func (m *BmMongodb) FindOne(ptr BmModel.BmModelBase, out BmModel.BmModelBase) error {
	session, err := mgo.Dial(m.Host + ":" + m.Port)
	if err != nil {
		return errors.New("dial db error")
	}
	defer session.Close()

	oid := m.resetId_WithId(ptr)
	v := reflect.ValueOf(ptr).Elem()
	cn := v.Type().Name()
	c := session.DB(m.Database).C(cn)

	err = c.Find(bson.M{ "_id": oid }).One(out)
	if err != nil {
		fmt.Println(err)
		return errors.New("query error")
	}

	m.resetIdWithId_(out)
	return nil
}

func (m *BmMongodb) generateModelId_(ptr BmModel.BmModelBase) bson.ObjectId {
	f := reflect.ValueOf(ptr).Elem().FieldByName("Id_")
	v := bson.NewObjectId()
	f.Set(reflect.ValueOf(v))
	return v
}

func (m *BmMongodb) resetIdWithId_(ptr BmModel.BmModelBase) string {
	f := reflect.ValueOf(ptr).Elem().FieldByName("Id_")
	t := f.Interface().(bson.ObjectId)
	fs := reflect.ValueOf(ptr).Elem().FieldByName("ID")
	fs.SetString(t.Hex())
	return t.Hex()
}

func (m *BmMongodb) resetId_WithId(ptr BmModel.BmModelBase) bson.ObjectId {
	fs := reflect.ValueOf(ptr).Elem().FieldByName("ID")
	t := fs.Interface().(string)
	f := reflect.ValueOf(ptr).Elem().FieldByName("Id_")
	v := bson.ObjectIdHex(t)
	f.Set(reflect.ValueOf(v))
	return v
}

func AttrValue(v reflect.Value) (interface{}, error) {
	switch v.Kind() {
	case reflect.Invalid:
		return nil, nil
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return v.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint(), nil
	case reflect.Float32, reflect.Float64:
		return v.Float(), nil
	case reflect.String:
		return v.String(), nil
	case reflect.Array, reflect.Slice:
		var rst []interface{}
		for i := 0; i < v.Len(); i++ {
			tmp, _ := AttrValue(v.Index(i))
			rst = append(rst, tmp)
		}
		return rst, nil
	case reflect.Map:
		rst := make(map[string]interface{})
		for _, key := range v.MapKeys() {
			kv := v.MapIndex(key)
			tmp, _ := AttrValue(kv)
			rst[key.String()] = tmp
		}
		return rst, nil
	case reflect.Interface:
		if bmmate.IsStruct(v) {
			if reflect.TypeOf(v.Interface()).Kind() == reflect.String {
				return AttrValue(reflect.ValueOf(v.Interface()))
			} else {
				return AttrValue(reflect.ValueOf(v.Interface()))
			}
		} else {
			return AttrValue(reflect.ValueOf(v.Interface()))
		}
	}

	return NoPtr{}, errors.New("not implement")
}

func Struct2map(v reflect.Value) (map[string]interface{}, error) {
	rst := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {

		fieldInfo := v.Type().Field(i) // a.reflect.struct.field
		fieldValue := v.Field(i)
		tag := fieldInfo.Tag // a.reflect.tag

		var name string
		if tag.Get(BMMongo) != "" {
			name = tag.Get(BMMongo)
		} else if tag.Get(BMMongo) == "-" {
			continue
		} else {
			//name = strings.ToLower(fieldInfo.Name)
			continue
		}

		if name == "id" || name == "_id" {
			continue
		}

		//ja, ok := tag.Lookup(BMJsonAPI)
		//if ok && ja == "relationships" {
			//NOTE: relationships
			//rst[name] = "TODO"
			//continue
		//}

		tmp, _ := AttrValue(fieldValue)
		rst[name] = tmp
	}

	return rst, nil
}