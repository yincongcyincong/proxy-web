package webtail

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"net/http"

	"github.com/alecthomas/template"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/julienschmidt/httprouter"
)

func home(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	filename := ps.ByName("name")
	if strings.Contains(filename, "/") || filename == "" {
		log.Printf("invalid log file name : %s", filename)
		r.Body.Close()
		return
	}
	homeTemplate.Execute(w, "ws://"+r.Host+"/log/"+filename)
}
func ViewLog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	filename := ps.ByName("name")
	if strings.Contains(filename, "/") || filename == "" {
		log.Printf("invalid log file name : %s", filename)
		r.Body.Close()
		return
	}
	logfile := fmt.Sprintf("%s/%s.log", basedir, filename)

	ft, err := os.Stat(logfile)
	if err != nil {
		log.Printf("can not access file %s , %s", logfile, err)
		r.Body.Close()
		return
	}
	conn, _, _, err := ws.UpgradeHTTP(r, w, nil)
	if err != nil {
		log.Printf("ws upgrade fail , %s", err)
		r.Body.Close()
		return
	}
	var (
		state  = ws.StateServerSide
		writer = wsutil.NewWriter(conn, state, ws.OpText)
	)
	log.Printf("client online , %s", conn.RemoteAddr())
	content, _ := TailN(logfile, 10)
	if content != "" {
		_, err := writer.Write([]byte(content))
		if err != nil {
			conn.Close()
			log.Printf("client offline with write , %s", err)
			return
		}
		err = writer.Flush()
		if err != nil {
			conn.Close()
			log.Printf("client offline with flush , %s", err)
			return
		}
	}
	file, err := os.Open(logfile)
	if err != nil {
		conn.Close()
		log.Printf("open log file fail , %s", err)
		return
	}
	file.Seek(ft.Size(), 0)
	if err != nil {
		conn.Close()
		file.Close()
		log.Printf("sedd log file stat fail , %s", err)
	}
	reader := bufio.NewReader(file)
	timer := time.NewTicker(time.Millisecond * 200)
	go func() {
		defer func() {
			conn.Close()
			file.Close()
			timer.Stop()
		}()
		for {
			select {
			case <-timer.C:
				line, err := reader.ReadString('\n')
				if line != "" {
					if err != nil && err == io.EOF && !strings.Contains(line, "\n") {
						line += "\n"
					}
					_, err = writer.Write([]byte(line))
					if err != nil {
						log.Printf("client offline with write, %s", err)
						return
					}
					err = writer.Flush()
					if err != nil {
						log.Printf("client offline with flush , %s", err)
						return
					}
				}
			}

		}
	}()
	go func() {
		_, _, err := wsutil.ReadClientData(conn)
		if err != nil {
			log.Printf("client offline with read , %s", err)
			conn.Close()
			file.Close()
			timer.Stop()
			return
		}
	}()
}

var (
	basedir string
)

func Serve(address, logDir string) (listener *net.Listener, err error) {
	basedir = logDir
	l, err := net.Listen("tcp", address)
	if err != nil {
		return
	}
	router := httprouter.New()
	router.GET("/log/:name", ViewLog)
	router.GET("/show/:name", home)
	log.Printf("WS Log Server on %s", l.Addr())
	go func() { log.Fatal(http.Serve(l, router)) }()
	listener = &l
	return
}

func TailN(filename string, numLines int) (string, error) {
	//MAKE SURE FILENAME IS GIVEN
	//actually, a path to the file
	if numLines == 0 {
		numLines = 10
	}
	if len(filename) == 0 {
		return "", errors.New("You must provide the path to a file")
	}

	//OPEN FILE
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	//SEEK BACKWARD CHARACTER BY CHARACTER ADDING UP NEW LINES
	//offset must start at "-1" otherwise we are already at the EOF
	//"-1" from numLines since we ignore "last" newline in a file
	numNewLines := 0
	var offset int64 = -1
	var finalReadStartPos int64
	for numNewLines <= numLines-1 {
		//seek to new position in file
		startPos, err := file.Seek(offset, 2)
		if err != nil {
			return "", err
		}

		//make sure start position can never be less than 0
		//aka, you cannot read from before the file starts
		if startPos == 0 {
			//set to -1 since we +1 to this below
			//the position will then start from the first character
			finalReadStartPos = -1
			break
		}

		//read the character at this position
		b := make([]byte, 1)
		_, err = file.ReadAt(b, startPos)
		if err != nil {
			return "", err
		}

		//ignore if first character being read is a newline
		if offset == int64(-1) && string(b) == "\n" {
			offset--
			continue
		}

		//if the character is a newline
		//add this to the number of lines read
		//and remember position in case we have reached our target number of lines
		if string(b) == "\n" {
			numNewLines++
			finalReadStartPos = startPos
		}

		//decrease offset for reading next character
		//remember, we are reading backward!
		offset--
	}

	//READ TO END OF FILE
	//add "1" here to move offset from the newline position to first character in line of text
	//this position should be the first character in the "first" line of data we want
	b := make([]byte, 4096)
	n, err := file.ReadAt(b, finalReadStartPos+1)
	if n > 0 {
		return string(b[:n]), nil
	}
	return "", err
}

var homeTemplate = template.Must(template.New("").Parse(`
	<!DOCTYPE html>
	<html>
	<head>
	<meta charset="utf-8">
	<style>
	#output{width:100%;height:300px;margin:5px;padding-bottom:1em;overflow:auto;}
	body{padding:0;font-size:0.8em;overflow:hidden;}
	</style>
	</head>
	<body>
	<div id="output"></div>
	<script>
	location.parseHash = function(){
		var hash = (this.hash ||'').replace(/^#/,'').split('&'),
			parsed = {};
	 
		for(var i =0,el;i<hash.length; i++ ){
			 el=hash[i].split('=')
			 parsed[el[0]] = el[1];
		}
		return parsed;
	};
	var para= location.parseHash();
	if(para.height){
		document.getElementById("output").style.height=para.height;
	}
	if(para.width){
		document.getElementById("output").style.width=para.width;
	}
	window.addEventListener("load", function(evt) {
		var output = document.getElementById("output");
		var ws;
		var print = function(message) {
			var d = document.createElement("div");
			message=message.replace(new RegExp(/\n/g),"<br>") ;
			message=message.replace(new RegExp(/\t/g),"&nbsp;&nbsp;&nbsp;&nbsp;") ;
			message=message.replace(new RegExp(/ /g),"&nbsp;") ;
			d.innerHTML = message
			output.appendChild(d);
			output.scrollTop = output.scrollHeight;
		};
		ws = new WebSocket("{{.}}");
		ws.onopen = function(evt) {
			print("OPEN_SUCCESS");
		}
		ws.onclose = function(evt) {
			print("CLOSED");
			ws = null;
		}
		ws.onmessage = function(evt) {
			print(evt.data);
		}
		ws.onerror = function(evt) {
			print("ERROR: " + evt.data);
		}
	});
	</script>
	</body>
	</html>
	`))
