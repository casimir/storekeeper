package d3

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var rawRecipe = []byte("{ \"id\" : \"GemCombine_Amethyst_01_02\", \"slug\" : \"flawed-amethyst\", \"name\" : \"Flawed Amethyst\", \"cost\" : 10, \"reagents\" : [ { \"quantity\" : 2, \"item\" : { \"id\" : \"Amethyst_01\", \"name\" : \"Chipped Amethyst\", \"icon\" : \"amethyst_01_demonhunter_male\", \"displayColor\" : \"blue\", \"tooltipParams\" : \"item/chipped-amethyst\" } } ], \"itemProduced\" : { \"id\" : \"Amethyst_02\", \"name\" : \"Flawed Amethyst\", \"icon\" : \"amethyst_02_demonhunter_male\", \"displayColor\" : \"blue\", \"tooltipParams\" : \"recipe/flawed-amethyst\" }}")

func TestUnmarshalRecipe(t *testing.T) {
	Convey("Given a recipe as json", t, func() {
		Convey("It should be unmarshallable", func() {
			var tmp Recipe
			So(json.Unmarshal(rawRecipe, &tmp), ShouldBeNil)
			So(tmp.Id, ShouldEqual, "GemCombine_Amethyst_01_02")
			So(tmp.Name, ShouldEqual, "Flawed Amethyst")
			So(len(tmp.Reagents), ShouldEqual, 1)
			So(tmp.ItemProduced.Id, ShouldEqual, "Amethyst_02")
			So(tmp.ItemProduced.Name, ShouldEqual, "Flawed Amethyst")

			recipe := tmp.normalize()
			So(recipe.ID, ShouldEqual, tmp.Id)
			So(recipe.Name, ShouldEqual, tmp.Name)
			So(len(recipe.Ingredients), ShouldEqual, len(tmp.Reagents))
			So(recipe.Out.Item.ID, ShouldEqual, tmp.ItemProduced.Id)
			So(recipe.Out.Item.Name, ShouldEqual, tmp.ItemProduced.Name)
		})
	})
}
