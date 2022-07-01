package _7_binary_tree

/**
 * 并查集 [Disjoint-set data structure]
 *
 * 定义[维基百科]:
 *      一种数据结构[树形结构], 用于处理一些不交集(Disjoint sets，一系列没有重复元素的集合)的合并及查询问题
 *      因为支持查询和合并两种操作, 并查集也常被称为联合-查找数据结构(Union-find data structure),
 *      或者联合-查找集合(Merge-find set)
 *
 * 支持方法:
 *      查询: 查询某个袁术属于哪个集合, 通常返回集合内的代表元素[一般用于判别两个元素是否属于同一个集合]
 *      合并: 将两个集合合并
 *      添加: 添加一个新的集合[只包含新元素本身]. 此方法基本被忽略, 故常见方法一般代指以上两种
 */