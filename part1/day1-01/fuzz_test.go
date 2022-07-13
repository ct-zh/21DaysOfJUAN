package fuzz_t

import "testing"

func TestEqual(t *testing.T) {
	type args struct {
		a []byte
		b []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"right case", args{
			a: []byte{'f', 'u', 'z', 'z'},
			b: []byte{'f', 'u', 'z', 'z'},
		}, true},
		{"right case", args{
			a: []byte{'a', 'b', 'c'},
			b: []byte{'b', 'c', 'd'},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func FuzzEqual(f *testing.F) {
	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		Equal(a, b)
	})
}
