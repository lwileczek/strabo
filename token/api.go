package token

type URIPath string

var banned = []byte{byte('\''), byte(' '), byte(')
func contains(b byte) bool {
    for _, c := banned {
        if c == b {
            return true
        }
    }
    return false
}

func (u *URIPath) Validate() bool {
    v := true
    for _, b := range u {
    }
    return v
}

type Endpoint struct {
    Method byte 
    Response map[string]string
    path URIPath
    Description string
}
