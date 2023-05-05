# 1. ArrayList

ArrayList是一个数组结构的存储容器。

ArrayList()会使用长度为0的数组，ArrayList(int initialCapacity)会使用指定容量的数组。

add(Object o)首次扩容为10，再次扩容为上次容量的1.5倍。

addAll(Collection c)没有元素时，扩容为Math.max(10,实际元素个数)，有元素时扩容为Math.max(原容量1.5倍,实际元素个数)

# 2. HashMap和HashTable

HashMap是采用数组+链表+红黑树实现的一种数据结构-------------

HashMap和HashTable都是基于hash表实现K-V结构的集合。

区别：

1. HashTable是线程安全的，它给所有的数据访问方法都加上了一个Synchronized同步锁，而HashMap是线程不安全的
2. HashTable内部采用数据+链表来实现，链表用来解决hash冲突的问题，而HashMap采用数组+链表+红黑树实现，当链表长度>8且数组长度>64的时候，就会把链表转换为红黑树，当链表长度>8而数组长度<64时数组会进行扩容。
3. HashMap的key可以为null，而HashTable的key不可以为null，在HashMap中会把null作为0存储。
4. key使用的散列算法不同，HashTable直接使用key的hashcode对数组长度取模，而HashMap对key的hashcode做了二次散列，从而避免key分布不均匀影响到查询性能。

# 3. TreeSet

