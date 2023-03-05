 # Soal String, Advance Function, Pointer, Struct, Method,  Interface, Package & Error Handling

## Soal Prioritas 1 
-  Buatlah sebuah *method* untuk menghitung perkiraan jarak yang bisa ditempuh berdasarkan bahan bakar yang sedang terisi (1.5 L / mill), method tersebut *receiver* kepada struct Car yang memiliki property type dan fuelIn
     <br>********************************Jawab :  [Sours Code](tugas/hitugjarak.go)********************************   
    Outuput 
    <br>![Alt Text](assets/2023-03-04_20-59.png)<br>
 
 - Buat sebuah struct dengan nama **Student** yang mempunyai properti **name** dan **score** dalam bentuk slice kemudian simpan data siswa sebanyak 5 siswa. Setelah 5 siswa dimasukkan maka program menunjukkan skor rata-rata, siswa yang memiliki skor minimum dan maksimal? (implementasikan method)
   <br>********************************Jawab :  [Sours Code](tugas/minmax.go)********************************   
    Outuput 
    <br>![Alt Text](assets/2023-03-04_21-34.png)<br>
 
    
- Buatlah program untuk menentukan substring yang sama di antara keduanya, dengan diberikan dua string A dan B!
    
    <br>********************************Jawab :  [Sours Code](tugas/minmax.go)********************************   
    Outuput 
    <br>![Alt Text](assets/2023-03-04_21-34.png)<br>
 
    
- Buatlah program di Golang untuk menemukan nilai maksimum serta minimum di antara 6 angka inputan. Gunakan multiple return fungsi, pointer untuk referencing maupun deferencing!
    <br>********************************Jawab :  [Sours Code](tugas/maksimum.go)********************************   
    Outuput 
    <br>![Alt Text](assets/2023-03-04_21-06.png)<br>

## Soal Prioritas 2
- **Caesar Cipher**

Buatlah sebuah method dengan parameter `offset` bertipe integer dan `input` bertipe string. Method ini menghasilkan sebuah string baru yang dimana setiap huruf dilakukan pergeseran berdasarkan `offset` yang merupakan jumlah pergeseran hurufnya. String `input` diasumsikan berisi huruf kecil saja dan spasi. Sebagai contoh, ketika terdapat huruf `z` yang digeser dengan `offset` sebesar 3 maka huruf hasil pergeseran adalah `c.`

Daftar karakter ASCII dapat dilihat di link [berikut](https://id.wikipedia.org/wiki/ASCII).

Berdasarkan referensi ASCII, huruf `a` memiliki kode 97, huruf `b` memiliki kode 98, `z` memiliki kode 122. Anda bisa menggunakan operator modulo jika diperlukan.
    <br>********************************Jawab :  [Sours Code](tugas/caesarchipter.go)********************************   
    Outuput 
    <br>![Alt Text](assets/2023-03-04_21-14.png)<br>

## Sumary
- Sting : di golang sendiri memiliki tipe data string yang dapat digunakan untuk merepresentasikan karakter atau teks.  
- Advance Function: Golang memiliki fitur function yang cukup advance, seperti function sebagai parameter, function sebagai return value, anonymous function (closure), dan variadic function.
- Pointer: Pointer merupakan alamat memori pada variabel yang dapat digunakan untuk mengakses nilai variabel tersebut. Di Golang, kita dapat menggunakan operator & untuk mendapatkan alamat memori variabel dan operator * untuk mengakses nilai variabel yang disimpan pada alamat tersebut.
- Struct: Struct merupakan tipe data khusus pada Golang yang memungkinkan kita untuk menyimpan beberapa tipe data yang berbeda dalam satu variabel. Struct sering digunakan untuk merepresentasikan sebuah objek atau data yang kompleks.
- Method: Method pada Golang adalah function yang terkait dengan tipe data tertentu. Method biasanya digunakan untuk melakukan operasi pada variabel yang bertipe struct atau tipe data lain yang sudah didefinisikan methodnya.