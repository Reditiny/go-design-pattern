package command

// Order 订单请求
type Order struct {
	TableId int
	FoodDir map[string]int
}
