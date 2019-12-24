package main

import (
	"github.com/golang/mock/gomock"
	"test_goconvey_gomock/Storage"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type User struct {
	Storage Storage.Storage
}

func NewUser(storage Storage.Storage) *User {
	return &User{
		Storage: storage,
	}
}

func TestRedisRepo(t *testing.T) {
	//Init mock control
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//Init mock redis storage
	mockRedis := Storage.NewMockStorage(mockCtrl)

	//Test convey
	Convey("Test Redis function", t, func() {
		Convey("Test Add", func() {
			Convey("Add to redis, check before and after add", func() {
				gomock.InOrder(
					mockRedis.EXPECT().Get("foo").Return("", nil),
					mockRedis.EXPECT().Add("foo", "bar").Return(nil),
					mockRedis.EXPECT().Get("foo").Return("bar", nil),
				)

				user := NewUser(mockRedis)
				value, err := user.Storage.Get("foo")
				So(value, ShouldEqual, "")
				So(err, ShouldBeNil)

				err = user.Storage.Add("foo", "bar")
				So(err, ShouldBeNil)

				value, err = user.Storage.Get("foo")
				So(value, ShouldEqual, "bar")
				So(err, ShouldBeNil)
			})
		})
	})
}
