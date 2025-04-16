package templates

import (
	"github.com/AlecAivazis/survey/v2"
)

type BuildOptions struct {
	Env     string
	Type    string
	Version string
	Build   string
}

// PromptBuildOptions displays interactive CLI questions and returns build configuration
func PromptBuildOptions() BuildOptions {
	var answers BuildOptions

	survey.Ask([]*survey.Question{
		{
			Name: "Env",
			Prompt: &survey.Select{
				Message: "Select environment:",
				Options: []string{"DEVELOPMENT", "STAGING", "PRODUCTION"},
			},
		},
		{
			Name: "Type",
			Prompt: &survey.Select{
				Message: "Build type:",
				Options: []string{"APK", "App Bundle"},
			},
		},
		{
			Name:     "Version",
			Prompt:   &survey.Input{Message: "Enter version (e.g. 1.0.0):"},
			Validate: survey.Required,
		},
		{
			Name:     "Build",
			Prompt:   &survey.Input{Message: "Enter build number (e.g. 5):"},
			Validate: survey.Required,
		},
	}, &answers)

	return answers
}
