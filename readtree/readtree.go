package readtree

import (
	"io/fs"
	"log"
	"mgit/catfile"
	"mgit/constants"
	"mgit/models"
	"mgit/utils"
	"os"
	"path/filepath"
	"strings"
)

func ReadTree(treeOid string) {
	err := cleanupFilesAndDirs(constants.DefaultFolderPath)
	utils.CheckErr(err)
	for path, oid := range getTree(treeOid, constants.DefaultFolderPath) {
		dir := filepath.Dir(path)
		err := os.MkdirAll(dir, os.ModePerm)
		utils.CheckErr(err)
		log.Println("Creating blob: " + path)
		utils.CreateFile(path, []byte(catfile.CatFile(oid, "")))
	}
}

func getEntries(oid string) []models.TreeEntry {
	log.Println("Getting entries")
	var stringEntries = catfile.CatFile(oid, "tree")
	lines := strings.Split(stringEntries, "\n")
	var entries []models.TreeEntry
	var parts []string
	//fmt.Println(lines)
	for _, line := range lines[0 : len(lines)-1] {
		parts = strings.Fields(line)
		//fmt.Println(parts)
		entries = append(entries, models.TreeEntry{ObjectType: parts[0], Oid: parts[1], Name: parts[2]})
	}
	return entries
}

func getTree(oid string, basePath string) map[string]string {
	entriesPath := make(map[string]string)
	var path string
	for _, entry := range getEntries(oid) {
		path = basePath + "/" + entry.Name
		if entry.ObjectType == "blob" {
			entriesPath[path] = entry.Oid
		} else if entry.ObjectType == "tree" {
			innerTree := getTree(entry.Oid, path)
			for k, v := range innerTree {
				entriesPath[k] = v
			}
		} else {
			log.Fatal("Unknown tree entry: ", entry)
		}
	}
	return entriesPath
}

func cleanupFilesAndDirs(root string) error {
	var files []string
	var dirs []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(root, path)
		if relPath == "." {
			return nil
		}

		if d.IsDir() {
			dirs = append(dirs, path)
		} else {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	for _, file := range files {
		relPath, _ := filepath.Rel(root, file)
		if utils.IsIgnoredPath(relPath) {
			continue
		}
		if info, err := os.Stat(file); err == nil && !info.IsDir() {
			_ = os.Remove(file)
		}
	}

	for i := len(dirs) - 1; i >= 0; i-- {
		dir := dirs[i]
		relPath, _ := filepath.Rel(root, dir)
		if utils.IsIgnoredPath(relPath) {
			continue
		}
		_ = os.Remove(dir)
	}
	return nil
}
