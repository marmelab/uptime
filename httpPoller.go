package poller

import "net/http" // package pour les requêtes http
import "fmt"	// package pour écrire
import "os"		// pacckage fonction système
import "strings"// package chaîne de caractères
import "net/url"



func (p *httpPoller) validateTarget(target string) boolean{
	targetUrl:=url.Parse(target)["protocole"]
	return (targetUrl=="http")
}