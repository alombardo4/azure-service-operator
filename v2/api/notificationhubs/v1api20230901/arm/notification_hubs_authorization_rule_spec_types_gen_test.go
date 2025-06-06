// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_NotificationHubsAuthorizationRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NotificationHubsAuthorizationRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNotificationHubsAuthorizationRule_Spec, NotificationHubsAuthorizationRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNotificationHubsAuthorizationRule_Spec runs a test to see if a specific instance of NotificationHubsAuthorizationRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNotificationHubsAuthorizationRule_Spec(subject NotificationHubsAuthorizationRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NotificationHubsAuthorizationRule_Spec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of NotificationHubsAuthorizationRule_Spec instances for property testing - lazily instantiated by
// NotificationHubsAuthorizationRule_SpecGenerator()
var notificationHubsAuthorizationRule_SpecGenerator gopter.Gen

// NotificationHubsAuthorizationRule_SpecGenerator returns a generator of NotificationHubsAuthorizationRule_Spec instances for property testing.
// We first initialize notificationHubsAuthorizationRule_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NotificationHubsAuthorizationRule_SpecGenerator() gopter.Gen {
	if notificationHubsAuthorizationRule_SpecGenerator != nil {
		return notificationHubsAuthorizationRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNotificationHubsAuthorizationRule_Spec(generators)
	notificationHubsAuthorizationRule_SpecGenerator = gen.Struct(reflect.TypeOf(NotificationHubsAuthorizationRule_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNotificationHubsAuthorizationRule_Spec(generators)
	AddRelatedPropertyGeneratorsForNotificationHubsAuthorizationRule_Spec(generators)
	notificationHubsAuthorizationRule_SpecGenerator = gen.Struct(reflect.TypeOf(NotificationHubsAuthorizationRule_Spec{}), generators)

	return notificationHubsAuthorizationRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNotificationHubsAuthorizationRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNotificationHubsAuthorizationRule_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNotificationHubsAuthorizationRule_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNotificationHubsAuthorizationRule_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SharedAccessAuthorizationRulePropertiesGenerator())
}
