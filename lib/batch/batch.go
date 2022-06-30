package batch

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	res = make([]user, 0, n)
	group := new(errgroup.Group)
	var mutex sync.Mutex

	var i int64
	for i = 0; i < pool; i++ {
		group.Go(func() error {
			for n > 0 {
				mutex.Lock()
				n--
				mutex.Unlock()
				u := getOne(n)
				mutex.Lock()
				res = append(res, u)
				mutex.Unlock()
			}
			return nil
		})
	}

	if err := group.Wait(); err != nil {
		fmt.Println("Get errors: ", err)
	}

	return
}
