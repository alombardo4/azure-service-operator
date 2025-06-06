/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"path/filepath"
)

// SupportedResourcesReport is configuration for the report that lists all the supported resources.
type SupportedResourcesReport struct {
	cfg *Configuration // Back reference to global configuration
	// OutputFolder is the destination folder for the report, relative to DestinationGoModuleFile
	OutputFolder string `yaml:"outputFolder,omitempty"`
	// FragmentPath is a folder path for markdown fragments to inject into the file
	FragmentPath string `yaml:"fragmentPath,omitempty"`
	// ResourceURLTemplate is a template for URL to the API docs for a resource
	// It may use the placeholders {group} {version} and {kind}
	ResourceURLTemplate string `yaml:"resourceUrlTemplate"`
	// ResourcePathTemplate is a template used for generating a file path for checking whether docs for a resource have been generated
	// specified relative to the directory of outputPath
	ResourcePathTemplate string `yaml:"resourcePathTemplate"`
	// CurrentRelease identifies the current release of ASO, allowing newer resources to be classified as Next Release
	CurrentRelease string `yaml:"currentRelease"`
}

// NewSupportedResourcesReport creates a new SupportedResourcesReport.
func NewSupportedResourcesReport(cfg *Configuration) *SupportedResourcesReport {
	return &SupportedResourcesReport{
		cfg: cfg,
	}
}

// FullOutputPath returns the fully qualified path to the output file.
func (srr *SupportedResourcesReport) FullOutputPath() string {
	return filepath.Join(
		filepath.Dir(srr.cfg.DestinationGoModuleFile),
		srr.OutputFolder,
		"_index.md")
}

// FullOutputPath returns the fully qualified path to the output file for a given group
func (srr *SupportedResourcesReport) GroupFullOutputPath(group string) string {
	return filepath.Join(
		filepath.Dir(srr.cfg.DestinationGoModuleFile),
		srr.OutputFolder,
		group,
		"_index.md")
}

// FullFragmentFolderPath returns the fully qualified path to our fragment folder
func (srr *SupportedResourcesReport) FullFragmentPath() string {
	return filepath.Join(
		filepath.Dir(srr.cfg.DestinationGoModuleFile),
		srr.FragmentPath)
}
