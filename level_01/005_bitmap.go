package level_01

// BitMap 位图
// 用少量的 64 位整形 slice 存入大量值
// slice 索引代表存入数据在那个段，如 索引 0 表示 0 ~ 63, 1 表示 64 ~ 127
// slice 索引值的 64 个比特位代表存入了哪些值, 哪个 bit 上有 1，则代表那个
type BitMap struct {
	bits []int64
}

// NewBitMap 设置需要存入最大的数
func NewBitMap(max int) *BitMap {
	// (max + 64) >> 6) 代表了最大值 max 需要多少个 int64 来存人, 如 0 ~ 63 需要 1 个, 64 ~ 127 需要 2 个
	return &BitMap{bits: make([]int64, (max + 64) >> 6)}
}

// add 添加一个数
func (b *BitMap) add(n int)  {
	// n >> 64 eq n / 64    代表 n 这个数落到 bit 的那个索引段上, 如 0 ~ 63 落到 0 上, 64 ~ 127 落到 1 上
	// n & 63 eq n % 64    代表 n 在其对应索引段的值的第几个 bit 位上, 如 0 就在 0 的第 1 位上，64 在 1 的第 1 位上
	// 将 int64(1) 左移 (n & 63) 位后, 与 b.bit[n >> 6] 的值做 | 运算, 将其对应的 bit 位上标 1, 代表这个数在 b.bits 里面了
	b.bits[n >> 6] |= int64(1) << (n & 63)
}

// delete 删除一个数
func (b *BitMap) delete(n int)  {
	// 同理 add, 删除就是将索引段上的对应 bit 位上的 1 标为 0, 所以将 1 左移对应的位数后取反，然后做 & 运算
	b.bits[n >> 6] &= ^(int64(1) << (n & 63))
}

// contain bits 是否包含这个数
func (b *BitMap) contain(n int) bool {
	// 同理 add, n 在 bits 索引段上对应 bit 位上的值和将 1 左移对应位数做 & 运算, 看结果是否为 0, 不等于 0 就代表包含
	return (b.bits[n >> 6] & (int64(1) << (n & 63))) != 0
}
