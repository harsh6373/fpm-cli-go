package boilerplate

import (
	"os"
	"os/exec"
)

func AddGetXBoilerplate() {
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
