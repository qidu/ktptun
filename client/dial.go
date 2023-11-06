package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
    "strconv"

	"github.com/pkg/errors"
	ktp "github.com/qidu/ktp-go/v6"
	"github.com/xtaci/kcptun/generic"
	"github.com/xtaci/tcpraw"
)

func dial(config *Config, block ktp.BlockCrypt) (*ktp.UDPSession, error) { 
    var key uint32
    keynum, keyerr := strconv.ParseUint(config.Key, 10, 32)
    if keyerr != nil {
        key = 0 // reset to 0x0
    } else {
        key = uint32(keynum)
    }
    log.Println("using *key*:", key)

	mp, err := generic.ParseMultiPort(config.RemoteAddr)
	if err != nil {
		return nil, err
	}

	var randport uint64
	err = binary.Read(rand.Reader, binary.LittleEndian, &randport)
	if err != nil {
		return nil, err
	}

	remoteAddr := fmt.Sprintf("%v:%v", mp.Host, uint64(mp.MinPort)+randport%uint64(mp.MaxPort-mp.MinPort+1))

	if config.TCP {
		conn, err := tcpraw.Dial("tcp", remoteAddr)
		if err != nil {
			return nil, errors.Wrap(err, "tcpraw.Dial()")
		}
		return ktp.NewConn(remoteAddr, block, config.DataShard, config.ParityShard, key, conn)
	}
	return ktp.DialWithOptions(remoteAddr, block, config.DataShard, config.ParityShard, key)

}
