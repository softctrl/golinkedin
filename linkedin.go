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
	"bytes"
	"io/ioutil"
	"net/http"

	l "log"

	req "github.com/softctrl/golinkedin/request"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	lnk "golang.org/x/oauth2/linkedin"
)

const (
	USER_INFO_URL = "https://api.linkedin.com/v1/people/~"
	SHARE_URL     = "https://api.linkedin.com/v1/people/~/shares?format=json"
)

type SCLinkedin struct {
	_ClientID     string
	_ClientSecret string
	_Config       *oauth2.Config
	_Ctx          context.Context
	_TokenAuth    *oauth2.Token
}

func NewSCLinkedin() *SCLinkedin {
	_res := SCLinkedin{
		_Ctx: context.Background(),
	}
	return _res.Configure()
}

func NewSCLinkedinWithValues(clientID, clientSecret string) *SCLinkedin {
	return &SCLinkedin{
		_ClientID:     clientID,
		_ClientSecret: clientSecret,
		_Ctx:          context.Background(),
	}
}

func (__obj *SCLinkedin) SetClientId(clientID string) *SCLinkedin {
	__obj._ClientID = clientID
	return __obj
}

func (__obj *SCLinkedin) SetClientSecret(clientSecret string) *SCLinkedin {
	__obj._ClientSecret = clientSecret
	return __obj
}

func (__obj *SCLinkedin) Configure() *SCLinkedin {

	__obj._Config = &oauth2.Config{
		ClientID:     __obj._ClientID,
		ClientSecret: __obj._ClientSecret,
		RedirectURL:  "https://localhost:8080", // TODO under development
		Endpoint: oauth2.Endpoint{
			AuthURL:  lnk.Endpoint.AuthURL,
			TokenURL: lnk.Endpoint.TokenURL,
		},
	}
	return __obj
}

func (__obj *SCLinkedin) GetPermissionUrl(state string) string {

	return __obj._Config.AuthCodeURL(string(state), oauth2.AccessTypeOffline)

}

func (__obj *SCLinkedin) Exchange(code string) error {

	_token, _err := __obj._Config.Exchange(__obj._Ctx, code)
	if _err == nil {
		__obj._TokenAuth = _token
	}
	return _err

}

func (__obj *SCLinkedin) Client() *http.Client {

	return __obj._Config.Client(__obj._Ctx, __obj._TokenAuth)

}

func (__obj *SCLinkedin) GetUserInfo() ([]byte, error) {

	client := __obj.Client()
	resp, err := client.Get(USER_INFO_URL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil

}

//
//
//
func (__obj *SCLinkedin) ShareContentURL(__url string) ([]byte, error) {

	_share := req.NewShare()
	_share.VisibleToAnyone() // TODO
	_share.SubmitedUrl(__url)

	if _json, _err := _share.ToJson(); _err != nil {
		return nil, _err
	} else {

		l.Printf("JSON(%s)\n", string(_json))
		l.Printf("_share(%+v)\n", _share)

		_reader := bytes.NewReader(_json)

		if _req, _err := http.NewRequest(http.MethodPost, SHARE_URL, _reader); _err != nil {
			return nil, _err
		} else {
			_req.Header.Set("Content-Type", "application/json")
			_req.Header.Set("x-li-format", "json")
			client := __obj.Client()
			if _resp, _err := client.Do(_req); _err != nil {
				return nil, _err
			} else {
				defer _resp.Body.Close()
				_body, _err := ioutil.ReadAll(_resp.Body)
				if _err != nil {
					return nil, _err
				}
				return _body, nil
			}

		}

	}

}
