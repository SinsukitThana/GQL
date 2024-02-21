package datafromgql

import (
	"encoding/json"
	"log"

	"github.com/SinsukitThana/GQL/model/person"
	"github.com/graphql-go/graphql"
	"github.com/uptrace/bun"
)

func DataFromGQL(db *bun.DB, persondata []person.Persons) []byte {

	personType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Persons",
			Fields: graphql.Fields{
				"personid": &graphql.Field{
					Type: graphql.Int,
				},
				"lastname": &graphql.Field{
					Type: graphql.String,
				},
				"firstname": &graphql.Field{
					Type: graphql.String,
				},
				"address": &graphql.Field{
					Type: graphql.String,
				},
				"city": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	fields := graphql.Fields{
		"person": &graphql.Field{
			Type:        graphql.NewList(personType),
			Description: "Get Person",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return persondata, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	query := `
		{
			person{
				personid
			}
		}
	`

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	return rJSON
}
