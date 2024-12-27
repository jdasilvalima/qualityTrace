package graphql

import (
    "github.com/graphql-go/graphql"
    "backend/internal/data"
)

func getResolvers() *graphql.Schema {
    ingredientType := graphql.NewObject(graphql.ObjectConfig{
        Name: "Ingredient",
        Fields: graphql.Fields{
            "id":              &graphql.Field{Type: graphql.Int},
            "ingredientNumber": &graphql.Field{Type: graphql.String},
            "name":            &graphql.Field{Type: graphql.String},
            "type":            &graphql.Field{Type: graphql.String},
            "unitType":        &graphql.Field{Type: graphql.String},
            "finalUnit":       &graphql.Field{Type: graphql.String},
        },
    })

    queryType := graphql.NewObject(graphql.ObjectConfig{
        Name: "Query",
        Fields: graphql.Fields{
            "ingredients": &graphql.Field{
                Type: graphql.NewList(ingredientType),
                Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                    return data.Ingredients, nil
                },
            },
        },
    })

    mutationType := graphql.NewObject(graphql.ObjectConfig{
        Name: "Mutation",
        Fields: graphql.Fields{
            "addIngredient": &graphql.Field{
                Type: ingredientType,
                Args: graphql.FieldConfigArgument{
                    "name":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
                    "type":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
                    "unitType":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
                    "finalUnit": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
                },
                Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                    newIngredient := data.Ingredient{
                        ID:              len(data.Ingredients) + 1,
                        IngredientNumber: "NEW00" + string(len(data.Ingredients)+1),
                        Name:            p.Args["name"].(string),
                        Type:            p.Args["type"].(string),
                        UnitType:        p.Args["unitType"].(string),
                        FinalUnit:       p.Args["finalUnit"].(string),
                    }
                    data.Ingredients = append(data.Ingredients, newIngredient)
                    return newIngredient, nil
                },
            },
        },
    })

    schema, err := graphql.NewSchema(graphql.SchemaConfig{
        Query:    queryType,
        Mutation: mutationType,
    })
    if err != nil {
        panic(err)
    }
    return &schema
}
