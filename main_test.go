package main

import "testing"

func TestAdd(t *testing.T) {
	t.Parallel()
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "two plus two",
			args: args{
				a: 2,
				b: 2,
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "big numbers",
			args: args{
				a: 99999999999,
				b: 99999999999,
			},
			want:    199999999998,
			wantErr: false,
		},
		{
			name: "negative a",
			args: args{
				a: -1,
				b: 2,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "negative b",
			args: args{
				a: 1,
				b: -2,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "a is zero",
			args: args{
				a: 0,
				b: 2,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "b is zero",
			args: args{
				a: 2,
				b: 0,
			},
			want:    2,
			wantErr: false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
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
