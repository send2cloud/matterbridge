// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package einterfaces

import (
	"github.com/mattermost/platform/model"
)

type SamlInterface interface {
	ConfigureSP() *model.AppError
	BuildRequest(relayState string) (*model.SamlAuthRequest, *model.AppError)
	DoLogin(encodedXML string, relayState map[string]string) (*model.User, *model.AppError)
	GetMetadata() (string, *model.AppError)
}

var theSamlInterface SamlInterface

func RegisterSamlInterface(newInterface SamlInterface) {
	theSamlInterface = newInterface
}

func GetSamlInterface() SamlInterface {
	return theSamlInterface
}
