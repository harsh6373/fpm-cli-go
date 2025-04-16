package boilerplate

import (
	"os"
	"os/exec"
)

func AddProviderBoilerplate() {
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
