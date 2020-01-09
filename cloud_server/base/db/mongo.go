package db

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"hackathon/conf"
	"time"
)

// MONGO
// // Mode constants
// const (
// 	_ Mode = iota
// 	// PrimaryMode indicates that only a primary is
// 	// considered for reading. This is the default
// 	// mode.
// 	PrimaryMode
// 	// PrimaryPreferredMode indicates that if a primary
// 	// is available, use it; otherwise, eligible
// 	// secondaries will be considered.
// 	PrimaryPreferredMode
// 	// SecondaryMode indicates that only secondaries
// 	// should be considered.
// 	SecondaryMode
// 	// SecondaryPreferredMode indicates that only secondaries
// 	// should be considered when one is available. If none
// 	// are available, then a primary will be considered.
// 	SecondaryPreferredMode
// 	// NearestMode indicates that all primaries and secondaries
// 	// will be considered.
// 	NearestMode
// )

var (
	client         *mongo.Client
	database       *mongo.Database
	ErrNoDocuments = mongo.ErrNoDocuments
)

// MongoConnect connects to mongodb
func MongoConnect() {
	logrus.WithFields(logrus.Fields{"module": "mongo_connect"}).Info("Initializing")
	url := conf.Configure.MongodbURL
	mode := conf.Configure.MongodbMode

	want, err := readpref.New(readpref.Mode(mode))
	if err != nil {
		panic(err)
	}
	wc := writeconcern.New(writeconcern.WMajority())
	opt := options.Client().ApplyURI(url)
	//opt.SetLocalThreshold(3 * time.Second)     //只使用与mongo操作耗时小于3秒的
	// opt.SetMaxConnIdleTime(5 * time.Second)    //指定连接可以保持空闲的最大毫秒数
	//opt.SetMaxPoolSize(200)                    //使用最大的连接数
	opt.SetReadPreference(want)
	opt.SetReadConcern(readconcern.Majority()) //指定查询应返回实例的最新数据确认为，已写入副本集中的大多数成员
	opt.SetWriteConcern(wc)                    //请求确认写操作传播到大多数mongod实例
	client, err = mongo.NewClient(opt)
	if err != nil {
		panic(err)
	}

	for {
		if func() bool {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			if err = client.Connect(ctx); err != nil {
				logrus.Errorf("dbmodel connect failed:%s", err.Error())
				<-ctx.Done()
				return false
			}
			return true
		}() {
			break
		}
	}
	for {
		if func() bool {
			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
			if err = client.Ping(ctx, readpref.Primary()); err != nil {
				logrus.Errorf("dbmodel ping failed:%s", err.Error())
				<-ctx.Done()
				return false
			}
			return true
		}() {
			break
		}
	}

	database = client.Database(opt.Auth.AuthSource)
	logrus.WithFields(logrus.Fields{"module": "mongo_connect"}).Info("Connected to ", url, " mode: ", mode)
}

func MongoDB() *mongo.Database {
	return database
}
