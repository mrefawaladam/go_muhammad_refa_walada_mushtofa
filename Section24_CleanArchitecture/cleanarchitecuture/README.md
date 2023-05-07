# Section 24
Project Structure yang saya buat adalah menggunakan arsitektur berikut ini 
untuk sourc file di folerd /cleararctecture mohon maaf masih belum bisa screenshot kan 
## Package Structure

```zsh
.
├── LICENSE
├── README.md
├── build                                     # Packaging and Continuous Integration
│   ├── Dockerfile
│   └── init.sql
├── cmd                                       # Main Application
│   └── app
│       └── main.go
├── internal                                  # Private Codes
│   └── app
│       ├── adapter
│       │   ├── controller.go                 # Controller
│       │   ├── postgresql                    # Database
│       │   │   ├── conn.go
│       │   │   └── model                     # Database Model
│       │   │       ├── card.go
│       │   │       ├── cardBrand.go
│       │   │       ├── order.go
│       │   │       ├── parameter.go
│       │   │       ├── payment.go
│       │   │       └── person.go
│       │   ├── repository                    # Repository Implementation
│       │   │   ├── order.go
│       │   │   └── parameter.go
│       │   ├── service                       # Application Service Implementation
│       │   │   └── bitbank.go
│       │   └── view                          # Templates
│       │       └── index.tmpl
│       ├── application
│       │   ├── service                       # Application Service Interface
│       │   │   └── exchange.go
│       │   └── usecase                       # Usecase
│       │       ├── addNewCardAndEatCheese.go
│       │       ├── ohlc.go
│       │       ├── parameter.go
│       │       ├── ticker.go
│       │       └── ticker_test.go
│       └── domain
│           ├── factory                       # Factory
│           │   └── order.go
│           ├── order.go                      # Entity
│           ├── parameter.go
│           ├── parameter_test.go
│           ├── person.go
│           ├── repository                    # Repository Interface
│           │   ├── order.go
│           │   └── parameter.go
│           └── valueobject                   # ValueObject
│               ├── candlestick.go
│               ├── card.go
│               ├── cardbrand.go
│               ├── pair.go
│               ├── payment.go
│               ├── ticker.go
│               └── timeunit.go
└── testdata                                  # Test Data
    └── exchange_mock.go
```


# Sumary


Dari struktur direktori di atas, terlihat bahwa Saya sedang menggunakan Clean Architecture dalam pengembangan aplikasi Golang Saya. Berikut adalah beberapa hal yang mungkin Saya pelajari terkait Clean Architecture:

Pemisahan kode: Dalam Clean Architecture, kode dibagi menjadi beberapa layer (atau modul) yang memiliki tugas spesifik. Layer-layer ini tidak boleh saling tergantung, dan harus beroperasi melalui antarmuka yang didefinisikan secara eksplisit.

Layer-domain: Layer-domain berisi entitas dan objek nilai, serta aturan bisnis. Entitas adalah objek yang memiliki identitas unik dan dapat berubah seiring waktu. Objek nilai adalah objek yang tidak memiliki identitas, dan hanya didefinisikan oleh nilainya. Aturan bisnis diterapkan di layer ini.

Layer-usecase: Layer-usecase berisi logika bisnis yang berhubungan dengan tindakan yang dilakukan dalam aplikasi. Setiap usecase adalah sebuah aturan bisnis tunggal, dan membutuhkan satu atau lebih input, dan menghasilkan satu atau lebih output.

Layer-interface: Layer-interface terdiri dari antarmuka dengan bagian lain dari aplikasi, seperti layer-infrastructure. Layer-interface mengimplementasikan aturan bisnis dan menggunakan usecase untuk memproses input dan menghasilkan output.

Layer-infrastructure: Layer-infrastructure berisi kode yang spesifik untuk lingkungan dan teknologi tertentu. Layer ini berisi implementasi dari antarmuka yang didefinisikan di layer-interface, dan memberikan akses ke database, jaringan, dan sumber daya eksternal lainnya.

Test: Clean Architecture mendorong penggunaan pengujian otomatis untuk memastikan bahwa kode berfungsi seperti yang diharapkan dan sesuai dengan aturan bisnis. Saya dapat menempatkan kode pengujian di dalam direktori testdata.

Dalam struktur direktori yang diberikan, dapat dilihat bahwa direktori cmd berisi kode untuk menjalankan aplikasi, sedangkan direktori internal berisi kode yang lebih spesifik, seperti aturan bisnis, implementasi database, dan kode antarmuka. Direktori testdata berisi data yang digunakan untuk menguji kode. Kode di dalam direktori internal diatur ke dalam subdirektori yang mewakili lapisan dalam Clean Architecture.