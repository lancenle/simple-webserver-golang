//*******************************************************************************
// Written by Lance N. Le
//
// Free to use.
// Free to distribute.
// No warranty, expressed or implied, comes with this program.
//*******************************************************************************

package main

import "flag"
import "fmt"
import "io"
import "log"
import "net/http"
import "os"




//*******************************************************************************
// Using a global variable doesn't require the creation of a special http handler
// with a third argument.  See https://golang.org/doc/articles/wiki/final.go
//
// For complex applications, global variables can make code difficult to debug
// and maintain.
//*******************************************************************************
var gstrIndexFile string;




func hdlSendIndexFile(w http.ResponseWriter, r *http.Request) {
   fileIndex, err := os.Open(gstrIndexFile)
   if err != nil {
      log.Fatal(err)
   } else {
      slBuffer := make([]byte, 1024)


      nBytesRead, err := fileIndex.Read(slBuffer)
      if err != nil {
         log.Fatal(err)
      }

      strToClient := string(slBuffer[:nBytesRead])

      io.WriteString(w, strToClient)
   }
}




func main() {
   //****************************************************************************
   // Process command line parameters.
   //****************************************************************************
   pflagnPort        := flag.Int("port", 80, "Host port number to listen")
   pflagbDebug       := flag.Bool("debug", false, "Print out debug messages")
   pflagstrListenIP  := flag.String("listenip", "127.0.0.1", "Host listen ip")
   pflagstrIndexFile := flag.String("indexfile", "./index.htm", "Index HTML file")


   flag.Parse()


   fmt.Printf("debug = %t\n",     *pflagbDebug)
   fmt.Printf("port = %d\n",      *pflagnPort)
   fmt.Printf("listenip = %s\n",  *pflagstrListenIP)
   fmt.Printf("indexfile = %s\n", *pflagstrIndexFile)


   //****************************************************************************
   // Host variable should have ip and port, e.g. 127.0.0.1:80.
   //****************************************************************************
   strHost := fmt.Sprintf("%s:%d", *pflagstrListenIP, *pflagnPort)
   if (*pflagbDebug) {fmt.Printf("DEBUG Calling http handler with %s\n", strHost)}


   //****************************************************************************
   // gstrIndexFile holds the path/filename of index.htm and used by
   // hdlSendIndexFile to read the file's contents and send to the client.
   //****************************************************************************
   gstrIndexFile = fmt.Sprintf("%s", *pflagstrIndexFile)
   if (*pflagbDebug) {fmt.Printf("DEBUG Global index file is %s\n", gstrIndexFile)}


   //****************************************************************************
   // Running on low port numbers will require root access.
   //****************************************************************************
   http.HandleFunc("/", hdlSendIndexFile)
   http.ListenAndServe(strHost, nil)
}
