package boilerplate

import (
	"os"
	"os/exec"
)

func AddBlocBoilerplate() {
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
