package suite

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/suite/fix"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Model struct {
	suite.Suite
	*require.Assertions
	DB       *gorm.DB
	popDB    *pop.Connection
	Fixtures packr.Box
}

func (m *Model) SetupTest() {
	t := m.T()
	m.Assertions = require.New(t)
	m.NotNil(m.DB)
	m.NoError(m.DB.Error)
	m.DB = m.DB.LogMode(testing.Verbose())
}

func (m *Model) TearDownTest() {
	m.NoError(m.DB.Error)
	m.NoError(m.popDB.TruncateAll())
}

func (m *Model) DBDelta(delta int, name string, fn func()) {
	var sc int
	err := m.DB.Table(name).Count(&sc)
	m.NoError(err.Error)
	fn()
	var ec int
	err = m.DB.Table(name).Count(&ec)
	m.NoError(err.Error)
	m.Equal(ec, sc+delta)
}

func (as *Model) LoadFixture(name string) {
	sc, err := fix.Find(name)
	as.NoError(err)

	for _, table := range sc.Tables {
		for _, row := range table.Row {
			q := "insert into " + table.Name
			keys := []string{}
			skeys := []string{}
			for k := range row {
				keys = append(keys, k)
				skeys = append(skeys, ":"+k)
			}

			q = q + fmt.Sprintf(" (%s) values (%s)", strings.Join(keys, ","), strings.Join(skeys, ","))
			derr := as.DB.Exec(q)
			as.NoError(derr.Error)
		}
	}
}

func NewModel() *Model {
	m := &Model{}
	c, err := pop.Connect(envy.Get("GO_ENV", "test"))
	if err == nil {
		deets := c.Dialect.Details()
		m.DB, err = gorm.Open(deets.Dialect, c.URL())
		if err != nil {
			m.NoError(err)
		}
		m.popDB = c
	}
	return m
}

func NewModelWithFixtures(box packr.Box) (*Model, error) {
	m := NewModel()
	m.Fixtures = box
	return m, fix.Init(box)
}

func Test_Model(t *testing.T) {
	model, err := NewModelWithFixtures(packr.NewBox("../fixtures"))
	if err != nil {
		t.Fatal(err)
	}
	suite.Run(t, model)
}
