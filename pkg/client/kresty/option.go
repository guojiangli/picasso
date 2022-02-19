package kresty

import (
	"net"
	"net/http"
	"runtime"
	"time"

	"github.com/guojiangli/picasso/pkg/klog"
	"github.com/guojiangli/picasso/pkg/klog/baselogger"
)

var defaultTransport = createTransport()

type Option struct {
	HTTPClient *http.Client
	Logger     baselogger.Logger
	Timeout    time.Duration
}

func NewOption() *Option {
	return new(Option)
}

func SetLogger(l baselogger.Logger) *Option {
	c := NewOption()
	c.Logger = l
	return c
}

func (c *Option) SetHTTPClient(l *http.Client) *Option {
	c.HTTPClient = l
	return c
}

func SetHTTPClient(l *http.Client) *Option {
	c := NewOption()
	c.HTTPClient = l
	return c
}

func SetTimeout(t time.Duration) *Option {
	c := NewOption()
	c.Timeout = t
	return c
}

func (c *Option) SetLogger(l baselogger.Logger) *Option {
	c.Logger = l
	return c
}

func (c *Option) SetTimeout(t time.Duration) *Option {
	c.Timeout = t
	return c
}

func (c *Option) MergeOption(opts ...*Option) *Option {
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if opt.Logger != nil {
			c.Logger = opt.Logger
		}
		if opt.HTTPClient != nil {
			c.HTTPClient = opt.HTTPClient
		}
		if opt.Timeout != 0 {
			c.Timeout = opt.Timeout
		}
	}
	return c
}

func defaultOption() *Option {
	l := klog.DefaultLogger().With().Str("kepler", "resty").Logger()
	return &Option{
		Logger: klog.NewLoggerWith(&l), HTTPClient: &http.Client{Transport: defaultTransport, Timeout: time.Second * 5},
	}
}

func createTransport() *http.Transport {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}
	return &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
	}
}
