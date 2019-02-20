package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/chedom/go_prog_lang/ch7/eval"
)

func scanExpr(scanner *bufio.Scanner) (eval.Expr, map[eval.Var]bool) {
	var expr eval.Expr
	vars := make(map[eval.Var]bool)
	for {
		var err error
		fmt.Println("Write expression ->")
		scanner.Scan()
		str := scanner.Text()
		if str == "" {
			fmt.Fprintf(os.Stderr, "error: string is empty, try again\n")
			continue
		}
		expr, err = eval.Parse(str)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v, try again\n", err)
			continue
		}
		if err := expr.Check(vars); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v, try again\n", err)
			continue
		}
		return expr, vars
	}
}

func scanEnv(scanner *bufio.Scanner, vars map[eval.Var]bool) eval.Env {
	env := make(eval.Env, len(vars))
	for k := range vars {
		for {
			fmt.Printf("Write variable: %q", k)
			scanner.Scan()
			s := scanner.Text()
			f, err := strconv.ParseFloat(s, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v, try again\n", err)
				continue
			}
			env[k] = f
			break
		}
	}
	return env
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	expr, vars := scanExpr(scanner)
	env := scanEnv(scanner, vars)
	res := expr.Eval(env)
	fmt.Printf("Result is %f", res)
}
