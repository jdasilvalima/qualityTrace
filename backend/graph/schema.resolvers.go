package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.61

import (
	"backend_qualitytrace/graph/model"
	"context"
	"fmt"
)

// CreateIngredient is the resolver for the createIngredient field.
func (r *mutationResolver) CreateIngredient(ctx context.Context, input model.CreateIngredientInput) (*model.Ingredient, error) {
	panic(fmt.Errorf("not implemented: CreateIngredient - createIngredient"))
}

// UpdateIngredient is the resolver for the updateIngredient field.
func (r *mutationResolver) UpdateIngredient(ctx context.Context, id string, input model.UpdateIngredientInput) (*model.Ingredient, error) {
	panic(fmt.Errorf("not implemented: UpdateIngredient - updateIngredient"))
}

// DeleteIngredient is the resolver for the deleteIngredient field.
func (r *mutationResolver) DeleteIngredient(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteIngredient - deleteIngredient"))
}

// Ingredients is the resolver for the ingredients field.
func (r *queryResolver) Ingredients(ctx context.Context) ([]*model.Ingredient, error) {
	panic(fmt.Errorf("not implemented: Ingredients - ingredients"))
}

// Ingredient is the resolver for the ingredient field.
func (r *queryResolver) Ingredient(ctx context.Context, id string) (*model.Ingredient, error) {
	dummyIngredient := &model.Ingredient{
		ID:              id,
		Type:            model.IngredientTypeRawMaterial,
		Number:          "12345",
		Name:            "Flour",
		UnitTypeReceived: "Bag",
		QuantityPerUnit:  25.0,
		FinalUnit:        "kg",
	}
	return dummyIngredient, nil
}

// Suppliers is the resolver for the suppliers field.
func (r *queryResolver) Suppliers(ctx context.Context) ([]*model.Supplier, error) {
	var suppliers []*model.Supplier
	dummySupplier := model.Supplier{
		Name: "EggySupplier",
		Email: "eggy@outlook.com",
	}
	suppliers = append(suppliers, &dummySupplier)
	return suppliers, nil
}

// Supplier is the resolver for the supplier field.
func (r *queryResolver) Supplier(ctx context.Context, id string) (*model.Supplier, error) {
	dummySupplier := &model.Supplier{
		ID:    id,
		Name:  "EggySupplier",
		Email: "eggy@outlook.com",
	}
	return dummySupplier, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	dummyUsers := []*model.User{
		{
			ID:        "1",
			FirstName: "John",
			LastName:  "Doe",
			Status:    model.UserStatusAdmin,
		},
		{
			ID:        "2",
			FirstName: "Jane",
			LastName:  "Smith",
			Status:    model.UserStatusOperator,
		},
	}
	return dummyUsers, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	dummyUser := &model.User{
		ID:        id,
		FirstName: "John",
		LastName:  "Doe",
		Status:    model.UserStatusAdmin,
	}
	return dummyUser, nil
}

// IngredientsReceived is the resolver for the ingredientsReceived field.
func (r *queryResolver) IngredientsReceived(ctx context.Context) ([]*model.IngredientReceived, error) {
	dummyIngredientsReceived := []*model.IngredientReceived{
		{
			ID:             "1",
			Ingredient:     &model.Ingredient{Name: "Flour"},
			UnitsReceived:  10,
			LotNumber:      "LOT123",
			ExpirationDate: "2025-01-01",
			UnitType:       "Bag",
			QuantityReceived: 250.0,
			UnitOfMeasure:  "kg",
			Supplier:       &model.Supplier{Name: "EggySupplier"},
			User:           &model.User{FirstName: "John", LastName: "Doe"},
			ReceiveDate:    "2024-12-26",
			InvoiceNumber:  "INV123",
		},
	}
	return dummyIngredientsReceived, nil
}

// IngredientReceived is the resolver for the ingredientReceived field.
func (r *queryResolver) IngredientReceived(ctx context.Context, id string) (*model.IngredientReceived, error) {
	dummyIngredientReceived := &model.IngredientReceived{
		ID:             id,
		Ingredient:     &model.Ingredient{Name: "Flour"},
		UnitsReceived:  10,
		LotNumber:      "LOT123",
		ExpirationDate: "2025-01-01",
		UnitType:       "Bag",
		QuantityReceived: 250.0,
		UnitOfMeasure:  "kg",
		Supplier:       &model.Supplier{Name: "EggySupplier"},
		User:           &model.User{FirstName: "John", LastName: "Doe"},
		ReceiveDate:    "2024-12-26",
		InvoiceNumber:  "INV123",
	}
	return dummyIngredientReceived, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}
*/
