package mm

//Sync Theater with Fesl

import (
	"github.com/NeonRG/Zovix_Fesl/inter/network"
)

var Games = make(map[string]*network.Client)

func FindGIDs() string {
	var gid string
	for ids := range Games {
		gid = ids		
	}
	return gid

}
