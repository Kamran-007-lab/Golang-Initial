package controllers

import (
	"fmt"
	"net/http"
	"server/config"
	"server/models"

	"github.com/labstack/echo/v4"
)

func CreateTeam(c echo.Context) error {
	fmt.Print("Hello 14")
	team := new(models.Team)
	if err := c.Bind(team); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	if err := config.DB.Create(team).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Could not create team",
		})
	}

	return c.JSON(http.StatusOK, team)
}

func AddUserToTeam(c echo.Context) error {
	fmt.Println("In add user to team controller after compiler daemon")
	var req struct {
		UserID uint `json:"user_id"`
		TeamID uint `json:"team_id"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}
	fmt.Println(req.UserID, " ", req.TeamID)
	var team models.Team
	if err := config.DB.Preload("Users").First(&team, req.TeamID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Team not found",
		})
	}

	var user models.User
	if err := config.DB.First(&user, req.UserID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "User not found",
		})
	}

	team.Users = append(team.Users, user)
	config.DB.Save(&team)

	return c.JSON(http.StatusOK, team)
}

func RemoveUserFromTeam(c echo.Context) error {
	var req struct {
		UserID uint `json:"user_id"`
		TeamID uint `json:"team_id"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	var team models.Team
	if err := config.DB.Preload("Users").First(&team, req.TeamID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Team not found",
		})
	}

	var updatedUsers []models.User
	for _, user := range team.Users {
		if user.ID != req.UserID {
			updatedUsers = append(updatedUsers, user)
		}
	}

	team.Users = updatedUsers
	config.DB.Save(&team)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User removed from team",
	})
}
