package handlers

import (
	"errors"
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/utils/api"
)

// @Summary Deletes an album
// @Tags albums
// @Produce  json
// @Success 200 {object} api.Response
// @Failure 500 {object} api.Response
// @Failure 400 {object} api.Response
// @Failure 404 {object} api.Response
// @Param id path string true "Album ID"
// @Param Authorization header string true "Authorization"
// @Router /api/v1/albums/{id} [delete]
func DeleteAlbumById(res http.ResponseWriter, req *http.Request) {
	id := api.PathVar(req, 1)
	var response api.Response

	if id == "" {
		api.Error(res, req, "ID is required", errors.New("ID is required"), http.StatusBadRequest)
		return
	}

	err := repositories.DeleteAlbumById(id)
	if err != nil {
		api.Error(res, req, "Error while deleting album", err, http.StatusInternalServerError)
		return
	}

	response = api.Response{
		Message: "Album deleted",
		Data:    "Provided ID: " + id,
	}
	api.JSON(res, response, http.StatusOK)
}
