package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_newLRUCache(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want LRUCache
	}{
		//
		{"case1",
			args{capacity: 2},
			LRUCache{m: make(map[int]*LinkNode), capacity: 2, head: &LinkNode{0, 0, nil, nil}, tail: &LinkNode{0, 0, nil, nil}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newLRUCache(tt.args.capacity); got.capacity != tt.args.capacity {
				t.Errorf("newLRUCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLRUCache_Get(t *testing.T) {
	lruCache := newLRUCache(2)
	lruCache.Put(1, 1)
	result := lruCache.Get(1)
	assert.Equal(t, 1, result)
	// get not exist key
	result = lruCache.Get(333)
	assert.Equal(t, -1, result)
}
