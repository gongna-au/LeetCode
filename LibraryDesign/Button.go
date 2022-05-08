package main

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"runtime"

	"strings"
	//"strconv"
	//"github.com/gin-vue-admin/server/model/system/response"
)

// Observable 被观察者
type Observable interface {
	Attach(observer ...ObserverInterface) Observable
	Detach(observer ObserverInterface) Observable
	Notify() error
}

// ObservableConcrete 一个具体的 订单状态变化的被观察者
type ObservableConcrete struct {
	observerList []ObserverInterface
}

// ObserverInterface 定义一个观察者的接口
type ObserverInterface interface {
	// 自身的业务
	Do(o Observable) error
}

//-----------------------Observable-------------------------
// Attach 注册观察者
// @param $observer ObserverInterface 观察者列表
// Attach 注册观察者
// @param $observer ObserverInterface 观察者列表
func (o *ObservableConcrete) Attach(observer ...ObserverInterface) Observable {
	o.observerList = append(o.observerList, observer...)
	return o
}

// Detach 注销观察者
// @param $observer ObserverInterface 待注销的观察者
func (o *ObservableConcrete) Detach(observer ObserverInterface) Observable {
	if len(o.observerList) == 0 {
		return o
	}
	for k, observerItem := range o.observerList {
		if observer == observerItem {
			fmt.Println(runFuncName(), "注销:", reflect.TypeOf(observer))
			o.observerList = append(o.observerList[:k], o.observerList[k+1:]...)
		}
	}
	return o
}

// Notify 通知观察者
func (o *ObservableConcrete) Notify() (err error) {
	// code ...
	for _, observer := range o.observerList {
		if err = observer.Do(o); err != nil {
			return err
		}
	}
	return nil
}

func ErrRespond(str string, err error) error {

	if err != nil {
		fmt.Println(str, "\tfunction Error")
		fmt.Printf("err=%v", err)
		return err
	} else {
		return nil
	}
	return err

}

//---------------------------------------------------------------------------------

//-----------------------------ObserverInterface---------------------
// OrderSt
type GeneralUserCreate struct {
	//用户类别  1 为超级用户 其他为普通用户
	Category           string
	DetailsInformation map[string]interface{}
}

//输入信息
func (u *GeneralUserCreate) Input() error {
	u.DetailsInformation = make(map[string]interface{})
	u.Category = "genera"
	var temp string
	fmt.Println("Please input userId:")
	fmt.Scan(&temp)
	_, err := regexp.MatchString("^[0-1]{1}$", temp)
	if err != nil {
		return errors.New("wrong happend in input user userId:")
	} else {
		u.DetailsInformation["userId"] = temp

	}
	fmt.Println("Please input user password:")
	fmt.Scan(&temp)
	_, err = regexp.MatchString("^[0-9a-zA-Z]{6,30}$", temp)
	if err != nil {
		return errors.New("wrong happend in input user password password length 6~30")
	}

	u.DetailsInformation["password"] = temp
	return nil

}

//插入数据库
func (u *GeneralUserCreate) Create() (err error) {
	//把数据和model绑定
	fmt.Println(" New user has been created!")
	return nil
}

// Do 具体业务
func (observer *GeneralUserCreate) Do(o Observable) (err error) {
	WaitingForLegalInput(observer)
	observer.Create()
	//observer输入完毕
	fmt.Println(runFuncName(), "新增用户相关的操作处理完毕...")
	return nil
}

//---------------------------------------------------------
type UserDelete struct {
	//用户类别  true 为超级用户 其他为普通用户
	userId   string
	password string
}

func (u *UserDelete) Input() (err error) {

	fmt.Println("Please input user Id:")
	fmt.Scan(&u.userId)
	_, err = regexp.MatchString("^[0-1]{10}$", u.userId)
	if err != nil {
		return errors.New("wrong happend in input user userId")
	}

	fmt.Println("Please input user password:")
	fmt.Scan(&u.password)
	_, err = regexp.MatchString("^[0-9a-zA-Z]{6,30}$", u.password)
	if err != nil {
		return errors.New("wrong happend in input user password")
	}

	return nil

}
func (u *UserDelete) Verify() (err error) {
	fmt.Println("Verify.....")
	return err
}
func (u *UserDelete) Search() (err error) {
	fmt.Println("User searching in DB...")

	return nil
}
func (u *UserDelete) Delete() (err error) {

	fmt.Println("User has been deleted!")
	return nil
}

// Do 具体业务
func (observer *UserDelete) Do(o Observable) (err error) {
	WaitingForLegalInput(observer)
	observer.Search()
	observer.Verify()
	observer.Delete()
	fmt.Println(runFuncName(), "删除用户相关的操作处理完毕...")
	return nil
}

// BookCreate-------------------------------------------------------
type BookCreate struct {
	DetailsInformation map[string]interface{}
}

/*
00马克思列宁主义、毛泽东思想
10哲学
20社会科学
21历史、历史学
27经济、经济学
31政治、社会生活
34法律、法学
36军事、军事学
37文化、科学、教育、体育
41语言、文字学
42文学
48艺术
49无神论、宗教学
50自然科学
51数学
52力学
53物理学
54化学
55天文学
56地质、地理科学
58生物科学
61医药、卫生
65农业科学
71技术科学
90综合性图书
*/
func (b *BookCreate) Input() (err error) {
	b.DetailsInformation = make(map[string]interface{})

	prompt := NewPrompt()
	prompt.Show("bookClassification")
	classificationMap := map[string]bool{"00": true, "10": true, "20": true, "21": true, "27": true, "31": true, "34": true, "36": true, "37": true, "41": true, "42": true, "48": true, "49": true, "50": true, "51": true, "52": true, "53": true, "54": true, "55": true, "56": true, "58": true, "61": true, "65": true, "71": true, "90": true}

	var temp string
	fmt.Print("Please input book classification:")
	fmt.Scanln(&temp)
	if classificationMap[temp] != true {
		return errors.New("wrong happend in input")
	} else {
		b.DetailsInformation["classification"] = temp
	}

	fmt.Print("Please input Book Author：")
	fmt.Scanln(&temp)
	b.DetailsInformation["bookAuthor"] = temp

	fmt.Print("Please input Book Price:：")
	fmt.Scanln(&temp)
	b.DetailsInformation["bookPrice"] = temp

	fmt.Print("Please enter whether to put on the shelf (Y|N）or(y|n)[ Y-Yes| N-No ]")
	fmt.Scanln(&temp)

	if strings.ToLower(string(temp[0])) == "y" {
		b.DetailsInformation["bookState"] = "y"
	} else if strings.ToLower(temp) == "n" {
		b.DetailsInformation["bookState"] = "n"
	} else {
		return errors.New("wrong happend in input")
	}

	return err

}
func (b *BookCreate) Create() (err error) {

	fmt.Println("Book has been created!")
	return err
}

// Do 具体业务
func (observer *BookCreate) Do(o Observable) (err error) {
	// code...
	WaitingForLegalInput(observer)
	fmt.Println(observer.DetailsInformation)
	observer.Create()
	fmt.Println(runFuncName(), "新增书籍相关的操作已处理完毕...")
	return nil
}

type Prompt struct {
	prompt map[string]interface{}
}

func NewPrompt() *Prompt {
	p := &Prompt{
		make(map[string]interface{}),
	}
	p.prompt["bookClassification"] = "\nA 马克思主义、列宁主义、毛泽东思想、邓小平理论\nB 哲学、宗教\nC 社会科学总论\nD	政治、法律\nE	军事\nF 经济\nG	文化、科学、教育、体育\nH	语言、文字\nI	文学\nJ 艺术\nK	历史、地理\nN	自然科学总论\nO 数理科学和化学\nP 天文学、地球科学\nQ	生物科学\nR 医药、卫生\nS 农业科学\nT	工业技术\nU 交通运输\nV	航空、航天\nX	环境科学、安全科学\nZ	综合性图书\n"
	p.prompt["bookClassificationMap"] = map[string]string{"A": "马克思主义、列宁主义、毛泽东思想、邓小平理论", "B": "哲学、宗教", "C": "社会科学总论", "D": "政治、法律", "E": "军事", "F": "经济", "G": "文化、科学、教育、体育", "H": "语言、文字", "I": "文学", "J": "艺术", "K": "历史、地理", "N": "自然科学总论", "O": "数理科学和化学", "P": "天文学、地球科学", "Q": "生物科学", "R": "医药、卫生", "S": "农业科学", "T": "工业技术", "U": "交通运输", "V": "航空、航天", "X": "环境科学、安全科学", "Z": "综合性图书"}
	//其他的提示信息
	p.prompt["bookSearchInfo"] = "\n1.bookName书籍名查找\n2.bookKey其他有关书籍的信息\n3.bookAuthor书籍作者\n4.bookclassification书籍分类查找\n5.bookState根据书籍的借阅状态查\n"
	p.prompt["bookSearchInfoMap"] = map[string]string{"1": "bookName", "2": "bookKey", "3": "bookAuthor", "4": "bookclassification", "5": "bookState"}
	return p
}
func (p *Prompt) Show(str string) {
	//classificationMap := map[string]bool{"00": true, "10": true, "20": true, "21": true, "27": true, "31": true, "34": true, "36": true, "37": true, "41": true, "42": true, "48": true, "49": true, "50": true, "51": true, "52": true, "53": true, "54": true, "55": true, "56": true, "58": true, "61": true, "65": true, "71": true, "90": true}
	fmt.Println(p.prompt[str])

}

type BookSearch struct {
	target string
	key    map[string]interface{}
}

func (b *BookSearch) Input() (err error) {
	p := NewPrompt()
	p.Show("bookSearchInfo")
	fmt.Print("Please input your choice:")
	var temp string
	fmt.Scanln(&temp)
	switch k := p.prompt["bookSearchInfoMap"].(type) {
	case map[string]string:
		v, ok := k[temp]

		if ok {
			b.target = v
			return nil
		} else {
			errors.New("Input illegal!")
		}
	}
	err = b.Search()
	if err != nil {
		return errors.New("Search Error!")
	} else {
		return nil
	}

}

func NewBookSerach() *BookSearch {
	b := &BookSearch{
		"",
		make(map[string]interface{}),
	}
	b.key["bookName"] = ""
	//关键字查找
	b.key["bookKey"] = ""
	//作者查
	b.key["bookAuthor"] = ""
	//书籍分类查找
	b.key["bookclassification"] = ""
	//根据书籍的借阅状态查
	b.key["bookState"] = ""

	return b

}
func (b *BookSearch) Search() (err error) {
	fmt.Println("Please input information about ", b.target, ":")
	var temp string
	fmt.Scanf("%30s", &temp)
	b.key[b.target] = temp
	fmt.Println("Searching in DB...")
	fmt.Println("Searching in DB success!")
	//打印查询结果
	return nil

}

// Do 具体业务
func (observer *BookSearch) Do(o Observable) (err error) {
	// code..
	observer = NewBookSerach()
	WaitingForLegalInput(observer)
	observer.Search()
	fmt.Println(runFuncName(), "查询书籍相关的操作已经处理完毕...")
	return
}

//---------------------------------------------------------
type Result struct {
	response string
	result   []interface{}
}

func NewResult() *Result {

	return &Result{}

}

func (b *Result) Show() (err error) {

	fmt.Println("Result are: ", b.result)
	return err
}

// Do 具体业务
func (observer *Result) Do(o Observable) (err error) {
	// code..
	observer = NewResult()
	//存储数据库信息
	fmt.Println("Result Stored")
	observer.Show()
	fmt.Println(runFuncName(), "展示相关结果.已经处理完毕...")
	return nil
}

// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

//----------------------------------------另外的设计

type Button struct {
	information string
	bePressed   int
}

func newButton(info string) *Button {
	return &Button{
		info,
		0,
	}
}
func (b *Button) showPageInformation() {
	fmt.Println(b.information)
}

func (b *Button) listenButtonBePressed() {
	fmt.Scan(&b.bePressed)
	fmt.Println(b.bePressed, "button is  pressed.....")

}

//--------------------------------------------

type UserLogin struct {
	userId   string
	password string
}

func (u *UserLogin) Login() error {
	fmt.Println("User login success!")
	return nil
}
func (u *UserLogin) Input() (err error) {
	fmt.Println("Please input user Id:")
	fmt.Scan(&u.userId)
	_, err = regexp.MatchString("^[0-1]{10}$", u.userId)
	if err != nil {
		return errors.New("wrong happend in input user userId")
	}

	fmt.Println("Please input user password:")
	fmt.Scan(&u.password)
	_, err = regexp.MatchString("^[0-9a-zA-Z]{6,30}$", u.password)
	if err != nil {
		return errors.New("wrong happend in input user password")
	}

	return nil

}

func (u *UserLogin) Search() (err error) {
	fmt.Println("User password and account search success!")
	return nil
}

// Do 具体业务
func (observer *UserLogin) Do(o Observable) (err error) {
	// code...
	WaitingForLegalInput(observer)
	observer.Login()
	fmt.Println(runFuncName(), "用户登陆的相关的操作已处理完毕...")
	return nil
}

// 客户端调用
func main() {
	app := NewApp()
	app.Run()
	// 未来可以快速的根据业务的变化 创建新的主题 从而快速构建新的业务接口
	fmt.Println("----------------------- 未来的扩展...")

}

//------------------------其他的函数

type Inputinterface interface {
	Input() error
}

func WaitingForLegalInput(w Inputinterface) {
	for {
		err := w.Input()
		if err == nil {
			break
		} else {
			fmt.Println("Illegal input, Please input again!")
		}

	}

}

type ButtonFunc func()

type App struct {
	funcList []ButtonFunc
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	a.funcList = append(a.funcList, Test1())
	for i, v := range a.funcList {
		fmt.Println("第 ", i, " 次执行") /*  */
		v()
		a.funcList = append(a.funcList, Test1())
	}
}

func Test1() ButtonFunc {
	breakTag := false
	for {
		if breakTag == true {
			break
		}
		button := newButton("\n>>1.用户注册\n>>2.用户登陆\n>>3.书籍查询\n")
		button.showPageInformation()
		button.listenButtonBePressed()
		switch button.bePressed {
		case 1:
			{
				// 创建 新增用户 “主题”
				fmt.Println("----------------------- 新增用户 “主题”")
				orderUnPaidCancelSubject := &ObservableConcrete{}
				orderUnPaidCancelSubject.Attach(
					&GeneralUserCreate{},
				)
				orderUnPaidCancelSubject.Notify()
				breakTag = true
				return Test1()

			}
		case 2:
			{
				// 创建 新增书籍 “主题”
				fmt.Println("----------------------- 用户登陆 “主题”")
				orderOverTimeSubject := &ObservableConcrete{}
				orderOverTimeSubject.Attach(
					&UserLogin{},
				)
				orderOverTimeSubject.Notify()
				breakTag = true
				return Test2()
			}
		case 3:
			{
				// 创建 查询结果的主题
				fmt.Println("----------------------- 新增查询书籍 “主题”")
				orderPaidCancelSubject := &ObservableConcrete{}
				orderPaidCancelSubject.Attach(
					&BookSearch{},
				)
				orderPaidCancelSubject.Notify()
				breakTag = true
				//返回下一次要执行的函数
				return Test1()

			}
		default:
			{
				fmt.Println("Illegal input, Please input again!")
				breakTag = false

			}

		}

	}
	return nil
}

func Test2() ButtonFunc {
	button := newButton("\n>>1.借书\n>>2.还书\n>>3.查询书籍\n>>4.查看借阅记录")
	button.showPageInformation()
	button.listenButtonBePressed()
	switch button.bePressed {
	case 1:
		{
			// 创建 新增用户 “主题”
			fmt.Println("----------------------- 用户借书 “主题”")
			orderUnPaidCancelSubject := &ObservableConcrete{}
			orderUnPaidCancelSubject.Attach(
				&GeneralUserCreate{},
			)
			orderUnPaidCancelSubject.Notify()
			fmt.Println("----------------------- 用户借书操作完毕")

			return Test2()
		}
	case 2:
		{
			// 创建 新增书籍 “主题”
			fmt.Println("----------------------- 用户还书 “主题”")
			orderOverTimeSubject := &ObservableConcrete{}
			orderOverTimeSubject.Attach(
				&UserLogin{},
			)
			orderOverTimeSubject.Notify()
			fmt.Println("----------------------- 用户还书操作完毕")
			return Test2()
		}
	case 3:
		{
			// 创建 查询结果的主题
			fmt.Println("----------------------- 查询书籍 “主题”")
			orderPaidCancelSubject := &ObservableConcrete{}
			orderPaidCancelSubject.Attach(
				&BookSearch{},
			)
			orderPaidCancelSubject.Notify()
			fmt.Println("----------------------- 用户查询书籍操作完毕")
			return Test2()
		}
	case 4:
		{
			// 创建 查询结果的主题
			fmt.Println("----------------------- 查询借阅记录 “主题”")
			orderPaidCancelSubject := &ObservableConcrete{}
			orderPaidCancelSubject.Attach(
				&BookSearch{},
			)
			orderPaidCancelSubject.Notify()
			fmt.Println("----------------------- 用户查询借阅记录操作完毕")
			return Test2()
		}

	}
	return nil

}
