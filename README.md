HART package
============

This is [HART](https://en.wikipedia.org/wiki/Highway_Addressable_Remote_Transducer_Protocol) communication protocol implementation in [go](golang.go)

Usage examples
--------------

#### Sending HART frame


``` go
package main

import (
    "fmt"
    "github.com/jszumigaj/hart"
)

func main() {

    port, err := hart.Open("COM1")
    if err != nil {
        panic(err)
    }

    defer port.Close()

    frame := hart.FrameZero
    tx := frame.Buffer()

    rx := make([]byte, 128)
    n, err := port.SendFrame(tx, rx)
    if err != nil {
        panic(err)
    }

    fmt.Printf("RxBuf %v: %02x\n", n, rx[:n])

    if reply, ok := hart.Parse(rx); ok {
        fmt.Printf("Reply: %v\n", *reply)
    }

    fmt.Printf("Done.")
}
```


#### Using HART commands

``` go
package main

import (
    "log"

    "github.com/jszumigaj/hart"
    "github.com/jszumigaj/hart/univrsl"
)

func main() {
    port, err := hart.Open("COM1")
    if err != nil {
        log.Fatal(err)
    }
    
    defer port.Close()
    
    master := hart.NewMaster(port)
    
    device := &univrsl.Device{}
    
    cmd0 := &univrsl.Command0{Device: device}
    
    log.Println("Executing:", cmd0.Description())
    status, err := master.Execute(cmd0, device)
    if err != nil {
        log.Fatal(err)
    }
    
    log.Println("Status:", status)
    log.Println("Result:", cmd0)
    
    cmd1 := &univrsl.Command1{}
    
    log.Println("Executing:", cmd1.Description())
    status, err = master.Execute(cmd1, device)
    if err != nil {
        log.Fatal(err)
    }
    
    log.Println("Command status:", status)
    log.Println("Device status:", device.Status())
    log.Printf("PV = %v [%v]", cmd1.PV, cmd1.Unit)
}
```

#### More complex scenario example

``` go
package main

import (
    "flag"
    "log"
    "time"

    "github.com/jszumigaj/hart"
    "github.com/jszumigaj/hart/univrsl"
)

var (
    portName = flag.String("c", "COM1", "Serial port name.")
    delay    = flag.Int("d", 10, "Delay between each commands set execution in seconds.")
)

func main() {
    flag.Parse()
    log.Printf("Start reading commands every %d sec", *delay)
    log.Printf("Using serial port %s", *portName)

    serial, err := hart.Open(*portName)
    if err != nil {
        log.Fatalln("ERROR:", err)
    }
    defer serial.Close()

    master := hart.NewMaster(serial)

    device := &univrsl.Device{}

    commands := []hart.Command{
        device.Command0(),
        &univrsl.Command1{},
        &univrsl.Command2{},
        &univrsl.Command3{},
        &univrsl.Command13{},
        &univrsl.Command15{},
    }

    executed := make(chan hart.Command)

    go executeCommands(master, device, commands, executed)

    // displayResults
    for command := range executed {
        log.Println("Command status:", command.Status())
        log.Println("Device status:", device.Status())

        switch cmd := command.(type) {
        case *univrsl.Command0:
            log.Printf("Cmd #0: Device %v", device)

        case *univrsl.Command1:
            log.Printf("Cmd #1: PV = %v [%v]\n", cmd.PV, cmd.Unit)

        case *univrsl.Command2:
            log.Printf("Cmd #2: Current = %v [mA]\n", cmd.Current)
            log.Printf("Cmd #2: PoR = %v [%%]\n", cmd.PercentOfRange)

        case *univrsl.Command3:
            log.Printf("Cmd #3: SV = %v [%v]\n", cmd.Sv, cmd.SvUnit)
            log.Printf("Cmd #3: TV = %v [%v]\n", cmd.Tv, cmd.TvUnit)
            log.Printf("Cmd #3: FV = %v [%v]\n", cmd.Fv, cmd.FvUnit)

        case *univrsl.Command13:
            log.Printf("Cmd #13: Tag: %v", cmd.Tag)
            log.Printf("Cmd #13: Descriptor: %v", cmd.Descriptor)
            log.Printf("Cmd #13: Date: %v", cmd.Date.Format("2006-01-02"))

        case *univrsl.Command15:
            unit := cmd.UpperAndLowerRangeValuesUnit.String()
            log.Printf("Cmd #15: %v", cmd)
            log.Printf("Cmd #15: Range = %v ... %v [%v]\n", cmd.LowerRangeValue, cmd.UpperRangeValue, unit)
            log.Printf("Cmd #15: Damping = %v [s]\n", cmd.Damping)
        }
    }

}

func executeCommands(master *hart.Master, device *univrsl.Device, commands []hart.Command, executed chan<- hart.Command) error {
    for {
        start := time.Now()
        for _, cmd := range commands {

            if _, err := master.Execute(cmd, device); err != nil {
                log.Println(cmd.Description(), "error:", err)
            } else {
                executed <- cmd
            }

        }
        elapsed := time.Now().Sub(start)
        time.Sleep(time.Duration(*delay)*time.Second - elapsed)
    }
}


```
