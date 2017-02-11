//
// Author:
//  Carlos Timoshenko
//  carlostimoshenkorodrigueslopes@gmail.com
//
//  https://github.com/softctrl
//
// This project is free software; you can redistribute it and/or
// modify it under the terms of the GNU Lesser General Public
// License as published by the Free Software Foundation; either
// version 3 of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// Lesser General Public License for more details.
//
package golinkedin

import (
	"net/http"

	l "log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	lnk "golang.org/x/oauth2/linkedin"
)

type SCLinkedin struct {
	_ClientID     string
	_ClientSecret string
	_Config       oauth2.Config
	_Ctx          context.Context
	_TokenAuth    oauth2.Token
}

func NewSCLinkedin() SCLinkedin {
	return SCLinkedin{
		_Ctx: context.Background(),
	}.Configure()
}

func NewSCLinkedinWithValues(clientID, clientSecret string) *SCLinkedin {
	return &SCLinkedin{
		_ClientID:     clientID,
		_ClientSecret: clientSecret,
		_Ctx:          context.Background(),
	}
}

func (__obj SCLinkedin) SetClientId(clientID string) SCLinkedin {
	__obj._ClientID = clientID
	return __obj
}

func (__obj SCLinkedin) SetClientSecret(clientSecret string) SCLinkedin {
	__obj._ClientSecret = clientSecret
	return __obj
}

func (__obj SCLinkedin) Configure() SCLinkedin {

	__obj._Config = oauth2.Config{
		ClientID:     __obj._ClientID,
		ClientSecret: __obj._ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  lnk.Endpoint.AuthURL,
			TokenURL: lnk.Endpoint.TokenURL,
		},
	}
	return __obj
}

func (__obj SCLinkedin) GetPermissionUrl(state string) string {

	return __obj._Config.AuthCodeURL(string(state), oauth2.AccessTypeOffline)

}

func (__obj SCLinkedin) Exchange(code string) error {

	_token, _err := __obj._Config.Exchange(__obj._Ctx, code)
	if _err == nil {
		l.Printf("Error(%s)", _err.Error())
		__obj._TokenAuth = *_token
	}
	return _err

}

func (__obj SCLinkedin) Client() *http.Client {

	return __obj._Config.Client(__obj._Ctx, &__obj._TokenAuth)

}
