package mahasiswa

import (
	"encoding/json"
	"net/http"
)

type Mahasiswa struct {
	Nama string "json:nama"
	NIM string "json:nim"
	Alamat string "json:alamat"
}

var mahasiswas = []Mahasiswa{
	Mahasiswa{
		Nama: "Nabil",
		NIM: "3279403",
		Alamat: "Malang",
	},
	Mahasiswa{
		Nama: "Joy",
		NIM: "3294538",
		Alamat: "Surabaya",
	},
}

func GetMahasiswaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	json.NewEncoder(w).Encode(mahasiswas)
}

func FindMahasiswaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	GetNameMahasiswa := Mahasiswa{
		Nama: r.FormValue("Nama"),
	}

	for _, value := range mahasiswas{
		if GetNameMahasiswa.Nama == value.Nama {
			response, err := json.Marshal(value)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)

			return
		}
	}

	http.Error(w, "Nama not found", http.StatusNotFound)
}

