package proctree

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Info struct {
	Processes []Process `json:"Processes"`
	Running   int       `json:"running"`
	Sleeping  int       `json:"sleeping"`
	Zombie    int       `json:"zombie"`
	Stopped   int       `json:"stopped"`
	Total     int       `json:"total"`
}

type Process struct {
	Pid   int            `json:"pid"`
	Name  string         `json:"name"`
	User  int            `json:"user"`
	State int            `json:"state"`
	Ram   int            `json:"ram"`
	Rss   int            `json:"rss"`
	Child []ChildProcess `json:"child"`
}

type ChildProcess struct {
	Pid      int    `json:"pid"`
	Name     string `json:"name"`
	State    int    `json:"state"`
	PidPadre int    `json:"pidPadre"`
	RssChild int    `json:"rssChild"`
	ChildUID int    `json:"childUID"`
}

func GetProc(pid int) string {
	data, err := os.ReadFile("/proc/procesos")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "digraph Arbol{\"-err obteniendo proc\" [fillcolor=red style=filled];}"
	}
	procJson := "{" + string(data)

	var info Info
	err = json.Unmarshal([]byte(procJson), &info)
	if err != nil {
		log.Fatal(err)
	}

	var tree string
	flag := false
	tree = "digraph Arbol{\n"
	for _, proc := range info.Processes {
		if proc.Pid == pid {
			flag = true
			fmt.Println(proc.Pid, proc.Name)
			if len(proc.Child) > 0 {
				for _, child := range proc.Child {
					tree += fmt.Sprintf("\"%d  %s\"\n", proc.Pid, proc.Name)
					tree += fmt.Sprintf("-> \"%d %s \"\n", child.Pid, child.Name)
					fmt.Println("  ", child.Pid, child.Name)
				}
			} else {
				tree += fmt.Sprintf("\"%d  %s\"", proc.Pid, proc.Name)
			}
		}
	}

	if !flag {
		return "digraph Arbol{\"-err proc no existe\" [fillcolor=red style=filled];}"
	}
	tree += "}"
	return tree
}

func GetProcList() string {
	data, err := os.ReadFile("/proc/procesos")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "err"
	}
	procJson := "{" + string(data)

	var procList string
	var info Info
	err = json.Unmarshal([]byte(procJson), &info)
	if err != nil {
		log.Fatal(err)
	}

	for _, proc := range info.Processes {
		procList += fmt.Sprintf("%d\n", proc.Pid)
		if len(proc.Child) > 0 {
			for _, child := range proc.Child {
				procList += fmt.Sprintf("%d\n", child.Pid)
			}
		}
	}
	return procList
}
