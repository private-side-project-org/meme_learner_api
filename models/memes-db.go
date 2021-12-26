package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

type DBModel struct {
	DB *sql.DB
}

// returns single meme and error, if any
func (m *DBModel) Get(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, name, password from users where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var user User

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Password,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Printf(`value is %v`, user)

	// get memes if any
	var memes []Meme

	query = `SELECT 
						memes.id, 
						memes.title, 
						memes.description, 
						memes.rating, 
						memes.created_at, 
						memes.updated_at 
					FROM 
						memes 
					INNER JOIN 
						users 
					ON 
						memes.user_id = users.id 
					WHERE 
						user_id = $1`

	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		fmt.Print("error", err)
		return nil, err
	}

	for rows.Next() {
		var meme Meme
		err := rows.Scan(
			&meme.ID,
			&meme.Title,
			&meme.Description,
			&meme.Rating,
			&meme.CreatedAt,
			&meme.UpdatedAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		memes = append(memes, meme)
	}

	user.Memes = memes

	return &user, nil
}

func (m *DBModel) Scraping() (*ScrapingResult, error) {
	sr := ScrapingResult{}
	geziyor.NewGeziyor(&geziyor.Options{
		// scraping resourse url
		StartURLs: []string{"https://knowyourmeme.com/random"},

		// parse function
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			// scrape `title`
			r.HTMLDoc.Find("header > section").Each(func(_ int, s *goquery.Selection) {
				parsedTitle := s.Find("article.entry header section.info h1").Text()
				sr.MemeTitle = strings.Replace(parsedTitle, "\n", "", -1)
			})

			// scrape `image`
			r.HTMLDoc.Find(".photo-wrapper").Each(func(_ int, s *goquery.Selection) {
				im, _ := s.Find("a > img").Attr("data-src")

				sr.Image = im
			})

			// scrape `about`, `origin`, `spread`, if any
			r.HTMLDoc.Find(".entry-section").Each(func(i int, s *goquery.Selection) {
				// `about`
				if i == 0 {
					sr.About = s.Find("h2").Text()
					sr.AboutText = s.Find(".bodycopy p").Text()
				}

				// `origin`
				if i == 1 {
					sr.Origin = s.Find("h2").Text()
					sr.OriginText = s.Find(".bodycopy p").Text()
				}

				// `spread`
				if i == 2 {
					sr.Spread = s.Find("h2").Text()
					sr.SpreadText = s.Find(".bodycopy p").Text()
				}
			})
		},
	}).Start()

	var scrapedMeme ScrapingResult

	scrapedMeme.MemeTitle = sr.MemeTitle
	scrapedMeme.Image = sr.Image
	scrapedMeme.About = sr.About
	scrapedMeme.AboutText = sr.AboutText
	scrapedMeme.Origin = sr.Origin
	scrapedMeme.OriginText = sr.OriginText
	scrapedMeme.Spread = sr.Spread
	scrapedMeme.SpreadText = sr.SpreadText

	fmt.Println(scrapedMeme)

	return &scrapedMeme, nil
}
