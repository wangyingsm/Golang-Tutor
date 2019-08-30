package main

import (
	"log"
)

const (
	opSave = iota
	opPay
	opTransit
)

type message struct {
	operation int
	account   string
	qty       uint
	target    string
}

func operate(msg <-chan message, result chan<- bool) {
	for m := range msg {
		switch m.operation {
		case opSave:
			log.Printf("Save %d to account %s\n", m.qty, m.account)
		case opPay:
			log.Printf("Pay %d from account %s\n", m.qty, m.account)
		case opTransit:
			log.Printf("transit %d from account %s to account %s",
				m.qty, m.account, m.target)
		}
		result <- true
	}
}

func main() {
	mchan := make(chan message)
	// mchan := make(chan message, 100)
	defer close(mchan)
	rchan := make(chan bool)
	defer close(rchan)
	// for i := 0; i < 10; i++ {
	// 	go operate(mchan, rchan)
	// }
	go operate(mchan, rchan)
	m := message{
		operation: opSave,
		account:   "tom",
		qty:       10000,
	}
	mchan <- m
	log.Printf("result: %v\n", <-rchan)
	m.operation = opPay
	m.qty = 20000
	mchan <- m
	log.Printf("result: %v\n", <-rchan)
	m.operation = opTransit
	m.qty = 300000
	m.target = "cat"
	mchan <- m
	log.Printf("result: %v\n", <-rchan)
}
