# API Data Wilayah Indonesia

REST API untuk membaca data provinsi, kabupaten/kota, kecamatan, desa, dan kode pos. Data disimpan di MySQL dan endpoint mengembalikan JSON.

## Sumber Data

Data yang digunakan API ini berasal dari repositori berikut:

| Data | Sumber |
| --- | --- |
| Kode dan nama wilayah administrasi Indonesia | [cahyadsn/wilayah](https://github.com/cahyadsn/wilayah) |
| Kode pos wilayah Indonesia | [cahyadsn/wilayah_kodepos](https://github.com/cahyadsn/wilayah_kodepos) |

## Menjalankan aplikasi

Persyaratan: Go sesuai versi pada `go.mod` dan MySQL.

Atur environment variable database berikut sebelum menjalankan aplikasi:

| Variable | Keterangan |
| --- | --- |
| `DB_HOST` | Host MySQL |
| `DB_PORT` | Port MySQL |
| `DB_USER` | Username MySQL |
| `DB_PASS` | Password MySQL |
| `DB_NAME` | Nama database |
| `IS_DEV` | Default `true`. Jika bernilai `true`, aplikasi menjalankan migrasi saat startup. |

Contoh menjalankan secara lokal:

```sh
export DB_HOST=127.0.0.1
export DB_PORT=3306
export DB_USER=root
export DB_PASS=password
export DB_NAME=wilayah
export IS_DEV=true

go run ./cmd/api
```

Server mendengarkan pada port `9876`, sehingga base URL lokal adalah `http://localhost:9876`.

## Endpoint

Semua endpoint menggunakan metode `GET` dan tidak memerlukan autentikasi.

| Endpoint | Query | Keterangan |
| --- | --- | --- |
| `/healthz` | — | Pemeriksaan kesehatan server. |
| `/wilayah/province` | — | Daftar provinsi. |
| `/wilayah/city` | `code` | Daftar kabupaten/kota di bawah kode provinsi. |
| `/wilayah/district` | `code` | Daftar kecamatan di bawah kode kabupaten/kota. |
| `/wilayah/subdistrict` | `code` | Daftar desa/kelurahan di bawah kode kecamatan. |
| `/wilayah/postalcode` | `code` | Daftar kode pos yang kode wilayahnya diawali nilai `code`. |

Kode wilayah menggunakan format bersegmen dengan titik, misalnya `11`, `11.01`, dan `11.01.01`. Gunakan kode level induk untuk mengambil anak wilayahnya:

```text
Provinsi       11
Kabupaten/kota 11.01
Kecamatan      11.01.01
Desa           11.01.01.2001
```

Contoh pemanggilan:

```sh
curl http://localhost:9876/wilayah/province
curl 'http://localhost:9876/wilayah/city?code=11'
curl 'http://localhost:9876/wilayah/district?code=11.01'
curl 'http://localhost:9876/wilayah/subdistrict?code=11.01.01'
curl 'http://localhost:9876/wilayah/postalcode?code=11.01.01.2001'
```

Parameter `code` sebaiknya diisi sesuai level wilayah yang diminta. Filter menggunakan awalan kode (`LIKE code%` untuk kode pos dan `LIKE code.%` untuk daftar wilayah bertingkat).

## Format respons

Respons endpoint wilayah (provinsi, kabupaten/kota, kecamatan, dan desa/kelurahan) berbentuk:

```json
{
  "status": "OK",
  "data": [
    {
      "code": "11",
      "nama": "Aceh",
      "parent_code": "",
      "tingkat": "Provinsi"
    }
  ]
}
```

Nilai `tingkat` adalah `Provinsi`, `Kabupaten/Kota`, `Kecamatan`, atau `Desa`. `parent_code` berisi kode induk yang dibentuk oleh handler. Untuk provinsi nilainya string kosong.

Respons endpoint kode pos menggunakan format berikut:

```json
{
  "status": "OK",
  "data": [
    {
      "kode": "11.01.01.2001",
      "kodepos": "23773"
    }
  ]
}
```

Jika terjadi error saat membaca data, API mengembalikan HTTP `404` dengan format:

```json
{
  "status": "!OK",
  "message_error": "pesan error"
}
```

## Status HTTP

- `200 OK`: request berhasil. Walaupun handler memberikan nilai status `202`, helper respons saat ini tidak menerapkan nilai tersebut sehingga Fiber mengirim status default `200`.
- `404 Not Found`: operasi pembacaan data menghasilkan error.

