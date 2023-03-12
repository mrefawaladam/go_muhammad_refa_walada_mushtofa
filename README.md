# Soal Problem Solving Paradigm - Brute Force, Greedy and Dynamic Programming

## Tugas

 1. Kode berikut ini dituliskan tanpa mengikuti kebiasaan penulisan yang disarankan.
    - Berapa banyak kekurangan dalam penulisan kode tersebut?
    - Bagian mana saja terjadi kekurangan tersebut?
    - Tuliskan alasan dari setiap kekurangan tersebut. Alasan bisa diberikan dalam bentuk komentar pada kode yang disediakan berikut.
      <br>********************************Jawab :  [Sours Code](tugas/../Section13_CleanCode/tugas/userservice.go)********************************   
    Penjelasan  
    <br>
    1. pertama saya melakukan penjelasan dengan memperjelasmaksut dari variabel contoh yang saya rubah pada method getuserbyid di looping itu menggunakan variabel yang tidak jelas lalu saya semputnakan dengan menggunakan variabel yang jelas dan saya berkomitmen untuk menggunakan bahasa inggris untuk penulisan code yang saya buat
    2. lalu saya melakukan perubahan tipe data pada struct saya ubah karena jika penulisan dari variabel yang tidak konsisten akan menimbul kan beberapa masalah contoh ketika memasukan data json dari api di parsing ke struc jika tipe dari variabel tidak sesuai akan menimbulkan kesalahan
    3. lalu saya merubah penamaan method yang tidak deskriptif karena jika menggunakan method yang deskriptif akan mempermudah pemahan dari sebuah method
    4. menghapus variabel yang tidak di gunakan

2. Kode berikut ini dituliskan tanpa mengikuti kebiasaan penulisan yang disarankan. Ubahlah penulisan kode berikut menjadi lebih terbaca dan lebih rapi!
   <br>********************************Jawab :  [Sours Code](tugas/../Section13_CleanCode/tugas/kendaraa.go)********************************   
   Penjelasan  
   <br>
   1. Di sini saya melakukan perubahan penamaan struct dan method karena penamaanya sendri kurang tepat saya merubah menjadi camelCase untuk penamaan variabel dan PascalCase untuk penamaan method dan sturct
   2. di method TambahKecepatan saya melakukan perubahan dengan mengubah oprator += karena untuk melakukan penambahan kecepatan baru supaya tidak boros memori dan kesimpelanya
  
## Sumary
### 3 Point Yang saya pelajari
1. saya bisa memahami kosep penggunaan clean code 
2. saya menyadari kebiasaan saya menambahkan banyak komentar itu kuran tepat
3. saya bisa membuat sebuah code yang mudah di pahami tanpa adanya komentar
   
### Resume
Clean code di Golang sendiri mengacu pada konsep  terbaik dalam menulis kode yang mudah dibaca, dipahami, dan dipelihara oleh programmer lain. Beberapa prinsip clean code di Golang antara lain:

1. Menggunakan nama variabel yang deskriptif dan bermakna, agar memudahkan programmer lain dalam memahami maksud dari variabel tersebut.

2. Menggunakan gaya penulisan camelCase untuk nama fungsi dan variabel, serta PascalCase untuk nama tipe data. Hal ini memudahkan programmer dalam membedakan tipe data dengan variabel atau fungsi.

3. Memiliki fungsi yang tidak terlalu panjang dan hanya melakukan satu tugas saja, sehingga memudahkan dalam debugging dan memperkecil kemungkinan terjadinya bug.

4. Menghindari penggunaan variabel global, kecuali jika benar-benar diperlukan. Hal ini memperkecil kemungkinan terjadinya konflik variabel pada program.

5. Menjaga agar kode memiliki indentasi yang konsisten, sehingga memudahkan programmer lain dalam membaca kode.

6. Menghindari penggunaan komentar yang tidak perlu atau tidak relevan, dan hanya menggunakan komentar ketika diperlukan untuk menjelaskan kode yang sulit dipahami.

Di golang sendiri Dalam menulis clean code sangat memperhatikan prinsip-prinsip di atas, juga perlu diperhatikan gaya penulisan kode yang konsisten dan sesuai dengan standar kode yang berlaku. Hal ini akan memudahkan programmer lain dalam memahami dan memodifikasi kode yang telah dibuat.