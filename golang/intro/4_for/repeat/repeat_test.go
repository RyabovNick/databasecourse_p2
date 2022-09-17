package repeat

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	exp := "aaaaa"

	if got != exp {
		t.Errorf("got %q, but exp %q", got, exp)
	}
}

// BenchmarkRepeat-12    	    8940	    112090 ns/op	  130913 B/op	     499 allocs/op
// BenchmarkRepeat-12    	       3	 357814733 ns/op	1352371986 B/op	   50062 allocs/op
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

// BenchmarkRepeatSl-12    	   78314	     15229 ns/op	   17392 B/op	      11 allocs/op
// BenchmarkRepeatSl-12    	  111925	     10399 ns/op	    9216 B/op	       2 allocs/op
// BenchmarkRepeatSl-12    	     681	   2425627 ns/op	  909318 B/op	       2 allocs/op
func BenchmarkRepeatSl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatSl("a", 5)
	}
}
