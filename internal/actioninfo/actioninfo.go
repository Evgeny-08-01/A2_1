package actioninfo

import (
	"fmt"
	"log"
	)

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _,v:=range dataset {
		err := dp.Parse(v)
		if err != nil { log.Println(err)} 
		st,err := dp.ActionInfo()		                       
		if err != nil {log.Println(err)
			}else{	fmt.Println(st)	}	
			                 }
							}
