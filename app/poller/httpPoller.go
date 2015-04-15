package poller

import (
	"basePoller.go"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func (p *httpPoller) validateTarget(target string) boolean {
	targetUrl := url.Parse(target)["protocole"]
	return (targetUrl == "http")
}
