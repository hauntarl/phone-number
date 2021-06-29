package normalize

import "testing"

func TestNumber(t *testing.T) {
	for _, test := range numberTests {
		actual, actualErr := Number(test.input)
		if !test.expectErr {
			if actualErr != nil {
				// if we don't expect an error and there is one
				var _ error = actualErr
				t.Errorf("FAIL: %s\nNumber(%q): expected no error, but error is: %s", test.description, test.input, actualErr)
			}
			if actual != test.number {
				t.Errorf("FAIL: %s\nNumber(%q): expected [%s], actual: [%s]", test.description, test.input, test.number, actual)
			}
		} else if actualErr == nil {
			// if we expect an error and there isn't one
			t.Errorf("FAIL: %s\nNumber(%q): expected an error, but error is nil", test.description, test.input)
		}
	}
}

func BenchmarkNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range numberTests {
			Number(test.input)
		}
	}
}
