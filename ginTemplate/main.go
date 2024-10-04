package main

import (
	"html/template"
	"net/http"
)

type Student struct {
	Name string
	Age  int
}

// 自定义模板函数
func add(a, b int) int {
	return a + b
}

func f1(w http.ResponseWriter, req *http.Request) {
	//html, err := template.ParseFiles("./template/test.html")
	//if err != nil {
	//	fmt.Println(err)
	//}

	// .代表传递的数据对象
	//html.Execute(w, "yangyang")

	funcMap := template.FuncMap{
		"add": add,
	}

	html, _ := template.New("test.html").Funcs(funcMap).ParseFiles("./template/test.html")

	data := make(map[string]interface{})

	mapData := map[string]interface{}{"name": "yangyang", "age": 18, "gender": "mal"}

	stu := Student{"yang", 18}

	data["student"] = stu
	data["mapData"] = mapData

	stu1 := Student{Name: "qq"}
	stu2 := Student{"ww", 18}
	stu3 := Student{"rr", 90}
	stu4 := Student{"ee", 70}
	stu5 := Student{Name: "kk"}
	stu6 := Student{Name: "ll"}

	stuList := []Student{stu1, stu2, stu3, stu4, stu5, stu6}

	data["stuList"] = stuList

	// .后面添加想要取出的数据
	html.Execute(w, data)
}

func main() {

	http.HandleFunc("/test", f1)

	http.ListenAndServe(":8080", nil)

}
