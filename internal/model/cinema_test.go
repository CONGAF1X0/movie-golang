package model

import (
	"github.com/jinzhu/gorm"
	"reflect"
	"testing"
)


func TestCinema_HotMovieList1(t *testing.T) {
	//fmt.Println(time.Date(2022,1,1,10,00,0,0,time.Local).Local().Unix())
	type fields struct {
		CinemaID   uint64
		CinemaName string
		province   string
		City       string
		District   string
		Location   string
		Longitude  float64
		Latitude   float64
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []MovieSimple
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "1",
			fields:  fields{City: "广州"},
			args:    args{db: NewDB()},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Cinema{
				CinemaID:   tt.fields.CinemaID,
				CinemaName: tt.fields.CinemaName,
				province:   tt.fields.province,
				City:       tt.fields.City,
				District:   tt.fields.District,
				Location:   tt.fields.Location,
				Longitude:  tt.fields.Longitude,
				Latitude:   tt.fields.Latitude,
			}
			got, err := c.HotMovieList(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("HotMovieList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HotMovieList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCinema_CinemaSessList(t *testing.T) {
	type fields struct {
		CinemaID   uint64
		CinemaName string
		province   string
		City       string
		District   string
		Location   string
		Longitude  float64
		Latitude   float64
	}
	type args struct {
		db         *gorm.DB
		session    Session
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
			fields:  fields{City: "广州"},
			args:    args{db: NewDB(),session: Session{MovieID: 1,StartTime: 1640933321}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Cinema{
				CinemaID:   tt.fields.CinemaID,
				CinemaName: tt.fields.CinemaName,
				province:   tt.fields.province,
				City:       tt.fields.City,
				District:   tt.fields.District,
				Location:   tt.fields.Location,
				Longitude:  tt.fields.Longitude,
				Latitude:   tt.fields.Latitude,
			}
			got, err := c.CinemaSessList(tt.args.db, tt.args.session, tt.args.pageOffset, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("CinemaSessList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CinemaSessList() got = %v, want %v", got, tt.want)
			}
		})
	}
}