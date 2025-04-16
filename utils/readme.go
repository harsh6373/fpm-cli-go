package utils

import (
	"fmt"
	"os"
)

func GenerateReadme(projectName, description string) {
	readme := fmt.Sprintf(`# %s

A Flutter project.

## Table of Contents

- [Introduction](#introduction)
- [Configuration (Environment Configurations)](#environment-configurations)
- [Running the Project](#running-the-project)
- [Generating Codes](#generating-codes)
- [Running Tests and Generating Coverage Report](#running-tests-and-generating-coverage-report)
- [Building the Project](#building-the-project)
- [Additional Information](#additional-information)

## Introduction

`+"`%s`"+` is a Flutter project.

%s

## Environment Configurations

This project contains 3 ENVIRONMENT configurations:

- DEVELOPMENT
- STAGGING
- PRODUCTION

## Running the Project

To run the desired ENVIRONMENT either use the launch configuration in VSCode or use the following commands:

### 1 DEVELOPMENT

`+"```sh"+`
flutter run --dart-define=ENVIRONMENT=DEVELOPMENT --no-enable-impeller
`+"```"+`

### 2 STAGGING

`+"```sh"+`
flutter run --dart-define=ENVIRONMENT=STAGGING --no-enable-impeller
`+"```"+`

### 3 PRODUCTION

`+"```sh"+`
flutter run --dart-define=ENVIRONMENT=PRODUCTION --no-enable-impeller
`+"```"+`

## Generating Codes

`+"```sh"+`
flutter pub run build_runner build --delete-conflicting-outputs
`+"```"+`

## Running Tests and Generating Coverage Report

`+"```sh"+`
flutter test --coverage
brew install lcov
genhtml coverage/lcov.info -o coverage/html
open coverage/html/index.html
`+"```"+`

## Building the Project

### Appbundle

`+"```sh"+`
flutter build appbundle --no-tree-shake-icons --dart-define=ENVIRONMENT=DEVELOPMENT
flutter build appbundle --no-tree-shake-icons --dart-define=ENVIRONMENT=STAGGING
flutter build appbundle --no-tree-shake-icons --dart-define=ENVIRONMENT=PRODUCTION
`+"```"+`

### APK

`+"```sh"+`
flutter build apk --release --no-tree-shake-icons --dart-define=ENVIRONMENT=DEVELOPMENT
flutter build apk --release --no-tree-shake-icons --dart-define=ENVIRONMENT=STAGGING
flutter build apk --release --no-tree-shake-icons --dart-define=ENVIRONMENT=PRODUCTION
`+"```"+`

## Additional Information

To run this app with different environments via VSCode, create a launch.json under .vscode/
`, projectName, projectName, description)

	if err := os.WriteFile("README.md", []byte(readme), 0644); err != nil {
		fmt.Println("⚠️ Failed to write README.md:", err)
	} else {
		fmt.Println("✅ README.md added to the project.")
	}
}
