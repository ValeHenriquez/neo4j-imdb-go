package database

import (
	"fmt"
	"reflect"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func executeQuery(query string, params map[string]interface{}) (*neo4j.EagerResult, error) {
	result, err := neo4j.ExecuteQuery(
		ctx,
		driver,
		query,
		params,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))

	if err != nil {
		fmt.Println("ERROR EN EXECUTE QUERY", err)
		return nil, err
	}

	return result, nil
}

func Create(obj interface{}) error {
	objValue := reflect.ValueOf(obj)

	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}

	objType := objValue.Type()

	query := fmt.Sprintf("MERGE (n:%s {Id: $Id}) SET", objType.Name())

	params := make(map[string]interface{})

	for i := 0; i < objValue.NumField(); i++ {
		field := objType.Field(i)
		fieldName := field.Name
		fieldValue := objValue.Field(i).Interface()

		comma := ","
		if i == objType.NumField()-1 {
			comma = ""
		}

		query += fmt.Sprintf(" n.%s = $%s%s", fieldName, fieldName, comma)
		params[fieldName] = fieldValue
	}

	_, err := executeQuery(query, params)

	if err != nil {
		fmt.Println("ERROR EN EXECUTE CREATE QUERY", err)
		return err
	}

	fmt.Println("Created node in database")
	//fmt.Printf("Created %v nodes in %+v.\n", result.Summary.Counters().NodesCreated(), result.Summary.ResultAvailableAfter())
	return err
}

func GetAll(obj interface{}) ([]map[string]interface{}, error) {
	objType := reflect.TypeOf(obj)

	query := fmt.Sprintf("MATCH (n:%s) RETURN n", objType.Name())

	result, err := neo4j.ExecuteQuery(
		ctx,
		driver,
		query,
		nil,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))

	if err != nil {
		fmt.Println("ERROR EN EXECUTE GETALL QUERY", err)
		return nil, err
	}

	var props []map[string]interface{}
	for _, record := range result.Records {
		node, _ := record.Get("n")
		props = append(props, node.(neo4j.Node).Props)
	}

	return props, nil
}

func GetOne(obj interface{}) (map[string]interface{}, error) {
	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)

	query := fmt.Sprintf("MATCH (n:%s {Id: $Id}) RETURN n", objType.Name())

	params := map[string]interface{}{
		"Id": objValue.FieldByName("Id").Interface(),
	}

	result, err := executeQuery(query, params)

	if err != nil {
		fmt.Println("ERROR EN EXECUTE GET QUERY", err)
		return nil, err
	}

	var props map[string]interface{}
	for _, record := range result.Records {
		node, _ := record.Get("n")
		props = node.(neo4j.Node).Props
	}

	if len(props) == 0 {
		return nil, nil //No se encontro el nodo pero verifico la existencia
	}

	return props, nil
}

func GetRelationships(obj interface{}, relation_type string) ([]map[string]interface{}, error) {
	objValue := reflect.ValueOf(obj)

	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}

	objType := objValue.Type()

	query := fmt.Sprintf("MATCH (n:%s {Id: $Id})-[r:%s]->(m) RETURN m", objType.Name(), relation_type)

	params := map[string]interface{}{
		"Id": objValue.FieldByName("Id").Interface(),
	}

	result, err := executeQuery(query, params)

	if err != nil {
		fmt.Println("ERROR EN EXECUTE GETRELATIONSHIPS QUERY", err)
		return nil, err
	}

	var props []map[string]interface{}
	for _, record := range result.Records {
		node, _ := record.Get("m")
		props = append(props, node.(neo4j.Node).Props)
	}

	if len(props) == 0 {
		return nil, nil //No se encontro el nodo pero verifico la existencia
	}

	return props, nil
}

func CreateRelation(obj1 interface{}, obj2 interface{}, relation string) error {
	obj1Value := reflect.ValueOf(obj1)
	if obj1Value.Kind() == reflect.Ptr {
		obj1Value = obj1Value.Elem() // Obtiene el valor apuntado si es un puntero
	}
	obj1Type := obj1Value.Type()

	obj2Value := reflect.ValueOf(obj2)
	if obj2Value.Kind() == reflect.Ptr {
		obj2Value = obj2Value.Elem()
	}
	obj2Type := obj2Value.Type()

	query := fmt.Sprintf(`
		MATCH (n:%s {Id: $Id1})
		MATCH (m:%s {Id: $Id2})
		MERGE (n)-[:%s]->(m)
	`, obj1Type.Name(), obj2Type.Name(), relation)

	params := map[string]interface{}{
		"Id1": obj1Value.FieldByName("Id").Interface(),
		"Id2": obj2Value.FieldByName("Id").Interface(),
	}

	_, err := executeQuery(query, params)

	if err != nil {
		fmt.Println("ERROR EN EXECUTE CREATE RELATION QUERY", err)
		return err
	}

	//fmt.Printf("Created %v nodes in %+v.\n", result.Summary.Counters().NodesCreated(), result.Summary.ResultAvailableAfter())
	return nil
}

func CreateBiRelation(obj1 interface{}, obj2 interface{}, relation string) error { //BIDIRECTIONAL
	obj1Value := reflect.ValueOf(obj1)
	if obj1Value.Kind() == reflect.Ptr {
		obj1Value = obj1Value.Elem()
	}
	obj1Type := obj1Value.Type()

	obj2Value := reflect.ValueOf(obj2)
	if obj2Value.Kind() == reflect.Ptr {
		obj2Value = obj2Value.Elem()
	}
	obj2Type := obj2Value.Type()

	query := fmt.Sprintf(`
		MATCH (n:%s {Id: $Id1})
		MATCH (m:%s {Id: $Id2})
		MERGE (n)-[:%s]->(m)
		MERGE (m)-[:%s]->(n)
	`, obj1Type.Name(), obj2Type.Name(), relation, relation)

	params := map[string]interface{}{
		"Id1": obj1Value.FieldByName("Id").Interface(),
		"Id2": obj2Value.FieldByName("Id").Interface(),
	}

	_, err := executeQuery(query, params)

	if err != nil {
		fmt.Println("ERROR EN EXECUTE CREATE RELATION QUERY", err)
		return err
	}

	//fmt.Printf("Created %v nodes in %+v.\n", result.Summary.Counters().NodesCreated(), result.Summary.ResultAvailableAfter())
	return nil
}

func DeleteOne(obj interface{}) error {
	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)

	query := fmt.Sprintf("MATCH (n:%s {Id: $Id}) DETACH DELETE n", objType.Name())

	params := map[string]interface{}{
		"Id": objValue.FieldByName("Id").Interface(),
	}

	result, err := executeQuery(query, params)

	if err != nil {
		fmt.Println("ERROR EN EXECUTE DELETE QUERY", err)
		return err
	}

	nodesDeleted := result.Summary.Counters().NodesDeleted()
	if nodesDeleted == 0 {
		return fmt.Errorf("no se encontro el nodo")
	}
	return nil
}

func DeleteAll() error {

	query := "MATCH (n) DETACH DELETE n"

	result, err := executeQuery(query, nil)

	if err != nil {
		fmt.Println("ERROR EN EXECUTE DELETEALL QUERY", err)
		return err
	}
	nodesDeleted := result.Summary.Counters().NodesDeleted()
	fmt.Printf("Deleted %v nodes in %+v.\n", nodesDeleted, result.Summary.ResultAvailableAfter())

	fmt.Print("Deleted all nodes")
	return nil
}
