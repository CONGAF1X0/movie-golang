package uid

import (
	"errors"
	"net"
	"sync"
	"time"
)

/**
雪花算法
	 40bit timestamp | 4bit typeID | 8bit machineID | 11bit sequenceBits
*/

const (
	timestampBits = 40
	machineIDBits = 8
	sequenceBits  = 11
	maxSequence   = 1<<sequenceBits - 1
	timeLeft      = 23
)

type Snowflake struct {
	mutex     *sync.Mutex
	StartTime int64
	LastStamp int64
	MachineID uint8
	Sequence  int64
}

var Sf *Snowflake

func init() {
	Sf = NewSnowflake()
}

func NewSnowflake() *Snowflake {
	st := time.Date(2019, 2, 21, 0, 0, 0, 0, time.Local).UnixNano() / 1e6
	mID, err := lower8BitPrivateIP()
	if err != nil {
		panic(err)
	}
	return &Snowflake{
		StartTime: st,
		MachineID: mID,
		LastStamp: 0,
		Sequence:  0,
		mutex:     new(sync.Mutex),
	}
}

func (sf *Snowflake) getMilliSeconds() int64 {
	return time.Now().UnixNano() / 1e6
}

// NextID userID:0, movieID:1
func (sf *Snowflake) NextID(typeID int) (uint64, error) {
	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	return sf.nextID(typeID)
}

func (sf *Snowflake) nextID(typeID int) (uint64, error) {
	if typeID > 15 {
		return 0,errors.New("typeID > 15")
	}
	timeStamp := sf.getMilliSeconds()
	if timeStamp < sf.LastStamp {
		return 0, errors.New("time is moving backwards")
	}
	if sf.LastStamp == timeStamp {
		sf.Sequence = (sf.Sequence + 1) & maxSequence
		if sf.Sequence == 0 {
			for timeStamp <= sf.LastStamp {
				timeStamp = sf.getMilliSeconds()
			}
		}
	} else {
		sf.Sequence = 0
	}
	sf.LastStamp = timeStamp

	id := ((timeStamp - sf.StartTime) << timeLeft) |
		int64(typeID) << (sequenceBits + machineIDBits)	|
		int64(sf.MachineID)<<sequenceBits |
		sf.Sequence

	return uint64(id), nil
}

func lower8BitPrivateIP() (uint8, error) {
	ip, err := privateIPv4()
	if err != nil {
		return 0, err
	}

	return ip[3], nil
}
func privateIPv4() (net.IP, error) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, a := range as {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}

		ip := ipnet.IP.To4()
		if isPrivateIPv4(ip) {
			return ip, nil
		}
	}
	return nil, errors.New("no private ip address")
}

func isPrivateIPv4(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}
