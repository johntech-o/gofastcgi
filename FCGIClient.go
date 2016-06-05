package main

const FCGI_LISTENSOCK_FILENO uint8 = 0
const FCGI_HEADER_LEN uint8 = 8
const VERSION_1 uint8 = 1
const FCGI_NULL_REQUEST_ID uint8 = 0

// Mask for flags component of FCGI_BeginRequestBody
const FCGI_KEEP_CONN uint8 = 1
const doubleCRLF = "\r\n\r\n"

// Values for type component of FCGI_Header
const (
	FCGI_BEGIN_REQUEST uint8 = iota + 1
	FCGI_ABORT_REQUEST
	FCGI_END_REQUEST
	FCGI_PARAMS
	FCGI_STDIN
	FCGI_STDOUT
	FCGI_STDERR
	FCGI_DATA
	FCGI_GET_VALUES
	FCGI_GET_VALUES_RESULT
	FCGI_UNKNOWN_TYPE
	FCGI_MAXTYPE = FCGI_UNKNOWN_TYPE
)

// Values for role component of FCGI_BeginRequestBody
const (
	FCGI_RESPONDER uint8 = iota + 1
	FCGI_AUTHORIZER
	FCGI_FILTER
)

// Values for protocolStatus component of FCGI_EndRequestBody
const (
	FCGI_REQUEST_COMPLETE uint8 = iota
	FCGI_CANT_MPX_CONN
	FCGI_OVERLOADED
	FCGI_UNKNOWN_ROLE
)

// Management Record Types
const (
	FCGI_MAX_CONNS  string = "MAX_CONNS"
	FCGI_MAX_REQS   string = "MAX_REQS"
	FCGI_MPXS_CONNS string = "MPXS_CONNS"
)

const (
	maxWrite = 65500 // 65535 may work, but for compatibility
	maxPad   = 255
)

type header struct {
	Version       uint8
	Type          uint8
	Id            uint16
	ContentLength uint16
	PaddingLength uint8
	Reserved      uint8
}

type record struct {
	h    header
	rbuf []byte
}

// padding so do not have to allocate for each request
var pad [maxPad]byte

type FCGIClient struct {
	appServers []string          // remote app server addresses
	envPairs   map[string]string // current web server env
	client     *Client
}

// return a fcgiclient connects to specify servers
// @validate remoteAddress
func NewFCGIClient(remoteAddress []string, env map[string]string) (*FCGIClient, error) {
	client := FCGIClient{envPairs: env, appServers: remoteAddress}
	return &client, nil
}

// get the capability of the application  server
func (fc *FCGIClient) GetAppServerCapability() (map[string]string, error) {

}

// send request to application server and get response
func (fc *FCGIClient) GetRequest(request *Request) (response *Response, err error) {
	
type header struct {
	Version       uint8
	Type          uint8
	Id            uint16
	ContentLength uint16
	PaddingLength uint8
	Reserved      uint8
}

type record struct {
	h    header
	rbuf []byte
}

// padding so do not have to allocate for each request
var pad [maxPad]byte

type FCGIClient struct {
	appServers []string          // remote app server addresses
	envPairs   map[string]string // current web server env
	client     *Client
}

// return a fcgiclient connects to specify servers
// @validate remoteAddress
func NewFCGIClient(remoteAddress []string, env map[string]string) (*FCGIClient, error) {
	client := FCGIClient{envPairs: env, appServers: remoteAddress}
	return &client, nil
}

// get the capability of the application  server
func (fc *FCGIClient) GetAppServerCapability() (map[string]string, error) {

}

// send request to application server and get response
func (fc *FCGIClient) GetRequest(request *Request) (response *Response, err error) {
	
type header struct {
	Version       uint8
	Type          uint8
	Id            uint16
	ContentLength uint16
	PaddingLength uint8
	Reserved      uint8
}

type record struct {
	h    header
	rbuf []byte
}

// padding so do not have to allocate for each request
var pad [maxPad]byte

type FCGIClient struct {
	appServers []string          // remote app server addresses
	envPairs   map[string]string // current web server env
	client     *Client
}

// return a fcgiclient connects to specify servers
// @validate remoteAddress
func NewFCGIClient(remoteAddress []string, env map[string]string) (*FCGIClient, error) {
	client := FCGIClient{envPairs: env, appServers: remoteAddress}
	return &client, nil
}

// get the capability of the application  server
func (fc *FCGIClient) GetAppServerCapability() (map[string]string, error) {

}

// send request to application server and get response
func (fc *FCGIClient) GetRequest(request *Request) (response *Response, err error) {
	
type header struct {
	Version       uint8
	Type          uint8
	Id            uint16
	ContentLength uint16
	PaddingLength uint8
	Reserved      uint8
}

type record struct {
	h    header
	rbuf []byte
}

// padding so do not have to allocate for each request
var pad [maxPad]byte

type FCGIClient struct {
	appServers []string          // remote app server addresses
	envPairs   map[string]string // current web server env
	client     *Client
}

// return a fcgiclient connects to specify servers
// @validate remoteAddress
func NewFCGIClient(remoteAddress []string, env map[string]string) (*FCGIClient, error) {
	client := FCGIClient{envPairs: env, appServers: remoteAddress}
	return &client, nil
}

// get the capability of the application  server
func (fc *FCGIClient) GetAppServerCapability() (map[string]string, error) {

}

type record struct {
	h    header
	rbuf []byte
}

// padding so do not have to allocate for each request
var pad [maxPad]byte

type FCGIClient struct {
	appServers []string          // remote app server addresses
	envPairs   map[string]string // current web server env
	client     *Client
}

// return a fcgiclient connects to specify servers
// @validate remoteAddress
func NewFCGIClient(remoteAddress []string, env map[string]string) (*FCGIClient, error) {
	client := FCGIClient{envPairs: env, appServers: remoteAddress}
	return &client, nil
}

// get the capability of the application  server
func (fc *FCGIClient) GetAppServerCapability() (map[string]string, error) {

}

func (this *FCGIClient) writeRecord(recType uint8, content []byte) ( err error) {
	this.h.init(recType, this.reqId, len(content))
	if err := binary.Write(&this.buf, binary.BigEndian, this.h); err != nil {
		return err
	}
	if _, err := this.buf.Write(content); err != nil {
		return err
	}
	if _, err := this.buf.Write(pad[:this.h.PaddingLength]); err != nil {
		return err
	}
	_, err = this.rwc.Write(this.buf.Bytes())
	return err
}


func (this *FCGIClient) writeBeginRequest(role uint16, flags uint8) error {
	b := [8]byte{byte(role >> 8), byte(role), flags}
	return this.writeRecord(FCGI_BEGIN_REQUEST, b[:])
}

type header struct {
	Version       uint8
	Type          uint8
	Id            []byte // 
	ContentLength uint16
	PaddingLength uint8
	Reserved      uint8
}

type BeginRequest struct {
	header
	body 

}


// send request to application server and get response
func (fc *FCGIClient) Call(service string,method string,args map[string]string) (response *Response, err error) {
	// p["REQUEST_METHOD"] = "GET"
	// p["CONTENT_LENGTH"] = "0"
	//	return this.Request(p, nil)
	
	err = this.writeBeginRequest(uint16(FCGI_RESPONDER), 0)	
	if err != nil {
		return
	}
  
	err = this.writePairs(FCGI_PARAMS, p)
	if err != nil {
		return
	}

	body := newWriter(this, FCGI_STDIN)  
	if req != nil {
		io.Copy(body, req)
	}
	body.Close()
  
	r = &streamReader{c:this}
	return 
	if err != nil {
		return
	}
	rb := bufio.NewReader(r)
	tp := textproto.NewReader(rb)
	resp = new(http.Response)
     
	// Parse the response headers.
	mimeHeader, err := tp.ReadMIMEHeader()
	if err != nil {
		return
	}
	resp.Header = http.Header(mimeHeader)

	// TODO: fixTransferEncoding ?
	resp.TransferEncoding = resp.Header["Transfer-Encoding"]
	resp.ContentLength,_ = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
  
	if chunked(resp.TransferEncoding) {
		resp.Body = ioutil.NopCloser(httputil.NewChunkedReader(rb))
	} else {
		resp.Body = ioutil.NopCloser(rb)
	}

	return

}

// call remote servers with specify env
func (fc *FCGIClient) GetRequestWithEnv(env map[string]string, request *Request) (response *Response, err error) {

}

func (fc *FCGIClient) call(controller string, method string, args interface{}, reply interface{}) error {

}
