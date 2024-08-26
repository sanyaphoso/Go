// package main

// import (
// 	"encoding/json"
// 	"io"
// 	"log"
// 	"net/http"
// )

// type Course struct {
// 	CourseId   int     `json:"id"`
// 	CourseName string  `json:"name"`
// 	Price      float64 `json:"price"`
// 	Instructor string  `json:"instructor"`
// }

// var CourseList []Course

// func init() {
// 	CourseJson := `[
// 		{
// 			"id": 1,
// 			"name": "Python",
// 			"price": 2590,
// 			"instructor": "Sanya"
// 		},
// 		{
// 			"id": 2,
// 			"name": "JavaScript",
// 			"price": 2590,
// 			"instructor": "Sanya"
// 		},
// 		{
// 			"id": 3,
// 			"name": "SQL",
// 			"price": 2590,
// 			"instructor": "Sanya"
// 		}
// 	]`
// 	err := json.Unmarshal([]byte(CourseJson), &CourseList)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func courseHandler(w http.ResponseWriter, r *http.Request) {
// 	courseJson, err := json.Marshal(CourseList)
// 	switch r.Method {
// 	case http.MethodGet:
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(courseJson)

// 	case http.MethodPost:
// 		var newCourse Course
// 		Bodybyte, err := io.ReadAll(r.Body)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		err = json.Unmarshal(Bodybyte, newCourse)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		CourseList = append(CourseList, newCourse)
// 		w.WriteHeader(http.StatusCreated)
// 		return
// 	}
// }

// func main() {
// 	http.HandleFunc("/course", courseHandler)
// 	http.ListenAndServe(":5000", nil)
// }

//examplerequest ``

package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Course struct {
	CourseId   int     `json:"id"`
	CourseName string  `json:"name"`
	Price      float64 `json:"price"`
	Instructor string  `json:"instructor"`
}

var CourseList []Course
var lastID int

func init() {
	CourseJson := `[
		{
			"id": 1,
			"name": "Python",
			"price": 2590,
			"instructor": "Sanya"
		},
		{
			"id": 2,
			"name": "JavaScript",
			"price": 2590,
			"instructor": "Sanya"
		},
		{
			"id": 3,
			"name": "SQL",
			"price": 2590,
			"instructor": "Sanya"
		}
	]`
	err := json.Unmarshal([]byte(CourseJson), &CourseList)
	if err != nil {
		log.Fatal(err)
	}

	// Set the lastID to the highest current CourseId in the list
	for _, course := range CourseList {
		if course.CourseId > lastID {
			lastID = course.CourseId
		}
	}
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		courseJson, err := json.Marshal(CourseList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(courseJson)

	case http.MethodPost:
		var newCourse Course
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		err = json.Unmarshal(bodyBytes, &newCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Automatically increment and set the new course ID
		lastID++
		newCourse.CourseId = lastID

		CourseList = append(CourseList, newCourse)
		w.WriteHeader(http.StatusCreated)
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/course", courseHandler)
	log.Println("Server started at :5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
