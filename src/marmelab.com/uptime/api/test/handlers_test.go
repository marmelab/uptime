package test

import (
	"../handlers"
	"net/http"
	"reflect"
	"testing"
)

func TestRetrieveTargetsShouldNotTriggerError(t *testing.T) {
	/*
	je monte un serveur
	j'appel la méthode
	je créer une request get et je controle le résultat
	*/
}

func TestShowTargetWithNullIdShouldTriggerError() {
	/*
	je monte un serveur
	j'appel la méthode avec id = 0
	je créer une request get et je controle le résultat
	*/	
}

func TestShowTargetWithNegatifIdShouldTriggerError() {
	/*
	je monte un serveur
	j'appel la méthode avec id = -1
	je créer une request get et je controle le résultat
	*/	
}

func TestShowTargetWithValidIdShouldNotTriggerError() {
	/*
	je monte un serveur
	j'appel la méthode avec id = 1
	je créer une request get et je controle le résultat
	*/
}

func TestShowTargetWithNBigIdShouldTriggerError() {
	/*
	je monte un serveur
	j'appel la méthode avec id = 20000
	je créer une request get et je controle le résultat
	*/	
}

func TestCreateTargetWithValidDataShouldNotTriggerError() {
	/*
	je monte un serveur
	j'appel la méthode avec une bonne target
	je créer une request get et je controle le résultat
	*/
}

func TestCreateTargetWithInvalidDataShouldNotTriggerError() {
	/*
	je monte un serveur
	j'appel la méthode avec nil
	je créer une request get et je controle le résultat
	*/
}

func TestUpdateTargetWithNullIdShouldTriggerError() {
	
}

func TestUpdateTargetWithNegatifIdShouldTriggerError() {
 	
}

func TestUpdateTargetWithValidIdShouldTriggerError() {
	
}

func TestUpdateTargetWithBigIdShouldTriggerError() {
	
}

func TestDeleteTargetWithNullIdShouldTriggerError() {
	
}

func TestDeleteTargetWithNegatifIdShouldTriggerError() {
 	
}

func TestDeleteTargetWithValidIdShouldTriggerError() {
	
}

func TestDeleteTargetWithBigIdShouldTriggerError() {
	
}


