package main

import (
	"context"
	"fmt"
)

type EmailSenderFake struct {
	Sent int
}
// service
// type Service struct {
// 	sender *EmailSenderFake //acomplamiento
// }


// service implementando la interfaz
type Service struct {
	sender Sender
}

type Sender interface {
	Send(ctx context.Context,to Email,body string) error
	Channel() string
}

type Logger interface {
	Info(msg string)

}

type ConsoleLogger struct {}

// embedding
type LoggedSender struct {
	Sender
	Log Logger
	}




func (sender *EmailSenderFake) Send(ctx context.Context,to Email,body string) error {
	_=ctx

	sender.Sent++

	fmt.Println("EMAIL SENT:" ,sender.Sent,"TO:",to)
	fmt.Println("BODY",body)
	return nil

}

// func NewSerice(sender *EmailSenderFake) *Service{
// 	return &Service{sender: sender}
// }
func NewSerice(sender Sender) *Service{
	return &Service{sender: sender}
}


func (service *Service) NotifyPaymentDue(ctx context.Context,event Event) error{
	body:= fmt.Sprintf(" [%s] Tienes un pago  pendiente de %s (event=%s)",service.sender.Channel(),event.Amount,event.ID)

	return service.sender.Send(ctx,event.Email,body)

}

func (sender *EmailSenderFake) Channel () string{
	return "email_fake"
}

// metodos para log
func (ConsoleLogger) Info (msg string){
	fmt.Println("Info",msg)
}

func (logSender LoggedSender) Send(ctx context.Context,to Email,body string) error{

	logSender.Log.Info("Enviado via: " +logSender.Channel())
	return logSender.Sender.Send(ctx,to,body)
}

