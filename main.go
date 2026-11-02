package main

import (
	"context"
	"fmt"
)

func main() {
	email, _ := NewEmail("koso@koso.com")
	amount, _ := NewMoneyFromCents(45000)

	fmt.Println(email,amount)

	ev:= NewPaymentDueEvent("evt_001",email,amount)

	fmt.Println(ev.Type.String(),ev.ID,ev.Amount)


	// probando send
	send :=&EmailSenderFake{}
	svc:= NewSerice(send)

		body:="Pago pendiente: " + ev.Amount.String()

		

	_= send.Send(context.Background(),ev.Email,body)
	_= send.Send(context.Background(),ev.Email,body)
	_= send.Send(context.Background(),ev.Email,body)
	_= send.Send(context.Background(),ev.Email,body)

	_= svc.NotifyPaymentDue(context.Background(),ev)

	base:= &EmailSenderFake{}

	sender:=LoggedSender{Sender: base,Log: ConsoleLogger{} }

	service:= NewSerice(sender)
	_= service.NotifyPaymentDue(context.Background(),ev)	
	// nil trap

	senderNil := returnsTypeNil()
	fmt.Println("Sender == nil",senderNil == nil)

}



func returnsTypeNil()Sender {
	var s *EmailSenderFake   = nil
	return s
}