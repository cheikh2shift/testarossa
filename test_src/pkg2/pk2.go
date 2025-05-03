package pkg2

import "github.com/cheikh2shift/testarossa/test_src/pkg1"

func GenTestWalks(count int) []pkg1.Tracker {

	result := []pkg1.Tracker{}

	for i := 0; i <= count; i++ {
		result = append(result, pkg1.Tracker{
			Walks: i + 2,
		})
	}

	return result
}
