package common

type Result struct {
	Success bool
	Error   error
}

// func SaveData(data *schema.RateItem, date string) Result {

// 	_, err := ConnDB().Exec(InsertCurrency, data.Fullname, data.Title, data.Description, date)
// 	if err != nil {
// 		log.Fatal(err)
// 		return Result{
// 			Success: false,
// 			Error:   err,
// 		}
// 	}

// 	return Result{
// 		Success: true,
// 		Error:   nil,
// 	}

// }
