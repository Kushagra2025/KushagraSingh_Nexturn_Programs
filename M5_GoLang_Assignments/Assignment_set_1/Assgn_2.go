package main

import "fmt"

type Accounts struct {
	ID      int
	Name    string
	Balance float64
}

type transaction_hisory struct {
	account_id int
	msg        string
	amount     float64
}

var accounts_data []Accounts
var history []transaction_hisory

func Deposit(id int, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Invalid Amount Value")
	} else {
		for i := range accounts_data {
			if accounts_data[i].ID == id {
				accounts_data[i].Balance = accounts_data[i].Balance + amount
				history = append(history, transaction_hisory{account_id: accounts_data[i].ID,
					msg: "Deposited Amount", amount: amount})
				return nil
			}
		}

		return fmt.Errorf("Invalid ID!")
	}
}

func Withdraw(id int, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Invalid Amount Value.")
	} else {
		for i := range accounts_data {
			if accounts_data[i].ID == id {
				if accounts_data[i].Balance >= amount {
					accounts_data[i].Balance = accounts_data[i].Balance - amount
					history = append(history, transaction_hisory{account_id: accounts_data[i].ID,
						msg: "Withdrawn Amount", amount: amount})
					return nil
				} else {
					return fmt.Errorf("Insufficient Balance.")
				}
			}
		}

		return fmt.Errorf("Invalid ID.")
	}
}

func ViewBalance(id int) (float64, error) {
	for i := range accounts_data {
		if accounts_data[i].ID == id {
			return accounts_data[i].Balance, nil
		}
	}
	return 0, fmt.Errorf("account not found")
}

func ViewTransactionHistory(id int) ([]transaction_hisory, error) {
	found := false
	var ans []transaction_hisory
	for _, t := range history {
		if t.account_id == id {
			// fmt.Printf("Account ID: %d, Message: %s, Amount: %.2f\n", t.account_id, t.msg, t.amount)
			ans = append(ans, t)
			found = true
		}
	}
	if !found {
		return nil, fmt.Errorf("no transaction history for account ID %d", id)
	}
	return ans, nil
}

// func main() {

// 	// accounts_data = append(accounts_data, Accounts{ID: 1, Name: "Jon", Balance: 123.12})
// 	// accounts_data = append(accounts_data, Accounts{ID: 2, Name: "Jane", Balance: 323.12})
// 	// accounts_data = append(accounts_data, Accounts{ID: 3, Name: "Lee", Balance: 325.12})

// 	// fmt.Println(accounts_data)

// 	// var err_1 = Deposit(1, 203.12)
// 	// if err_1 != nil {
// 	// 	fmt.Println("Error:", err_1)
// 	// } else {
// 	// 	fmt.Println(accounts_data[0])
// 	// }

// 	// var err_2 = Withdraw(2, 20.12)
// 	// if err_2 != nil {
// 	// 	fmt.Println("Error:", err_2)
// 	// } else {
// 	// 	fmt.Println(accounts_data[1])
// 	// }

// 	// var err_3 = Withdraw(2, 20.12)
// 	// if err_3 != nil {
// 	// 	fmt.Println("Error:", err_2)
// 	// } else {
// 	// 	fmt.Println(accounts_data[1])
// 	// }

// 	// var result, err_4 = ViewTransactionHistory(2)
// 	// if err_4 != nil {
// 	// 	fmt.Println("Error:", err_2)
// 	// } else {
// 	// 	fmt.Println(result)
// 	// }

// 	// // fmt.Println(history)

// }

func main_2() {
	// Initializing some sample accounts
	accounts_data = append(accounts_data, Accounts{ID: 1, Name: "Jon", Balance: 123.12})
	accounts_data = append(accounts_data, Accounts{ID: 2, Name: "Jane", Balance: 323.12})
	accounts_data = append(accounts_data, Accounts{ID: 3, Name: "Lee", Balance: 325.12})

	// Menu-driven system
	for {
		// Displaying the menu options
		fmt.Println("===== Menu =====")
		fmt.Println("1. Deposit Money")
		fmt.Println("2. Withdraw Money")
		fmt.Println("3. View Account Balance")
		fmt.Println("4. View Transaction History")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		DepositOption := 1
		WithdrawOption := 2
		ViewBalanceOption := 3
		ViewHistoryOption := 4
		ExitOption := 5

		switch choice {
		case DepositOption:
			var id int
			var amount float64
			fmt.Print("Enter Account ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter Amount to Deposit: ")
			fmt.Scanln(&amount)

			err := Deposit(id, amount)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful!")
			}

		case WithdrawOption:
			var id int
			var amount float64
			fmt.Print("Enter Account ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter Amount to Withdraw: ")
			fmt.Scanln(&amount)

			err := Withdraw(id, amount)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful!")
			}

		case ViewBalanceOption:
			var id int
			fmt.Print("Enter Account ID: ")
			fmt.Scanln(&id)

			balance, err := ViewBalance(id)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Account Balance: %.2f\n", balance)
			}

		case ViewHistoryOption:
			var id int
			fmt.Print("Enter Account ID: ")
			fmt.Scanln(&id)

			result, err := ViewTransactionHistory(id)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(result)
			}

		case ExitOption:
			fmt.Println("Exiting program...")
			return

		default:
			fmt.Println("Invalid option! Please try again.")
		}
	}
}
