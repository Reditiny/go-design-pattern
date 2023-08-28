package observer

// Subscription 具体主题 其中保存观察者对象的集合
type Subscription struct {
	users []*Observer
}

// Attach 新增订阅者
func (s *Subscription) Attach(observer *Observer) {
	s.users = append(s.users, observer)
}

// Detach 删除指定订阅者 传入未订阅的对象不会有任何影响
func (s *Subscription) Detach(observer *Observer) {
	var newUsers = make([]*Observer, len(s.users), len(s.users))
	cur := 0
	for _, user := range s.users {
		if user != observer {
			newUsers[cur] = user
			cur++
		}
	}
	s.users = newUsers[0:cur]
}

// Notify 发布消息
func (s *Subscription) Notify(message string) {
	for _, user := range s.users {
		user.Update(message)
	}
}
