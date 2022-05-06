package memory

import (
	"container/list"
	"sync"
	"time"

	"github.com/astaxie/session"
)

var p = &Provider{list: list.New()}

type SessionStore struct {
	id            string
	timeLastVisit time.Time
	value         map[interface{}]interface{}
}

func (s *SessionStore) Set(key interface{}, value interface{}) error {
	s.value[key] = value
	//
	return nil
}
func (s *SessionStore) Get(key interface{}) interface{} {
	//
	if k, ok := s.value[key]; ok {
		return v
	} else {
		return nil
	}

}
func (s *SessionStore) Delete(key interface{}) error {
	delete(s.value, key)
	//
	return nil
}
func SessionID() string {
	return s.id

}

type Provider struct {
	lock     sync.Mutex
	sessions map[string]*list.Element
	list     *list.List
}

func (p *Provider) SessionInit(sid string) (session.Session, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	v := make(map[interface{}]interface{})
	newsess := &SessionStore{sid: sid, timeLastVisit: time.Now(), value: v}
	element := pder.list.PushFront(newsess)
	p.sessions[sid] = element
	return newsess, nil

}
func (p *Provider) SessionRead(sid string) (SessionStore, error) {
	if element, ok := p.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := p.SessionInit(sid)
		return sess, err
	}
	return nil, nil

}
func (p *Provider) SessionExist(sid string) bool {

}
func (p *Provider) SessionRegenerate(oldsid, sid string) (SessionStore, error) {

}
func (p *Provider) SessionDestroy(sid string) error {

}
func (p *Provider) SessionAll() int {

}
func (p *Provider) SessionGC() {

}
