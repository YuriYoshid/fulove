import 'dart:async';
import 'dart:convert';
import 'dart:html' as html;
import 'package:http/http.dart' as http;
import './upload_service_interface.dart';

class UploadServiceWeb implements UploadService {
  final String baseUrl = 'http://127.0.0.1:8080';

  @override
  Future<String> uploadImage(dynamic file) async {
    if (file is! html.File) throw Exception('Invalid file type');
    
    try {
      // FormDataの作成
      final formData = html.FormData();
      
      // ファイルを追加
      formData.appendBlob('image', file);

      // XMLHttpRequestの作成
      final request = html.HttpRequest();
      request.open('POST', '$baseUrl/upload/image');
      
      // リクエストの送信
      final completer = Completer<String>();
      request.onLoad.listen((e) {
        if (request.status == 200) {
          final response = jsonDecode(request.responseText!);
          completer.complete(response['url']);
        } else {
          completer.completeError('Failed to upload image: ${request.statusText}');
        }
      });

      request.onError.listen((e) {
        completer.completeError('Error uploading image: $e');
      });

      request.send(formData);
      return await completer.future;
    } catch (e) {
      print('Error uploading image: $e');
      rethrow;
    }
  }
}