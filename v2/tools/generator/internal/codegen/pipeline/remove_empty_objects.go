/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/go-logr/logr"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/embeddedresources"
)

const RemoveEmptyObjectsStageID = "removeEmptyObjects"

// RemoveEmptyObjects removes Definitions which are empty (an object type with no properties)
func RemoveEmptyObjects(log logr.Logger) *Stage {
	return NewStage(
		RemoveEmptyObjectsStageID,
		"Remove empty Objects",
		func(ctx context.Context, state *State) (*State, error) {
			updatedDefs, err := embeddedresources.RemoveEmptyObjects(state.Definitions(), log)
			if err != nil {
				return nil, err
			}

			return state.WithDefinitions(updatedDefs), nil
		})
}
