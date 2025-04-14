package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Flutter project with state management boilerplate",
	Run: func(cmd *cobra.Command, args []string) {
		var answers struct {
			ProjectName  string
			PackageName  string
			Description  string
			StateManager string
			ProjectPath  string
		}

		qs := []*survey.Question{
			{
				Name:     "ProjectName",
				Prompt:   &survey.Input{Message: "Enter project name:"},
				Validate: survey.Required,
			},
			{
				Name:     "PackageName",
				Prompt:   &survey.Input{Message: "Enter package name (e.g. com.example.app):"},
				Validate: survey.Required,
			},
			{
				Name:   "Description",
				Prompt: &survey.Input{Message: "Project description:"},
			},
			{
				Name: "ProjectPath",
				Prompt: &survey.Input{
					Message: "Enter the full path where the project should be created:",
					Default: "./",
				},
				Validate: survey.Required,
			},
			{
				Name: "StateManager",
				Prompt: &survey.Select{
					Message: "Select a state management solution:",
					Options: []string{"GetX", "BLoC", "Provider", "Riverpod"},
				},
			},
		}

		if err := survey.Ask(qs, &answers); err != nil {
			fmt.Println("Prompt failed:", err)
			return
		}

		projectName := answers.ProjectName
		packageName := answers.PackageName
		description := answers.Description
		stateManager := answers.StateManager
		projectBase := answers.ProjectPath

		// Ensure base path exists
		if err := os.MkdirAll(projectBase, os.ModePerm); err != nil {
			fmt.Println("❌ Failed to create base path:", err)
			return
		}

		// Check if path is writable
		testFile := filepath.Join(projectBase, ".fpm_write_test")
		if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
			fmt.Println("❌ Cannot write to the specified path:", projectBase)
			fmt.Println("💡 Tip: Try a different directory like $HOME/dev or ./projects")
			return
		} else {
			os.Remove(testFile) // cleanup
		}

		// Move to base path
		if err := os.Chdir(projectBase); err != nil {
			fmt.Println("❌ Cannot switch to base project directory:", err)
			return
		}

		// Flutter Create
		fmt.Println("🚀 Creating Flutter project...")
		createFlutterCmd := exec.Command("flutter", "create",
			"--org", packageName,
			"--project-name", projectName,
			"--description", description,
			projectName,
		)
		createFlutterCmd.Stdout = os.Stdout
		createFlutterCmd.Stderr = os.Stderr
		if err := createFlutterCmd.Run(); err != nil {
			fmt.Println("❌ Failed to create Flutter project:", err)
			return
		}

		// Change to project dir
		if err := os.Chdir(projectName); err != nil {
			fmt.Println("❌ Failed to change into new project directory:", err)
			return
		}

		fmt.Println("📦 Adding state management:", stateManager)

		switch stateManager {
		case "GetX":
			addGetXBoilerplate()
		case "BLoC":
			addBlocBoilerplate()
		case "Provider":
			addProviderBoilerplate()
		case "Riverpod":
			addRiverpodBoilerplate()
		default:
			fmt.Println("Invalid selection.")
		}

		fmt.Println("✅ Project setup complete.")
		generateReadme(projectName, description)

		// Git integration
		var initGit bool
		survey.AskOne(&survey.Confirm{
			Message: "Would you like to initialize this project with Git?",
			Default: true,
		}, &initGit)

		if initGit {
			fmt.Println("🔧 Setting up Git...")
			gitSteps := [][]string{
				{"git", "init"},
				{"git", "add", "."},
				{"git", "commit", "-m", "initial commit"},
				{"git", "branch", "-M", "main"},
			}

			for _, step := range gitSteps {
				cmd := exec.Command(step[0], step[1:]...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					fmt.Printf("❌ Git step failed: %s %v\n", step[0], step[1:])
					return
				}
			}

			var remoteURL string
			survey.AskOne(&survey.Input{
				Message: "Enter remote repository URL (leave blank to skip):",
			}, &remoteURL)

			if remoteURL != "" {
				remoteCmds := [][]string{
					{"git", "remote", "add", "origin", remoteURL},
					{"git", "push", "-u", "origin", "main"},
				}
				for _, cmdArgs := range remoteCmds {
					cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					if err := cmd.Run(); err != nil {
						fmt.Printf("❌ Failed to run: %s %v\n", cmdArgs[0], cmdArgs[1:])
						break
					}
				}
				fmt.Println("✅ Git remote set and pushed.")
			} else {
				fmt.Println("ℹ️ Skipping remote push.")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func addGetXBoilerplate() {
	exec.Command("flutter", "pub", "add", "get").Run()

	main := `import 'package:flutter/material.dart';
import 'package:get/get.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return GetMaterialApp(home: Home());
  }
}

class Home extends StatelessWidget {
  final Controller c = Get.put(Controller());

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("GetX Example")),
      body: Center(child: Obx(() => Text("Clicks: ${c.count}"))),
      floatingActionButton: FloatingActionButton(
        onPressed: c.increment,
        child: Icon(Icons.add),
      ),
    );
  }
}

class Controller extends GetxController {
  var count = 0.obs;
  void increment() => count++;
}`
	os.WriteFile("lib/main.dart", []byte(main), 0644)
}

func addBlocBoilerplate() {
	exec.Command("flutter", "pub", "add", "flutter_bloc").Run()
	exec.Command("flutter", "pub", "add", "equatable").Run()

	os.MkdirAll("lib/bloc", os.ModePerm)

	main := `import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'bloc/counter_bloc.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: BlocProvider(
        create: (_) => CounterBloc(),
        child: Home(),
      ),
    );
  }
}

class Home extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final counterBloc = BlocProvider.of<CounterBloc>(context);

    return Scaffold(
      appBar: AppBar(title: Text("BLoC Example")),
      body: Center(child: BlocBuilder<CounterBloc, int>(
        builder: (context, count) => Text("Clicks: $count"),
      )),
      floatingActionButton: FloatingActionButton(
        onPressed: () => counterBloc.add(CounterIncrementPressed()),
        child: Icon(Icons.add),
      ),
    );
  }
}`

	blocCode := `import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';

class CounterBloc extends Bloc<CounterEvent, int> {
  CounterBloc() : super(0) {
    on<CounterIncrementPressed>((event, emit) => emit(state + 1));
  }
}

abstract class CounterEvent extends Equatable {
  @override
  List<Object> get props => [];
}

class CounterIncrementPressed extends CounterEvent {}`
	os.WriteFile("lib/main.dart", []byte(main), 0644)
	os.WriteFile("lib/bloc/counter_bloc.dart", []byte(blocCode), 0644)
}

func addProviderBoilerplate() {
	exec.Command("flutter", "pub", "add", "provider").Run()

	main := `import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

void main() {
  runApp(ChangeNotifierProvider(
    create: (_) => Counter(),
    child: MyApp(),
  ));
}

class Counter with ChangeNotifier {
  int _count = 0;
  int get count => _count;
  void increment() {
    _count++;
    notifyListeners();
  }
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final counter = Provider.of<Counter>(context);
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(title: Text("Provider Example")),
        body: Center(child: Text("Clicks: ${counter.count}")),
        floatingActionButton: FloatingActionButton(
          onPressed: counter.increment,
          child: Icon(Icons.add),
        ),
      ),
    );
  }
}`
	os.WriteFile("lib/main.dart", []byte(main), 0644)
}

func addRiverpodBoilerplate() {
	exec.Command("flutter", "pub", "add", "flutter_riverpod").Run()

	main := `import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

final counterProvider = StateProvider<int>((ref) => 0);

void main() {
  runApp(ProviderScope(child: MyApp()));
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(home: Home());
  }
}

class Home extends ConsumerWidget {
  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final count = ref.watch(counterProvider);
    return Scaffold(
      appBar: AppBar(title: Text("Riverpod Example")),
      body: Center(child: Text("Clicks: $count")),
      floatingActionButton: FloatingActionButton(
        onPressed: () => ref.read(counterProvider.notifier).state++,
        child: Icon(Icons.add),
      ),
    );
  }
}`
	os.WriteFile("lib/main.dart", []byte(main), 0644)
}

func generateReadme(projectName, description string) {
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
