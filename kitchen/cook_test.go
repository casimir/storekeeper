package kitchen

import (
	"github.com/casimir/storekeeper/storage"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCook(t *testing.T) {
	bag := makeTestingBag()
	cook := Cook{bag: &bag}
	recipe := makeCookableRecipe()

	Convey("Given a recipe", t, func() {
		Convey("It should consume the ingredients and make the result", func() {
			So(cook.Cook(recipe), ShouldBeNil)
			So(cook.bag.Count("1"), ShouldBeZeroValue)
			So(cook.bag.Count("10"), ShouldEqual, 1)
		})
		Convey("Errors should be returned on illegal actions", func() {
			So(cook.Cook(recipe), ShouldNotBeNil)
		})
	})
}

func TestRecipeCookable(t *testing.T) {
	bag := makeTestingBag()
	cook := Cook{bag: &bag}

	Convey("Given a recipe", t, func() {
		Convey("When the recipe can be cook it should return true", func() {
			So(cook.IsCookable(makeCookableRecipe()), ShouldBeTrue)
		})
		Convey("When the recipe can not be cooked it should return false", func() {
			So(cook.IsCookable(makeNotCookableRecipe()), ShouldBeFalse)
		})
	})
}

func TestMissingIngredients(t *testing.T) {
	bag := makeTestingBag()
	cook := Cook{bag: &bag}

	Convey("Given a recipe", t, func() {
		Convey("When the recipe can be cook it should return an empty slice", func() {
			So(len(cook.MissingIngredients(makeCookableRecipe())), ShouldBeZeroValue)
		})
		Convey("When the recipe can not be cooked it should return a slice containing the missing ingredients", func() {
			expected := []storage.Stack{{97, storage.Item{Id: "1"}}}
			So(cook.MissingIngredients(makeNotCookableRecipe()), ShouldResemble, expected)
		})
	})
}

func makeTestingBag() (b storage.Bag) {
	b.AddItem(2, "1")
	b.AddItem(5, "2")
	return
}

func makeCookableRecipe() Recipe {
	return Recipe{
		Id: "cookable",
		Ingredients: []storage.Stack{
			{2, storage.Item{Id: "1"}},
			{3, storage.Item{Id: "2"}},
		},
		Out: storage.Stack{1, storage.Item{Id: "10"}},
	}
}

func makeNotCookableRecipe() (r Recipe) {
	return Recipe{
		Id: "not_cookable",
		Ingredients: []storage.Stack{
			{99, storage.Item{Id: "1"}},
			{3, storage.Item{Id: "2"}},
		},
		Out: storage.Stack{1, storage.Item{Id: "10"}},
	}
}
