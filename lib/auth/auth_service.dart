import 'dart:convert';
import 'package:http/http.dart' as http;

class AuthService {
  final String baseUrl = 'http://127.0.0.1:8080';

  Future<Map<String, dynamic>> signUp({
    required String email,
    required String password,
  }) async {
    try {
      print('Attempting to sign up with URL: $baseUrl/signup');
      
      final response = await http.post(
        Uri.parse('$baseUrl/signup'),
        headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/json',
        },
        body: jsonEncode({
          'email': email,
          'password': password,
        }),
      );

      print('Response status: ${response.statusCode}');
      print('Response headers: ${response.headers}');
      print('Response body: ${response.body}');

      if (response.statusCode == 201) {
        return jsonDecode(response.body);
      } else {
        throw Exception('サインアップに失敗しました。ステータスコード: ${response.statusCode}, レスポンス: ${response.body}');
      }
    } catch (e) {
      print('Error during sign up: $e');
      rethrow;
    }
  }
  Future<Map<String, dynamic>> signIn({
    required String email,
    required String password,
  }) async {
    final response = await http.post(
      Uri.parse('$baseUrl/signin'),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({
        'email': email,
        'password': password,
      }),
    );

    if (response.statusCode == 200) {
      return jsonDecode(response.body);
    } else {
      throw Exception('Failed to sign in: ${response.body}');
    }
  }

  Future<Map<String, dynamic>> updateProfile({
    required int userId,
    required String username,
    required String? iconUrl,
    required Map<String, String> bathTimes,
  }) async {
    final response = await http.put(
      Uri.parse('$baseUrl/users/$userId/profile'),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({
        'username': username,
        'icon_url': iconUrl ?? '',
        'bath_times': bathTimes,
      }),
    );

    if (response.statusCode == 200) {
      return jsonDecode(response.body);
    } else {
      throw Exception('Failed to update profile: ${response.body}');
    }
  }
}