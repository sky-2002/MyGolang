package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for courses - file
type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"` // keeping Author a pointer
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware , helper - file
func (c *Course) IsEmpty() bool {
	// return c.CourseID == "" && c.CourseName == "" // we can create more such functions
	return c.CourseName == ""
}

func main() {
	// we will build an api, say to deal with courses for a website.

	fmt.Println("Handling APIs in Golang")

	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseID: "2", CourseName: "ReactJS", CoursePrice: 100, Author: &Author{FullName: "James bond", Website: "jb.com"}})
	courses = append(courses, Course{CourseID: "4", CourseName: "MERN", CoursePrice: 199, Author: &Author{FullName: "James bond", Website: "mernjb.com"}})
	courses = append(courses, Course{CourseID: "5", CourseName: "Python", CoursePrice: 499, Author: &Author{FullName: "Sky", Website: "sky.com/Py"}})
	courses = append(courses, Course{CourseID: "17", CourseName: "Golang", CoursePrice: 499, Author: &Author{FullName: "Sky", Website: "sky.com/Go"}})
	courses = append(courses, Course{CourseID: "20", CourseName: "JS", CoursePrice: 150, Author: &Author{FullName: "Sky", Website: "jb.com/JS"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET") // if we call id as course_id, make changes in function
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to api in gloang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	// http header contains metadata about the request or response,
	// specifies media format too
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")

	// grab id from request
	params := mux.Vars(r)

	fmt.Println("Params looks like this:", params)

	// loop through courses, find matching id and return resposne
	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create all courses")
	w.Header().Set("Content-type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what if: data is like this {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data in json")
	}

	// TODO: check only if title is duplicate, or maybe we can check if author is also duplicate
	// further we can check for price and do something if all match to an existing course

	// generate random course id, string
	rand.Seed(time.Now().Unix())
	course.CourseID = strconv.Itoa(rand.Intn(100))

	// push the course in the slice
	courses = append(courses, course)

	json.NewEncoder(w).Encode("Added course successfully !!")
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create all courses")
	w.Header().Set("Content-type", "application/json")

	// grab the id from request
	params := mux.Vars(r)

	// loop through courses, remove the course, add with the same ID which was before
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			course.CourseID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode("Value updated successfully !!")
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// send a response when ID is not found
	json.NewEncoder(w).Encode("ID not found")
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting one course")
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	// loop through the courses, find the ID, remove that course
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully !!")
			break
		}
	}
}
