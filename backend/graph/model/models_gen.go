// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type CreateIngredientInput struct {
	Type             IngredientType `json:"type"`
	Number           string         `json:"number"`
	Name             string         `json:"name"`
	UnitTypeReceived string         `json:"unitTypeReceived"`
	QuantityPerUnit  float64        `json:"quantityPerUnit"`
	FinalUnit        string         `json:"finalUnit"`
}

type Ingredient struct {
	ID               string         `json:"id"`
	Type             IngredientType `json:"type"`
	Number           string         `json:"number"`
	Name             string         `json:"name"`
	UnitTypeReceived string         `json:"unitTypeReceived"`
	QuantityPerUnit  float64        `json:"quantityPerUnit"`
	FinalUnit        string         `json:"finalUnit"`
}

type IngredientReceived struct {
	ID               string      `json:"id"`
	Ingredient       *Ingredient `json:"ingredient"`
	UnitsReceived    int32       `json:"unitsReceived"`
	LotNumber        string      `json:"lotNumber"`
	ExpirationDate   string      `json:"expirationDate"`
	UnitType         string      `json:"unitType"`
	QuantityReceived float64     `json:"quantityReceived"`
	UnitOfMeasure    string      `json:"unitOfMeasure"`
	Supplier         *Supplier   `json:"supplier"`
	User             *User       `json:"user"`
	ReceiveDate      string      `json:"receiveDate"`
	InvoiceNumber    string      `json:"invoiceNumber"`
}

type Mutation struct {
}

type Query struct {
}

type Supplier struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateIngredientInput struct {
	Type             *IngredientType `json:"type,omitempty"`
	Number           *string         `json:"number,omitempty"`
	Name             *string         `json:"name,omitempty"`
	UnitTypeReceived *string         `json:"unitTypeReceived,omitempty"`
	QuantityPerUnit  *float64        `json:"quantityPerUnit,omitempty"`
	FinalUnit        *string         `json:"finalUnit,omitempty"`
}

type User struct {
	ID        string     `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Status    UserStatus `json:"status"`
}

type IngredientType string

const (
	IngredientTypeRawMaterial IngredientType = "RAW_MATERIAL"
	IngredientTypePackaging   IngredientType = "PACKAGING"
)

var AllIngredientType = []IngredientType{
	IngredientTypeRawMaterial,
	IngredientTypePackaging,
}

func (e IngredientType) IsValid() bool {
	switch e {
	case IngredientTypeRawMaterial, IngredientTypePackaging:
		return true
	}
	return false
}

func (e IngredientType) String() string {
	return string(e)
}

func (e *IngredientType) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IngredientType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid IngredientType", str)
	}
	return nil
}

func (e IngredientType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserStatus string

const (
	UserStatusAdmin    UserStatus = "ADMIN"
	UserStatusOperator UserStatus = "OPERATOR"
)

var AllUserStatus = []UserStatus{
	UserStatusAdmin,
	UserStatusOperator,
}

func (e UserStatus) IsValid() bool {
	switch e {
	case UserStatusAdmin, UserStatusOperator:
		return true
	}
	return false
}

func (e UserStatus) String() string {
	return string(e)
}

func (e *UserStatus) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserStatus", str)
	}
	return nil
}

func (e UserStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}