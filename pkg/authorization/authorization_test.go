package authorization

import (
	"fmt"
	"testing"

	"github.com/equipindustry/visanetgo/test"
	"github.com/stretchr/testify/assert"
)

func TestAuthorization_CreateAuthorization(t *testing.T) {
	authorizationFactory := NewAuthorizationSuccessResponseFactory()
	authorizationFactory.SetAuthorization(test.AuthorizationResponse)
	assert.NotNil(t, authorizationFactory.GetAuthorization())
	assert.EqualValues(t, authorizationFactory.GetAuthorization().DataMap.ActionCode, "000")
	assert.EqualValues(t, authorizationFactory.GetAuthorization().DataMap.ActionDescription, "Aprobado y completado con exito")
	assert.EqualValues(t, authorizationFactory.GetAuthorization().DataMap.Status, "Verified")
}

func TestAuthorization_CreateAuthorizationFailed(t *testing.T) {
	authorizationFactory := NewFailAuthorizationSuccessResponseFactory()
	authorizationFactory.SetFailAuthorization(test.AuthorizationFailResponse)
	fmt.Println(authorizationFactory.GetFailAuthorization())
	assert.NotNil(t, authorizationFactory.GetFailAuthorization())
	assert.EqualValues(t, authorizationFactory.GetFailAuthorization().Data.ActionCode, "180")
	assert.EqualValues(t, authorizationFactory.GetFailAuthorization().Data.ActionDescription, "Tarjeta no valida")
	assert.EqualValues(t, authorizationFactory.GetFailAuthorization().Data.Status, "Not Authorized")
}

func TestAuthorization_CreateAuthorizationTokenFailed(t *testing.T) {
	authorizationFactory := NewFailAuthorizationSuccessResponseFactory()
	authorizationFactory.SetFailAuthorization(test.AuthorizationTokenNotExists)
	fmt.Println(authorizationFactory.GetFailAuthorization())
	assert.NotNil(t, authorizationFactory.GetFailAuthorization())
	assert.EqualValues(t, authorizationFactory.GetFailAuthorization().ErrorMessage, "Token id does not exist")
}

func TestAuthorization_CreateAuthorizationPayloadFailed(t *testing.T) {
	authorizationFactory := NewFailAuthorizationSuccessResponseFactory()
	authorizationFactory.SetFailAuthorization(test.AuthorizationFailPayload)
	fmt.Println(authorizationFactory.GetFailAuthorization())
	assert.NotNil(t, authorizationFactory.GetFailAuthorization())
	assert.EqualValues(t, authorizationFactory.GetFailAuthorization().ErrorMessage, "PurchaseNumber is wrong or empty.")
}
