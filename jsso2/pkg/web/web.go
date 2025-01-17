// Package web generates hyperlinks for the web app.
package web

import (
	"fmt"
	"net/url"
	"strings"
)

type Linker struct {
	BaseURL *url.URL
}

func NewLinker(base string) (*Linker, error) {
	baseURL, err := url.Parse(base)
	if err != nil {
		return nil, fmt.Errorf("parse base url: %w", err)
	}
	baseURL.Path = strings.TrimSuffix(baseURL.Path, "/")
	return &Linker{
		BaseURL: baseURL,
	}, nil
}

func (l *Linker) Origin() string {
	return l.BaseURL.Scheme + "://" + l.BaseURL.Host
}

func (l *Linker) Domain() string {
	return l.BaseURL.Hostname()
}

func (l *Linker) RPID() string {
	return l.Domain()
}

func (l *Linker) Base() string {
	return l.BaseURL.String()
}

func (l *Linker) EnrollmentPage(token string) string {
	return l.BaseURL.String() + "/#/enroll/" + token
}

func (l *Linker) LoginPage() string {
	return l.BaseURL.String() + "/#/login"
}

func (l *Linker) LoginPageWithRedirect(destination string) string {
	return l.BaseURL.String() + "/#/login/" + url.PathEscape(destination)
}

func (l *Linker) SetCookie(cookie string) string {
	return l.BaseURL.String() + "/set-cookie?set=" + cookie
}
