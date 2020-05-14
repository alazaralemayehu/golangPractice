package main

type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

type Queue []*Order

func (o *Order) New(p int, q int, pr string, c string) {
	o.priority = p
	o.quantity = q
	o.product = pr
	o.customerName = c
}

func (queue *Queue) Add(order *Order) {
	appended := false
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {
		var rightHalf Queue
		for i := 0; i < len(*queue); i++ {
			if (*queue)[i].priority > order.priority {
				rightHalf = append(Queue{order}, (*queue)[i:]...)
				*queue = append((*queue)[:i], rightHalf...)
				appended = true
				break
			}
		}
	}
	if (!appended) {
		*queue = append(*queue, order)
	}
}