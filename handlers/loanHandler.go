package handlers

import (
	"net/http"

	"encoding/json"
	"workshop/crud"
	"workshop/types"

	"github.com/gorilla/mux"
)

//LoanHandler will handle the user loans
func LoanHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]
	loanID := params["loanId"]

	var loan types.Loan
	err := json.NewDecoder(r.Body).Decode(&loan)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	item, err := crud.GetOne("users", ID, w)

	user := types.User{}
	json.Unmarshal(item, &user)

	if err != nil {
		return
	}

	if user.Loans == nil {
		user.Loans = make(map[string]types.Loan)
	}

	user.Loans[loanID] = loan

	resp, err := crud.Update("users", ID, user, w)

	if err != nil {
		return
	}

	json.NewEncoder(w).Encode(resp)

}

//AmortizationHandler handles the logic for calculating the user's Amortization
func AmortizationHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]
	loanID := params["loanId"]

	item, err := crud.GetOne("users", ID, w)
	if err != nil {
		return
	}

	user := types.User{}
	json.Unmarshal(item, &user)

	if loan, ok := user.Loans[loanID]; ok {
		aSched := loan.CalculateASchedule()
		ret, _ := json.MarshalIndent(aSched, "", "   ")
		w.Write(ret)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("The loan does not exist"))
}
