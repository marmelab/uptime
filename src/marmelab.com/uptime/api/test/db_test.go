package test

import (
	"../Database.go"
	"reflect"
	"testing"
)

func TestGetDbShouldNotTriggerError(t *testing.T) {
	
}

func TestAddValidTargetShouldNotTriggerError(t *testing.T) {
	/*
	se connecter à base, 
	créer une target
	appeler addTarget
	faire un select 
	comparer le resultat avec la target ajoutée
	*/
}

func TestAddInvalidTargetShouldTriggerError(t *testing.T) {
	/*
	créer target invalid
	appeler add
	récupérer le resultat
	*/
}

func TestGetValidTargetShouldNotTriggerError(t *testing.T) {
		/*
	 se connecter a la base 
	 créer une fausse target
	 insert dans la table de test
	 appeler get
	 comparer le résultat avec les données de l'insert
	*/
}

func TestGetInvalidTargetShouldTriggerError(t *testing.T) {
			/*
	 se connecter a la base 
	 créer une fausse target
	 insert dans la table de test
	 appeler get avec un autre id
	 comparer le résultat avec les données de l'insert
	*/
}

func TestGetTargetsShouldNotTriggerError(t *testing.T) {
	/*
	se connecter a la base, 
	ajouter deux fausses targets, 
	faire le get,
	comparéer le resultat avec les targets ajoutées
	*/
}

func TestUpdateValideTargetShouldNotTriggerError(t *testing.T) {
	/*
	se connecter a la base,
	insert d'une target
	créer nouvelle target
	appeler update avec cete target et une nouvelle destination 
	faire un select et comparer le résultat avec la target créer
	*/	
}

func TestUpdateInvalideTargetShouldTriggerError(t *testing.T) {
	/*
	se connecter a la base,
	insert d'une target
	créer nouvelle target
	appeler update avec une target invalide et une nouvelle destination 
	*/	
}

func TestDeleteValidTargetShouldNotTriggerError(t *testing.T) {
	/*
	se connecter a la base
	insert d'une target
	appel méthode delete  de cette target
	faire un select de la target
	comparer 
	*/
}

func TestDeleteInvalidTargetShouldNotTriggerError(t *testing.T) {
		/*
	se connecter a la base
	insert d'une target
	appel méthode delete  d'une target invalide'
	faire un select de la target
	comparer 
	*/
}
