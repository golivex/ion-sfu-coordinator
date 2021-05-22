package coordinator

import (
	"fmt"
	"strings"
)

type SfuScaler struct {
	origin  map[string][]HostSession // this will map the original session to all hosts and their names
	hostmap map[string]Host          //this will map all scaled session and their assigned hosts
}

var sfu = SfuScaler{
	origin:  make(map[string][]HostSession, 0),
	hostmap: make(map[string]Host, 0),
}

func (sfu SfuScaler) GetSessionHost(session string) Host {
	return sfu.hostmap[session]
}

func (sfu SfuScaler) IsScaledSession(session string) bool {
	return strings.Index(session, "-scale-") != -1
}

func (sfu SfuScaler) AssignHostToSession(session string, nhost Host) string {
	_, ok := sfu.origin[session]
	scaledsessionname := ""
	if !ok {
		sfu.origin[session] = []HostSession{}
		scaledsessionname = fmt.Sprintf("%v-scale-%v", session, 0)
	} else {
		scaledsessionname = fmt.Sprintf("%v-scale-%v", session, len(sfu.origin[session]))
	}
	sfu.origin[session] = append(sfu.origin[session], HostSession{
		ip:   nhost.Ip,
		name: scaledsessionname,
	})
	sfu.hostmap[scaledsessionname] = nhost
	return scaledsessionname
}
