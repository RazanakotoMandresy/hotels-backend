package service

import (
	"testing"
)

func TestDate(t *testing.T) {
	hotelsDate := []string{
		"2024-10-28->2024-10-30",
		"2024-10-31->2024-11-02",
		"2024-11-03->2024-11-05",
	}
	r := ReserveParams{Starting_date: "2024-11-29", Ending_date: "2024-11-31"}
	t.Run("validDate", func(t *testing.T) {
		err := validDate(r, hotelsDate)
		if err != nil {
			t.Errorf("errors %v ", err)
		}
	})
}
