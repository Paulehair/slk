package middlewares

import (
	"fmt"

	"github.com/blyndusk/salika-pagination/pkg"
)

type Movie struct {
	Title       string `json:"title"`
	Category    string `json:"category"`
	TotalRental uint   `json:"total_rental"`
}

func GetMoviesWithPages(asc string, orderby string, limit, offset int) ([]Movie, error) {
	dbclient := pkg.GetDBClient()
	rows, err := dbclient.Query(fmt.Sprintf(`
		WITH film_table AS (
			SELECT film.title, film.film_id, category.name AS category
			FROM category,
				film_category,
				film
			WHERE film_category.category_id = category.category_id
			AND film_category.film_id = film.film_id
		)
		SELECT film_table.title, film_table.category, count(rental.rental_id) AS total_rental
		FROM film_table,
		inventory,
		rental
		WHERE film_table.film_id = inventory.film_id
		AND inventory.inventory_id = rental.inventory_id
		GROUP BY film_table.title
		ORDER BY %s %s
		LIMIT %d OFFSET %d;
	`, orderby, asc, limit, offset))

	if err != nil {
		return nil, err
	}

	var movies []Movie

	for rows.Next() {
		var movie Movie
		err = rows.Scan(&movie.Title, &movie.Category, &movie.TotalRental)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func CountPages(limit int) (int, error) {
	dbclient := pkg.GetDBClient()
	rows, err := dbclient.Query(fmt.Sprintf(`select ceil(count(*)/%d) as total_pages from film;`, limit))

	if err != nil {
		return 0, err
	}

	var count int
	for rows.Next() {
		err = rows.Scan(&count)
	}

	return count, err
}
