package simpleaccessmodel

import (
	"testing"

	"github.com/zapililirad/securedor/accessmodel"
)

func TestSimpleAccessModel_IsAccessValid(t *testing.T) {
	type fields struct {
		required SimpleAccessType
	}
	type args struct {
		current accessmodel.AccessType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "requires read access, read provided",
			fields: fields{Read},
			args:   args{accessmodel.AccessType(Read)},
			want:   true,
		},
		{
			name:   "requires read access, readwrite provided",
			fields: fields{Read},
			args:   args{accessmodel.AccessType(ReadWrite)},
			want:   true,
		},
		{
			name:   "requires read access, write provided",
			fields: fields{Read},
			args:   args{accessmodel.AccessType(Write)},
			want:   false,
		},
		{
			name:   "requires write access, read provided",
			fields: fields{Write},
			args:   args{accessmodel.AccessType(Read)},
			want:   false,
		},
		{
			name:   "requires write access, readwrite provided",
			fields: fields{Write},
			args:   args{accessmodel.AccessType(ReadWrite)},
			want:   true,
		},
		{
			name:   "requires write access, write provided",
			fields: fields{Write},
			args:   args{accessmodel.AccessType(Write)},
			want:   true,
		},
		{
			name:   "requires readwrite access, read provided",
			fields: fields{ReadWrite},
			args:   args{accessmodel.AccessType(Read)},
			want:   false,
		},
		{
			name:   "requires readwrite access, readwrite provided",
			fields: fields{ReadWrite},
			args:   args{accessmodel.AccessType(ReadWrite)},
			want:   true,
		},
		{
			name:   "requires readwrite access, write provided",
			fields: fields{ReadWrite},
			args:   args{accessmodel.AccessType(Write)},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SimpleAccessModel{
				required: tt.fields.required,
			}
			if got := m.IsAccessValid(tt.args.current); got != tt.want {
				t.Errorf("SimpleAccessModel.IsAccessValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
