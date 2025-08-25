Dukcapil Web Desa Dasan Lekong

Aplikasi web Dukcapil Desa Dasan Lekong dibuat untuk menampilkan data kependudukan desa secara digital. Sistem ini menggunakan Golang (Gin Framework) di sisi backend dan HTML template (Jinja-like) di sisi frontend.

✨ Fitur

📊 Menampilkan data keluarga (kepala keluarga, alamat, jumlah anggota).

👨‍👩‍👧‍👦 Menampilkan data penduduk.

🔎 Pencarian & filter data.

🖥️ Tampilan berbasis web yang responsif.

🛠️ Teknologi yang Digunakan

Backend: Golang + Gin Framework

Frontend: HTML Template + Bootstrap/Tailwind (opsional)

Utils: HTTP JSON Fetcher

Config: API base URL disimpan di config/

📂 Struktur Folder
dukcapil-web/
│── config/           # Konfigurasi aplikasi (API Base URL)
│── controllers/      # Controller (logic untuk routing)
│── templates/        # Template HTML (frontend)
│   └── pages/keluarga/
│── utils/            # Helper (misalnya HTTP JSON fetch)
│── main.go           # Entry point aplikasi

🚀 Cara Menjalankan

Clone repo ini:

git clone https://github.com/EkyGalih/dukcapil-web.git
cd dukcapil-web


Install dependencies:

go mod tidy


Jalankan aplikasi:

go run main.go


Akses di browser:

http://localhost:8080

⚙️ Konfigurasi

File konfigurasi ada di folder config/. Pastikan API_BASE_URL sudah mengarah ke endpoint API Dukcapil. Contoh:

package config

const API_BASE_URL = "http://localhost:8000/api"

📸 Screenshot (opsional)

(tambahkan gambar UI kalau sudah ada)

📜 Lisensi

Proyek ini dibuat untuk kebutuhan internal Desa Dasan Lekong.