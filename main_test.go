package main

import "testing"

func TestAdd(t *testing.T) {
	t.Parallel()
	type args struct {
		a int
		b int
	}
	tests := map[string]struct {
		args    args
		want    int
		wantErr bool
	}{
		"two plus two": {
			args: args{
				a: 2,
				b: 2,
			},
			want:    4,
			wantErr: false,
		},
		"big numbers": {
			args: args{
				a: 99999999999,
				b: 99999999999,
			},
			want:    199999999998,
			wantErr: false,
		},
		"negative a": {
			args: args{
				a: -1,
				b: 2,
			},
			want:    0,
			wantErr: true,
		},
		"negative b": {
			args: args{
				a: 1,
				b: -2,
			},
			want:    0,
			wantErr: true,
		},
		"a is zero": {
			args: args{
				a: 0,
				b: 2,
			},
			want:    2,
			wantErr: false,
		},
		"b is zero": {
			args: args{
				a: 2,
				b: 0,
			},
			want:    2,
			wantErr: false,
		},
	}
	for name, testCase := range tests {
		name := name
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := Add(testCase.args.a, testCase.args.b)
			if testCase.wantErr && err == nil {
				t.Error("wanted an error, but got nil")
			}
			if got != testCase.want {
				t.Errorf("got %v, want %v", got, testCase.want)
			}

		})
	}
}
