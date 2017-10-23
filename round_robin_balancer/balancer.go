package round_robin_balancer

type RoundRobinBalancer struct {
	nodes    []int
	nodeChan chan int
	resChan  chan int
}

func (this *RoundRobinBalancer) Init(count int) {
	this.nodes = make([]int, count)
	for i := range this.nodes {
		this.nodes[i] = 0
	}
	this.nodeChan = make(chan int, 128)
	this.resChan = make(chan int, 1)

}

func (this *RoundRobinBalancer) GiveNode() int {
	this.nodeChan <- 1
	for {
		select {
		case res := <-this.resChan:
			return res
		default:

		}
	}
}

func (this *RoundRobinBalancer) GiveStat() []int {
	return []int{}
}
