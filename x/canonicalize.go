package x

import (
	"github.com/pkg/errors"
	"golang.org/x/net/idna"
	"net/url"
	"strings"
)

func CanonicalizeEmail(email string) (string, error) {
	components := strings.Split(email, "@")

	if len(components) < 2 {
		return "", errors.Errorf("expected an email address but got: %s", email)
	}

	domain := components[len(components)-1]
	asciiDomain, err := idna.ToASCII(domain)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return strings.Join(components[:len(components)-1], "@") + "@" + asciiDomain, nil
}

func CanonicalizeURI(uri string) (string, error) {
	parsed, err := url.ParseRequestURI(uri)
	if err != nil {
		return "", errors.WithStack(err)
	}

	asciiHost, err := idna.ToASCII(parsed.Hostname())
	if err != nil {
		return "", errors.WithStack(err)
	}

	parsed.Host = asciiHost

	if parsed.Port() != "" {
		parsed.Host += ":" + parsed.Port()
	}

	return parsed.String(), nil
}
