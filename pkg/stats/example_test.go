package stats

import (
	"fmt"

	"github.com/Qurbonov-AA/bank/v2/pkg/types"
)

func ExampleAvg_positive() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10,
			Category: "Food",
		},
		{
			ID:       2,
			Amount:   20,
			Category: "Food",
		},
		{
			ID:       3,
			Amount:   30,
			Category: "Carshering",
		},
		{
			ID:       4,
			Amount:   40,
			Category: "Credit",
		},
	}

	fmt.Println(Avg(payments))

	// Output:
	// 25
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10,
			Category: "Food",
		},
		{
			ID:       2,
			Amount:   20,
			Category: "Food",
		},
		{
			ID:       3,
			Amount:   30,
			Category: "Carshering",
		},
		{
			ID:       4,
			Amount:   40,
			Category: "Credit",
		},
	}

	fmt.Println(TotalInCategory(payments, "Food"))

	//Output:
	// 30
}

func ExampleTotalInCategory_empty() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10,
			Category: "Food",
			Status:   types.StatusOK,
		},
		{
			ID:       2,
			Amount:   20,
			Category: "Food",
			Status:   types.StatusOK,
		},
		{
			ID:       3,
			Amount:   30,
			Category: "Carshering",
			Status:   types.StatusFail,
		},
		{
			ID:       4,
			Amount:   40,
			Category: "Credit",
			Status:   types.StatusOK,
		},
	}
	fmt.Println(TotalInCategory_empty(payments))

	//Output:
	// 70

}
