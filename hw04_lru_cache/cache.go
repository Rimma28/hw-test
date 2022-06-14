package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
func (l *lruCache) Clear() {
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

func (l *lruCache) Get(k Key) (interface{}, bool) {
	// пытаемся получить элемент из словаря по ключу
	if mapItem, ok := l.items[k]; ok {
		// элемент присутствует в словаре, перемещаем в начало очереди
		l.queue.MoveToFront(mapItem)
		// и возвращаем
		return mapItem.Value.(*cacheItem).value, true
	}
	return nil, false
}

func (l *lruCache) Set(k Key, value interface{}) bool {
	// формируем элемент кеша
	cItem := &cacheItem{k, value}
	// пытаемся получить элемент из словаря по ключу
	mapItem, ok := l.items[k]

	if ok {
		// элемент присутствует в словаре, меняем значение и перемещаем в начало очереди
		mapItem.Value = cItem
		l.queue.MoveToFront(mapItem)
	} else {
		// элемент отсутствует
		// проверка переполнения кеша
		if l.queue.Len() >= l.capacity {
			// не хватает места для нового элемента
			// получаем последний элемент очереди
			lastItem := l.queue.Back()

			// удаляем его из очереди
			l.queue.Remove(lastItem)

			// и из словаря
			delete(l.items, lastItem.Value.(*cacheItem).key)
		}

		// добавляем в очередь и в словарь
		l.items[k] = l.queue.PushFront(cItem)
	}
	return ok
}
