package common

const (
	InsertCurrency string = `INSERT INTO R_CURRENCY (TITLE, CODE, VALUE, A_DATE)
			VALUES ($1, $2, $3, $4)`

	GetCurrency string = `
				SELECT TITLE, CODE, VALUE, A_DATE FROM R_CURRENCY WHERE A_DATE = ? 
`
)
