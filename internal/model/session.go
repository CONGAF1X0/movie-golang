package model

import (
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

type Session struct {
	SessionID int     `gorm:"primaryKey:session_id"`
	MovieID   int     `json:"movie_id"`
	CinemaID  uint64  `json:"cinema_id"`
	HallID    int     `json:"hall_id"`
	StartTime int64     `json:"start_time"`
	Price     float32 `json:"price"`
}

type SessHall struct {
	Session
	Hall
}

type SessApi struct {
	movie []MovieSimple
	Sess  [][]SessHall
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

func (s Session) List(db *gorm.DB) ([]*Session, error) {
	var sess []*Session
	var err error
	if err = db.Where(&s).Find(&sess).Error; err != nil {
		return nil, err
	}
	return sess, nil
}
func (s Session) ListGroupByMovie(db *gorm.DB) (SessApi, error) {
	var (
		res SessApi
		ids []Session
		err error
	)
	if err = db.Raw("SELECT movie_id FROM `session` WHERE cinema_id=? GROUP BY movie_id", s.CinemaID).Scan(&ids).Error; err != nil {
		return res, err
	}
	for i := 0; i < len(ids); i++ {
		var mov MovieSimple
		mov, err = Movie{MovieID: ids[i].MovieID}.GetSimple(db)
		res.movie = append(res.movie, mov)
		if err != nil {
			return res, err
		}
		var ses []SessHall
		if err = db.Raw("select * from `session` as s,hall as h where s.hall_id=h.hall_id and "+
			"movie_id=? and s.cinema_id=? and start_time>?", ids[i].MovieID, s.CinemaID,time.Now().Unix()).Scan(&ses).Error; err != nil {
			return res, err
		}
		res.Sess = append(res.Sess, ses)
	}
	return res, nil
}

func (s Session) StrTime(db *gorm.DB) (string, error) {
	var (
		str string
		err error
		arr []Session
	)

	if err = db.Where(&Session{MovieID: s.MovieID,CinemaID: s.CinemaID}).Where("to_days(FROM_UNIXTIME(start_time)) = to_days(FROM_UNIXTIME(?))",s.StartTime).Order("start_time").Find(&arr).Error; err != nil {
		return "", err
	}
	for i := 0; i < len(arr); i++ {
		str += time.Unix(arr[i].StartTime,0).Format("15:04") + " | "
	}
	return strings.TrimRight(str," | "),nil
}

func (s Session) Create(db *gorm.DB) error {
	return db.Create(&s).Error
}

func (s Session) Update(db *gorm.DB) error {
	return db.Model(&Session{}).Where("session_id = ?", s.SessionID).Update(s).Error
}

func (s Session) Delete(db *gorm.DB) error {
	return db.Where("session_id = ?", s.SessionID).Delete(&s).Error
}
