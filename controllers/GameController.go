package controllers

import (
	"github.com/kataras/iris/v12"

	"aatrox/chess/models"
	"aatrox/chess/services"
)

type GameController struct {
	Ctx iris.Context
}

func (c *GameController) Get() {
	var games []models.Game
	services.GamesSelectAll(&games)

	if len(games) == 0 {
		c.Ctx.ViewData("msg", "No games found")
	}
	c.Ctx.ViewData("games", games)

	c.Ctx.View("games/games.html")
}
