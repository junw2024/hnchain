package distributedid

import (
	"sync"
	"time"
)

const (
	timestampBits  = 41                         // 时间戳位数
	machineIDBits  = 10                         // 机器ID位数
	sequenceBits   = 12                         // 序列号位
	maxMachineID   = -1 ^ (-1 << machineIDBits) // 最大机器ID
	maxSequenceNum = -1 ^ (-1 << sequenceBits)  // 最大序列号
)

type Snowflake struct {
	timestamp   int64       // 时间戳
	machineID   int64       // 机器ID
	sequenceNum int64       // 序列号
	lock        *sync.Mutex //锁
}

var snowflake *Snowflake
var once sync.Once

func NewSnowflake(machineID int64) *Snowflake {
	if machineID < 0 || machineID > maxMachineID {
		panic("Invalid machine ID")
	}

	once.Do(func() {
		snowflake = &Snowflake{
			timestamp:   time.Now().Unix() / 1e6,
			machineID:   machineID,
			sequenceNum: 0,
			lock:        &sync.Mutex{},
		}
	})

	return snowflake
}

func (s *Snowflake) GenerateId() int64 {
	s.lock.Lock()
	defer s.lock.Unlock()
	currentTimestamp := time.Now().UnixNano() / 1e6

	if currentTimestamp == s.timestamp {
		s.sequenceNum = (s.sequenceNum + 1) & maxSequenceNum
		if s.sequenceNum == 0 {
			currentTimestamp = s.waitNextMillis()
		}
	} else {
		s.sequenceNum = 0
	}

	s.timestamp = currentTimestamp

	id := (currentTimestamp << (machineIDBits + sequenceBits)) | (s.machineID << sequenceBits) | s.sequenceNum

	return id
}

func (s *Snowflake) waitNextMillis() int64 {
	currentTimestamp := time.Now().UnixNano() / 1e6
	for currentTimestamp <= s.timestamp {
		currentTimestamp = time.Now().UnixNano() / 1e6
	}
	return currentTimestamp
}
