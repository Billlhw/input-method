package loader

import (
	"bufio"
	"fmt"
	"input_method/trie"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type WorkerPool struct {
	TaskCh chan func() // channel of type func
	num    int         // number of concurrent workers
}

// Start worker goroutines
func StartWorker(p *WorkerPool) {
	for i := 0; i < p.num; i++ {
		go func() {
			for task := range p.TaskCh {
				task()
			}
		}()
	}
}

// Add task to worker pool
func (p *WorkerPool) AddTask(f func()) {
	p.TaskCh <- f
}

func (p *WorkerPool) ProcessTasks(rootNode *trie.TrieNode) {
	//read files from the data directory
	dataDirPath := "data"
	entries, err := os.ReadDir(dataDirPath)
	if err != nil {
		fmt.Println("Error: cannot open data directory.")
		return
	}

	StartWorker(p)
	var wg sync.WaitGroup

	for _, entry := range entries {
		entryPath := filepath.Join(dataDirPath, entry.Name())
		info, err := os.Stat(entryPath)
		if err != nil || info.IsDir() {
			fmt.Println("Error: invalid file type in data directory", entryPath)
			continue
		}
		wg.Add(1)
		taskFunc := func(fileName string) func() {
			return func() {
				defer wg.Done()
				processFile(fileName, rootNode)
			}
		}
		p.AddTask(taskFunc(entry.Name())) //send to worker
	}

	wg.Wait()
}

// read file content, and add valid entries to the trie tree
func processFile(fileName string, rootNode *trie.TrieNode) {
	filePath := filepath.Join("data", fileName)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error: cannot open file", filePath, ":", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineElem := strings.Split(line, " ")
		if len(lineElem) != 2 {
			fmt.Println("Error: line is not properly formatted: ", line)
			continue
		}
		weight, err := strconv.Atoi(lineElem[1])
		if err != nil {
			fmt.Println("Error: weight value is not an integer: ", line)
			continue
		}

		//fmt.Printf("Processing character %s, weight %d\n", lineElem[0], weight)
		spelling := strings.TrimSuffix(fileName, ".txt")
		rootNode.Insert(spelling, []rune(lineElem[0])[0], weight)
	}
}

// create a trie tree, and insert data
func StartLoader() *trie.TrieNode {
	p := &WorkerPool{
		TaskCh: make(chan func()),
		num:    10,
	}
	trieRoot := trie.NewTrieNode()
	p.ProcessTasks(trieRoot)
	return trieRoot
}
