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
		{"misc", "hElLo woRld", "Llo", "l10", "hEl10 woRld"},
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
	rm0 := map[string]string{
		"abc":            "123",
		"github.com":     "gitlab.com",
		"github.com/abc": "bitbucket.org/def",
	}
	rm1 := map[string]string{
		"pluralsight.com":   "hello.org",
		"pluralsight":       "whaohelo",
		"conference":        "meeting",
		"Coin Sciences Ltd": "Hello Network Ltd",
		"MultiChain":        "SingleChain",
		"Multi Chain":       "Single Chain",
		"2019":              "2021",
	}
	rm2 := map[string]string{
		"Jay Stevens":                      "John Doe",
		"Jay2645":                          "John1234",
		"Unreal-Polygonal-Map-Gen":         "MapTabs",
		"Jay2645/Unreal-Polygonal-Map-Gen": "John1234/MapTabs",
		"Jay2645/IslandGenerator":          "John1234/MapTabs",
	}
	rm3 := map[string]string{
		"ABCD": "1234",
		"123":  "def",
		"ef":   "XX",
	}
	tests := []struct {
		name string
		s    string
		rm   map[string]string
		want string
	}{
		{"no change", "https://gitlab.com/ab3C/repo", rm0, "https://gitlab.com/ab3C/repo"},
		{"check order", "https://github.com/abc/repo", rm0, "https://bitbucket.org/def/repo"},
		{"self repeat", "abcd-123-ef", rm3, "1234-def-xx"},
		{"infinite bug", `                        help="MultiChain path prefix (default: %(default)s)")`, rm1, `                        help="SingleChain path prefix (default: %(default)s)")`},
		{"longer fix", `www.pluralsight.com`, rm1, `www.hello.org`},
		{"test 1", `Jay Stevens is Jay2645`, rm2, `John Doe is John1234`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceAllLike(tt.s, tt.rm); got != tt.want {
				t.Errorf("ReplaceAllLike() ori = %q, got = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func BenchmarkReplaceAllLike(b *testing.B) {
	rm := map[string]string{
		"abc":            "123",
		"github.com":     "gitlab.com",
		"github.com/abc": "bitbucket.org/def",
	}
	s := "Hello World"
	for i := 0; i < b.N; i++ {
		_ = ReplaceAllLike(s, rm)
	}
}
