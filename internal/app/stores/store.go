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
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

const (
	charset  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // Набор символов для генерации ID
	idLength = 10                                                     // Длина генерируемого идентификатора
)

// Store Структура Store с Read/Write Mutex и map для хранения коротких и полных URL.
// И вторая мапа для быстрого поиска.
type Store struct {
	sync.RWMutex
	shortToOriginal map[string]string
	originalToShort map[string]string
}

// NewStore Функция для создания нового Store с пустой map.
func NewStore() *Store {
	return &Store{
		shortToOriginal: make(map[string]string),
		originalToShort: make(map[string]string),
	}
}

// Set Функция для добавления нового URL в Store. Мы генерируем короткий ID, добавляем его в map,
// а затем возвращаем этот ID. Блокировка обеспечивает, что добавление происходит атомарно.
func (s *Store) Set(url string) (string, bool) {
	s.Lock()
	defer s.Unlock()

	// Проверяем, существует ли уже оригинальный URL
	if shortURL, exists := s.originalToShort[url]; exists {
		return shortURL, false // URL уже существует, возвращаем существующий короткий URL и false
	}

	// Генерируем и добавляем новый сокращенный URL
	shortURL := generateID() // Предполагается, что эта функция генерирует уникальный ID
	s.originalToShort[url] = shortURL
	s.shortToOriginal[shortURL] = url

	return shortURL, true // Возвращаем новый короткий URL и true
}

// Get Функция для получения полного URL по короткому идентификатору из Store.
// Если идентификатор найден, возвращается полный URL и флаг exists будет true.
// Если идентификатор не найден, возвращается пустая строка и флаг exists будет false.
// Read Lock используется для обеспечения безопасного доступа на чтение.
func (s *Store) Get(shortURL string) (string, bool) {
	s.RLock()
	defer s.RUnlock()
	url, exists := s.shortToOriginal[shortURL]

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

func (s *Store) Save(filepath string) error {
	s.RLock()
	defer s.RUnlock()

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	// Начинаем с 1 для генерации уникальных идентификаторов (UUID)
	uuid := 1

	// Итерируем по каждой паре короткий URL -> оригинальный URL в карте
	for shortURL, originalURL := range s.shortToOriginal {
		// Пытаемся закодировать данные в формате JSON и записать их в файл
		if err := encoder.Encode(map[string]string{
			"uuid":         fmt.Sprintf("%d", uuid), // Преобразуем числовой ID в строку для UUID
			"short_url":    shortURL,
			"original_url": originalURL,
		}); err != nil {
			return err // Возвращаем ошибку, если не удается закодировать или записать данные
		}
		uuid++ // Увеличиваем UUID для следующей записи
	}

	return nil
}

func (s *Store) LoadFromFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // Ничего не делаем, если файл не существует
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var urlEntry map[string]string
	for decoder.More() {
		if err := decoder.Decode(&urlEntry); err != nil {
			return err
		}
		// заполняем мапы данными.
		shortURL := urlEntry["short_url"]
		originalURL := urlEntry["original_url"]
		s.shortToOriginal[shortURL] = originalURL
		s.originalToShort[originalURL] = shortURL
	}

	return nil
}
