package cache

import "github.com/daidokoro/go-git/plumbing"

const (
	Byte FileSize = 1 << (iota * 10)
	KiByte
	MiByte
	GiByte
)

type FileSize int64

type Object interface {
	Add(o plumbing.EncodedObject)
	Get(k plumbing.Hash) plumbing.EncodedObject
	Clear()
}
