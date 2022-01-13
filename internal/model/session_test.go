package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
	"testing"
	"time"
)

func TestSession_List(t *testing.T) {
	db,err := gorm.Open("mysql","root:123456@tcp(127.0.0.1:3306)/movie?charset=utf8mb4&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	type fields struct {
		SessionID int
		MovieID   int
		CinemaID  uint64
		HallID    int
		StartTime int64
		Price     float32
	}
	type args struct {
		db         *gorm.DB
		city       string
		pageOffset int
		pageSize   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []CinemaSess
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "1",
			fields:  fields{MovieID: 1},
			args:    args{db: db},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Session{
				SessionID: tt.fields.SessionID,
				MovieID:   tt.fields.MovieID,
				CinemaID:  tt.fields.CinemaID,
				HallID:    tt.fields.HallID,
				StartTime: tt.fields.StartTime,
				Price:     tt.fields.Price,
			}
			got, err := s.CinemaSessList(tt.args.db,Cinema{City: "广州"}, tt.args.pageOffset, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_List2(t *testing.T) {
	db,err := gorm.Open("mysql","root:123456@tcp(127.0.0.1:3306)/movie?charset=utf8mb4&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	type fields struct {
		SessionID int
		MovieID   int
		CinemaID  uint64
		HallID    int
		StartTime int64
		Price     float32
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Session
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "1",
			fields:  fields{CinemaID: 1,MovieID: 1},
			args:    args{db: db},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Session{
				SessionID: tt.fields.SessionID,
				MovieID:   tt.fields.MovieID,
				CinemaID:  tt.fields.CinemaID,
				HallID:    tt.fields.HallID,
				StartTime: tt.fields.StartTime,
				Price:     tt.fields.Price,
			}
			got, err := s.List(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("List2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List2() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_MovieList(t *testing.T) {
	db,err := gorm.Open("mysql","root:123456@tcp(127.0.0.1:3306)/movie?charset=utf8mb4&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	type fields struct {
		SessionID int
		MovieID   int
		CinemaID  uint64
		HallID    int
		StartTime int64
		Price     float32
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    [][]SessHall
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "1",
			fields:  fields{CinemaID: 1},
			args:    args{db: db},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Session{
				SessionID: tt.fields.SessionID,
				MovieID:   tt.fields.MovieID,
				CinemaID:  tt.fields.CinemaID,
				HallID:    tt.fields.HallID,
				StartTime: tt.fields.StartTime,
				Price:     tt.fields.Price,
			}
			got, err := s.ListGroupByMovie(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("MovieList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovieList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_TimeList(t *testing.T) {
	fmt.Println(time.Now().Unix())
	type fields struct {
		SessionID int
		MovieID   int
		CinemaID  uint64
		HallID    int
		StartTime int64
		Price     float32
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "1",
			fields:  fields{CinemaID: 1,MovieID: 1,StartTime: time.Date(2022,1,1,0,0,0,0,time.Local).Unix()},
			args:    args{db: NewDB()},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Session{
				SessionID: tt.fields.SessionID,
				MovieID:   tt.fields.MovieID,
				CinemaID:  tt.fields.CinemaID,
				HallID:    tt.fields.HallID,
				StartTime: tt.fields.StartTime,
				Price:     tt.fields.Price,
			}
			got, err := s.StrTime(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TimeList() got = %v, want %v", got, tt.want)
			}
		})
	}
}