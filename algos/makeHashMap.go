package algos

type MyHashMap struct {
	nums map[int]int
}

func Constructor() MyHashMap {
	return MyHashMap{
		nums: make(map[int]int),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	this.nums[key] = value
	return
}

func (this *MyHashMap) Get(key int) int {
	if num, ok := this.nums[key]; ok {
		return num
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	delete(this.nums, key)
}
