package gladeapi

import (
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	unmarshalData(w, r)
	var data Data

	emailQuery := fmt.Sprintf(`select * from user_data where email="%v"`, data.Email)
	userQuery := fmt.Sprintf(`select * from user_data where username="%v"`, data.Username)
	passwordQuery := fmt.Sprintf(`select * from user_data where username="%v"`, data.Password)

	queryDb(emailQuery)
	queryDb(userQuery)
	queryDb(passwordQuery)
}
