package compiler

import (
    "log"
    "os"
)


var file* os.File

type token struct {
    tokentype string
    val string
}


func readchar() byte {
    buf := make([]byte, 1, 1)
    file.Read(buf)
    return buf[0]
}


func readword() string {
    return "not implm"

}

func maketoken(tokentype, value string) token {
    validtokentypes := []string {
        "States",
        "Edges",
    }

    for _, vtokens := range validtokentypes {
        if vtokens == tokentype {
            return token{tokentype, value}
        }
    }
    
    log.Fatal("Did not find expected keyword, found " + tokentype)
    return token{}
}

func scanner() token {
    // TODO clear whitespace
    chr := readchar()
    if chr  == 'S' { // Must be States
        
        return maketoken("S" + readword(), nil)
    }
    return 
}


func Parse(filename string) {
    filelocal, err := os.Open(filename)
    file = filelocal

    if err != nil {
        log.Fatal("Not existing")
    }

    state := 0
    for true {
        token := scanner()
    

        if state == 0 { // We are before States:
        } else if state == 1 {

        } else if state == 2 {

        } else {
            log.Fatal("Something went wrong while parsing")
        }
    }
}
