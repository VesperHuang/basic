package nonrepeatingsubstr

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		//Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		//Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbb", 1},
		{"abcabcabcd", 4},

		//chinese
		{"中文測試", 4},
		{"一二三二一", 3},
		{"黑化肥發灰，灰化肥發黑 黑化肥發灰會揮發，灰化肥揮發會發黑。", 8},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmakSubstr(b *testing.B) {
	s := "黑化肥發灰，灰化肥發黑 黑化肥發灰會揮發，灰化肥揮發會發黑。"
	ans := 8

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expected %d", actual, s, ans)
		}
	}
}
