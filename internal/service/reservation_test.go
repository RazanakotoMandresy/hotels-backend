package service

import (
	"testing"
)

func TestDate(t *testing.T) {
	r := ReserveParams{Starting_date: "2024-10-27", Ending_date: "2024-11-26"}
	t.Run("validDate", func(t *testing.T) {
		err := validDate(r)
		if err != nil {
			t.Errorf("errors %v ", err)
		}
	})
}
