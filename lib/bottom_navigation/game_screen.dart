import 'package:flutter/material.dart';

class GameScreen extends StatelessWidget {
  const GameScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('ゲーム'),
        backgroundColor: Colors.transparent,
        elevation: 0,
      ),
      body: GridView.count(
        padding: const EdgeInsets.all(16),
        crossAxisCount: 2,
        mainAxisSpacing: 16,
        crossAxisSpacing: 16,
        children: [
          _buildGameCard('温泉経営', Icons.hot_tub),
          _buildGameCard('バスレース', Icons.directions_boat),
          _buildGameCard('Coming Soon', Icons.lock),
          _buildGameCard('Coming Soon', Icons.lock),
        ],
      ),
    );
  }

  Widget _buildGameCard(String title, IconData icon) {
    return Card(
      child: InkWell(
        onTap: () {
          // ゲーム画面への遷移
        },
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Icon(icon, size: 48),
            const SizedBox(height: 8),
            Text(title),
          ],
        ),
      ),
    );
  }
}