package kafka

import (
	"github.com/xdg-go/scram"
)

type XDGSCRAMClient struct {
    *scram.Client  
    *scram.ClientConversation
    HashGeneratorFcn scram.HashGeneratorFcn
}

var SHA512 scram.HashGeneratorFcn = scram.SHA512

func (x *XDGSCRAMClient) Begin(userName, password, authzID string) error {
    var err error
    x.Client, err = x.HashGeneratorFcn.NewClient(userName, password, authzID)
    if err != nil {
        return err
    }
    x.ClientConversation = x.Client.NewConversation()
    return nil
}

func (x *XDGSCRAMClient) Step(challenge string) (string, error) {
    return x.ClientConversation.Step(challenge)
}

func (x *XDGSCRAMClient) Done() bool {
    return x.ClientConversation.Done()
}