package cycletls

import (
	"time"

	http "github.com/useflyent/fhttp"

	"golang.org/x/net/proxy"
)

type browser struct {
	// Return a greeting that embeds the name in a message.
	JA3       string
	UserAgent string
	Cookies   []Cookie
}

// newClient creates a new http client
func newClient(browser browser, allowRedirect bool, proxyURL ...string) (http.Client, error) {
	//fix check PR

	if len(proxyURL) > 0 && len(proxyURL[0]) > 0 {
		dialer, err := newConnectDialer(proxyURL[0])
		if err != nil {
			if allowRedirect {
				return http.Client{
					Timeout: time.Second * 6,
				}, err
			}
			return http.Client{
				Timeout: time.Second * 6,
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
			}, err
		}
		if allowRedirect {
			return http.Client{
				Transport: newRoundTripper(browser, dialer),
				Timeout:   time.Second * 6,
			}, nil
		}
		return http.Client{
			Transport: newRoundTripper(browser, dialer),
			Timeout:   time.Second * 6,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}, nil
	}
	if allowRedirect {
		return http.Client{
			Transport: newRoundTripper(browser, proxy.Direct),
			Timeout:   time.Second * 6,
		}, nil
	}
	return http.Client{
		Transport: newRoundTripper(browser, proxy.Direct),
		Timeout:   time.Second * 6,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}, nil

}
