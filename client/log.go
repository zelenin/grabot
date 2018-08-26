package client

import (
    "os"
    "log"
    "io/ioutil"
)

var NullLogger = log.New(ioutil.Discard, "", log.LstdFlags)
var StdLoger = log.New(os.Stderr, "", log.LstdFlags)
