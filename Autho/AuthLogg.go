package Autho

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorReset  = "\033[0m"
)

type GlobalLogger struct {
	outputDir string
	Info      *ColorLog
	Debug     *ColorLog
	Warning   *ColorLog
	Fatal     *ColorLog
	openfiles []*os.File
}
type logFile struct {
	file *os.File
}
type ColorLog struct {
	logger *log.Logger
	color  string
}

func (l *logFile) WriteLog(prefix string, v ...interface{}) {
	logMessage := fmt.Sprintf("[%s] %s ", time.Now().Format("01/02/2006 15:04:05"), prefix)
	logMessage += fmt.Sprintf(v[0].(string), v[1:]...)
	logMessage += "\n"
	l.file.Write([]byte(logMessage))

}

/*

 */

func (g *GlobalLogger) InitLoggDir(Directory string) error {
	g.outputDir = Directory
	err := os.MkdirAll(Directory, 0755)
	if err != nil {
		return err
	}
	return nil
}

func (g *GlobalLogger) initInfoLogger(filename string) *os.File {
	filepath := filepath.Join(g.outputDir, filename)
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	//	defer file.Close()

	firstLine := []byte("\n" + "  < ServerStart >" + time.Now().Format("2006-01-02 15:04:05") + " <--------------> " + "\n")
	file.Write(firstLine)
	return file

}

func (g *GlobalLogger) initDebugLogger(filename string) *os.File {
	filepath := filepath.Join(g.outputDir, filename)
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	//	defer file.Close()
	firstLine := []byte("\n" + "  < ServerStart >" + time.Now().Format("2006-01-02 15:04:05") + " <--------------> " + "\n")
	file.Write(firstLine)
	return file

}

func (g *GlobalLogger) initWarningLogger(filename string) *os.File {

	filepath := filepath.Join(g.outputDir, filename)
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	//	defer file.Close()
	firstLine := []byte("\n" + "  < ServerStart >" + time.Now().Format("2006-01-02 15:04:05") + " <--------------> " + "\n")
	file.Write(firstLine)
	return file

}

func (g *GlobalLogger) initFataloLogger(filename string) *os.File {
	filepath := filepath.Join(g.outputDir, filename)
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	//	defer file.Close()
	firstLine := []byte("\n" + "  < ServerStart >" + time.Now().Format("2006-01-02 15:04:05") + " <--------------> " + "\n")
	file.Write(firstLine)
	return file
}

/*
Grab each file pointer for each file and then create output to those files using the inbuilt loger lib
Each logger level take the data,time,and line number in file that triggered the call
Then pass these loggers into ColorLog struct which allows for colorful termnial responses
Global Logger servers as a simple interface to use loggin throughout application
*/

func NewGlobalLogger(outputDirectory string) *GlobalLogger {
	glob := GlobalLogger{outputDir: outputDirectory}
	glob.InitLoggDir(glob.outputDir)
	infoLogger := glob.initInfoLogger("InfomationLogs")
	debugLogger := glob.initDebugLogger("DebuggerLogs")
	warningLogger := glob.initWarningLogger("WarningLogs")
	fatalLogger := glob.initFataloLogger("FatalLogs")

	toclose := []*os.File{infoLogger, debugLogger, warningLogger, fatalLogger}

	InfoLog := log.New(infoLogger, "Info ", log.Ldate|log.Ltime|log.Lshortfile)
	Debuglog := log.New(debugLogger, "Debug ", log.Ldate|log.Ltime|log.Lshortfile)
	Warninglog := log.New(warningLogger, "Warning ", log.Ldate|log.Ltime|log.Lshortfile)
	Fatallog := log.New(fatalLogger, "Fatal ", log.Ldate|log.Ltime|log.Lshortfile)

	GreenInfo := newColorLog(InfoLog, ColorGreen)
	BlueDebug := newColorLog(Debuglog, ColorBlue)
	YellowWarning := newColorLog(Warninglog, ColorYellow)
	RedFatal := newColorLog(Fatallog, ColorRed)

	GlobalLogs := GlobalLogger{outputDir: outputDirectory, Info: GreenInfo, Debug: BlueDebug, Warning: YellowWarning, Fatal: RedFatal, openfiles: toclose}
	return &GlobalLogs
}

func newColorLog(logger *log.Logger, color string) *ColorLog {
	return &ColorLog{
		logger: logger,
		color:  color,
	}
}

func (c *ColorLog) Output(message string) {
	_, filePath, line, _ := runtime.Caller(1)
	fileName := filepath.Base(filePath)
	fmt.Printf("%s[%s:%d]%s %s%s\n", c.color, fileName, line, message, "\x1b[0m", c.color)
	c.logger.Output(2, message) // Write to file
}

/*
Files are kept open throughout the lifecycle of program. This is to close them at the termniation of the server
wrapper function over
*/
func (g *GlobalLogger) CleanUp() {
	defer g.closeFiles()
	currentTime := time.Now()

	timestamp := currentTime.Format("2006-01-02 15:04:05")

	fmt.Printf("Finish closing Globallogs at %s\n", timestamp)
}

func (g *GlobalLogger) closeFiles() {

	for _, file := range g.openfiles {
		file.Close()
	}

}
