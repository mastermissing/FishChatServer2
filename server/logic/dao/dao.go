package dao

import (
	"github.com/golang/glog"
	// "github.com/oikomi/FishChatServer2/server/logic/conf"
	// "github.com/oikomi/FishChatServer2/server/logic/model"
)

// const

type Dao struct {
	MongoDB       *MongoDB
	KafkaProducer *KafkaProducer
}

func NewDao() (dao *Dao, err error) {
	m, err := NewMongoDB()
	if err != nil {
		glog.Error(err)
		return
	}
	KafkaProducer := NewKafkaProducer()
	dao = &Dao{
		MongoDB:       m,
		KafkaProducer: KafkaProducer,
	}
	go dao.KafkaProducer.HandleError()
	go dao.KafkaProducer.HandleSuccess()
	go dao.KafkaProducer.Process()
	return
}
