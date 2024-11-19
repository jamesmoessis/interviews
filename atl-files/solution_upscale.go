package main

import "sync"

/**

In this extension we propose that instead of receiving the total list of files as an input,
we instead receive “file events” and want to maintain our report in real-time so it can be queried by a web front-end.

These “file events” are:
File added
File deleted

Candidates may tackle this in a couple of ways:
Maintain an event store of incoming events and generate a report at query time; or
Maintain a report state and update it as events come in

Ideally the candidate recognizes these options and can discuss the pros and cons of each.
Which one they choose may be influenced by their existing implementation.
*/

type fileManager struct {
	sync.Mutex
}

func (fm *fileManager) addFile(f *file) {
	fm.Lock()
	defer fm.Unlock()
}

func (fm *fileManager) deleteFile(name string) {
	fm.Lock()
	defer fm.Unlock()

}

func (fm *fileManager) getReport() *report {
	return &report{}
}
