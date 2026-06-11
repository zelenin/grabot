package client

import (
	"io/ioutil"
	"log"
	"os"
)

var NullLogger = log.New(ioutil.Discard, "", log.LstdFlags)
var StdLoger = log.New(os.Stderr, "", log.LstdFlags)
