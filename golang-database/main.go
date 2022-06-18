package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/jcelliott/lumber"
)

const Version = "1.0.0"

type (

	// Logger interface will be used for logging different types of messages
	Logger interface {
		// These functions below can take any number of empty interfaces as arguments.
		// That means we can pass any type as an argument

		//------------------------------------------------------------------------
		// ----------We may need to implement these-------------------------------
		//------------------------------------------------------------------------
		Fatal(string, ...interface{}) // ...interface{} means that after passing a string, we can pass any number of interfaces
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}

	Driver struct {
		mutex   sync.Mutex             // a driver will have its mutex to lock and unlock objects.
		mutexes map[string]*sync.Mutex // it will also have a map whose values are
		dir     string
		log     Logger
	}
)

type Options struct {
	Logger
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

func New(directory string, options *Options) (*Driver, error) {
	directory = filepath.Clean(directory) // Clean applies some rules iteratively on the string until no further processing can be done

	opts := Options{}

	if options != nil {
		opts = *options // set as opts as the options passed by the user
	}

	if opts.Logger == nil {
		opts.Logger = lumber.NewConsoleLogger((lumber.INFO))
	}

	myDriver := Driver{
		dir:     directory,
		mutexes: make(map[string]*sync.Mutex), // make creates a slice of the type of argument passed
		log:     opts.Logger,
	}

	if _, err := os.Stat(directory); err == nil {
		opts.Logger.Debug("Using '%s' (Database already exists)\n", directory)
		return &myDriver, nil
	}

	opts.Logger.Debug("Creating the database at '%s'\n", directory)

	return &myDriver, os.MkdirAll(directory, 0755)
}

func (d *Driver) Write(collection string, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("Missing collection - no place to save record")
	}
	if resource == "" {
		return fmt.Errorf("Missing resource - unable to save record (no name) !")
	}

	mutex := d.getOrCreateMutex(collection)
	mutex.Lock()
	// defer mutex.Unlock()

	dir := filepath.Join(d.dir, collection)
	finalPath := filepath.Join(dir, resource, ".json")
	tempPath := finalPath + ".tmp"

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}

	b = append(b, byte('\n'))

	if err = ioutil.WriteFile(tempPath, b, 0644); err != nil {
		return err
	}
	defer mutex.Unlock()
	return os.Rename(tempPath, finalPath)
}

func (d *Driver) Read(collection string, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("Missing collection - unable to read")
	}
	if resource == "" {
		return fmt.Errorf("Missing resource - unable to read record (no name) !")
	}

	record := filepath.Join(d.dir, collection, resource)

	if _, err := stat(record); err != nil {
		return err
	}

	b, err := ioutil.ReadFile(record + ".json")
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &v)
}

func (d *Driver) ReadAll(collection string) ([]string, error) {
	if collection == "" {
		return nil, fmt.Errorf("Missing collection - unable to read")
	}

	dir := filepath.Join(d.dir, collection)

	if _, err := stat(dir); err != nil {
		return nil, err
	}

	files, _ := ioutil.ReadDir(dir)

	var records []string

	for _, file := range files {
		b, err := ioutil.ReadFile(filepath.Join(dir, file.Name())) // maybe we need to add .json as well
		if err != nil {
			return nil, err
		}
		records = append(records, string(b))
	}
	return records, nil
}

func (d *Driver) Delete(collection string, resource string) error {
	path := filepath.Join(collection, resource)
	mutex := d.getOrCreateMutex(collection)

	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, path)

	switch fi, err := stat(dir); {
	case fi == nil, err != nil:
		return fmt.Errorf("Unable to find file or directory named %v\n", path)
	case fi.Mode().IsDir():
		return os.RemoveAll(dir)
	case fi.Mode().IsRegular():
		return os.RemoveAll(dir + ".json")
	}
	return nil
}

func (d *Driver) getOrCreateMutex(collection string) *sync.Mutex {

	d.mutex.Lock()
	defer d.mutex.Unlock()

	m, ok := d.mutexes[collection]
	if !ok {
		m := &sync.Mutex{}
		d.mutexes[collection] = m
	}
	return m
}

func stat(path string) (fi os.FileInfo, err error) {
	if fi, err = os.Stat(path); os.IsNotExist(err) {
		fi, err = os.Stat(path + ".json")
	}
	return
}

func main() {
	dir := "./"

	db, err := New(dir, nil)

	if err != nil {
		fmt.Println("Error", err)
	}

	employees := []User{
		User{
			Name:    "Aakash",
			Age:     "20",
			Contact: "945689",
			Company: "SkyTech",
			Address: Address{City: "Nagpur", State: "Maharashtra", Country: "India"},
		},
		User{
			Name:    "John",
			Age:     "25",
			Contact: "988689",
			Company: "SkyTech",
			Address: Address{City: "New York", State: "New York State", Country: "USA"},
		},
		User{
			Name:    "Zhang",
			Age:     "28",
			Contact: "444689",
			Company: "SkyTech",
			Address: Address{City: "Beijing", State: "", Country: "China"},
		},
		User{
			Name:    "Rahul",
			Age:     "22",
			Contact: "945559",
			Company: "SkyTech",
			Address: Address{City: "Nagpur", State: "Maharashtra", Country: "India"},
		},
		User{
			Name:    "Rohit",
			Age:     "26",
			Contact: "933339",
			Company: "SkyTech",
			Address: Address{City: "Nagpur", State: "Maharashtra", Country: "India"},
		},
	}

	for _, employee := range employees {
		db.Write("users", employee.Name, User{
			Name:    employee.Name,
			Age:     employee.Age,
			Contact: employee.Contact,
			Company: employee.Company,
			Address: employee.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(records)

	allUsers := []User{}
	for _, record := range records {
		employeeFound := User{}
		if err := json.Unmarshal([]byte(record), &employeeFound); err != nil {
			fmt.Println("Error", err)
		}
		allUsers = append(allUsers, employeeFound)
	}
	fmt.Println(allUsers)

	// if err:= db.Delete("users","john"); err !=nil{
	// 	fmt.Println("Error",err)
	// }

	// if err:= db.Delete("users",""); err !=nil{
	// 	fmt.Println("Error",err)
	// }

}
