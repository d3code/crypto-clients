package client

import (
    "encoding/json"
    "github.com/d3code/zlog"
    "net/http"
    "os"
)

func DoRequest(request *http.Request, replay bool, basePath string, removeHeader []string, removeQueryParam []string) (*Response, error) {
    if basePath == "" {
        return getResponse(request)
    }

    requestEntry := getRequestEntry(request, removeQueryParam, removeHeader)

    // File paths
    basePath = basePath + "/" + requestEntry.Host
    requestEntryPath := basePath + "/response" + requestEntry.Path

    index := getIndex(basePath)

    for _, indexEntry := range index {

        if replay &&
            indexEntry.Method == request.Method &&
            indexEntry.Host == requestEntry.Host &&
            indexEntry.Path == requestEntry.Path &&
            indexEntry.Params == requestEntry.Params &&
            indexEntry.Headers == requestEntry.Headers {

            file := indexEntry.File + "_response.json"
            name := requestEntryPath + "/" + file

            if fileInfo, err := os.Stat(name); err != nil || fileInfo.IsDir() {
                zlog.Log.Warnf("File %s not found", name)
                continue
            } else {

                fileBytes, errRead := os.ReadFile(name)
                if errRead != nil {
                    zlog.Log.Errorf("Error while reading file: %v", errRead)
                    continue
                }

                var response Response
                errUnmarshalFile := json.Unmarshal(fileBytes, &response)
                if errUnmarshalFile != nil {
                    zlog.Log.Errorf("Error while unmarshalling file: %v", errUnmarshalFile)
                    return nil, errUnmarshalFile
                }

                if !(response.StatusCode >= 200 && response.StatusCode < 300) {
                    zlog.Log.Warnf("(skip) Status code %v for saved response %s", response.StatusCode, file)
                    continue
                }

                response.Stored = true
                return &response, nil
            }
        }
    }

    if replay {
        zlog.Log.Warnf("No stored response found for %s %s", request.Method, request.URL)
    }

    response, err := getResponse(request)
    if err != nil {
        return response, err
    }

    // Saving all with response details
    responseFile, errJson := json.Marshal(response)
    if errJson != nil {
        zlog.Log.Errorf("Error while marshalling response: %v", errJson)
        return response, errJson
    }

    saveFile(requestEntryPath, requestEntry.File+"_response.json", responseFile)
    saveFile(requestEntryPath, requestEntry.File+"_body.json", response.Body)

    index = append(index, requestEntry)

    indexFile, errIndex := json.Marshal(index)
    if errIndex != nil {
        zlog.Log.Errorf("Error while marshalling index: %v", errIndex)
        return nil, errIndex
    }

    saveFile(basePath, "/index.json", indexFile)

    return response, err
}
