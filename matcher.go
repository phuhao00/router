package router

import _ "github.com/gorilla/mux"

type Matcher struct {
	handle map[uint64]func()
}

func (m *Matcher) GetForwardHandle() {

}
