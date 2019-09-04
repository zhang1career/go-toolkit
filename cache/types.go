package cache

type Cache struct {
	value interface{}
}


func (this *Cache) Get() interface{} {
	return this.value
}

func (this *Cache) Set(value interface{}) interface{} {
	this.value = value
	return this.value
}

func (this *Cache) Inc() interface{} {
	return this.value
}