package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("id is ", id)

	movie, err := app.models.DB.Get(id)

	// movie := models.Movie{
	// 	ID:          id,
	// 	Title:       "Some movie",
	// 	Description: "Some desc",
	// 	Year:        2021,
	// 	ReleaseDate: time.Date(2020, 01, 01, 01, 0, 0, 0, time.Local),
	// 	Runtime:     100,
	// 	Rating:      5,
	// 	MPAARating:  "PG-13",
	// 	CreatedAt:   time.Now(),
	// 	UpdatedAt:   time.Now(),
	// }

	err = app.writJSON(w, http.StatusOK, movie, "movie")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllMovie(w http.ResponseWriter, r *http.Request) {
	movies, err := app.models.DB.All()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}
