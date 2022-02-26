package item_test

import (
	"testing"

	"github.com/dannegm/spendzer-service/domain/item"
)

func TestItem_NewItem(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty Name validation",
			name:        "",
			expectedErr: item.ErrInvalidItem,
		}, {
			test:        "Valid Name",
			name:        "Banana",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := item.NewItem(tc.name)

			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
