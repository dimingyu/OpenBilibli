// Code generated by $GOPATH/src/go-common/app/tool/cache/gen. DO NOT EDIT.

/*
  Package dao is a generated cache proxy package.
  It is generated from:
  type _cache interface {
		// cache: -nullcache=&model.OpenBindInfo{Mid:-1} -check_null_code=$!=nil&&$.Mid==-1
		BindInfoByMid(c context.Context, mid int64, appID int64) (*model.OpenBindInfo, error)
		// cache: -nullcache=&model.OpenInfo{Mid:-1} -check_null_code=$!=nil&&$.Mid==-1
		OpenInfoByOpenID(c context.Context, openID string, appID int64) (*model.OpenInfo, error)
	}
*/

package dao

import (
	"context"

	"go-common/app/service/main/vip/model"
	"go-common/library/stat/prom"
)

var _ _cache

// BindInfoByMid get data from cache if miss will call source method, then add to cache.
func (d *Dao) BindInfoByMid(c context.Context, id int64, appID int64) (res *model.OpenBindInfo, err error) {
	addCache := true
	res, err = d.CacheBindInfoByMid(c, id, appID)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if res != nil && res.Mid == -1 {
			res = nil
		}
	}()
	if res != nil {
		prom.CacheHit.Incr("BindInfoByMid")
		return
	}
	prom.CacheMiss.Incr("BindInfoByMid")
	res, err = d.RawBindInfoByMid(c, id, appID)
	if err != nil {
		return
	}
	miss := res
	if miss == nil {
		miss = &model.OpenBindInfo{Mid: -1}
	}
	if !addCache {
		return
	}
	d.cache.Do(c, func(c context.Context) {
		d.AddCacheBindInfoByMid(c, id, miss, appID)
	})
	return
}

// OpenInfoByOpenID get data from cache if miss will call source method, then add to cache.
func (d *Dao) OpenInfoByOpenID(c context.Context, id string, appID int64) (res *model.OpenInfo, err error) {
	addCache := true
	res, err = d.CacheOpenInfoByOpenID(c, id, appID)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if res != nil && res.Mid == -1 {
			res = nil
		}
	}()
	if res != nil {
		prom.CacheHit.Incr("OpenInfoByOpenID")
		return
	}
	prom.CacheMiss.Incr("OpenInfoByOpenID")
	res, err = d.RawOpenInfoByOpenID(c, id, appID)
	if err != nil {
		return
	}
	miss := res
	if miss == nil {
		miss = &model.OpenInfo{Mid: -1}
	}
	if !addCache {
		return
	}
	d.cache.Do(c, func(c context.Context) {
		d.AddCacheOpenInfoByOpenID(c, id, miss, appID)
	})
	return
}