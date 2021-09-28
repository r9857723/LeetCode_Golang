package gotest

import (
	"reflect"
	"testing"
)

func Test_DefangIPaddr(t *testing.T) {
	type args struct {
		address string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1:",
			args: args{
				"1.1.1.1",
			},
			want: "1[.]1[.]1[.]1",
		},
		{
			name: "Example 2:",
			args: args{
				"255.100.50.0",
			},
			want: "255[.]100[.]50[.]0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defangIPaddr(tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defangIPaddr() = %v, wnat %v", got, tt.want)
			}
		})
	}

}
