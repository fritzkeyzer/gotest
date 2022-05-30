package stats

import "testing"

func TestSMA(t *testing.T) {
	type args struct {
		data   []*float64
		window int
	}
	tests := []struct {
		name string
		args args
		want *float64
	}{
		// TODO: Add test cases.
		{
			name: "0 - 3",
			args: args{
				data:   []*float64{f(0), f(1), f(2)},
				window: 3,
			},
			want: f(1),
		},
		{
			name: "0 - 9",
			args: args{
				data:   []*float64{nil, f(1), f(2), nil, f(4), f(5), f(6), f(7), nil, f(9)},
				window: 3,
			},
			want: f(8),
		},
		{
			name: "0 - 9",
			args: args{
				data:   []*float64{f(0), f(1), f(2), f(3), f(4), f(5), f(6), f(7), f(8), f(9)},
				window: 10,
			},
			want: f(4.5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SMA(tt.args.data, tt.args.window)

			if got == nil {
				if tt.want != nil {
					t.Errorf("SMA() = <nil>, want = %v", *tt.want)
				}
				return
			}

			if tt.want == nil {
				if got != nil {
					t.Errorf("SMA() = %v, want = <nil>", *got)
				}
				return
			}

			if *got != *tt.want {
				t.Errorf("SMA() = %v, want = %v", *got, *tt.want)
			}
		})
	}
}

func f(v float64) *float64 {
	return &v
}
