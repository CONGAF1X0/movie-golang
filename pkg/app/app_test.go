package app

import (
	"TicketSales/pkg/aliyun"
	"TicketSales/pkg/uid"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

/*
350534107421253632
350534107421253633
350534107421253634
350534107421253635
350534107421253636
350534107421253637
350534107421253638
350534107421253639

350534107421253640
350534107421253641
350534107421253642
*/
func TestGenSnowFlake(t *testing.T) {
	//fmt.Println(time.Date(2019, 4, 21, 0, 0, 0, 0, time.UTC).UTC().UnixNano()/1e6)
	ch := make(chan uint64, 10000)
	count := 10000
	wg.Add(count)
	defer close(ch)
	//并发 count个goroutine 进行 snowFlake ID 生成
	f := uid.NewSnowflake()
	//f := sonyflake.NewSonyflake(sonyflake.Settings{})

	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			rand.Seed(time.Now().UnixNano())
			t := rand.Intn(4)
			id, _ := f.NextID(t)
			ch <- id
		}()
	}
	wg.Wait()
	m := make(map[uint64]int)
	for i := 0; i < count; i++ {
		id := <-ch
		// 如果 map 中存在为 id 的 key, 说明生成的 snowflake ID 有重复
		_, ok := m[id]
		if ok {
			fmt.Printf("repeat id %d\n", id)

		}
		// 将 id 作为 key 存入 map
		m[id] = i
		fmt.Println(id)
	}
	// 成功生成 snowflake ID
	fmt.Println("All", len(m), "snowflake ID Get success!")
}

func Test(t *testing.T) {
	f := uid.NewSnowflake()
	id, _ := f.NextID(1)
	fmt.Printf("%b\n%v\n", id, id)
}

func TestM(t *testing.T) {
	aliyun.SendMobileCaptcha("18665726192")
}
