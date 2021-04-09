package dtc
import (
    "sync"
/*
    "errors"
    log "github.com/sirupsen/logrus"
    "github.com/golang/protobuf/proto"
    "encoding/json"
    "google.golang.org/protobuf/encoding/protojson"

    "github.com/RileyR387/sc-dtc-client/dtcproto"
)
*/
)

type Account struct {

}

type AccountStore struct {
    acctIds []string
    acctUpdateMutex sync.Mutex
}

