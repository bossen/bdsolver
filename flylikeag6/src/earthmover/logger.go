package earthmover

import (
    "log"
    "os"
)

var logger *log.Logger

func init() {
    logger = log.New(os.Stdout, "logger: ", log.Lshortfile)
}
