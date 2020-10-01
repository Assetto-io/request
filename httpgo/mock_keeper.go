package httpgo

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"sync"
)

type MockKeeper struct {
	keeperMutex sync.Mutex

	mocks map[string]*Mock
}

func (m *MockKeeper) DeleteMocks() {
	m.keeperMutex.Lock()
	defer m.keeperMutex.Unlock()

	m.mocks = make(map[string]*Mock)
}

func (m *MockKeeper) AddMock(mock Mock) {
	m.keeperMutex.Lock()
	defer m.keeperMutex.Unlock()

	key := m.getMockKey(mock.Method, mock.Url, mock.RequestBody)
	m.mocks[key] = &mock
}

func (m *MockKeeper) getMockKey(method, url, body string) string {
	hasher := md5.New()
	hasher.Write([]byte(method + url + m.cleanBody(body)))
	key := hex.EncodeToString(hasher.Sum(nil))
	return key
}

func (m *MockKeeper) cleanBody(body string) string {
	body = strings.TrimSpace(body)
	if body == "" {
		return ""
	}
	body = strings.ReplaceAll(body, "\t", "")
	body = strings.ReplaceAll(body, "\n", "")
	return body
}
