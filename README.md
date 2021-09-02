# log
```
func NewLogger(logfile,LogLevel string)*log.Logger{
    var Logger *log.Logger
    file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
            panic("Failed to open configuration file.")
    }
    switch LogLevel{
    case "debug":
            Logger = log.New(file,"",log.Ldate|log.Ltime|log.Lshortfile,0)
    case "info":
            Logger = log.New(file,"",log.Ldate|log.Ltime|log.Lshortfile,1)
    case "warning":
            Logger = log.New(file,"",log.Ldate|log.Ltime|log.Lshortfile,2)
    case "error":
            Logger = log.New(file,"",log.Ldate|log.Ltime|log.Lshortfile,3)
    case "fatal":
            Logger = log.New(file,"",log.Ldate|log.Ltime|log.Lshortfile,4)
    default:
            Logger = log.New(file,"",log.Ldate|log.Ltime|log.Lshortfile,1)
    }
    return Logger
}
```
