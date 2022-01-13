package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
	"testing"
)

func TestActor_List(t *testing.T) {
	db,err := gorm.Open("mysql","root:123456@tcp(127.0.0.1:3306)/movie?charset=utf8mb4&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	type fields struct {
		ActorID      int
		Name1        string
		Name2        string
		Birthday     string
		Introduction string
		Image        string
	}
	type args struct {
		db  *gorm.DB
		ids []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Actor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "1",
			fields: fields{},
			args: args{db: db,ids: []int{1,3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Actor{
				ActorID:      tt.fields.ActorID,
				Name1:        tt.fields.Name1,
				Name2:        tt.fields.Name2,
				Birthday:     tt.fields.Birthday,
				Introduction: tt.fields.Introduction,
				Image:        tt.fields.Image,
			}
			got, err := a.List(tt.args.db, tt.args.ids)
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
