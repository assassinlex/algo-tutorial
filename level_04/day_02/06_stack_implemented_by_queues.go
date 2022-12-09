package day_02

/*
	两个队列实现栈
		准备两个队列[不是之前的双端队列, 只能从一个方向出], 一个数据队列[dataQueue], 一个辅助队列[helpQueue]
		这两个队列角色可以互换
		入栈时, 直接将元素入队到数据队列, 出栈时, 将数据队列所有元素依次出队 & 入队到辅助队列, 直至数据队列只有一个元素
		此时数据队列只有最后一次入队的元素, 出队时就是这个元素, 就是栈先进先出的特点
		核心点, 在每一次出栈操作过程中, 在数据队列最后一个元素出队后, 数据队列 & 辅助队列互换地址
*/

func NewStackByQueues() *stackByQueues {
	return &stackByQueues{
		dataQueue: NewQueueByList(),
		helpQueue: NewQueueByList(),
	}
}

type stackByQueues struct {
	dataQueue *queueByList
	helpQueue *queueByList
}

func (s *stackByQueues) Push(v int) {
	s.dataQueue.Push(v)
}

func (s *stackByQueues) Pop() (int, error) {
	for s.dataQueue.Size() > 1 {
		if v, err := s.dataQueue.Poll(); err != nil {
			return 0, err
		} else {
			s.helpQueue.Push(v)
		}
	}
	v, err := s.dataQueue.Poll()
	if err != nil {
		return 0, err
	}
	s.dataQueue, s.helpQueue = s.helpQueue, s.dataQueue
	return v, nil
}

func (s *stackByQueues) Peek() (int, error) {
	for s.dataQueue.Size() > 1 {
		if v, err := s.dataQueue.Poll(); err != nil {
			return 0, err
		} else {
			s.helpQueue.Push(v)
		}
	}
	v, err := s.dataQueue.Poll()
	if err != nil {
		return 0, err
	}
	s.helpQueue.Push(v) // 将出栈的数据再次入栈[helpQueue 即将变为 dataQueue]
	s.dataQueue, s.helpQueue = s.helpQueue, s.dataQueue
	return v, nil
}

func (s *stackByQueues) IsEmpty() bool {
	return s.dataQueue.Size() == 0
}
