package user

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/naggie/dsnet"
)

type addUserParams struct {
	Owner       string
	Hostname    string
	Description string
}

func addNewUser(c *gin.Context) (*dsnet.PeerConfig, error) {
	var newUser addUserParams
	if err := c.BindJSON(&newUser); err != nil {
		return nil, err
	}

	privateKey := dsnet.GenerateJSONPrivateKey()
	publicKey := privateKey.PublicKey()

	peer := dsnet.PeerConfig{
		Owner:        newUser.Owner,
		Hostname:     newUser.Hostname,
		Description:  newUser.Description,
		Added:        time.Now(),
		PublicKey:    publicKey,
		PrivateKey:   privateKey, // omitted from server config JSON!
		PresharedKey: dsnet.GenerateJSONKey(),
		Networks:     []dsnet.JSONIPNet{},
	}

	if len(conf.Network.IPNet.Mask) > 0 {
		peer.IP = conf.MustAllocateIP()
	}

	if len(conf.Network6.IPNet.Mask) > 0 {
		peer.IP6 = conf.MustAllocateIP6()
	}

	conf.MustAddPeer(peer)
	conf.MustSave()
	dsnet.ConfigureDevice(conf)
	return &peer, nil
}
