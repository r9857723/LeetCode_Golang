package toLowerCase

import (
	"reflect"
	"testing"
)

func Test_DefangIPaddr(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1:",
			args: args{
				"Hello",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toLowerCase(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defangIPaddr() = %v, wnat %v", got, tt.want)
			}
		})
	}

}
