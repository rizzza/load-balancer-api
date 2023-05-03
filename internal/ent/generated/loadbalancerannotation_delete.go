// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/loadbalancerannotation"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/predicate"
)

// LoadBalancerAnnotationDelete is the builder for deleting a LoadBalancerAnnotation entity.
type LoadBalancerAnnotationDelete struct {
	config
	hooks    []Hook
	mutation *LoadBalancerAnnotationMutation
}

// Where appends a list predicates to the LoadBalancerAnnotationDelete builder.
func (lbad *LoadBalancerAnnotationDelete) Where(ps ...predicate.LoadBalancerAnnotation) *LoadBalancerAnnotationDelete {
	lbad.mutation.Where(ps...)
	return lbad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lbad *LoadBalancerAnnotationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, LoadBalancerAnnotationMutation](ctx, lbad.sqlExec, lbad.mutation, lbad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (lbad *LoadBalancerAnnotationDelete) ExecX(ctx context.Context) int {
	n, err := lbad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lbad *LoadBalancerAnnotationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(loadbalancerannotation.Table, sqlgraph.NewFieldSpec(loadbalancerannotation.FieldID, field.TypeString))
	if ps := lbad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, lbad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	lbad.mutation.done = true
	return affected, err
}

// LoadBalancerAnnotationDeleteOne is the builder for deleting a single LoadBalancerAnnotation entity.
type LoadBalancerAnnotationDeleteOne struct {
	lbad *LoadBalancerAnnotationDelete
}

// Where appends a list predicates to the LoadBalancerAnnotationDelete builder.
func (lbado *LoadBalancerAnnotationDeleteOne) Where(ps ...predicate.LoadBalancerAnnotation) *LoadBalancerAnnotationDeleteOne {
	lbado.lbad.mutation.Where(ps...)
	return lbado
}

// Exec executes the deletion query.
func (lbado *LoadBalancerAnnotationDeleteOne) Exec(ctx context.Context) error {
	n, err := lbado.lbad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{loadbalancerannotation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lbado *LoadBalancerAnnotationDeleteOne) ExecX(ctx context.Context) {
	if err := lbado.Exec(ctx); err != nil {
		panic(err)
	}
}
