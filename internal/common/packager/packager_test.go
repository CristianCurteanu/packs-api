package packager

import (
	"reflect"
	"testing"
)

func Test_packager_PackItems(t *testing.T) {
	p := NewPackager(DefaultPackSizes)

	type fields struct {
		packSizes PackSizes
	}
	type args struct {
		items int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Pack
	}{
		{
			args:   args{1},
			fields: fields{DefaultPackSizes},
			want: []Pack{
				{Capacity: 250, StoredItems: 1},
			},
		},
		{
			args:   args{250},
			fields: fields{DefaultPackSizes},
			want: []Pack{
				{Capacity: 250, StoredItems: 250},
			},
		},
		{
			args:   args{251},
			fields: fields{DefaultPackSizes},
			want: []Pack{
				{Capacity: 500, StoredItems: 251},
			},
		},
		{
			args:   args{501},
			fields: fields{DefaultPackSizes},
			want: []Pack{
				{Capacity: 500, StoredItems: 500},
				{Capacity: 250, StoredItems: 1},
			},
		},
		{
			args:   args{12001},
			fields: fields{DefaultPackSizes},
			want: []Pack{
				{Capacity: 5000, StoredItems: 5000},
				{Capacity: 5000, StoredItems: 5000},
				{Capacity: 2000, StoredItems: 2000},
				{Capacity: 250, StoredItems: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := p.PackItems(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("packager.PackItems() = %v, want %v", got, tt.want)
			}
		})
	}
}
