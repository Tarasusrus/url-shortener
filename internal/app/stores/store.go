package stores

// Импортируем нужные нам пакеты
import (
	"math/rand"
	"sync"
	"time"
)

// Определяем константу символов для генерации короткого ID
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Store Структура Store с Read/Write Mutex и map для хранения коротких и полных URL
type Store struct {
	sync.RWMutex
	Urls map[string]string
}

// NewStore Функция для создания нового Store с пустой map
func NewStore() *Store {
	return &Store{
		Urls: make(map[string]string),
	}
}

// Set Функция для добавления нового URL в Store. Мы генерируем короткий ID, добавляем его в map,
// а затем возвращаем этот ID. Блокировка обеспечивает, что добавление происходит атомарно.
func (s *Store) Set(url string) string {
	s.Lock()
	defer s.Unlock()
	id := generateID()
	s.Urls[id] = url
	return id
}

// Get Функция для получения полного URL по короткому идентификатору из Store.
// Если идентификатор найден, возвращается полный URL и флаг exists будет true.
// Если идентификатор не найден, возвращается пустая строка и флаг exists будет false.
// Read Lock используется для обеспечения безопасного доступа на чтение.
func (s *Store) Get(id string) (string, bool) {
	s.RLock()
	defer s.RUnlock()
	url, exists := s.Urls[id]
	return url, exists
}

// Функция для генерации случайного идентификатора из заданного набора символов (charset).
// Идентификатор генерируется как строка длиной 10 символов.
func generateID() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
	return string(b)
}
