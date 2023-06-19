package discovery

import (
	"context"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc/resolver"
)

const scheme = "etcd"

//grpc client
type Resolver struct {
	scheme      string
	EtcdAddrs   []string //etcd服务器地址
	DialTimeout int      //连接超时(秒)

	closeCh      chan struct{}      //关闭通道
	watchCh      clientv3.WatchChan //监听通道
	cli          *clientv3.Client
	keyPrifix    string             //key前缀
	srvAddrsList []resolver.Address //服务地址

	cc     resolver.ClientConn
	logger *zap.Logger
}

//新建一个解析器
func NewResolver(etcdAddrs []string, timeOut int, logger *zap.Logger) *Resolver {
	return &Resolver{
		scheme:      scheme,
		EtcdAddrs:   etcdAddrs,
		DialTimeout: timeOut,
		logger:      logger,
	}
}

//继承resolver.Builder interface
func (r *Resolver) Scheme() string {
	return r.scheme
}

//继承resolver.Builder interface
//build方法 当调用grpc.Dial时候则会调用
func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r.cc = cc

	r.keyPrifix = BuildPrefix(Server{Name: target.Endpoint()})
	if _, err := r.start(); err != nil {
		return nil, err
	}
	return r, nil
}

//继承resolver.Resolver interface
func (c *Resolver) ResolveNow(o resolver.ResolveNowOptions) {}

//继承resolver.Resolver interface
func (r *Resolver) Close() {
	r.closeCh <- struct{}{}
}

func (r *Resolver) start() (chan<- struct{}, error) {
	var err error
	r.cli, err = clientv3.New(clientv3.Config{
		Endpoints:   r.EtcdAddrs,
		DialTimeout: time.Duration(r.DialTimeout) * time.Second,
	})
	if err != nil {
		return nil, err
	}

	resolver.Register(r)

	r.closeCh = make(chan struct{})
	if err = r.sync(); err != nil {
		return nil, err
	}

	go r.watch()
	return r.closeCh, nil
}

//根据keyPrifix监听 events
func (r *Resolver) watch() {
	ticker := time.NewTicker(time.Minute)
	r.watchCh = r.cli.Watch(context.Background(), r.keyPrifix, clientv3.WithPrefix())

	for {
		select {
		case <-r.closeCh:
			return
		case res, ok := <-r.watchCh:
			if ok {
				r.update(res.Events)
			}
		case <-ticker.C:
			if err := r.sync(); err != nil {
				r.logger.Error("sync faild", zap.Error(err))
			}
		}
	}
}

//更新状态
func (r *Resolver) update(events []*clientv3.Event) {
	for _, ev := range events {
		var info Server
		var err error

		switch ev.Type {
		case mvccpb.PUT:
			info, err = ParseValue(ev.Kv.Value)
			if err != nil {
				continue
			}
			addr := resolver.Address{Addr: info.Addr, Metadata: info.Weight}
			if !Exist(r.srvAddrsList, addr) {
				r.srvAddrsList = append(r.srvAddrsList, addr)
				r.cc.UpdateState(resolver.State{Addresses: r.srvAddrsList})
			}
		case mvccpb.DELETE:
			info, err := ParseValue(ev.Kv.Value)
			if err != nil {
				continue
			}
			addr := resolver.Address{Addr: info.Addr}
			if s, ok := Remove(r.srvAddrsList, addr); ok {
				r.srvAddrsList = s
				r.cc.UpdateState(resolver.State{Addresses: r.srvAddrsList})
			}
		}
	}
}

//同步获取所有地址信息
func (r *Resolver) sync() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.DialTimeout)*time.Second)
	defer cancel()
	res, err := r.cli.Get(ctx, r.keyPrifix, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	r.srvAddrsList = []resolver.Address{}
	for _, v := range res.Kvs {
		info, err := ParseValue(v.Value)
		if err != nil {
			continue
		}
		addr := resolver.Address{Addr: info.Addr, Metadata: info.Weight}
		r.srvAddrsList = append(r.srvAddrsList, addr)
	}
	r.cc.UpdateState(resolver.State{Addresses: r.srvAddrsList})
	return nil
}
