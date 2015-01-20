package kitchen

import (
	"testing"

	"github.com/casimir/doable"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCook(t *testing.T) {
	recipe := makeCookableRecipe()
	stock := makeTestingStock()
	cook := NewCook(recipe, stock)

	Convey("Cook should consume the ingredients and make the result", t, func() {
		So(cook.Cook(), ShouldBeNil)
		So(cook.stock.Count(makeItem("1")), ShouldBeZeroValue)
		So(cook.stock.Count(makeItem("10")), ShouldEqual, 1)
	})
	Convey("Cook should return errors on illegal actions", t, func() {
		So(cook.Cook(), ShouldNotBeNil)
	})
}

func TestRecipeCookable(t *testing.T) {
	Convey("When the recipe can be cook it should return true", t, func() {
		cook := NewCook(makeCookableRecipe(), makeTestingStock())
		So(cook.IsCookable(), ShouldBeTrue)
	})
	Convey("When the recipe can not be cooked it should return false", t, func() {
		cook := NewCook(makeNotCookableRecipe(), makeTestingStock())
		So(cook.IsCookable(), ShouldBeFalse)
	})
}

func TestMissingIngredients(t *testing.T) {
	Convey("When the recipe can be cook it should return an empty slice", t, func() {
		cook := NewCook(makeCookableRecipe(), makeTestingStock())
		So(cook.MissingIngredients().Size(), ShouldEqual, 0)
	})
	Convey("When the recipe can not be cooked it should return a slice containing the missing ingredients", t, func() {
		cook := NewCook(makeNotCookableRecipe(), makeTestingStock())
		expected := doable.NewList()
		expected.AddN(makeItem("1"), 97)
		So(cook.MissingIngredients(), ShouldResemble, expected)
	})
}

type mockItem struct {
	ID   string
	Name string
}

func (i mockItem) UID() string {
	return i.ID
}

func (i mockItem) Match(other doable.Item) bool {
	return i.ID == other.(mockItem).ID
}

func makeItem(n string) mockItem {
	return mockItem{ID: n, Name: n}
}

func makeTestingStock() *doable.List {
	s := doable.NewList()
	s.AddN(makeItem("1"), 2)
	s.AddN(makeItem("2"), 5)
	return s
}

func makeCookableRecipe() *Recipe {
	ret := &Recipe{
		ID:   "cookable",
		Name: "cookable",
		Node: &doable.Node{
			Item: makeItem("10"),
			Nb:   1,
		},
	}
	ret.Node.AddDep(&doable.Node{
		Item: makeItem("1"),
		Nb:   2,
	})
	ret.Node.AddDep(&doable.Node{
		Item: makeItem("2"),
		Nb:   3,
	})
	return ret
}

func makeNotCookableRecipe() *Recipe {
	ret := &Recipe{
		ID:   "not_cookable",
		Name: "not_cookable",
		Node: &doable.Node{
			Item: makeItem("10"),
			Nb:   1,
		},
	}
	ret.Node.AddDep(&doable.Node{
		Item: makeItem("1"),
		Nb:   99,
	})
	ret.Node.AddDep(&doable.Node{
		Item: makeItem("2"),
		Nb:   3,
	})
	return ret
}
