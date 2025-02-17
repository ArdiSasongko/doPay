# Setup & Development
---------------------------------------------
1. Package structure
2. Define Primary Ports, Secondary Ports
3. Define Models & Data Structure 
4. Proto Contract
5. Generate Proto File
6. Define Ports & Models
7. Implement Primary & Secondary Adapter
8. Implement Core Business Logic
9. Testing

# Cara Generate Protobuf
------------------------------------------------
1. Di repo ini, menggunakan libprotoc 26.1, untuk cek versi protoc yang terinstall bisa menggunakan command berikut
   ```shell
   protoc --version
   ```
2. Definisikan spec proto di folder `lib/protos/v1/xxxx/xxx.proto`
2. Misal untuk wallet di `lib/protos/v1/wallet/wallet.proto`
3. Lalu jalankan command berikut
   ```shell
    protoc -I . --proto_path=. \
 	       --go_out=paths=source_relative:. \
 	       --go-http_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative:. \
	       --openapiv2_out ./lib/protos/openapiv2 --openapiv2_opt logtostderr=true \
			./lib/protos/v1/wallet/wallet.proto
   ```
4. File *.go akan tergenerate sesuai lokasi folder proto yang kita define, dan akan melakukan generate ke spec open API di folder `lib/protos/openapiv2/lib/protos/v1/wallet/wallet.swagger.json`
5. Isi dari makefile, pada dasarnya merupakan command di atas yang lebih di rapikan
6. Untuk generate melalui makefile, misal untuk proto wallet yang ada di folder `lib/protos/v1/wallet/wallet.proto` bisa menggunakan command berikut
   ```shell
      MODULE=v1/wallet make generate
   ```
7. Untuk generate file go di semua proto, bisa langsung jalankan command berikut
   ```shell
   make generate
   ```

# Cara Generate Migrations
------------------------------------------------
1. jalankan command berikut
   ```shell
   make migrations (name_file_yang_dihasilkan)
   ```
2. buat file config baru misal (config.name.json) untuk menyimpan config
3. Lalu jalankan command berikut untuk up ke database
   ```shell
      MODULE=name make migrate-up
   ```
   name dari (config."name".json)
4. jalankan command berikut untuk down dari database
   ```shell
      MODULE=name make migrate-down
   ```
   name dari (config."name".json)