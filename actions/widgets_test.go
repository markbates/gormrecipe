package actions

import (
	"github.com/markbates/going/randx"
	"github.com/markbates/gormrecipe/models"
)

func (as *ActionSuite) BuildWidget() *models.Widget {
	return &models.Widget{
		Name: randx.String(20),
		Body: randx.String(20),
	}
}

func (as *ActionSuite) CreateWidget() *models.Widget {
	w := as.BuildWidget()
	verrs, err := w.ValidateAndCreate(as.Gorm)
	as.NoError(err)
	as.False(verrs.HasAny())
	return w
}

func (as *ActionSuite) Test_WidgetsResource_List() {
	w1 := as.CreateWidget()
	w2 := as.CreateWidget()
	res := as.HTML("/widgets").Get()
	as.Equal(200, res.Code)
	body := res.Body.String()
	as.Contains(body, w1.Name)
	as.Contains(body, w2.Name)
}

func (as *ActionSuite) Test_WidgetsResource_Show() {
	w := as.CreateWidget()
	res := as.HTML("/widgets/%s", w.ID).Get()
	as.Equal(200, res.Code)
	body := res.Body.String()
	as.Contains(body, w.Name)
}

func (as *ActionSuite) Test_WidgetsResource_New() {
	res := as.HTML("/widgets/new").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_WidgetsResource_Create() {
	w := as.BuildWidget()
	as.GormDelta(1, "widgets", func() {
		res := as.HTML("/widgets").Post(w)
		as.Equal(302, res.Code)
	})
}

func (as *ActionSuite) Test_WidgetsResource_Edit() {
	w := as.CreateWidget()
	res := as.HTML("/widgets/%s/edit", w.ID).Get()
	as.Equal(200, res.Code)
	body := res.Body.String()
	as.Contains(body, w.Name)
}

func (as *ActionSuite) Test_WidgetsResource_Update() {
	w := as.CreateWidget()
	w.Name = "Updated"
	res := as.HTML("/widgets/%s", w.ID).Put(w)
	as.Equal(302, res.Code)

	derr := as.Gorm.Where("id = ?", w.ID).First(w)
	as.NoError(derr.Error)
	as.Equal("Updated", w.Name)
}

func (as *ActionSuite) Test_WidgetsResource_Destroy() {
	w := as.CreateWidget()
	as.GormDelta(-1, "widgets", func() {
		res := as.HTML("/widgets/%s", w.ID).Delete()
		as.Equal(302, res.Code)
	})
}
