package model

import (
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

type Session struct {
	SessionID int     `gorm:"primaryKey:session_id" json:"session_id"`
	MovieID   int     `json:"movie_id"`
	CinemaID  uint64  `json:"cinema_id"`
	HallID    int     `json:"hall_id"`
	StartTime int64   `json:"start_time"`
	EndTime   int64   `json:"end_time"`
	Price     float32 `json:"price"`
}

type SessHall struct {
	SessionID int     `json:"session_id"`
	MovieID   int     `json:"movie_id"`
	CinemaID  uint64  `json:"cinema_id"`
	HallID    int     `json:"hall_id"`
	StartTime int64   `json:"start_time"`
	EndTime   int64   `json:"end_time"`
	Price     float32 `json:"price"`
	HallName  string  `json:"hall_name"`
	Capacity  string  `json:"capacity"`
}
type SessInfo struct {
	Session
	Cinema
	Hall
	MovieSimple
}

type MovieSess struct {
	MovieSimple
	Session []SessHall
}

func (s Session) TableName() string {
	return "session"
}

func (s Session) CinemaSessList(db *gorm.DB, cinema Cinema, pageOffset, pageSize int) ([]CinemaSess, error) {
	var sess []CinemaSess
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if err = db.Raw("select * from `session` as s,cinema as c WHERE s.cinema_id=c.cinema_id AND "+
		"movie_id= ? AND city= ?",
		s.MovieID, cinema.City).Group("c.cinema_id").Scan(&sess).Error; err != nil {
		return nil, err
	}

	return sess, nil
}
func (s Session) Get(db *gorm.DB) (SessInfo, error) {
	var (
		res SessInfo
		err error
	)

	if err = db.Raw("select * from session as s,hall as h,movie as m,cinema as c where s.cinema_id=c.cinema_id and  s.hall_id = h.hall_id and s.movie_id = m.movie_id and session_id = ?", s.SessionID).First(&res).Error; err != nil {
		return SessInfo{}, err
	}
	return res, nil
}
func (s Session) List(db *gorm.DB) ([]*Session, error) {
	var sess []*Session
	var err error
	if err = db.Where(&s).Find(&sess).Error; err != nil {
		return nil, err
	}
	return sess, nil
}

func (s Session) ListGroupByMovie(db *gorm.DB) ([]MovieSess, error) {
	var (
		res []MovieSess
		ids []Session
		err error
	)
	if err = db.Raw("SELECT movie_id FROM `session` WHERE cinema_id=? GROUP BY movie_id", s.CinemaID).Scan(&ids).Error; err != nil {
		return res, err
	}
	for i := 0; i < len(ids); i++ {
		var mov MovieSimple
		mov, err = Movie{MovieID: ids[i].MovieID}.GetSimple(db)
		if err != nil {
			return res, err
		}
		ms := MovieSess{mov, []SessHall{}}

		if err = db.Raw("select * from `session` as s,hall as h where s.hall_id=h.hall_id and "+
			"movie_id=? and s.cinema_id=? and start_time>? order by start_time ", ids[i].MovieID, s.CinemaID, s.StartTime).Scan(&ms.Session).Error; err != nil {
			return res, err
		}
		if len(ms.Session) != 0 {
			res = append(res, ms)
		}
	}
	return res, nil
}

func (s Session) StrTime(db *gorm.DB) (string, error) {
	var (
		str string
		err error
		arr []Session
	)

	if err = db.Where(&Session{MovieID: s.MovieID, CinemaID: s.CinemaID}).Where("to_days(FROM_UNIXTIME(start_time/1000)) = to_days(FROM_UNIXTIME(?))", s.StartTime/1000).Order("start_time").Find(&arr).Error; err != nil {
		return "", err
	}
	for i := 0; i < len(arr); i++ {
		str += time.Unix(arr[i].StartTime/1000, 0).Format("15:04") + " | "
	}
	return strings.TrimRight(str, " | "), nil
}

func (s Session) Create(db *gorm.DB) error {
	return db.Create(&s).Error
}

func (s Session) Update(db *gorm.DB) error {
	return db.Model(&Session{}).Where("session_id = ? and cinema_id = ?", s.SessionID, s.CinemaID).Update(s).Error
}

func (s Session) Delete(db *gorm.DB) error {
	return db.Where("session_id = ?", s.SessionID).Delete(&s).Error
}
