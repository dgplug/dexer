package indexer

import (
	"os"
	"path/filepath"

	"github.com/blevesearch/bleve"
	"github.com/farhaanbukhsh/file-indexer/conf"
	"github.com/farhaanbukhsh/file-indexer/logger"
	"github.com/farhaanbukhsh/file-indexer/utility"
)

// FileIndexer is a data structure to hold the content of the file
type FileIndexer struct {
	FileName    string
	FileContent string
}

type FileIndexerArray struct {
	IndexerArray    []FileIndexer
	FileIndexLogger *logger.Logger
}

func Search(indexFilename string, searchWord string) *bleve.SearchResult {
	index, _ := bleve.Open(indexFilename)
	defer index.Close()
	query := bleve.NewQueryStringQuery(searchWord)
	request := bleve.NewSearchRequest(query)
	result, _ := index.Search(request)
	return result
}

func fileIndexing(fileIndexer FileIndexerArray, c conf.Configuration) error {
	err := utility.DeleteExistingIndex(c.IndexFilename)
	if err != nil {
		return err
	}
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(c.IndexFilename, mapping)
	if err != nil {
		return err
	}
	for _, fileIndex := range fileIndexer.IndexerArray {
		index.Index(fileIndex.FileName, fileIndex.FileContent)
	}
	defer index.Close()
	return nil
}

func fileNameContentMap(c conf.Configuration, lg *logger.Logger) FileIndexerArray {
	var root = c.RootDirectory
	var files []string
	fileIndexer := FileIndexerArray{
		FileIndexLogger: lg,
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	fileIndexer.FileIndexLogger.Must(err, "Successfully traversed "+root)
	for _, filename := range files {
		content, err := utility.GetContent(filename)
		fileIndexer.FileIndexLogger.Must(err, "Successfully obtained content from "+filename)
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
func NewIndex(c conf.Configuration, lg *logger.Logger) error {
	fileIndexer := fileNameContentMap(c, lg)
	err := fileIndexing(fileIndexer, c)
	if err != nil {
		return err
	}
	return nil
}
