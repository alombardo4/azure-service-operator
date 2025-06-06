// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

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

func Test_FlexibleServersAdministrator_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersAdministrator via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersAdministrator, FlexibleServersAdministratorGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersAdministrator runs a test to see if a specific instance of FlexibleServersAdministrator round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersAdministrator(subject FlexibleServersAdministrator) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersAdministrator
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

// Generator of FlexibleServersAdministrator instances for property testing - lazily instantiated by
// FlexibleServersAdministratorGenerator()
var flexibleServersAdministratorGenerator gopter.Gen

// FlexibleServersAdministratorGenerator returns a generator of FlexibleServersAdministrator instances for property testing.
func FlexibleServersAdministratorGenerator() gopter.Gen {
	if flexibleServersAdministratorGenerator != nil {
		return flexibleServersAdministratorGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFlexibleServersAdministrator(generators)
	flexibleServersAdministratorGenerator = gen.Struct(reflect.TypeOf(FlexibleServersAdministrator{}), generators)

	return flexibleServersAdministratorGenerator
}

// AddRelatedPropertyGeneratorsForFlexibleServersAdministrator is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersAdministrator(gens map[string]gopter.Gen) {
	gens["Spec"] = FlexibleServersAdministrator_SpecGenerator()
	gens["Status"] = FlexibleServersAdministrator_STATUSGenerator()
}

func Test_FlexibleServersAdministratorOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersAdministratorOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersAdministratorOperatorSpec, FlexibleServersAdministratorOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersAdministratorOperatorSpec runs a test to see if a specific instance of FlexibleServersAdministratorOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersAdministratorOperatorSpec(subject FlexibleServersAdministratorOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersAdministratorOperatorSpec
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

// Generator of FlexibleServersAdministratorOperatorSpec instances for property testing - lazily instantiated by
// FlexibleServersAdministratorOperatorSpecGenerator()
var flexibleServersAdministratorOperatorSpecGenerator gopter.Gen

// FlexibleServersAdministratorOperatorSpecGenerator returns a generator of FlexibleServersAdministratorOperatorSpec instances for property testing.
func FlexibleServersAdministratorOperatorSpecGenerator() gopter.Gen {
	if flexibleServersAdministratorOperatorSpecGenerator != nil {
		return flexibleServersAdministratorOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	flexibleServersAdministratorOperatorSpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServersAdministratorOperatorSpec{}), generators)

	return flexibleServersAdministratorOperatorSpecGenerator
}

func Test_FlexibleServersAdministrator_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersAdministrator_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersAdministrator_STATUS, FlexibleServersAdministrator_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersAdministrator_STATUS runs a test to see if a specific instance of FlexibleServersAdministrator_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersAdministrator_STATUS(subject FlexibleServersAdministrator_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersAdministrator_STATUS
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

// Generator of FlexibleServersAdministrator_STATUS instances for property testing - lazily instantiated by
// FlexibleServersAdministrator_STATUSGenerator()
var flexibleServersAdministrator_STATUSGenerator gopter.Gen

// FlexibleServersAdministrator_STATUSGenerator returns a generator of FlexibleServersAdministrator_STATUS instances for property testing.
// We first initialize flexibleServersAdministrator_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServersAdministrator_STATUSGenerator() gopter.Gen {
	if flexibleServersAdministrator_STATUSGenerator != nil {
		return flexibleServersAdministrator_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersAdministrator_STATUS(generators)
	flexibleServersAdministrator_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServersAdministrator_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersAdministrator_STATUS(generators)
	AddRelatedPropertyGeneratorsForFlexibleServersAdministrator_STATUS(generators)
	flexibleServersAdministrator_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServersAdministrator_STATUS{}), generators)

	return flexibleServersAdministrator_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersAdministrator_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersAdministrator_STATUS(gens map[string]gopter.Gen) {
	gens["AdministratorType"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IdentityResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Login"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Sid"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServersAdministrator_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersAdministrator_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_FlexibleServersAdministrator_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersAdministrator_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersAdministrator_Spec, FlexibleServersAdministrator_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersAdministrator_Spec runs a test to see if a specific instance of FlexibleServersAdministrator_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersAdministrator_Spec(subject FlexibleServersAdministrator_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersAdministrator_Spec
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

// Generator of FlexibleServersAdministrator_Spec instances for property testing - lazily instantiated by
// FlexibleServersAdministrator_SpecGenerator()
var flexibleServersAdministrator_SpecGenerator gopter.Gen

// FlexibleServersAdministrator_SpecGenerator returns a generator of FlexibleServersAdministrator_Spec instances for property testing.
// We first initialize flexibleServersAdministrator_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServersAdministrator_SpecGenerator() gopter.Gen {
	if flexibleServersAdministrator_SpecGenerator != nil {
		return flexibleServersAdministrator_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersAdministrator_Spec(generators)
	flexibleServersAdministrator_SpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServersAdministrator_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersAdministrator_Spec(generators)
	AddRelatedPropertyGeneratorsForFlexibleServersAdministrator_Spec(generators)
	flexibleServersAdministrator_SpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServersAdministrator_Spec{}), generators)

	return flexibleServersAdministrator_SpecGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersAdministrator_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersAdministrator_Spec(gens map[string]gopter.Gen) {
	gens["AdministratorType"] = gen.PtrOf(gen.AlphaString())
	gens["Login"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Sid"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServersAdministrator_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersAdministrator_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(FlexibleServersAdministratorOperatorSpecGenerator())
}
