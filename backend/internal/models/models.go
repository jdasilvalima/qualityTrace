package models

import "time"

type UserStatus string

const (
  AdminStatus    UserStatus = "ADMIN"
  OperatorStatus UserStatus = "OPERATOR"
)

type User struct {
  ID        string     `json:"id"`
  FirstName string     `json:"firstName"`
  LastName  string     `json:"lastName"`
  Status    UserStatus `json:"status"`
}

type Supplier struct {
  ID    string `json:"id"`
  Name  string `json:"name"`
  Email string `json:"email"`
}

type IngredientType string

const (
  RawMaterial IngredientType = "RAW_MATERIAL"
  Packaging   IngredientType = "PACKAGING"
)

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
  ID              string      `json:"id"`
  IngredientID    string      `json:"ingredientId"`
  UnitsReceived   int         `json:"unitsReceived"`
  LotNumber       string      `json:"lotNumber"`
  ExpirationDate  time.Time   `json:"expirationDate"`
  UnitType        string      `json:"unitType"`
  QuantityReceived float64    `json:"quantityReceived"`
  UnitOfMeasure   string      `json:"unitOfMeasure"`
  SupplierID      string      `json:"supplierId"`
  UserID          string      `json:"userId"`
  ReceiveDate     time.Time   `json:"receiveDate"`
  InvoiceNumber   string      `json:"invoiceNumber"`
}