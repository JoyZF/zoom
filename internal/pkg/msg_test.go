package pkg

import "testing"

func TestGetMsg(t *testing.T) {
	type args struct {
		code int
		lang string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMsg(tt.args.code, tt.args.lang); got != tt.want {
				t.Errorf("GetMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}
