package personstore

import (
	"reflect"
	"time"
)

type Object interface {
	GetKind() string
	GetId() string
	GetName() string
	SetID(id string)
	SetName(name string)
}

type Person struct {
	Name      string    `json:"name"`
	ID        string    `json:"id"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birthday"`
}

func (p *Person) GetKind() string {
	return reflect.TypeOf(p).String()
}

func (p *Person) GetId() string {
	return p.ID
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) SetID(id string) {
	p.ID = id
}

func (p *Person) SetName(name string) {
	p.Name = name
}

type Animal struct {
	Name    string `json:"name"`
	ID      string `json:"id"`
	Type    string `json:"type"`
	OwnerID string `json:"owner_id"`
}

func (a *Animal) GetKind() string {
	return reflect.TypeOf(a).String()
}

func (a *Animal) GetId() string {
	return a.ID
}

func (a *Animal) GetName() string {
	return a.Name
}

func (a *Animal) SetID(id string) {
	a.ID = id
}

func (a *Animal) SetName(name string) {
	a.Name = name
}
