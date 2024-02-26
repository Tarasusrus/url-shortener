// Package stores предоставляет реализацию хранилища для управления короткими и полными URL.
// Он включает в себя структуру Store, которая использует синхронизированный доступ к данным
// для обеспечения потокобезопасности при работе с URL.
//
// В хранилище для каждого полного URL генерируется уникальный короткий идентификатор,
// который затем может быть использован для получения полного URL. Генерация идентификатора
// осуществляется с использованием случайного выбора символов из предопределенного набора.
package stores

// Импортируем нужные нам пакеты.
import (
	"math/rand"
	"sync"
	"time"
)

const (
	charset  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // Набор символов для генерации ID
	idLength = 10                                                     // Длина генерируемого идентификатора
)

// Store Структура Store с Read/Write Mutex и map для хранения коротких и полных URL.
type Store struct {
	sync.RWMutex
	urls map[string]string
}

// NewStore Функция для создания нового Store с пустой map.
func NewStore() *Store {
	return &Store{
		urls: make(map[string]string),
	}
}

// Set Функция для добавления нового URL в Store. Мы генерируем короткий ID, добавляем его в map,
// а затем возвращаем этот ID. Блокировка обеспечивает, что добавление происходит атомарно.
func (s *Store) Set(url string) string {
	s.Lock()
	defer s.Unlock()

	shortURLID := generateID()
	s.urls[shortURLID] = url

	return shortURLID
}

// Get Функция для получения полного URL по короткому идентификатору из Store.
// Если идентификатор найден, возвращается полный URL и флаг exists будет true.
// Если идентификатор не найден, возвращается пустая строка и флаг exists будет false.
// Read Lock используется для обеспечения безопасного доступа на чтение.
func (s *Store) Get(id string) (string, bool) {
	s.RLock()
	defer s.RUnlock()
	url, exists := s.urls[id]

	return url, exists
}

// Функция для генерации случайного идентификатора из заданного набора символов (charset).
// Идентификатор генерируется как строка длиной 10 символов.
func generateID() string {
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)

	idBytes := make([]byte, idLength)
	for i := range idBytes {
		idBytes[i] = charset[randomizer.Intn(len(charset))]
	}

	return string(idBytes)
}
