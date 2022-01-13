package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Cinema struct {
	CinemaID   uint64  `json:"cinema_id"`
	CinemaName string  `json:"cinema_name"`
	province   string  `json:"province"`
	City       string  `json:"city"`
	District   string  `json:"district"`
	Location   string  `json:"location"`
	Longitude  float64 `json:"longitude"`
	Latitude   float64 `json:"latitude"`
}

func (c Cinema) TableName() string {
	return "cinema"
}

type CinemaSess struct {
	Cinema
	Time string `json:"time"`
}

func (c Cinema) Get(db *gorm.DB) (Cinema, error) {
	var cin Cinema
	var err error
	if err = db.Where(&c).Find(&cin).Error; err != nil {
		return cin, err
	}
	return cin, nil
}

func (c Cinema) List(db *gorm.DB, pageOffset, pageSize int) ([]*Cinema, error) {
	var cinema []*Cinema
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if c.CinemaName != "" {
		query := "%" + c.CinemaName + "%"
		db = db.Where("(cinema_name like ? or location like ?)", query,query)
	}
	if c.District != "" {
		db = db.Where("district = ?", c.District)
	}
	db = db.Where("city = ?", c.City)
	if err = db.Find(&cinema).Error; err != nil {
		return nil, err
	}
	return cinema, nil
}

func (c Cinema) HotMovies(db *gorm.DB) ([]MovieSimple, error) {
	var res []MovieSimple
	var err error
	if db.Raw("SELECT movie_id FROM `session`as s,cinema as c WHERE c.cinema_id=s.cinema_id AND city=? GROUP BY movie_id ORDER BY count(1) DESC",
		c.City).Scan(&res); err != nil {
		return nil, err
	}
	for i := 0; i < len(res); i++ {
		res[i], err = Movie{MovieID: res[i].MovieID}.GetSimple(db)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}
func (c Cinema) CinemaSessList(db *gorm.DB, session Session, pageOffset, pageSize int) ([]CinemaSess, error) {
	var cin []CinemaSess
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if session.StartTime==0{
		session.StartTime=time.Now().Unix()
	}
	if c.CinemaName != "" {
		query := "%" + c.CinemaName + "%"
		err = db.Raw("select * from `session` as s,cinema as c WHERE s.cinema_id=c.cinema_id AND "+
			"movie_id= ? AND city= ? AND start_time > ? AND (cinema_name like ? or location like ?)",
			session.MovieID, c.City, session.StartTime,query,query).Group("c.cinema_id").Scan(&cin).Error
	} else {
		if time.Now().Day() == time.Unix(session.StartTime, 0).Day() {
			err = db.Raw("select * from `session` as s,cinema as c WHERE s.cinema_id=c.cinema_id AND "+
				"movie_id= ? AND city= ? AND to_days(FROM_UNIXTIME(start_time)) = to_days(FROM_UNIXTIME(?)) AND start_time > ?",
				session.MovieID, c.City, session.StartTime, session.StartTime).Group("c.cinema_id").Scan(&cin).Error
		} else {
			err = db.Raw("select * from `session` as s,cinema as c WHERE s.cinema_id=c.cinema_id AND "+
				"movie_id= ? AND city= ? AND to_days(FROM_UNIXTIME(start_time)) = to_days(FROM_UNIXTIME(?))",
				session.MovieID, c.City, session.StartTime).Group("c.cinema_id").Scan(&cin).Error
		}
		if err != nil {
			return nil, err
		}
		for i := 0; i < len(cin); i++ {
			cin[i].Time, _ = Session{MovieID: session.MovieID, CinemaID: cin[i].CinemaID, StartTime: session.StartTime}.StrTime(db)
		}
	}
	return cin, nil
}

func (c Cinema) Create(db *gorm.DB) error {
	return db.Create(&c).Error
}

func (c Cinema) Update(db *gorm.DB) error {
	return db.Model(&Cinema{}).Where("cinema_id = ?", c.CinemaID).Update(c).Error
}

func (c Cinema) Delete(db *gorm.DB) error {
	return db.Where("cinema_id = ?", c.CinemaID).Delete(&c).Error
}
