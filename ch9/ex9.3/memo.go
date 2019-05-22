// Package memo provides a concurrency-unsafe
// memoization of a function of type Func.
package memo1

import "errors"

// Func is the type of the function to memoize
type Func func(key string, done <-chan struct{}) (interface{}, error)
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

type request struct {
	key      string
	response chan<- result
	done     <-chan struct{}
}

// A Memo caches the result of calling a Func
type Memo struct {
	requests, cancels chan request
}

var CanceledContext = errors.New("context has been canceled")

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request), cancels: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, done <-chan struct{}) (interface{}, error) {
	response := make(chan result)
	req := request{key, response, done}

	memo.requests <- req
	res := <-response

	select {
	case <-done:
		memo.cancels <- req
	default:

	}
	return res.value, res.err
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)

	for {

		for {
		Cancel:
			for {
				select {
				case req := <-memo.cancels:
					delete(cache, req.key)
				default:
					break Cancel
				}
			}
		}

		select {
		case req, ok := <-memo.requests:
			if !ok {
				return
			}
			e := cache[req.key]
			if e == nil {
				e = &entry{ready: make(chan struct{})}
				cache[req.key] = e
				go e.call(f, req.key, req.done) // call f(key)
			}
			go e.deliver(req.response)

		case req := <-memo.cancels:
			delete(cache, req.key)
		}
	}

}

func (e *entry) call(f Func, key string, done <-chan struct{}) {
	e.res.value, e.res.err = f(key, done)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}
