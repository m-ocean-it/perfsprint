package analyzer

import "testing"

func TestReplaceWith(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		s        string
		old      string
		news     []string
		expected string
	}{
		{
			name:     "two %%s",
			s:        "%s, %s",
			old:      "%s",
			news:     []string{`"a"`, `"b"`},
			expected: `"a" + ", " + "b"`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := ReplaceWith(tc.s, tc.old, tc.news)

			if got != tc.expected {
				t.Fatalf("\ngot\n\t%s\ninstead of\n\t%s", got, tc.expected)
			}
		})
	}

}
