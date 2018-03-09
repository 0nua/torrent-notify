package db

import (
	"time"
	"database/gobdecorator"
	"config"
)

type Buffer struct {
	Data     map[int]map[int]int
	SaveTime time.Time
}

var buffer = newBuffer()

func newBuffer() Buffer {
	buffer := new(Buffer)
	buffer.Data = gobdecorator.Get()
	buffer.SaveTime = time.Now()
	return *buffer
}

func flushBuffer(buffer *Buffer) bool {
	result := gobdecorator.Save(buffer.Data)
	buffer.SaveTime = time.Now()
	return result;
}

func GetData(userId int) (map[int]int, bool) {
	data, isset := buffer.Data[userId]
	return data, isset
}

func SetData(userId int, data map[int]int) bool {
	buffer.Data[userId] = data;
	processFlushing()
	return true;
}

func processFlushing() {
	diff := time.Now().Sub(buffer.SaveTime)
	if diff.Seconds() > config.GetSaveTimeout() {
		flushBuffer(&buffer)
	}
}