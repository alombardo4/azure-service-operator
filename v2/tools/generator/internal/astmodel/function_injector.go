/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// FunctionInjector is a utility for injecting function definitions into resources and objects
type FunctionInjector struct {
	// visitor is used to do the actual injection
	visitor TypeVisitor[any]
}

// NewFunctionInjector creates a new function injector for modifying resources and objects
func NewFunctionInjector() *FunctionInjector {
	result := &FunctionInjector{}

	result.visitor = TypeVisitorBuilder[any]{
		VisitObjectType:   result.injectFunctionIntoObject,
		VisitResourceType: result.injectFunctionIntoResource,
	}.Build()

	return result
}

// Inject modifies the passed type definition by injecting the passed function
func (fi *FunctionInjector) Inject(def TypeDefinition, fns ...Function) (TypeDefinition, error) {
	result := def

	for _, fn := range fns {
		var err error
		result, err = fi.visitor.VisitDefinition(result, fn)
		if err != nil {
			return TypeDefinition{}, err
		}
	}

	return result, nil
}

// injectFunctionIntoObject takes the function provided as a context and includes it on the
// provided object type
func (*FunctionInjector) injectFunctionIntoObject(
	_ *TypeVisitor[any], ot *ObjectType, ctx any,
) (Type, error) {
	fn := ctx.(Function)
	return ot.WithFunction(fn), nil
}

// injectFunctionIntoResource takes the function provided as a context and includes it on the
// provided resource type
func (*FunctionInjector) injectFunctionIntoResource(
	_ *TypeVisitor[any], rt *ResourceType, ctx any,
) (Type, error) {
	fn := ctx.(Function)
	return rt.WithFunction(fn), nil
}
