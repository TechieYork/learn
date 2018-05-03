package test

import (
	"sync"
	"context"
	"time"
	"runtime/debug"
	"errors"

	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/resolver"
)

const (
	testSchema = "test"
)

var (
	updateDuration = 10 * time.Second
)

//init function to register builder to resolver
func init() {
	resolver.Register(NewBuilder())
}

func NewBuilder() resolver.Builder {
	return &testBuilder{}
}

/*******************************
 * Test Builder
 *******************************/
// Builder creates a resolver that will be used to watch name resolution updates.
type testBuilder struct {

}

// Scheme returns the scheme supported by this resolver.
// Scheme is defined at https://github.com/grpc/grpc/blob/master/doc/naming.md.
func (b *testBuilder) Scheme() string {
	grpclog.Infof("Scheme called! schema:%v", testSchema)
	return testSchema
}

// Build creates a new resolver for the given target.
//
// gRPC dial calls Build synchronously, and fails if the returned error is not nil.
func (b *testBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	grpclog.Infof("Build called!")

	ctx, cancel := context.WithCancel(context.Background())

	r := &testResolver{
		target: target,
		ctx: ctx,
		cancel: cancel,
		cc: cc,
		resolveNowChannel: make(chan bool, 1),
		updateTimer: time.NewTimer(0),
	}

	r.waitGroup.Add(1)

	go r.StartWatch()

	return r, nil
}

/*******************************
 * Test Resolver
 *******************************/
// Resolver watches for the updates on the specified target.
// Updates include address updates and service config updates.
type testResolver struct {
	target resolver.Target

	ctx context.Context
	cancel context.CancelFunc

	cc resolver.ClientConn

	resolveNowChannel chan bool		//This channel is used by ResolveNow() to force an immediate resolution of the target
	waitGroup sync.WaitGroup		//Wait for watch() to be finished
	updateTimer *time.Timer			//Timer to update
}

// ResolveNow will be called by gRPC to try to resolve the target name again.
// It's just a hint, resolver can ignore this if it's not necessary.
func (r *testResolver) ResolveNow(option resolver.ResolveNowOption) {
	grpclog.Infof("ResolveNow called!")

	select {
	case r.resolveNowChannel <- true:
	default:
	}
}

// Close closes the resolver.
func (r *testResolver) Close() {
	grpclog.Infof("Close called!")

	r.cancel()
	r.waitGroup.Wait()
}

func (r *testResolver) StartWatch() {
	grpclog.Infof("StartWatch called!")

	for {
		err := r.watch()

		if err != nil {
			grpclog.Errorf("watch failed, error:%v", err.Error())
		}

		select {
		case <- time.After(time.Second):
		}
	}
}

func (r *testResolver) watch() (err error) {
	grpclog.Infof("Begin watch!")

	defer r.waitGroup.Done()
	defer func() {
		if r:= recover(); r != nil {
			grpclog.Errorf("watch recovered from panic, error:%v, stack:%v", r.(error), string(debug.Stack()))
			err = errors.New("watch recoverd from panic!")
		}
	}()

	for {
		select {
		case <- r.ctx.Done():
			grpclog.Infof("Recv ctx.Done() event")
			return nil
		case <- r.resolveNowChannel:
			grpclog.Infof("Recv resolveNow event")
		case <- r.updateTimer.C:
			grpclog.Infof("Recv timer event")
		}

		r.updateTimer.Reset(updateDuration)

		err := r.handleResolve()

		if err != nil {
			grpclog.Warningf("handle resolve failed! error:%v", err.Error())
		}
	}

	return nil
}

func (r *testResolver) handleResolve() error {
	grpclog.Infof("handleResolve called")

	r.cc.NewAddress([]resolver.Address{{Addr:"localhost:58888"}})

	return nil
}
