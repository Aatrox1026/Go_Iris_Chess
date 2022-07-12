package controllers

import (
	"github.com/kataras/iris/v12"

	"aatrox/chess/models"
	"aatrox/chess/services"
)

type GameController struct {
	Ctx     iris.Context
	service services.GameService
}

func (c *GameController) Get() {
	var games *[]models.Game
	games = c.service.GetAll()

	if len(*games) == 0 {
		c.Ctx.ViewData("msg", "No data found")
	}
	c.Ctx.ViewData("games", *games)

	c.Ctx.View("games/games.html")
}

func (c *GameController) GetBy(id int64) {
	var game *models.Game
	game = c.service.GetByID(id)

	if game == nil {
		c.Ctx.StatusCode(iris.StatusNotFound)
		return
	}

	c.Ctx.JSON(game)
}
