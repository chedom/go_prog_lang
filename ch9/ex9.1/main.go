package bank1

var deposits = make(chan int) // send amount to deposits
var balances = make(chan int) // receive balane
var withdrawRes = make(chan bool)
var withdraw = make(chan withdrawMes) // send amount to withdraw

type withdrawMes struct {
	ch     chan bool
	amount int
}

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdraw <- withdrawMes{ch: ch, amount: amount}
	return <-ch
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case m := <-withdraw:
			if m.amount > balance {
				m.ch <- false
				continue
			}
			balance -= m.amount
			m.ch <- true
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
