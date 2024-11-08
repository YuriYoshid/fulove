import 'package:flutter/material.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leading: Builder(
          builder: (context) => IconButton(
            icon: const Icon(Icons.more_vert),  // 3点リーダーアイコン
            onPressed: () => Scaffold.of(context).openDrawer(),  // ドロワーを開く
          ),
        ),
        title: const Text('Fulove'),
        backgroundColor: Colors.transparent,
        elevation: 0,
      ),
      drawer: Drawer(  // ドロワーをここで定義
        child: ListView(
          children: [
            const DrawerHeader(
              decoration: BoxDecoration(color: Colors.blue),
              child: Text('設定'),
            ),
            ListTile(
              leading: const Icon(Icons.settings),
              title: const Text('設定'),
              onTap: () {
                // 設定画面へ遷移
                Navigator.pop(context); // ドロワーを閉じる
                // TODO: 設定画面への遷移を実装
              },
            ),
            ListTile(
              leading: const Icon(Icons.person_add),
              title: const Text('フレンド検索'),
              onTap: () {
                // フレンド検索画面へ遷移
                Navigator.pop(context); // ドロワーを閉じる
                // TODO: フレンド検索画面への遷移を実装
              },
            ),
          ],
        ),
      ),
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // ステータスカード
            Card(
              child: Padding(
                padding: const EdgeInsets.all(16),
                child: Column(
                  children: [
                    Text(
                      '現在のポイント: 100 pt',
                      style: Theme.of(context).textTheme.headlineSmall,
                    ),
                    const SizedBox(height: 8),
                    const Text('最終入浴: 今日 19:00'),
                    const Text('連続入浴: 3日'),
                  ],
                ),
              ),
            ),
            const SizedBox(height: 24),
            
            // 入浴記録ボタン（大きめのボタン）
            SizedBox(
              width: double.infinity,
              child: ElevatedButton.icon(
                icon: const Icon(Icons.bathtub),
                label: const Text('入浴を記録する'),
                style: ElevatedButton.styleFrom(
                  padding: const EdgeInsets.all(16),
                ),
                onPressed: () {
                  // TODO: 入浴記録画面への遷移を実装
                },
              ),
            ),
            
            const SizedBox(height: 24),
            const Text(
              '今日の入浴状況',
              style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
            ),
            const SizedBox(height: 8),
            
            // フレンドリスト
            ListView.builder(
              shrinkWrap: true,
              physics: const NeverScrollableScrollPhysics(),
              itemCount: 3,
              itemBuilder: (context, index) {
                return ListTile(
                  leading: const CircleAvatar(
                    child: Icon(Icons.person),
                  ),
                  title: Text('フレンド ${index + 1}'),
                  subtitle: Text(index % 2 == 0 ? '入浴済み' : '未入浴'),
                  trailing: Text(index % 2 == 0 ? '19:30' : '予定: 21:00'),
                );
              },
            ),
          ],
        ),
      ),
    );
  }
}