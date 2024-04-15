package common

type GetResult struct {
	Data string `json:"data"`
}

// func SelectDataByID(date string) (GetResult, Result) {
// 	// Define your SQL query for selection
// 	//query := "SELECT column1 FROM YourTableName WHERE id = ?"

// 	// Prepare the SQL statement
// 	s, err := ConnDB().Prepare(GetCurrency)
// 	if err != nil {
// 		return GetResult{}, Result{
// 			Success: false,
// 			Error:   err,
// 		}
// 	}
// 	defer s.Close()

// 	var result string
// 	err = s.QueryRowContext(context.Background(), date).Scan(&result)
// 	if err != nil {
// 		return GetResult{}, Result{
// 			Success: false,
// 			Error:   err,
// 		}
// 	}
// 	return GetResult{
// 			result,
// 		},
// 		Result{
// 			Success: true,
// 			Error:   nil,
// 		}
// }
