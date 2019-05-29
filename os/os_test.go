package main

import (
	"fmt"
	"os"
	"testing"
	"errors"
)

//=============================================os base=======================================

//打开文件
func TestOsOpen(t *testing.T){
	file, err := os.Open("../test.txt")
	if err != nil {
		t.Errorf("%v",err)
		return
	}
	t.Logf("%v",file.Name())
	//open file: no such file or directory
}

//环境变量操作
func TestOsSetenv(t *testing.T){
	//设置环境变量(并没有设置到环境变量中,应该是临时的)
	err := os.Setenv("TMPDOR","D:/logtest/")
	if err != nil{
		t.Errorf("%v",err)
		return
	}

	//获取环境变量
	show := func(key string){
		val,ok := os.LookupEnv(key)
		if !ok{
			fmt.Printf("%s not set\n", key)
		}else{
			fmt.Printf("%s=%s\n", key, val)
		}
	}

	show("GOPATH")
	show("USER")
	show("TMPDOR")

	//删除环境变量
	err = os.Unsetenv("TMPDOR")
	if err != nil{
		t.Errorf("%v",err)
		return
	}
}

//TestOsHostName - 返回主机名
func TestOsHostName(t *testing.T){
	hostname,err := os.Hostname()
	if err != nil{
		t.Errorf("%v",err)
		return
	}
	t.Logf("%v",hostname)
}

//TestOsExit -退出当前程序
func TestOsExit(t *testing.T){
	os.Exit(0) //表示成功,退出
	os.Exit(1) //表示因为错误,退出
}

//TestOsStdErr -错误输出到控制台
func TestOsStdErr(t *testing.T){
	//os.Stderr io.Writer
	err := errors.New("请求错误")
	fmt.Fprint(os.Stderr,err)
}


//=============================================os more=======================================
//TestFileBasedPipe-管道操作
func TestFileBasedPipe(t *testing.T){
	//注意r,w类型都是*File
	r,w,err := os.Pipe()
	if err != nil{
		t.Errorf("%v",err)
		return
	}

	//管道操作会阻塞在read or write,具体阻塞在哪取决于哪一个先被求值，这就是开启协程没有结束掉的原因
	go func() {
		content := make([]byte,6)
		var c int
		c,err = r.Read(content)
		if err != nil{
			t.Errorf("%v",err)
			return
		}

		t.Logf("%v",c)
		t.Logf("%v",string(content))
	}()

	var n int
	n,err =w.Write([]byte("你好"));
	if err != nil{
		t.Errorf("%v",err)
		return
	}

	t.Logf("%v",n)
}








