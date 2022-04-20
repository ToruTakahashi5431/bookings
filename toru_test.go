package main

/*
func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "normal_1", args: args{a: 1, b: 2}, want: 3},
		{name: "normal_2", args: args{a: 2, b: 3}, want: 5},
		{name: "normal_3", args: args{a: 3, b: 4}, want: 7},
	}

	for _, tt := range tests {

		tt := tt

		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Fatalf("add() = %v, but want %v", got, tt.want)
			}

		})
	}
}

func add(a, b int) int {
	time.Sleep(time.Duration(a+b) * time.Second)
	return a + b
}
*/
