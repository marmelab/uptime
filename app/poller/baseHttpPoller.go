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

type baseHttpPoller struct {
	basePoller
}
