package model

import (
	"TicketSales/pkg/convert"
	"github.com/jinzhu/gorm"
	"strings"
)

type Movie struct {
	MovieID   int    `json:"movie_id"`
	MovieName string `json:"movie_name"`
	Director  string `json:"director"`
	StarsIds  string `json:"stars_ids"`
	Genre     string `json:"genre"`
	Storyline string `json:"storyline"`
	Runtime   string `json:"runtime"`
	Release   int64  `json:"release"`
	Rating    string `json:"rating"`
	BoxOffice string `json:"box_office"`
	Cover     string `json:"cover"`
}

type MovieSimple struct {
	MovieID   int    `json:"movie_id"`
	MovieName string `json:"movie_name"`
	Director  string `json:"director"`
	StarsIds  string `json:"stars_ids"`
	Stars     string `json:"stars"`
	Runtime   string `json:"runtime"`
	Genre     string `json:"genre"`
	Rating    string `json:"rating"`
	Cover     string `json:"cover"`
}

func (m Movie) TableName() string {
	return "movie"
}

type MovieApi struct {
	Movie
	Stars []*Actor `json:"stars"`
}

func (m Movie) Get(db *gorm.DB) (MovieApi, error) {
	var movie MovieApi
	var err error
	if err = db.Where("movie_id = ?", m.MovieID).First(&movie.Movie).Error; err != nil {
		return movie, err
	}
	ids := convert.Str2IntSlice(strings.Split(movie.Movie.StarsIds, ","))
	var actor Actor
	movie.Stars, err = actor.List(db, ids)
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func (m Movie) GetRuntime(db *gorm.DB) (string, error) {
	var movie Movie
	if err := db.Raw("select runtime from movie where movie_id = ?", m.MovieID).First(&movie).Error; err != nil {
		return "", err
	}
	return movie.Runtime, nil
}

func (m Movie) GetSimple(db *gorm.DB) (MovieSimple, error) {
	var movie MovieSimple
	var err error
	if err = db.Table("movie").Where("movie_id = ?", m.MovieID).First(&movie).Error; err != nil {
		return movie, err
	}
	ids := convert.Str2IntSlice(strings.Split(movie.StarsIds, ","))
	if ids[0] != 0 {
		movie.Stars, err = Actor{}.StrList(db, ids)
		if err != nil {
			return movie, err
		}
	}
	return movie, nil
}

func (m Movie) List(db *gorm.DB, pageOffset, pageSize int) ([]*Movie, error) {
	var movies []*Movie
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if m.MovieName != "" {
		db = db.Where("movie_name like ?", "%"+m.MovieName+"%")
	}
	if err = db.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (m Movie) Create(db *gorm.DB) error {
	return db.Create(&m).Error
}

func (m Movie) Update(db *gorm.DB) error {
	return db.Model(&Movie{}).Where("movie_id = ?", m.MovieID).Update(m).Error
}

func (m Movie) Delete(db *gorm.DB) error {
	return db.Where("movie_id = ?", m.MovieID).Delete(&m).Error
}
