import 'dart:io';
import 'dart:html' as html;


import 'package:flutter/material.dart';
import 'package:fulove/auth/auth_service.dart';
import 'package:fulove/upload_service/upload_service_web.dart';


class ProfileSettingScreen extends StatefulWidget {
  final int userId;

  const ProfileSettingScreen({super.key, required this.userId});

  @override
  _ProfileSettingScreenState createState() => _ProfileSettingScreenState();
}

class _ProfileSettingScreenState extends State<ProfileSettingScreen> {
  final _formKey = GlobalKey<FormState>();
  final _authService = AuthService();
  final _uploadService = UploadServiceWeb();

  String? _iconUrl;
  String? _previewUrl;
  
  String _username = '';
  final Map<String, TimeOfDay?> _bathTimes = {
    '月曜日': null,
    '火曜日': null,
    '水曜日': null,
    '木曜日': null,
    '金曜日': null,
    '土曜日': null,
    '日曜日': null,
  };
  bool _isLoading = false;

  Future<void> _selectTime(String day) async {
    final TimeOfDay? picked = await showTimePicker(
      context: context,
      initialTime: _bathTimes[day] ?? TimeOfDay.now(),
    );
    if (picked != null) {
      setState(() {
        _bathTimes[day] = picked;
      });
    }
  }

  Future<void> _pickImage() async {
    // input要素の作成
    final input = html.FileUploadInputElement()..accept = 'image/*';
    input.click();

    await input.onChange.first;
    if (input.files?.isEmpty ?? true) return;

    final file = input.files![0];
    // プレビューの作成
    final reader = html.FileReader();
    reader.readAsDataUrl(file);

    await reader.onLoad.first;
    setState(() {
      _previewUrl = reader.result as String;
    });

    try {
      // 画像のアップロード
      final imageUrl = await _uploadService.uploadImage(file);
      setState(() {
        _iconUrl = imageUrl;
      });
      
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('画像のアップロードが完了しました')),
        );
      }
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('画像のアップロードに失敗しました: $e')),
        );
      }
    }
  }

Widget _buildAvatar() {
    return Stack(
      children: [
        CircleAvatar(
          radius: 50,
          backgroundImage: _previewUrl != null 
              ? NetworkImage(_previewUrl!)
              : (_iconUrl != null ? NetworkImage(_iconUrl!) : null),
          child: (_previewUrl == null && _iconUrl == null) 
              ? const Icon(Icons.person, size: 50)
              : null,
        ),
        Positioned(
          bottom: 0,
          right: 0,
          child: Container(
            decoration: BoxDecoration(
              color: Theme.of(context).primaryColor,
              shape: BoxShape.circle,
            ),
            child: IconButton(
              onPressed: _pickImage,
              icon: const Icon(Icons.camera_alt),
              color: Colors.white,
            ),
          ),
        ),
      ],
    );
  }  Future<void> _saveProfile() async {
  if (_formKey.currentState!.validate()) {
    _formKey.currentState!.save();
    setState(() => _isLoading = true);

    try {
      // timeOfDayを文字列に変換する関数
      String timeToString(TimeOfDay? time) {
        if (time == null) return '';
        return '${time.hour.toString().padLeft(2, '0')}:${time.minute.toString().padLeft(2, '0')}';
      }

      // 入浴時間をMap<String, String>に変換
      final Map<String, String> bathTimesStr = _bathTimes.map(
        (key, value) => MapEntry(key.toLowerCase(), timeToString(value))
      );

      await _authService.updateProfile(
        userId: widget.userId,
        username: _username,
        iconUrl: _iconUrl,
        bathTimes: bathTimesStr,
      );
      
      if (mounted) {
        Navigator.pushReplacementNamed(context, '/home');
      }
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('エラー: $e')),
        );
      }
    } finally {
      if (mounted) {
        setState(() => _isLoading = false);
      }
    }
  }
}

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('プロフィール設定'),
        automaticallyImplyLeading: false,
      ),
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(16.0),
        child: Form(
          key: _formKey,
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: [
              
              Center(
                child: Stack(
                  children: [
                    CircleAvatar(
                      radius: 50,
                      backgroundImage: _iconUrl != null ? NetworkImage(_iconUrl!) : null,
                      child: _iconUrl == null ? const Icon(Icons.person, size: 50) : null,
                    ),
                    Positioned(
                      bottom: 0,
                      right: 0,
                      child: IconButton(
                        onPressed: _pickImage,
                        icon: const Icon(Icons.camera_alt),
                        style: IconButton.styleFrom(
                          backgroundColor: Theme.of(context).primaryColor,
                          foregroundColor: Colors.white,
                        ),
                      ),
                    ),
                  ],
                ),
              ),
              const SizedBox(height: 24),
              TextFormField(
                decoration: const InputDecoration(
                  labelText: 'ユーザーネーム',
                  border: OutlineInputBorder(),
                ),
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return 'ユーザーネームを入力してください';
                  }
                  return null;
                },
                onSaved: (value) => _username = value!,
              ),
              const SizedBox(height: 24),
              const Text('希望入浴時間', style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold)),
              const SizedBox(height: 16),
              ..._bathTimes.entries.map((entry) => Padding(
                padding: const EdgeInsets.only(bottom: 8.0),
                child: ListTile(
                  title: Text(entry.key),
                  trailing: Text(
                    entry.value?.format(context) ?? '未設定',
                    style: TextStyle(
                      color: entry.value == null ? Colors.grey : Colors.black,
                    ),
                  ),
                  onTap: () => _selectTime(entry.key),
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(8),
                    side: BorderSide(color: Colors.grey.shade300),
                  ),
                ),
              )),
              const SizedBox(height: 24),
              ElevatedButton(
                onPressed: _isLoading ? null : _saveProfile,
                style: ElevatedButton.styleFrom(
                  padding: const EdgeInsets.symmetric(vertical: 16),
                ),
                child: _isLoading
                    ? const CircularProgressIndicator()
                    : const Text('設定を完了'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}