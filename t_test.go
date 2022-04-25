package t

import (
	"reflect"
	"testing"
)

func TestT(t *testing.T) {
	type args struct {
		cond bool
		a    func() string
		b    func() string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "false ",
			args: args{
				cond: false,
				a:    L("true"),
				b:    L("false"),
			},
			want: "false",
		},
		{
			name: "true ",
			args: args{
				cond: true,
				a:    L("true"),
				b:    L("false"),
			},
			want: "true",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := T(tt.args.cond, tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("T() = %v, want %v", got, tt.want)
			}
		})
	}
}
