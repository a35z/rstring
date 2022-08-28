package rstring

import "testing"

func TestReplaceLike(t *testing.T) {
	tests := []struct {
		name string
		s    string
		old  string
		new  string
		want string
	}{
		{"normal", "hello world", "lo", "10", "hel10 world"},
		{"all lower", "how are you", "are", "is", "how is you"},
		{"all upper", "HOW ARE YOU", "are", "is", "HOW IS YOU"},
		{"all title", "How Are You", "are", "is", "How Is You"},
		{"hybrid lower upper", "HOW ARE are Are YOU", "are", "is", "HOW IS is Is YOU"},
		{"overlap", "Hello World", "hello", "hello-first", "Hello-First World"},
		{"overlap small", "iiiiii", "i", "IO", "ioioioioioio"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceLike(tt.s, tt.old, tt.new); got != tt.want {
				t.Errorf("ReplaceLike() ori = %q, got = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func BenchmarkReplaceLike(b *testing.B) {
	s, o, n := "Hello World", "hello", "hello-first"
	for i := 0; i < b.N; i++ {
		_ = ReplaceLike(s, o, n)
	}
}

func TestReplaceAllLike(t *testing.T) {
	rm1 := map[string]string{
		"pluralsight.com":   "hello.org",
		"pluralsight":       "nuage",
		"conference":        "meeting",
		"Coin Sciences Ltd": "TrustFi Network Ltd",
		"MultiChain":        "TrustFi",
		"Multi Chain":       "TrustFi",
		"2019":              "2021",
	}
	rm2 := map[string]string{
		"Jay Stevens":                      "Herman Chia",
		"Jay2645":                          "moichia",
		"Unreal-Polygonal-Map-Gen":         "MapGenerator",
		"Jay2645/Unreal-Polygonal-Map-Gen": "moichia/MapGenerator",
		"Jay2645/IslandGenerator":          "moichia/MapGenerator",
	}
	tests := []struct {
		name string
		s    string
		rm   map[string]string
		want string
	}{
		{"infinite bug", `                        help="MultiChain path prefix (default: %(default)s)")`, rm1, `                        help="TrustFi path prefix (default: %(default)s)")`},
		{"longer fix", `www.pluralsight.com`, rm1, `www.hello.org`},
		{"bad case 1", `[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com) [![Maintenance](https://img.shields.io/badge/Maintained%3F-no-red.svg)](https://github.com/Jay2645/Unreal-Polygonal-Map-Gen/graphs/commit-activity)`, rm2, `[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com) [![Maintenance](https://img.shields.io/badge/Maintained%3F-no-red.svg)](https://github.com/moichia/MapGenerator/graphs/commit-activity)`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceAllLike(tt.s, tt.rm); got != tt.want {
				t.Errorf("ReplaceAllLike() ori = %q, got = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}
