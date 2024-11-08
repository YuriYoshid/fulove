import 'package:flutter/material.dart';

class NotificationScreen extends StatelessWidget {
  const NotificationScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('通知'),
        backgroundColor: Colors.transparent,
        elevation: 0,
      ),
      body: ListView.builder(
        itemCount: 5, // 仮のデータ数
        itemBuilder: (context, index) {
          return ListTile(
            leading: const CircleAvatar(
              backgroundColor: Colors.blue,
              child: Icon(Icons.notifications, color: Colors.white),
            ),
            title: Text('通知 ${index + 1}'),
            subtitle: Text(
              index % 2 == 0 
                ? 'フレンドが入浴を完了しました！'
                : '入浴時間になりました！',
            ),
            trailing: const Text('30分前'),
          );
        },
      ),
    );
  }
}