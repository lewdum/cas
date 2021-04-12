package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Key")
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	// info(r, "Starting response...")
	// defer info(r, "Finished response.")

	enableCORS(w)

	switch r.Method {
	default:
		fail(w, r, http.StatusMethodNotAllowed, "Invalid HTTP Method")

	case http.MethodGet:
		key := strings.TrimPrefix(r.URL.Path, "/")
		if key == "" {
			fail(w, r, http.StatusBadRequest, "Missing Key")
			return
		}

		info(r, fmt.Sprintf("Requested file: %s", key))

		base := filepath.Join(*flagStorage, key)
		path := base + ".data"
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			fail(w, r, http.StatusNotFound, "Not Found")
			return
		}

		mutex.RLock()
		defer mutex.RUnlock()

		file, err := os.Open(path)
		if err != nil {
			failInternal(w, r, err)
			return
		}
		defer file.Close()

		w.Header().Add("Content-Type", "application/octet-stream")

		_, err = io.Copy(w, file)
		if err != nil {
			failInternal(w, r, err)
			return
		}

	case http.MethodPost:
		source, fileInfo, err := r.FormFile("data")
		if err != nil {
			fail(w, r, http.StatusBadRequest, "Missing File")
			return
		}

		info(r, fmt.Sprintf("Posting file: %s", fileInfo.Filename))

		// TODO
		// if info.Size > ... {
		// }

		data, err := io.ReadAll(source)
		if err != nil {
			failInternal(w, r, err)
			return
		}

		meta := Metadata{
			Name:     fileInfo.Filename,
			Size:     fileInfo.Size,
			Uploader: r.RemoteAddr,
			Uploaded: time.Now().UTC(), // Time saved in UTC.
		}

		key := hashFile(data)
		base := filepath.Join(*flagStorage, key)
		path := base + ".data"
		if _, err := os.Stat(path); err == nil {
			w.Header().Add("Key", key)
			fail(w, r, http.StatusConflict, "Conflict")
			return
		}

		mutex.Lock()
		defer mutex.Unlock()

		file, err := os.Create(path)
		if err != nil {
			failInternal(w, r, err)
			return
		}
		defer file.Close()

		_, err = file.Write(data)
		if err != nil {
			failInternal(w, r, err)
			return
		}

		metaPath := base + ".meta"
		metaFile, err := os.Create(metaPath)
		if err != nil {
			failInternal(w, r, err)
			return
		}
		defer metaFile.Close()

		enc := json.NewEncoder(metaFile)
		// jsonEnc.SetIndent("", "    ")
		err = enc.Encode(meta)
		if err != nil {
			failInternal(w, r, err)
			return
		}

		info(r, fmt.Sprintf("Created file: %s", key))

		w.Write([]byte(key))

	}
}

func handleList(w http.ResponseWriter, r *http.Request) {
	// info(r, "Starting response...")
	// defer info(r, "Finished response.")

	enableCORS(w)

	switch r.Method {
	default:
		fail(w, r, http.StatusMethodNotAllowed, "Invalid HTTP Method")

	case http.MethodGet:
		info(r, "Requested index.")

		files := make([]string, 0, 256)

		mutex.RLock()
		defer mutex.RUnlock()

		entries, err := os.ReadDir(*flagStorage)
		if err != nil {
			failInternal(w, r, err)
			return
		}

		for _, entry := range entries {
			name := entry.Name()
			ext := filepath.Ext(name)
			if ext != ".data" {
				continue
			}
			files = append(files, strings.TrimSuffix(name, ext))
		}

		sort.Strings(files)
		joined := strings.Join(files, "\n")
		w.Write([]byte(joined))

	}
}
