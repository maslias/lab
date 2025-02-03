
package main

import "context"

type store struct{}

func NewStore() *store {
	return &store{}
}

func (st *store) Create(ctx context.Context) error {
	return nil
}
