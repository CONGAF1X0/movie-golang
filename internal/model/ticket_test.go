package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
	"testing"
)

func TestTicket_List(t1 *testing.T) {
	db,err := gorm.Open("mysql","root:123456@tcp(127.0.0.1:3306)/movie?charset=utf8mb4&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	type fields struct {
		TicketID int
		UID uint64
		SessionID  uint64
		Seat     string
	}
	type args struct {
		db         *gorm.DB
		pageOffset int
		pageSize   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Ticket
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "1",
			fields: fields{UID: 1},
			args: args{db: db},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Ticket{
				TicketID: tt.fields.TicketID,
				UID:     tt.fields.UID,
				SessionID:  tt.fields.SessionID,
				Seat:     tt.fields.Seat,
			}
			got, err := t.List(tt.args.db, tt.args.pageOffset, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t1.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}
