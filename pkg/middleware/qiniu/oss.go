package qiniu

import (
	"bytes"
	"context"
	"sync"
	"time"

	"github.com/grin-ch/grin-utils/log"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var (
	policy *qiniuPolicy
)

func InitOSS(bucket, accessKey, secretKey string, expires int) {
	policy = &qiniuPolicy{
		mac: qbox.NewMac(accessKey, secretKey),
		policy: storage.PutPolicy{
			Scope:   bucket,
			Expires: uint64(expires),
		},
		cfg: &storage.Config{
			UseHTTPS:      false,
			UseCdnDomains: false,
		},
	}
}

type qiniuPolicy struct {
	mac    *qbox.Mac
	policy storage.PutPolicy
	cfg    *storage.Config
	mutex  sync.Mutex

	upToken string
	expires int64
}

func (p *qiniuPolicy) RefreshUpToken() {
	if int64(p.expires) <= time.Now().Unix() {
		p.resetUpToken()
	}
}
func (p *qiniuPolicy) resetUpToken() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.upToken = p.policy.UploadToken(p.mac)
	p.expires = time.Now().Unix() + int64(p.policy.Expires)
}

func (q *qiniuPolicy) UploadBytes(key string, data []byte) (*storage.PutRet, error) {
	q.RefreshUpToken()
	form := storage.NewFormUploader(q.cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	dataLen := int64(len(data))
	err := form.Put(context.Background(), &ret, q.upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	return &ret, err
}

func UploadBytes(key string, data []byte) {
	if ret, err := policy.UploadBytes(key, data); err != nil {
		log.Logger.Errorf("upload bytes by oss err:%v", err)
	} else {
		log.Logger.Warnf("upload bytes, hash:%s, key:%s", ret.Hash, ret.Key)
	}
}
