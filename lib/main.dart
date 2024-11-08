import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:fulove/bottom_navigation/bath_screen.dart';
import 'package:fulove/bottom_navigation/game_screen.dart';
import 'package:fulove/bottom_navigation/home_screen.dart';
import 'package:fulove/bottom_navigation/notification_screen.dart';
import 'package:fulove/bottom_navigation/timeline_screen.dart';

void main() {
  runApp(const ProviderScope(child: MyApp()));
}

final indexProvider = StateProvider<int>((ref) => 0);

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Fulove',
      theme: ThemeData(
        primarySwatch: Colors.blue,
        visualDensity: VisualDensity.adaptivePlatformDensity,
      ),
      home: const Root(),
    );
  }
}

class Root extends ConsumerWidget {
  const Root({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final index = ref.watch(indexProvider);

    return Scaffold(
      body: IndexedStack(
        index: index,
        children: const [
          HomeScreen(),
          TimelineScreen(),
          BathScreen(),
          GameScreen(),
          NotificationScreen(),
        ],
      ),
      floatingActionButton: index == 2 ? null : FloatingActionButton(
        onPressed: () {
          ref.read(indexProvider.notifier).state = 2;
        },
        backgroundColor: Colors.blue,
        child: const Icon(Icons.bathtub, size: 32),
      ),
      floatingActionButtonLocation: FloatingActionButtonLocation.centerDocked,
      bottomNavigationBar: CustomBottomNavigationBar(
        currentIndex: index,
        onTap: (index) {
          ref.read(indexProvider.notifier).state = index;
        },
      ),
    );
  }
}

class CustomBottomNavigationBar extends StatelessWidget {
  final int currentIndex;
  final Function(int) onTap;

  const CustomBottomNavigationBar({
    super.key,
    required this.currentIndex,
    required this.onTap,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
        color: Colors.white,
        boxShadow: [
          BoxShadow(
            color: Colors.black.withOpacity(0.1),
            blurRadius: 10,
            offset: const Offset(0, -5),
          ),
        ],
      ),
      child: BottomAppBar(
        notchMargin: 8,
        shape: const CircularNotchedRectangle(),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceAround,
          children: [
            _buildNavItem(Icons.home_outlined, 0),
            _buildNavItem(Icons.timeline_outlined, 1),
            const SizedBox(width: 40),  // 中央のスペース
            _buildNavItem(Icons.sports_esports_outlined, 3),
            _buildNavItem(Icons.notifications_outlined, 4),
          ],
        ),
      ),
    );
  }

  Widget _buildNavItem(IconData icon, int index) {
    final isSelected = currentIndex == index;
    return IconButton(
      icon: Icon(
        icon,
        color: isSelected ? Colors.blue : Colors.grey,
        size: 28,
      ),
      onPressed: () => onTap(index),
    );
  }
}