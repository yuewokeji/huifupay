package autoassign

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Name struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

type Person struct {
	Name       string `json:"name,omitempty"`
	NameObject Name   `json:"-" autoassign:"Name"`
	Age        int    `json:"age"`
}

func (n Name) IsEmpty() bool {
	return n.First == "" && n.Last == ""
}

func Test_ObjectToJSONString(t *testing.T) {
	name := Name{
		First: "first name",
		Last:  "last name",
	}

	p := &Person{
		Name:       "",
		NameObject: name,
		Age:        18,
	}
	err := ObjectToJSONString(p)
	assert.Nil(t, err)
	assert.Equal(t, `{"first":"first name","last":"last name"}`, p.Name)
}

func Test_ObjectToJSONString_EmptyValue(t *testing.T) {
	name := Name{
		First: "",
		Last:  "",
	}

	p := &Person{
		Name:       "",
		NameObject: name,
		Age:        18,
	}
	err := ObjectToJSONString(p)
	assert.Nil(t, err)
	data, err := json.Marshal(p)
	assert.Nil(t, err)
	assert.Equal(t, `{"age":18}`, string(data))
}

func Test_JSONStringToObject(t *testing.T) {
	p := &Person{
		Name: `{"first":"first name","last":"last name"}`,
		Age:  18,
	}
	err := JSONStringToObject(p)
	assert.Nil(t, err)
	assert.Equal(t, "first name", p.NameObject.First)
	assert.Equal(t, "last name", p.NameObject.Last)
}

type PersonPtr struct {
	Name       *string `json:"name"`
	NameObject *Name   `json:"-" autoassign:"Name"`
	Age        int     `json:"age"`
}

func Test_ObjectToJSONString_Ptr(t *testing.T) {
	name := Name{
		First: "first name",
		Last:  "last name",
	}

	s := ""
	p := &PersonPtr{
		Name:       &s,
		NameObject: &name,
		Age:        18,
	}
	err := ObjectToJSONString(p)
	assert.Nil(t, err)
	assert.Equal(t, `{"first":"first name","last":"last name"}`, *p.Name)
}

func Test_JSONStringToObject_Ptr(t *testing.T) {
	s := `{"first":"first name","last":"last name"}`
	p := &PersonPtr{
		Name: &s,
		Age:  18,
	}
	err := JSONStringToObject(p)
	assert.Nil(t, err)
	assert.Equal(t, "first name", p.NameObject.First)
	assert.Equal(t, "last name", p.NameObject.Last)
}
