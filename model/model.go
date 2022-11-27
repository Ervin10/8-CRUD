package model

import (
	"log"
	"strconv"
	"time"
)

// Create struct, struct berfungsi untuk membuat struktur dari tipe data
type Project struct {
	Id          int
	ProjectName string
	StartDate   string
	EndDate     string
	Desc        string
	Tech        []string
	Image       string
}

// Buat array of object sebagai local storage
var DataProjects = []Project{
	{
		ProjectName: "Dumbways 2022",
		StartDate:   "2022-11-24",
		EndDate:     "2022-12-24",
		Desc:        "Halo Dumbways",
		Tech:        []string{"node", "angular", "react", "typescript"},
		Image:       "public/assets/img/saitama.png",
	},
}

// Function render time
func (p Project) RenderTime(date string) string { // parameter didapatkan dari pemanggilan funtion di file project-detail.html
	// method time.Parse() berfungsi untuk memparsing date. par1: layout format dari waktu yang diparsing, par2: string yang ingin diparsing
	time, err := time.Parse("2006-01-02", date)

	if err != nil {
		log.Fatal(err)
	}

	// Buat slice yang akan digunakan untuk format date yang akan diparsing
	Months := [...]string{"Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Des"}

	// strconv.Itoa() berfungsi untuk mengkonversi int menjadi string
	return strconv.Itoa(time.Day()) + " " + Months[time.Month()-1] + " " + strconv.Itoa(time.Year())
}

func (p Project) DurationTime(startDate string, endDate string) string {
	start, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)

	duration := end.Sub(start).Hours() // Selisih waktu akan dikonversi menjadi jam
	day := 0
	month := 0
	year := 0

	for duration >= 24 {
		day += 1
		duration -= 24
	}

	for day >= 30 {
		month += 1
		day -= 30
	}

	for month >= 12 {
		year += 1
		month -= 12
	}

	if year != 0 && month != 0 {
		return strconv.Itoa(year) + " year, " + strconv.Itoa(month) + " month, " + strconv.Itoa(day) + " day"
	} else if month != 0 {
		return strconv.Itoa(month) + " month, " + strconv.Itoa(day) + " day"
	} else {
		return strconv.Itoa(day) + " Day"
	}
}

// func HandleEditProject(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 		tmpt, err := template.ParseFiles("views/edit-project.html")

// 		if err != nil {
// 			w.Write([]byte("Message :" + err.Error()))
// 		}

// 		// Tangkap id dari halaman index
// 		id, _ := strconv.Atoi(mux.Vars(r)["id"])

// 		// Panggil struct lalu tampung kedalam variabel
// 		var data = model.Project{}

// 		// Panggil array DataProject lalu looping dan simpan ke variabel data
// 		for i, d := range model.DataProjects {
// 			if i == id {
// 				data = model.Project{
// 					ProjectName: d.ProjectName,
// 					StartDate:   d.StartDate,
// 					EndDate:     d.EndDate,
// 					Desc:        d.Desc,
// 					Tech:        d.Tech,
// 					Image:       d.Image,
// 				}
// 			}
// 		}

// 		// Kemudian buat map lalu simpan variabel data kedalam map
// 		var dataProject = map[string]interface{}{
// 			"DataProject": data,
// 		}
// 		tmpt.Execute(w, dataProject)
// 	} else if r.Method == "POST" {
// // Tangkap id dan value dari input
// id, _ := strconv.Atoi(mux.Vars(r)["id"])
// ProjectName := r.PostForm.Get("ProjectName")
// StartDate := r.PostForm.Get("startDate")
// EndDate := r.PostForm.Get("endDate")
// Desc := r.PostForm.Get("desc")

// //  Buat variabel untuk menampung data checkbox
// var checkboxs []string

// // Jika didalam form checkboxs ada value-nya, maka append ke array checkboxs
// if r.FormValue("node") != "" {
// 	checkboxs = append(checkboxs, r.FormValue("node"))
// }
// if r.FormValue("angular") != "" {
// 	checkboxs = append(checkboxs, r.FormValue("angular"))
// }
// if r.FormValue("react") != "" {
// 	checkboxs = append(checkboxs, r.FormValue("react"))
// }
// if r.FormValue("typescript") != "" {
// 	checkboxs = append(checkboxs, r.FormValue("typescript"))
// }

// DataProjects := model.DataProjects

// http.Redirect(w, r, "/", http.StatusMovedPermanently)
// 	}
// }
