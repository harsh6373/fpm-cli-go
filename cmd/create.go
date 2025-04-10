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
		}

		// Questions
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

		// Flutter Create
		fmt.Println("🚀 Creating Flutter project...")
		createCmd := exec.Command("flutter", "create",
			"--org", packageName,
			"--project-name", projectName,
			"--description", description,
			projectName,
		)
		createCmd.Stdout = os.Stdout
		createCmd.Stderr = os.Stderr
		if err := createCmd.Run(); err != nil {
			fmt.Println("❌ Failed to create Flutter project:", err)
			return
		}

		projectPath := filepath.Join(".", projectName)
		if err := os.Chdir(projectPath); err != nil {
			fmt.Println("❌ Cannot switch to project directory:", err)
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
