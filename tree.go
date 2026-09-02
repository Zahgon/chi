package chi

import (
	"net/http"
	"regexp"
)

type methodTyp uint

const (
	mSTUB methodTyp = 1 << iota
	mCONNECT
	mDELETE
	mGET
	mHEAD
	mOPTIONS
	mPATCH
	mPOST
	mPUT
	mQUERY
	mTRACE
)

var mALL = mCONNECT | mDELETE | mGET | mHEAD |
	mOPTIONS | mPATCH | mPOST | mPUT | mQUERY | mTRACE

const methodQuery = "QUERY"

var methodMap = map[string]methodTyp{
	http.MethodConnect: mCONNECT,
	http.MethodDelete:  mDELETE,
	http.MethodGet:     mGET,
	http.MethodHead:    mHEAD,
	http.MethodOptions: mOPTIONS,
	http.MethodPatch:   mPATCH,
	http.MethodPost:    mPOST,
	http.MethodPut:     mPUT,
	methodQuery:        mQUERY,
	http.MethodTrace:   mTRACE,
}

var reverseMethodMap = map[methodTyp]string{
	mCONNECT: http.MethodConnect,
	mDELETE:  http.MethodDelete,
	mGET:     http.MethodGet,
	mHEAD:    http.MethodHead,
	mOPTIONS: http.MethodOptions,
	mPATCH:   http.MethodPatch,
	mPOST:    http.MethodPost,
	mPUT:     http.MethodPut,
	mQUERY:   methodQuery,
	mTRACE:   http.MethodTrace,
}

func RegisterMethod(method string) { _ = "STUB: not implemented"; return }

type nodeTyp uint8

const (
	ntStatic nodeTyp = iota
	ntRegexp
	ntParam
	ntCatchAll
)

type node struct {
	subroutes Routes

	rex *regexp.Regexp

	endpoints endpoints

	prefix string

	children [ntCatchAll + 1]nodes

	tail byte

	typ nodeTyp

	label byte
}

type endpoints map[methodTyp]*endpoint

type endpoint struct {
	handler http.Handler

	pattern string

	paramKeys []string
}

func (s endpoints) Value(method methodTyp) *endpoint { _ = "STUB: not implemented"; return nil }

func (n *node) InsertRoute(method methodTyp, pattern string, handler http.Handler) *node {
	_ = "STUB: not implemented"
	return nil
}

func (n *node) addChild(child *node, prefix string) *node { _ = "STUB: not implemented"; return nil }

func (n *node) replaceChild(label, tail byte, child *node) { _ = "STUB: not implemented"; return }

func (n *node) getEdge(ntyp nodeTyp, label, tail byte, prefix string) *node {
	_ = "STUB: not implemented"
	return nil
}

func (n *node) setEndpoint(method methodTyp, handler http.Handler, pattern string) {
	_ = "STUB: not implemented"
	return
}

func (n *node) FindRoute(rctx *Context, method methodTyp, path string) (*node, endpoints, http.Handler) {
	_ = "STUB: not implemented"
	return nil, *new(endpoints), *new(http.Handler)
}

func (n *node) findRoute(rctx *Context, method methodTyp, path string) *node {
	_ = "STUB: not implemented"
	return nil
}

func (n *node) findEdge(ntyp nodeTyp, label byte) *node { _ = "STUB: not implemented"; return nil }

func (n *node) isLeaf() bool { _ = "STUB: not implemented"; return false }

func (n *node) findPattern(pattern string) bool { _ = "STUB: not implemented"; return false }

func (n *node) routes() []Route { _ = "STUB: not implemented"; return nil }

func equalHandlers(a, b http.Handler) bool { _ = "STUB: not implemented"; return false }

func (n *node) walk(fn func(eps endpoints, subroutes Routes) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func patNextSegment(pattern string) (nodeTyp, string, string, byte, int, int) {
	_ = "STUB: not implemented"
	return *new(nodeTyp), "", "", 0, 0, 0
}

func patParamKeys(pattern string) []string { _ = "STUB: not implemented"; return nil }

func longestPrefix(k1, k2 string) (i int) { _ = "STUB: not implemented"; return 0 }

type nodes []*node

func (ns nodes) Sort()              { _ = "STUB: not implemented"; return }
func (ns nodes) Len() int           { _ = "STUB: not implemented"; return 0 }
func (ns nodes) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (ns nodes) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (ns nodes) tailSort() { _ = "STUB: not implemented"; return }

func (ns nodes) findEdge(label byte) *node { _ = "STUB: not implemented"; return nil }

type Route struct {
	SubRoutes Routes
	Handlers  map[string]http.Handler
	Pattern   string
}

type WalkFunc func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error

func Walk(r Routes, walkFn WalkFunc) error { _ = "STUB: not implemented"; return nil }

func walk(r Routes, walkFn WalkFunc, parentRoute string, parentMw ...func(http.Handler) http.Handler) error {
	_ = "STUB: not implemented"
	return nil
}
