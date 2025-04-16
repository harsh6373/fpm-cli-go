package boilerplate

import (
	"os"
	"os/exec"
)

func AddRiverpodBoilerplate() {
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
