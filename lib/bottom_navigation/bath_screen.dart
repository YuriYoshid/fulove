import 'package:flutter/material.dart';

class BathScreen extends StatelessWidget {
  const BathScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('入浴記録'),
        backgroundColor: Colors.transparent,
        elevation: 0,
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Icon(
              Icons.bathtub,
              size: 100,
              color: Colors.blue,
            ),
            const SizedBox(height: 24),
            ElevatedButton(
              onPressed: () {
                // 入浴記録の処理
              },
              style: ElevatedButton.styleFrom(
                padding: const EdgeInsets.symmetric(
                  horizontal: 48,
                  vertical: 16,
                ),
              ),
              child: const Text('入浴を記録する'),
            ),
          ],
        ),
      ),
    );
  }
}