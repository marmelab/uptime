package poller

import "net/http" // package pour les requêtes http
import "fmt"	// package pour écrire
import "os"		// pacckage fonction système
import "strings"// package chaîne de caractères
import "net/url"
import "time"
import "basePoller.go"


type baseHttpPoller{
	basePoller
}