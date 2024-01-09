package fluentsql

import "testing"

// TestDeleteTable
func TestDeleteTable(t *testing.T) {
	testCases := map[string]*DeleteBuilder{
		"DELETE TABLE customers WHERE contact_name = 'Alfred Schmidt' AND city = 'Frankfurt' AND customer_id = 1": DeleteInstance().
			Delete("customers").
			Where("contact_name", Eq, "Alfred Schmidt").
			Where("city", Eq, "Frankfurt").
			Where("customer_id", Eq, 1),
	}

	for expected, query := range testCases {
		if query.String() != expected {
			t.Fatalf(`Query %s != %s`, query.String(), expected)
		}
	}
}
