package client

import (
    "encoding/json"
    "fmt"
    "github.com/d3code/zlog"
    "net/http"
    "os"
    "sort"
    "strings"
    "time"
)

type RequestEntry struct {
    Host      string    `json:"host"`
    Path      string    `json:"path"`
    Params    string    `json:"params"`
    Headers   string    `json:"headers"`
    Method    string    `json:"method"`
    Timestamp time.Time `json:"timestamp"`
    File      string    `json:"file"`
}

func getHeaders(req *http.Request, removeHeader []string) string {
    filteredHeaders := req.Header.Clone()
    for _, y := range removeHeader {
        filteredHeaders.Del(y)
    }
    formattedHeaders := fmt.Sprintf("%v", filteredHeaders)
    return formattedHeaders
}

func getParams(req *http.Request, removeQuery []string) string {
    copiedParams := req.URL.Query()
    for _, y := range removeQuery {
        copiedParams.Del(y)
    }
    params := fmt.Sprintf("%v", copiedParams)
    return params
}

func getIndex(basePath string) []RequestEntry {

    // Create base path if not exist
    err := os.MkdirAll(basePath+"/response", 0755)
    if err != nil {
        zlog.Log.Errorf("Error creating path: %s", err.Error())
        return make([]RequestEntry, 0)
    }

    // Create index if not exist
    index := basePath + "/index.json"
    if _, existError := os.Stat(index); os.IsNotExist(existError) {

        newIndex, marshalError := json.Marshal(make([]RequestEntry, 0))
        if marshalError != nil {
            zlog.Log.Errorf("Marshalling new index: %s", marshalError.Error())
            return make([]RequestEntry, 0)
        }

        saveFile(basePath, "/index.json", newIndex)
    }

    // Return index from path
    indexByteArray, err := os.ReadFile(index)
    if err != nil {
        zlog.Log.Errorf("Reading index %s", err.Error())
        return make([]RequestEntry, 0)
    }

    var fileIndex []RequestEntry
    unmarshalError := json.Unmarshal(indexByteArray, &fileIndex)
    if unmarshalError != nil {
        zlog.Log.Errorf("Unmarshalling index %s", unmarshalError.Error())
        return make([]RequestEntry, 0)
    }

    sort.Slice(fileIndex, func(i, j int) bool {
        return fileIndex[i].Timestamp.Before(fileIndex[j].Timestamp)
    })

    return fileIndex
}

func saveFile(path string, fileName string, byteArray []byte) {
    pathFile := path + "/" + fileName

    // Normalize path
    file := strings.Replace(pathFile, "//", "/", -1)
    //log.Log.Debugf("Saving file %s", file)

    // Create path if not exist
    err := os.MkdirAll(path, 0755)
    if err != nil {
        zlog.Log.Errorf("Error creating path: %s", err.Error())
        return
    }

    // Write file
    err = os.WriteFile(file, byteArray, 0666)
    if err != nil {
        zlog.Log.Errorf("Error writing file: %s", err.Error())
    }
}
