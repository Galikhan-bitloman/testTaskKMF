package common

// func ConnDB() *sql.DB {

// 	c := ReadConfig()
// 	connInfo := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
// 		c.ConnectionString.Server, c.ConnectionString.Port, c.ConnectionString.User, c.ConnectionString.Password, c.ConnectionString.Database)

// 	db, err := sql.Open("sqlserver", connInfo)

// 	defer db.Close()

// 	if err != nil {
// 		log.Fatal("Error opening database connection:", err)

// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal("Error pinging database:", err)

// 	}

// 	fmt.Println("Connected")

// 	return db
// }
