package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Widget struct {
	ID        uuid.UUID `json:"id" gorm:"col:id,primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"col:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"col:updated_at"`
	Name      string    `json:"name" gorm:"col:name"`
	Body      string    `json:"body" gorm:"col:body"`
}

// String is not required by pop and may be deleted
func (w Widget) String() string {
	jw, _ := json.Marshal(w)
	return string(jw)
}

// Widgets is not required by pop and may be deleted
type Widgets []Widget

// String is not required by pop and may be deleted
func (w Widgets) String() string {
	jw, _ := json.Marshal(w)
	return string(jw)
}

func (w *Widget) BeforeCreate(tx *gorm.DB) error {
	var err error
	w.ID, err = uuid.NewV4()
	return err
}

func (w *Widget) Validate(tx *gorm.DB) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: w.Name, Name: "Name"},
		&validators.StringIsPresent{Field: w.Body, Name: "Body"},
	), nil
}

func (w *Widget) ValidateAndCreate(tx *gorm.DB) (*validate.Errors, error) {
	verrs, err := w.Validate(tx)
	if err != nil {
		return verrs, errors.WithStack(err)
	}
	e := tx.Create(w)
	return verrs, e.Error
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (w *Widget) ValidateAndUpdate(tx *gorm.DB) (*validate.Errors, error) {
	verrs, err := w.Validate(tx)
	if err != nil {
		return verrs, errors.WithStack(err)
	}
	e := tx.Model(w).Updates(w)
	return verrs, e.Error
}
