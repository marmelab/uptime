package poller

import "net/http" // package pour les requêtes http
import "fmt"	// package pour écrire
import "os"		// pacckage fonction système
import "strings"// package chaîne de caractères
import "net/url"
import "time"

type httpPoller struct{
	target,timeout,isDebugEnabled,timer
}

func (basePoller *basePoller) Init(t string, ti Time) {
	basePoller.target = t
	basePoller.timeout = ti || 5000
	isDebugEnabled = false
}

func (basePoller *basePoller) SetDebug(b boolean) {
	basePoller.isDebugEnabled = b
}

func (basePoller *basePoller) Debug(s string) {
	fmt.Println(s)
}

func (basePoller *basePoller) Poll(d Duration) {
	basePoller.timer = new Timer(d)
}

func (basePoller *basePoller) GetTimer() (t Timer){
	t :=basePoller.timer
	return t
}

func (basePoller *basePoller) ValidateTarget(target string) (b boolean){
	return false
}