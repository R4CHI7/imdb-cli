package main

import (
	"fmt"
	"github.com/R4CHI7/go-omdb"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	l := log.New(os.Stdout, "[imdb-cli] ", 0)
	app := cli.NewApp()
	app.Version = "1.0"
	app.Name = "imdb-cli"
	app.Usage = "Get information about movies/series straight to your command line!"
	app.EnableBashCompletion = true
	app.Action = func(c *cli.Context) {
		if c.NArg() < 1 {
			println("IMDb CLI tool. Please run with -h to see options.")
		}

		searchTitle := c.Args().Get(0)
		l.Printf("Fetching information for %s", searchTitle)
		response, err := omdb.SearchByTitle(searchTitle)
		if err != nil {
			l.Print("Oops! Some error occurred while fetching movie details. Please try again later.")
			return
		}
		l.Print("Information fetched successfully!")
		displayMovieInfo(response)
	}
	app.Run(os.Args)
}

func displayMovieInfo(response omdb.SingleResultResponse) {
	fmt.Println("Title:", response.Title)
	fmt.Println("Year of release:", response.Year)
	fmt.Println("Runtime:", getDisplayableRuntime(response.Runtime))
	fmt.Println("Genre(s):", response.Genre)
	fmt.Println("Directed By:", response.Director)
	fmt.Println("Actors:", response.Actors)
	fmt.Println("Plot:", response.Plot)
	fmt.Println("Language:", response.Language)
	fmt.Println("Rating on IMDb:", response.ImdbRating)
	fmt.Println("IMDb Link:", getImdbLink(response.ImdbId))
}
