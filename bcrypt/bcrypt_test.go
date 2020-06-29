package bcrypt

import (
	"reflect"
	"testing"
)

func Test_bcryptClientImpl_GenerateFromPassword(t *testing.T) {
	client := bcryptClientImpl{}
	type args struct {
		password []byte
		cost     int
	}
	tests := []struct {
		name    string
		c       *bcryptClientImpl
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "normal case",
			c:    &client,
			args: args{
				password: []byte("ahihihi"),
				cost:     14,
			},
			want:    []byte("ahihihi"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &bcryptClientImpl{}
			got, err := c.GenerateFromPassword(tt.args.password, tt.args.cost)
			if (err != nil) != tt.wantErr {
				t.Errorf("bcryptClientImpl.GenerateFromPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bcryptClientImpl.GenerateFromPassword() = %v, want %v", string(got), tt.want)
			}
		})
	}
}
