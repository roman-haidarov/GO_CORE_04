package reversestr

import "testing"

func TestRev(t *testing.T) {
	tests := []struct {
		name string
		s string
		want string
	}{
		{
			name: "test1",
			s: "ABC",
			want: "CBA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rev(tt.s); got != tt.want {
				t.Errorf("Rev() = %v, want %v", got, tt.want)
			}
		})
	}
}
