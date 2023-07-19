package conf

import (
	"fmt"
	"os"
	"path"
)

type LoadCB func(e map[string]string, cover bool)

// Load loads data from a file and parses it based on the file extension(.json/.env/.yml/.yaml).
//
// It takes a LoadCB callback function, a boolean flag to determine if the data should be covered, and a string representing the file path.
// It returns an error if there is any issue reading or parsing the file.
func Load(cb LoadCB, cover bool, fn string) error {
	data, err := os.ReadFile(fn)
	if err != nil {
		return err
	}
	var items map[string]string
	switch path.Ext(fn) {
	case ".env":
		items = ParseDotENV(string(data))
	case ".json":
		items = ParseJson(string(data))
	case ".yml", ".yaml":
		items = ParseYML(string(data))
	default:
		return fmt.Errorf("unknow ext:%s", path.Ext(fn))
	}
	cb(items, true)
	return nil
}

// LoadFiles loads files and applies them to a callback function.
//
// The function takes a callback function `cb` of type `LoadCB` which is called with the loaded files as a map[string]string and the `cover` boolean value. The `cover` parameter determines whether the loaded files should overwrite existing values in the map.
// The `files` parameter is a variadic argument that accepts a list of file names to load.
// The function does not return anything.
func LoadFiles(cb LoadCB, cover bool, files ...string) {
	out := make(map[string]string)
	for _, fn := range files {
		Load(func(e map[string]string, cover bool) {
			for k, v := range e {
				if _, ok := out[k]; !ok || cover {
					out[k] = v
				}
			}
		}, cover, fn)
	}
	cb(out, cover)
}

// LoadAnyOne loads any one of the given files and returns the file name that was successfully loaded.
//
// Parameters:
// - cb: A callback function that handles the loaded file.
// - cover: A boolean flag indicating whether to cover the existing data or not.
// - files: A variadic parameter representing the file names to be loaded.
//
// Return:
// - string: The file name that was successfully loaded. If no file was successfully loaded, an empty string is returned.
func LoadAnyOne(cb LoadCB, cover bool, files ...string) string {
	for _, fn := range files {
		err := Load(cb, cover, fn)
		if err == nil {
			return fn
		}
	}
	return ""
}

// LoadOneToENV loads one config file into the environment variables.
//
// If no files are provided, it loads the default set of files: ".env", "conf.json", "config.json", "conf.yml", and "config.yml".
//
// The function returns a string representing the result of loading the files into the environment variables.
func LoadOneToENV(files ...string) string {
	if len(files) == 0 {
		files = append(files, ".env", "conf.json", "config.json", "conf.yml", "config.yml")
	}
	return LoadAnyOne(func(e map[string]string, cover bool) {
		SaveToENV(e, true)
	}, true, files...)
}
