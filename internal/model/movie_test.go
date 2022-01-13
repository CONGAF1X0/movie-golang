package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
	"testing"
)

func TestMovie_Get(t *testing.T) {
	db,err := gorm.Open("mysql","root:123456@tcp(127.0.0.1:3306)/movie?charset=utf8mb4&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	type fields struct {
		MovieID   int
		MovieName string
		Director  string
		StarsIds     string
		Genre     string
		Storyline string
		Runtime   string
		Release   int
		Rating    string
		BoxOffice string
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    MovieApi
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "1",
			fields: fields{MovieID: 1},
			args:   args{db},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Movie{
				MovieID:   tt.fields.MovieID,
				MovieName: tt.fields.MovieName,
				Director:  tt.fields.Director,
				StarsIds:     tt.fields.StarsIds,
				Genre:     tt.fields.Genre,
				Storyline: tt.fields.Storyline,
				Runtime:   tt.fields.Runtime,
				Release:   tt.fields.Release,
				Rating:    tt.fields.Rating,
				BoxOffice: tt.fields.BoxOffice,
			}
			got, err := m.Get(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got.Stars[0], tt.want)
			}
		})
	}
}
