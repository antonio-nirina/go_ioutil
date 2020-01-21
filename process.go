package main

// https://stackoverflow.com/questions/47055584/how-to-start-and-monitor-and-kill-process-with-go
import (
    "fmt"
    "os/exec"
    "strconv"
    "runtime"

    log "github.com/sirupsen/logrus"
)

type iBrowserHandler interface {
    Start(processListener chan bool) error
    Stop() error
    KillProcess() error
}

type browserHandler struct {
    cmd            *exec.Cmd
    pathToChromium string
    pdvEndpoint    string
}

func newBrowserHandler(pathToChromium string, pdvEndpoint string) iBrowserHandler {
    b := &browserHandler{}
    b.pathToChromium = pathToChromium
    b.pdvEndpoint = pdvEndpoint

    return b
}

func (b *browserHandler) Start(processListener chan bool) error {
    endpoint := fmt.Sprintf("--app=%s", b.pdvEndpoint)

    b.cmd = exec.Command(b.pathToChromium, endpoint, "--kiosk")
    err := b.cmd.Run()
    if err != nil {
        log.WithError(err).Error("Error with the browser process")
        processListener <- false
    } else {
        processListener <- true
    }

    return err
}

func (b *browserHandler) Stop() error {
    err := b.cmd.Process.Release()
    if err != nil {
        log.WithError(err).Fatal("Error shutting down chromium")
    }

    return err
}

// Correct KillProcess
func (b *browserHandler) KillProcess() error {
    log.Info("Killing browser process")
    kill := exec.Command("taskkill", "/T", "/F", "/PID", strconv.Itoa(b.cmd.Process.Pid))
    err := kill.Run()
    if err != nil {
        log.WithError(err).Error("Error killing chromium process")
    }
    log.Info("Browser process was killed")

    return err
}

func InitP() {
   // var pathToChromium string
 //   var c = config.GetInstance()
    var os = runtime.GOOS

    if os == "windows" {
    //    pathToChromium = "chromium-browser\\ChromiumPortable.exe"
    } else {
    //    pathToChromium = "chrome"
    }

    //handler := newBrowserHandler(pathToChromium, c.GetConfig().PDVUrl)
}

func StartBrowser(done chan bool) {
    browserProcessListener := make(chan bool)
    defer close(browserProcessListener)
    // go Start(browserProcessListener)

    var tryReopenBrowser = true
    for {
        select {
        case <-browserProcessListener:
            if tryReopenBrowser {
                log.Warn("Browser process is stopped. Attempting to restart")
               // go handler.Start(browserProcessListener)
            } else {
                log.Warn("Browser process is stopped. Will not attempt to restart")
            }

        case <-done:
            log.Info("Shutting down browser")
            tryReopenBrowser = false
//            handler.KillProcess()
            return

        default:
        }
    }
}