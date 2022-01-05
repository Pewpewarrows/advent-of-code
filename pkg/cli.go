package advent

import (
    "bufio"
    "flag"
    "log"
    "os"
)

// TODO: docs
// TODO: maybe have this curry and return another func type?
type Callback func (scanner *bufio.Scanner, inputDataPtr interface{})

// TODO: docs
func Execute(callback Callback, inputDataPtr interface{}) {
    livePtr := flag.Bool("live", false, "whether to fetch the input live from the internet")
    // TODO: usage instructions for args: single path to input file
    // TODO: also allow input to come via stdin, to use curl pipe
    flag.Parse()
    live := *livePtr
    filePaths := flag.Args()

    if live {
        if len(filePaths) > 0 {
            log.Fatalf("invalid args, may only fetch input live or supply file arg, not both: %s", filePaths)
        }
    } else {
        if len(filePaths) == 0 {
            log.Fatal("invalid args, expected one input file path, found none")
        } else if len(filePaths) > 1 {
            log.Fatalf("invalid args, expected only one input file path, found multiple: %s", filePaths)
        }
    }

    if live {
        // TODO: write me
        return
    }

    f, err := os.Open(filePaths[0])
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    callback(scanner, inputDataPtr)

	if err := scanner.Err(); err != nil {
        log.Fatal("could not read input:", err)
	}
}
