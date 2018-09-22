package validate

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func ExampleStruct() {
	u := &UserForm{
		Name: "inhere",
	}

	v := Struct(u)
	ok := v.Validate()

	fmt.Println(ok)
}

func TestValidation(t *testing.T) {
	is := assert.New(t)

	m := GMap{
		"name":  "inhere",
		"age":   100,
		"oldSt": 1,
		"newSt": 2,
		"email": "some@e.com",
	}

	v := New(m)
	v.AddRule("name", "required")
	v.AddRule("name", "minLen", 7)
	v.AddRule("age", "max", 99)
	v.AddRule("age", "min", 1)

	v.WithScenes(SValues{
		"create": []string{"name", "email"},
		"update": []string{"name"},
	})

	ok := v.Validate()
	is.False(ok)
	is.Equal("name value min length is 7", v.Errors.Get("name"))
}

// UserForm struct
type UserForm struct {
	Name     string    `json:"name" validate:"required|minLen:7"`
	Email    string    `json:"email" validate:"email"`
	CreateAt int       `json:"createAt" validate:"email"`
	Safe     int       `json:"safe" validate:"-"`
	UpdateAt time.Time `json:"updateAt" validate:"required"`
	Code     string    `json:"code" validate:"customValidator"`
}

// custom validator in the source struct.
func (f *UserForm) CustomValidator(val string) bool {
	return len(val) == 4
}

// Messages you can custom define validator error messages.
func (f *UserForm) Messages() map[string]string {
	return SMap{
		"required":      "oh! the {field} is required",
		"Name.required": "message for special field",
	}
}

// Translates you can custom field translates.
func (f *UserForm) Translates() map[string]string {
	return SMap{
		"Name":  "User Name",
		"Email": "User Email",
	}
}

func TestStruct(t *testing.T) {
	is := assert.New(t)
	u := &UserForm{
		Name: "inhere",
	}

	v := Struct(u)
	ok := v.Validate()

	is.False(ok)
}