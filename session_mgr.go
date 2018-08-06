package main

import (
	"sync"
	"sync/atomic"
)

type CoreSessionManager struct {
	sesByID sync.Map

	sesIDGen int64 // 记录已经生成的会话id流水号
}

func (self *CoreSessionManager) Add(sess Session) {
	id := atomic.AddInt64(&self.sesIDGen, 1)

	sess.(interface {
		SetID(int64)
	}).SetID(id)

	self.sesByID.Store(id, sess)
}

func (self *CoreSessionManager) Remove(sess Session) {
	self.sesByID.Delete(sess.ID())
}

func (self *CoreSessionManager) VisitSession(callback func(Session) bool) {
	self.sesByID.Range(func(key, value interface{}) bool {
		return callback(value.(Session))
	})
}
