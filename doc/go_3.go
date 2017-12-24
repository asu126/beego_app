package main

import "fmt"

type S map[string][]string

func Summary(paramstring string) (s *S) {
	s = &S{
		"name":           []string{paramstring},
		"profession":     []string{"Javaprogrammer", "ProjectManager"},
		"interest(lang)": []string{"Clojure", "Python", "Go"},
		"focus(project)": []string{"UE", "AgileMethodology", "SoftwareEngineering"},
		"hobby(life)":    []string{"Basketball", "Movies", "Travel"},
	}
	return s
}

func main() {
	s := Summary("Harry")
	fmt.Printf("Summary(address):%v\r\n", s)
	fmt.Printf("Summary(content):%v\r\n", *s)
}

/*
Summary(address):&map[hobby(life):[Basketball Movies Travel] name:[Harry] profession:[Javaprogrammer ProjectManager] interest(lang):[Clojure Python Go] focus(project):[UE AgileMethodology SoftwareEngineering]]
Summary(content):map[name:[Harry] profession:[Javaprogrammer ProjectManager] interest(lang):[Clojure Python Go] focus(project):[UE AgileMethodology SoftwareEngineering] hobby(life):[Basketball Movies Travel]]
*/
