// demomiddleware
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
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

func findID(ID int) (*Course, int) {
	for i, course := range CourseList {
		if course.CourseId == ID {
			return &course, i
		}
	}
	return nil, 0
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegment := strings.Split(r.URL.Path, "course/")
	if len(urlPathSegment) < 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ID, err := strconv.Atoi(urlPathSegment[len(urlPathSegment)-1])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	course, listItemIndex := findID(ID)
	if course == nil {
		http.Error(w, fmt.Sprintf("no course with id %d", ID), http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		courseJson, err := json.Marshal(course)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(courseJson)
	case http.MethodPut:
		var updatedCourse Course
		byteBody, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		err = json.Unmarshal(byteBody, &updatedCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if updatedCourse.CourseId != ID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		CourseList[listItemIndex] = updatedCourse
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func coursesHandler(w http.ResponseWriter, r *http.Request) {
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

func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before handler middle start")

		handler.ServeHTTP(w, r)

		fmt.Println("after handler middle finised")
	})
}

func main() {

	courseItemHandler := http.HandlerFunc(courseHandler)
	courseListHandler := http.HandlerFunc(coursesHandler)
	http.Handle("/course/", middlewareHandler(courseItemHandler))
	http.Handle("/course", middlewareHandler(courseListHandler))
	log.Println("Server started at :5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
