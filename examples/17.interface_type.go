package main

import (
	"errors"
	"fmt"
	"sync"
)

type Payer interface {
	Pay(uint) error
	Transit(Payer, uint) error
	Save(uint) error
}

const SAVE_LIMIT = 2000000

var (
	NotEnoughError = errors.New("deposit not enough")
	QuotationError = errors.New("exceed saving quotation")
)

type BankAccount struct {
	AccountID string
	Deposit   uint
	sync.Mutex
}

type PaypalAccount struct {
	ID     string
	Amount uint
	Credit uint
	sync.Mutex
}

func (b *BankAccount) Pay(payment uint) error {
	if payment > b.Deposit {
		return NotEnoughError
	}
	b.Lock()
	defer b.Unlock()
	b.Deposit -= payment
	return nil
}

func (b *BankAccount) Transit(account Payer, trans uint) error {
	if trans > b.Deposit {
		return NotEnoughError
	}
	b.Lock()
	defer b.Unlock()
	b.Deposit -= trans
	if err := account.Save(trans); err != nil {
		b.Deposit += trans
		return err
	}
	return nil
}

func (b *BankAccount) Save(qty uint) error {
	if qty > SAVE_LIMIT {
		return QuotationError
	}
	b.Lock()
	defer b.Unlock()
	b.Deposit += qty
	return nil
}

func (p *PaypalAccount) Pay(payment uint) error {
	if payment > p.Credit && payment > p.Amount {
		return NotEnoughError
	}
	p.Lock()
	defer p.Unlock()
	if payment <= p.Credit {
		p.Credit -= payment
		return nil
	}
	p.Amount -= payment
	return nil
}

func (p *PaypalAccount) Transit(account Payer, trans uint) error {
	if trans > p.Amount {
		return NotEnoughError
	}
	p.Lock()
	defer p.Unlock()
	p.Amount -= trans
	if err := account.Save(trans); err != nil {
		p.Amount += trans
		return err
	}
	return nil
}

func (p *PaypalAccount) Save(qty uint) error {
	if qty > SAVE_LIMIT {
		return QuotationError
	}
	p.Lock()
	defer p.Unlock()
	p.Amount += qty
	return nil
}

func PayBill(p Payer, bill uint) {
	if err := p.Pay(bill); err != nil {
		fmt.Println(err)
	}
}

func main() {
	tom := &PaypalAccount{
		"tom",
		20000000,
		40000000,
		sync.Mutex{},
	}
	cat := &BankAccount{
		"cat",
		1000000,
		sync.Mutex{},
	}
	fmt.Println(tom, cat)
	PayBill(tom, 10000)
	PayBill(cat, 30000)
	fmt.Println(tom, cat)
	cat.Transit(tom, 100000)
	fmt.Println(tom, cat)
	fmt.Println(tom.Transit(cat, 20110000))
	fmt.Println(tom, cat)
	fmt.Println(cat.Save(10000000))
	fmt.Println(tom.Transit(cat, 10000000))
	cat.Save(2000000)
	fmt.Println(tom, cat)
}
