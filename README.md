# PHP Filter Chain Generator (Versi Go)

Proyek ini merupakan implementasi ulang dari [PHP Filter Chain Generator milik Synacktiv](https://github.com/synacktiv/php_filter_chain_generator/), namun ditulis ulang dalam bahasa Go. Proyek ini memberikan fungsionalitas yang sama, tetapi dioptimalkan untuk digunakan sebagai binary command-line di sistem Linux.

## Fitur:
- **Binary Eksekutabel**: Setelah dikompilasi, kode Go ini dapat dijalankan sebagai binary dengan menempatkannya di `/usr/bin/phpchain` untuk akses mudah.
- **Tanpa Ketergantungan**: Tidak seperti versi Python, versi Go ini tidak memerlukan ketergantungan apa pun untuk dijalankan, sehingga lebih ringan dan cepat.
- **Fungsionalitas yang Sama**: Menyediakan kemampuan generasi rantai filter PHP yang sama seperti proyek aslinya.

## Instalasi:
1. Clone repository ini.
   ```bash
   git clone https://github.com/NumeXx/PHP-Filter-Chain-Generator.git
   cd PHP-Filter-Chain-Generator
   ```
2. Kompilasi kode menjadi binary
   ```bash
   go build -o phpchain phpchain.go
   ```
3. Pindahkan binary ke `/usr/bin`
   ```bash
   sudo mv phpchain /usr/bin/
   ```

## Penggunaan
```bash
phpchain --chain "payload_anda_disini"
```

## Kredit:
Proyek ini terinspirasi oleh [PHP Filter Chain Generator oleh Synacktiv](https://github.com/synacktiv/php_filter_chain_generator/).
