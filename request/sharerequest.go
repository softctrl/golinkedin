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
package request

import (
	"encoding/json"
)

//
//
//
type _Content struct {
	Title            string `json:"title,omitempty"`
	Description      string `json:"description,omitempty"`
	SubmitedUrl      string `json:"submitted-url,omitempty"`
	SubmitedUrlImage string `json:"submitted-image-url,omitempty"`
}

//
//
//
type Visibility string

const (
	ANYONE           = Visibility("anyone")
	CONNECTIONS_ONLY = Visibility("connections-only")
)

//
//
//
type _Visibility struct {
	Code Visibility `json:"code,omitempty"`
}

//
// Based on:
// https://developer.linkedin.com/docs/share-on-linkedin
//
// {
//   "comment": "Check out developer.linkedin.com!",
//   "content": {
//     "title": "LinkedIn Developers Resources",
//     "description": "Leverage LinkedIn's APIs to maximize engagement",
//     "submitted-url": "https://developer.linkedin.com",
//     "submitted-image-url": "https://example.com/logo.png"
//   },
//   "visibility": {
//     "code": "anyone"
//   }
// }
//
type Share struct {
	Comment    string       `json:"comment,omitempty"`
	Content    *_Content    `json:"content,omitempty"`
	Visibility *_Visibility `json:"visibility,omitempty"`
}

func NewShare() *Share {
	return &Share{
		Content: &_Content{},
		Visibility: &_Visibility{
			Code: CONNECTIONS_ONLY,
		},
	}
}

func (__obj *Share) VisibleToAnyone() *Share {
	__obj.Visibility.Code = ANYONE
	return __obj
}

func (__obj *Share) VisibleToConnectionsOnly() *Share {
	__obj.Visibility.Code = CONNECTIONS_ONLY
	return __obj
}

func (__obj *Share) Title(__title string) *Share {
	__obj.Content.Title = __title
	return __obj
}

func (__obj *Share) Description(__description string) *Share {
	__obj.Content.Description = __description
	return __obj
}

func (__obj *Share) SubmitedUrl(__url string) *Share {
	__obj.Content.SubmitedUrl = __url
	return __obj
}

func (__obj *Share) SubmitedImageUrl(__url string) *Share {
	__obj.Content.SubmitedUrlImage = __url
	return __obj
}

//
//
//
func (__obj *Share) ToJson() ([]byte, error) {
	return json.Marshal(__obj)
}

//
//
//
func FromJson(__data []byte) (*Share, error) {
	var _res Share
	_err := json.Unmarshal(__data, &_res)
	return &_res, _err
}
