package indexer

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/blevesearch/bleve"
	"github.com/dgplug/dexer/lib/conf"
	"github.com/dgplug/dexer/lib/logger"
	"github.com/dgplug/dexer/lib/watcher"
)

// FileIndexer is a data structure to hold the content of the file
type FileIndexer struct {
	FileName    string
	FileContent string
}

type FileIndexerArray struct {
	IndexerArray []FileIndexer
	fLogger      *logger.Logger
}

func (farray *FileIndexerArray) Must(e error, logstring string) {
	farray.fLogger.Must(e, logstring)
}

// Search function is used to search the string in the file and return the index
func Search(indexFilename string, searchWord string) *bleve.SearchResult {
	// opens the index file using bleve
	index, _ := bleve.Open(indexFilename)
	// closes file after the function completes its execution
	defer index.Close()
	// makes query to search the string
	query := bleve.NewQueryStringQuery(searchWord)
	request := bleve.NewSearchRequest(query)
	// matches the keyword if any from the index created and returns
	result, _ := index.Search(request)
	return result
}

// fileIndexing function is used to create an index
func fileIndexing(fileIndexer FileIndexerArray, c conf.Configuration) {
	// check if previously index exists and delete if present
	err := DeleteExistingIndex(c.IndexFilename)
	fileIndexer.Must(err, "Successfully deleted previous index")
	// maps new index
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(c.IndexFilename, mapping)
	fileIndexer.Must(err, "Successfully ran bleve for indexing")
	// updates the value of index in the IndexerArray
	for _, fileIndex := range fileIndexer.IndexerArray {
		index.Index(fileIndex.FileName, fileIndex.FileContent)
	}
	defer index.Close()
}

// fileNameContentMap function populates the index
func fileNameContentMap(c conf.Configuration) FileIndexerArray {
	var root = c.RootDirectory
	var files []string
	fileIndexer := FileIndexerArray{
		fLogger: c.GetLogger(),
	}
	// visits each file starting from root directory and adds path to files array
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	// traverses each file and adds index and returns it
	fileIndexer.Must(err, "Successfully traversed "+root)
	for _, filename := range files {
		content, err := GetContent(filename)
		fileIndexer.Must(err, "Successfully obtained content from "+filename)
		filesIndex := NewFileIndexer(filename, content)
		fileIndexer.IndexerArray = append(fileIndexer.IndexerArray, filesIndex)
	}
	return fileIndexer
}

// NewFileIndexer is a function to create a new File Indexer
func NewFileIndexer(fname, fcontent string) FileIndexer {
	temp := FileIndexer{
		FileName:    fname,
		FileContent: fcontent,
	}

	return temp
}

// NewIndex is a function to create new indexes
func NewIndex(c conf.Configuration) {

	fileIndexer := fileNameContentMap(c)
	fileIndexing(fileIndexer, c)

	c.Must(nil, "Refreshing the index")
	// adds a watcher w which watches files change
	w := watcher.NewWatcher(c, time.Millisecond*100)

	action := func(path string, status watcher.FileStatus) {
		fileIndexer := fileNameContentMap(c)
		fileIndexing(fileIndexer, c)
		fmt.Println("yo")
	}

	go w.Start(action)
}

// GetContent is a function for retrieving data from file
func GetContent(name string) (string, error) {
	data, err := ioutil.ReadFile(name)
	return string(data), err
}

// DeleteExistingIndex checks if the index exist if it does, then flushes it off
func DeleteExistingIndex(name string) error {
	_, err := os.Stat(name)
	if !os.IsNotExist(err) {
		if err := os.RemoveAll(name); err != nil {
			return fmt.Errorf("Can't Delete file: %v", err)
		}
	}
	return nil
}
