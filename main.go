package main

import (
	"context"
	"errors"
	"fmt"
)

func main() {
	email, err := NewEmail("koso@koso.com")

	if err != nil {
			if errors.Is(err,ErrInvalidEmail){
				fmt.Println("email invalido corrige el formato")
				return
			}

		fmt.Print("Error al crear el email",err)
		return
	}	
	amount, err := NewMoneyFromCents(-45000)

		if err != nil {
		fmt.Print("Error al crear el monto",err)
		return
	}	
	

	fmt.Println(email,amount)

		if err:= WriteReport("report.txt","Reporte de la semana"); err !=nil  {
			fmt.Println("Al crear el reporte. Error: ",err)
		}

		if err:= ShadowingError(); err!=nil{
			fmt.Println("Shadowing error: ",err)
		}else{
			fmt.Println("Shadowing error retorno nil",err)
		}

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


	// MustTemplate("payment_do")
		if err:= RunSafely(func() {MustTemplate("Payment_do")}); err!= nil {
			fmt.Println("Safe error: ",err)
		}
}



func returnsTypeNil()Sender {
	var s *EmailSenderFake   = nil
	return s
}