package logger

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func Fatal(msg ...interface{}) {
	fmt.Println(append([]interface{}{color.RedString("Fatal error:")}, msg...)...)
	os.Exit(1)
}

func Error(msg ...interface{}) {
	fmt.Println(append([]interface{}{color.RedString("Error:")}, msg...)...)
}

func Warn(msg ...interface{}) {
	fmt.Println(append([]interface{}{color.YellowString("Warning:")}, msg...)...)
}

func Info(msg ...interface{}) {
	fmt.Println(append([]interface{}{color.GreenString("Info:")}, msg...)...)
}

func Fatalf(f string, v ...interface{}) {
	fmt.Printf(color.RedString("\nFatal error: ")+f+"\n", v...)
	os.Exit(1)
}

func Errorf(f string, v ...interface{}) {
	fmt.Printf(color.RedString("Error: ")+f+"\n", v...)
}

func Warnf(f string, v ...interface{}) {
	fmt.Printf(color.YellowString("Warning: ")+f+"\n", v...)
}

func Infof(f string, v ...interface{}) {
	fmt.Printf(color.GreenString("Info: ")+f+"\n", v...)
}

func WarnWithBackGround(msg ...interface{}) {
	st := color.New(warnFontColor).Add(warnBackgroundColor).SprintFunc()
	fmt.Println(st(msg...))
}

func WarnfWithBackGround(f string, v ...interface{}) {
	st := color.New(warnFontColor).Add(warnBackgroundColor).SprintFunc()

	fmt.Printf(st(f)+"\n", v...)
}

func FatalWithBackGround(msg ...interface{}) {
	st := color.New(fatalFontColor).Add(fatalBackgroundColor).SprintFunc()

	fmt.Println(st(msg...))
	os.Exit(1)
}

func FatalfWithBackGround(f string, v ...interface{}) {
	st := color.New(fatalFontColor).Add(fatalBackgroundColor).SprintFunc()

	fmt.Printf(st(f)+"\n", v...)
	os.Exit(1)
}

func InfoWithBackGround(msg ...interface{}) {
	st := color.New(infoFontBackgroundColor).Add(infoBackgroundColor).SprintFunc()

	fmt.Println(st(msg...))
}

func InfofWithBackground(f string, v ...interface{}) {
	st := color.New(infoFontBackgroundColor).Add(infoBackgroundColor).SprintFunc()

	fmt.Printf(st(f)+"\n", v...)
}

func CustomColorPrint(fontColor color.Attribute, backGroundColor color.Attribute, msg ...interface{}) {
	st := color.New(fontColor).Add(backGroundColor).SprintFunc()

	fmt.Println(st(msg...))
}

func CustomColorPrintf(f string, fontColor color.Attribute, backGroundColor color.Attribute, v ...interface{}) {
	st := color.New(fontColor).Add(backGroundColor).SprintFunc()

	fmt.Printf(st(f)+"\n", v...)
}

func InfoBold(msg ...interface{}) {
	st := color.New(infoFontColor).Add(color.Bold).SprintFunc()

	fmt.Println(st(msg...))
}

func InfofBold(f string, v ...interface{}) {
	st := color.New(infoFontColor).Add(color.Bold).SprintFunc()

	fmt.Printf(st(f)+"\n", v...)
}
