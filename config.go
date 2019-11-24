package tonicpow

import (
	"net"
	"net/http"
	"time"
)

// APIEnvironment is used internally to represent the possible values
type APIEnvironment string

// Package global constants and configuration
const (
	// LiveEnvironment is where we POST queries to (live)
	LiveEnvironment APIEnvironment = "https://api.tonicpow.com/"

	// TestEnvironment is where we POST queries to (testing)
	//TestEnvironment APIEnvironment = "https://test.tonicpow.com/"

	// LocalEnvironment is where we POST queries to (local)
	LocalEnvironment APIEnvironment = "http://localhost:3000/"

	// ConnectionExponentFactor backoff exponent factor
	ConnectionExponentFactor float64 = 2.0

	// ConnectionInitialTimeout initial timeout
	ConnectionInitialTimeout = 2 * time.Millisecond

	// ConnectionMaximumJitterInterval jitter interval
	ConnectionMaximumJitterInterval = 2 * time.Millisecond

	// ConnectionMaxTimeout max timeout
	ConnectionMaxTimeout = 10 * time.Millisecond

	// ConnectionRetryCount retry count
	ConnectionRetryCount int = 2

	// ConnectionWithHTTPTimeout with http timeout
	ConnectionWithHTTPTimeout = 10 * time.Second

	// ConnectionTLSHandshakeTimeout tls handshake timeout
	ConnectionTLSHandshakeTimeout = 5 * time.Second

	// ConnectionMaxIdleConnections max idle http connections
	ConnectionMaxIdleConnections int = 10

	// ConnectionIdleTimeout idle connection timeout
	ConnectionIdleTimeout = 20 * time.Second

	// ConnectionExpectContinueTimeout expect continue timeout
	ConnectionExpectContinueTimeout = 3 * time.Second

	// ConnectionDialerTimeout dialer timeout
	ConnectionDialerTimeout = 5 * time.Second

	// ConnectionDialerKeepAlive keep alive
	ConnectionDialerKeepAlive = 20 * time.Second

	// DefaultUserAgent is the default user agent for all requests
	DefaultUserAgent string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.80 Safari/537.36"
)

// HTTP and Dialer connection variables
var (
	// _Dialer net dialer for ClientDefaultTransport
	_Dialer = &net.Dialer{
		KeepAlive: ConnectionDialerKeepAlive,
		Timeout:   ConnectionDialerTimeout,
	}

	// ClientDefaultTransport is the default transport struct for the HTTP client
	ClientDefaultTransport = &http.Transport{
		DialContext:           _Dialer.DialContext,
		ExpectContinueTimeout: ConnectionExpectContinueTimeout,
		IdleConnTimeout:       ConnectionIdleTimeout,
		MaxIdleConns:          ConnectionMaxIdleConnections,
		Proxy:                 http.ProxyFromEnvironment,
		TLSHandshakeTimeout:   ConnectionTLSHandshakeTimeout,
	}
)
