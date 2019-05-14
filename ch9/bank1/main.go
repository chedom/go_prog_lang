package bank1

var deposits = make(chan int) // send amount to deposits
var balances = make(chan int) // receive balane

func Deposit(amount int) { deposits <- amount }
func Balance() int { return <- balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <- deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
