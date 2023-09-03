package task

import "time"

type Task struct {
	Username 	string
	ActName  	string
	TaskName     	string
	Link     	string
	DateTime	time.Time	
}