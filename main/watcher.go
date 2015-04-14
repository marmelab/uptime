package main

import "net/http" // package pour les requêtes http
import "fmt"	// package pour écrire
import "os"		// pacckage fonction système
import "strings"// package chaîne de caractères

type Request struct{
	url string
	method string
	headers []string
}
/* Fonction principale qui récupère le domaine à Ping à partir 
des variables d'environnment fournis par Docker-compose du container nginx
puis Ping la cible 
*/
func main() {
	domainName := os.Getenv("NGINX_PORT_80_TCP")
	if(domainName != ""){
		domainName2 := strings.Split(domainName,"p")
		if(domainName2[1] != ""){
			resp,err:=sendPing("http"+domainName2[1])
			if(err==nil){
				fmt.Println("Status :" + resp.Status)
				fmt.Println("Protocole : " + resp.Proto)
			} else {
				fmt.Println(err)
			}
		}	
	}
}



/* Fonction sendPing qui prend en paramètre le nom de domaine
de la cible puis lui envois une requête http de typ GET
*/
func sendPing(domainName string) (resp *http.Response, err error){
	resp, err = http.Get(domainName)
	fmt.Println("Get sur " + domainName)
	return resp,err
}

func ErrorAnalyse() { // selon le code status de la requete ?
	
}