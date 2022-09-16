package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockChecker(url string) bool {
	return url != "c"
}

func TestCheckWebsites(t *testing.T) {

	type args struct {
		wc   Checker
		urls []string
	}
	tests := []struct {
		name string
		args args
		want map[string]bool
	}{
		{
			name: "1",
			args: args{
				wc: mockChecker,
				urls: []string{
					"a",
					"b",
					"c",
				},
			},
			want: map[string]bool{
				"a": true,
				"b": true,
				"c": false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckWebsites(tt.args.wc, tt.args.urls); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckWebsites() = %v, want %v", got, tt.want)
			}
		})
	}
}

func slowChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowChecker, urls)
	}
}
