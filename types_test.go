package main

import (
	"errors"
	"testing"
)


func TestNewEmailInvalid(test *testing.T){
	_,err:=NewEmail("not-an-email")
	if err == nil {
		test.Fatalf("se esperaba un erorr, pero se obtuvo un nil")
	}


	if !errors.Is(err,ErrInvalidEmail){
		test.Fatalf("Esperabamos errors.Is(err,ErrInvalidEmail) = true , obtuvimos=%v",err)
	}


}

func TestMoneyFromCentsNegative(test *testing.T){
	_,err:=NewMoneyFromCents(-1)

	if err == nil{
		test.Fatalf("Esperabamos un error y obtuvimos  un nil")
	}
	
	if !errors.Is(err,ErrNegativeMoney){
		test.Fatalf("Esperabamos un error errors.Is(err,ErrNegativeMoney)= true  obtuvimos=%v",err)
		
	}

}

func TestNewEmailOk(test *testing.T) {
	email,err:=NewEmail("KosO@correo.com")
	
	if err != nil{
		test.Fatalf("Esperabamos un nil, obtuvimos un: %v",err)
	}
	
	if email != "koso@correo.com"{
		test.Fatalf("No se esperaba el correo: %q",email)
		
	}

}


func TestMoneyFromCents(test *testing.T){
	money,err:= NewMoneyFromCents(45000)

	if err != nil{
		test.Fatalf("Esperabamos un nil, obtuvimos un: %v",err)
	}
	
	if money.String() != "$450.00" {
		test.Fatalf("No se esperaba el formato: %s",money.String())

	}

}
