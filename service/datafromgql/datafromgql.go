package datafromgql

import (
	"encoding/json"
	"log"

	"github.com/SinsukitThana/GQL/model/person"
	"github.com/SinsukitThana/GQL/model/workflow"
	"github.com/graphql-go/graphql"
	"github.com/uptrace/bun"
)

func DataFromGQL(db *bun.DB, persondata []person.Persons, getquery string) []byte {

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

	query := getquery

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	return rJSON
}

func DataWorkFlowFromGQL(db *bun.DB, workflowdata []workflow.Workflow, getquery string) []byte {

	workflowType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "workflows",
			Fields: graphql.Fields{
				"WorkflowID": &graphql.Field{
					Type: graphql.String,
				},
				"WorkflowName": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	fields := graphql.Fields{
		"workflow": &graphql.Field{
			Type:        graphql.NewList(workflowType),
			Description: "Get workflow",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return workflowdata, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	query := getquery

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	return rJSON
}
