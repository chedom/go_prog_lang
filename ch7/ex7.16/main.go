package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chedom/go_prog_lang/ch7/eval"
)

func main() {
	http.HandleFunc("/", calculate)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	return expr, nil
}

func calculate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// ordinary calculator doesnt have any vars
	fmt.Fprintln(w, expr.Eval(eval.Env{}))
}
