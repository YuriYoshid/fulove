import 'package:fulove/upload_service/upload_service_interface.dart';

class UploadServiceStub implements UploadService {
  @override
  Future<String> uploadImage(dynamic file) {
    throw UnimplementedError('uploadImage not implemented on this platform');
  }
}