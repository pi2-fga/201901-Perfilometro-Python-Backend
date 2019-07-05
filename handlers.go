package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"mime/multipart"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todoId int
	var err error
	if todoId, err = strconv.Atoi(vars["todoId"]); err != nil {
		panic(err)
	}
	todo := RepoFindTodo(todoId)
	if todo.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

/*
curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos
*/
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

/*
curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos
*/
func CreatePathLocation(w http.ResponseWriter, r *http.Request) {
	var location Location
	body, err := ioutil.ReadAll(r.Body)

	fmt.Printf("%s\n", string(body))
	
	if err != nil { panic(err)} 
	if err := r.Body.Close(); err != nil { panic(err) }

	if err := json.Unmarshal(body, &location); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	l := RepoCreatePath(location)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(l); err != nil {
		panic(err)
	}
}

func Upload(client *http.Client, url string, values map[string]io.Reader) (err error) {

    var b bytes.Buffer
    w := multipart.NewWriter(&b)
    for key, r := range values {
        var fw io.Writer
        if x, ok := r.(io.Closer); ok {
            defer x.Close()
        }
      
        if x, ok := r.(*os.File); ok {
            if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
                return
            }
        } else {
            // Add other fields
            if fw, err = w.CreateFormField(key); err != nil {
                return
            }
        }
        if _, err = io.Copy(fw, r); err != nil {
            return err
        }

    }
    // Don't forget to close the multipart writer.
    // If you don't close it, your request will be missing the terminating boundary.
    w.Close()

    // Now that you have a form, you can submit it to your handler.
    req, err := http.NewRequest("POST", url, &b)
    if err != nil {
        return
    }
	
	// Don't forget to set the content type, this will contain the boundary.
    req.Header.Set("Content-Type", w.FormDataContentType())
    res, err := client.Do(req)
    if err != nil {
        return
    }

    // Check the response
    if res.StatusCode != http.StatusOK {
        err = fmt.Errorf("bad status: %s", res.Status)
    }
    return
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	
	r.ParseMultipartForm(10 << 20) // Limit file size to 10
   
    sensorFile, sensorHandler, sensorErr := r.FormFile("sensors_data")
	if sensorErr != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(sensorErr)
        return
	}
	defer sensorFile.Close()

	sensorFileBytes, err := ioutil.ReadAll(sensorFile)
    if err != nil {
        fmt.Println(err)
	}

	str := string(sensorFileBytes)
	fmt.Println(str)
	
    fmt.Printf("Uploaded File: %+v\n", sensorHandler.Filename)
    fmt.Printf("File Size: %+v\n", sensorHandler.Size)
	fmt.Printf("MIME Header: %+v\n", sensorHandler.Header)
	
	locationFile, locationHandler, locationErr := r.FormFile("location_data")
    if locationErr != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(locationErr)
        return
	}
	defer locationFile.Close()

	locationFileBytes, err := ioutil.ReadAll(locationFile)
    if err != nil {
        fmt.Println(err)
	}

	s := string(locationFileBytes)
	fmt.Println(s)

	fmt.Printf("Uploaded File: %+v\n", locationHandler.Filename)
    fmt.Printf("File Size: %+v\n", locationHandler.Size)
	fmt.Printf("MIME Header: %+v\n", locationHandler.Header)
	
    fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func GetPath(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(paths); err != nil {
		panic(err)
	}

	
}