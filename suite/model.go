package suite

import (
	"testing"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/pop"
	gbsuite "github.com/gobuffalo/suite"
	"github.com/gobuffalo/suite/fix"
	"github.com/jinzhu/gorm"
)

type Model struct {
	*gbsuite.Model
	Gorm *gorm.DB
}

func (m *Model) SetupTest() {
	m.Model.SetupTest()
	m.NotNil(m.Gorm)
	m.NoError(m.Gorm.Error)
	m.Gorm = m.Gorm.LogMode(testing.Verbose())
}

func (m *Model) TearDownTest() {
	m.Model.TearDownTest()
	m.NoError(m.Gorm.Error)
}

func (m *Model) GormDelta(delta int, name string, fn func()) {
	var sc int
	err := m.Gorm.Table(name).Count(&sc)
	m.NoError(err.Error)
	fn()
	var ec int
	err = m.Gorm.Table(name).Count(&ec)
	m.NoError(err.Error)
	m.Equal(ec, sc+delta)
}

func NewModel() *Model {
	gbm := gbsuite.NewModel()
	m := &Model{
		Model: gbm,
	}
	c, err := pop.Connect(envy.Get("GO_ENV", "test"))
	if err == nil {
		deets := c.Dialect.Details()
		m.Gorm, err = gorm.Open(deets.Dialect, c.URL())
		if err != nil {
			m.NoError(err)
		}
		m.DB = c
	}
	return m
}

func NewModelWithFixtures(box packr.Box) (*Model, error) {
	m := NewModel()
	m.Fixtures = box
	return m, fix.Init(box)
}
