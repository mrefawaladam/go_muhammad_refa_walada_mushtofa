# 16 - Soal ORM & Code Structure (MVC)
## Soal Prioritas 1 (80)
1. Buat API CRUD User dengan spesifikasi  
   - Route : /users  Method : GET
    <br>![Alt Text](assets/2023-04-01_18-32_1.png)<br>
   -  Route : /users:id  Method : GET
   <br>![Alt Text](assets/2023-04-01_18-33.png)<br>
    -  Route : /users  Method : POST (disini saya melakukan explorasi dengan melakukan encrypt password yang di kirim menggukanan bcrypt)
   <br>![Alt Text](assets/2023-04-04_20-19.png)<br>
    -  Route : /users:id  Method : PUT
       <br>![Alt Text](assets/2023-04-04_21-02.png)<br>
   -  Route : /users:id  Method : DELETE
    <br>![Alt Text](assets/2023-04-04_21-03.png)<br>

   -  Hendling ketika menambahkan sesuatu yang tidak harus di kirim
       <br>![Alt Text](assets/2023-04-04_20-24.png)<br>
2. Berdasarkan soal prioritas 1 yang telah dikerjakan, tambahkan fitur CRUD untuk data buku dengan spesifikasi sebagai berikut.
   - Route : /books  Method : GET
   <br>![Alt Text](assets/2023-04-04_21-42.png)<br>
   -  Route : /books:id  Method : GET
   <br>![Alt Text](assets/2023-04-04_21-44.png)<br>
   -  Route : /books  Method : POST 
   <br>![Alt Text](assets/2023-04-04_21-24.png)<br>
   -  Route : /books:id  Method : PUT
    <br>![Alt Text](assets/2023-04-04_21-47.png)<br>
   -  Route : /books:id  Method : DELETE
   <br>![Alt Text](assets/2023-04-04_21-47_1.png)<br>

## Resume
ketika Menggunakan ORM dan struktur MVC dapat membantu proeses developing untuk mempercepat proses pengembangan aplikasi dengan mengurangi waktu dan kompleksitas dalam membuat operasi basis data, memisahkan tanggung jawab dari setiap komponen, dan memudahkan pemeliharaan kode. di Golang sendiri menawarkan fleksibilitas dalam mengorganisasi kode dengan struktur file dan package yang dapat disesuaikan dengan kebutuhan proyek. Struktur file dan package yang baik membantu untuk memudahkan debugging dan pemeliharaan kode dan Dalam aplikasi Golang, struktur MVC umumnya diimplementasikan dengan menggunakan package yang berbeda untuk setiap komponen. Model berisi representasi data dan akses ke database, View menampilkan data untuk user, dan Controller mengatur logika bisnis dan pengolahan data. di golang sendiri mempunyai ORM (Object-Relational Mapping) adalah teknik pemetaan antara basis data relasional dan model objek. ORM memungkinkan developer untuk memanipulasi data pada basis data dengan menggunakan objek dan metode yang mudah dipahami.

3 poin yang dapat saya pelajari 
1. saya bisa menggunakan MVC di golang
2. saya bisa melakukan interaksi database di golang
3. mengentahui komudahan golang dalam membuat mvc