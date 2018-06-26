package models

import (
	"github.com/markbates/going/randx"
)

func (as *ModelSuite) CreateWidget() *Widget {
	w := &Widget{
		Name: randx.String(20),
		Body: randx.String(20),
	}

	verrs, err := w.ValidateAndCreate(as.Gorm)
	as.NoError(err)
	as.False(verrs.HasAny())
	return w
}

func (m *ModelSuite) Test_Widget_ValidateAndCreate() {
	w := &Widget{}
	verrs, err := w.ValidateAndCreate(m.Gorm)
	m.NoError(err)
	m.True(verrs.HasAny())

	w.Name = "Foo"
	w.Body = "Bar"

	verrs, err = w.ValidateAndCreate(m.Gorm)
	m.NoError(err)
	m.False(verrs.HasAny())
}

func (m *ModelSuite) Test_Widget_ValidateAndUpdate() {
	w := m.CreateWidget()
	w.Name = "Updated"

	verrs, err := w.ValidateAndUpdate(m.Gorm)
	m.NoError(err)
	m.False(verrs.HasAny())

	derr := m.Gorm.Where("id = ?", w.ID).First(w)
	m.NoError(derr.Error)
	m.Equal("Updated", w.Name)
}
