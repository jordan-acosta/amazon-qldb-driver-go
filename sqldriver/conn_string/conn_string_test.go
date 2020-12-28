package sqldriver

import "testing"

func TestParseConnString(t *testing.T) {
	testCases := []struct {
		input  string
		region string
		ledger string
		err    error
	}{
		{
			input:  "us-east-1:my-ledger",
			region: "us-east-1",
			ledger: "my-ledger",
			err:    nil,
		},
	}

	for _, tc := range testCases {
		region, ledger, err := parseConnString(tc.input)

		if region != tc.region ||
			ledger != tc.ledger ||
			err != tc.err {
			t.Logf("input: %v\n", tc.input)
			t.Errorf("expected: %v, %v, %v\n", tc.region, tc.ledger, tc.err)
			t.Errorf("got: %v, %v, %v\n", region, ledger, err)
		}
	}
}
