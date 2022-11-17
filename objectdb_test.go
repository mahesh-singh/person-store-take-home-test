package personstore

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

func TestStore(t *testing.T) {
	// Create a instance of redis
	// Create a new person object and store
	// Create a new amimal object and store
	// Veirfy that both person and object are available in the redis
	// Clean up

	objDB, err := NewRedisDB(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})
	if err != nil {
		t.Errorf("Error while connecting Redis %v", err)
	}
	p := &Person{Name: "John", LastName: "Sharma", BirthDate: time.Now()}
	objDB.Store(context.Background(), p)
	a := &Animal{Name: "Jony", Type: "Dog", OwnerID: p.ID}
	objDB.Store(context.Background(), a)

	outputPerson, err := objDB.GetObjectByID(context.Background(), p.GetId())

	if err != nil {
		t.Errorf("Expected: Person Name, output: %v", err)
	} else if outputPerson.GetName() != p.Name || outputPerson.GetKind() != p.GetId() {
		t.Errorf("Expected: %v, Output %v", p, outputPerson)
	}
}
