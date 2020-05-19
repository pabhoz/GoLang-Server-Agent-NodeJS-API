package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	ps "github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
)

// Max length 60 => saved to database as varchar
var agentUID string = os.Args[1]

// Ex: 8080
var port string = os.Args[2]

// Ex: http://localhost:3000
var api string = os.Args[3]

func dealwithErr(err error) {
	if err != nil {
		fmt.Println(err)
		//os.Exit(-1)
	}
}

// PostBody Structure
type PostBody struct {
	AgetUID   string      `json:"agentUID"`
	AgentName string      `json:"agentName"`
	Data      interface{} `json:"data"`
}

func preparePostBody(data interface{}) PostBody {
	hostInfo := getHostInfo()
	postBody := PostBody{
		AgetUID:   agentUID,
		AgentName: hostInfo.Hostname,
		Data:      data,
	}
	return postBody
}

func postRequest(endpoint string, requestBody []byte) []byte {
	req, err := http.NewRequest("POST", api+endpoint, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Agent-Key", "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918")
	dealwithErr(err)

	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	resp, err := client.Do(req)
	dealwithErr(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	dealwithErr(err)

	return body
}

// ProcessorInfo Structure
type ProcessorInfo struct {
	CPUIndex              int64            `json:"cpuIndex"`
	VendorID              string           `json:"vendorId"`
	Family                string           `json:"family"`
	NumberOfCores         int64            `json:"numberOfCores"`
	ModelName             string           `json:"modelName"`
	Speed                 string           `json:"speed"`
	CurrentCPUUtilization []CPUUtilization `json:"currentCPUUtilization"`
}

// CPUUtilization structure
type CPUUtilization struct {
	Index       int    `json:"index"`
	Utilization string `json:"utilization"`
}

func getProcessorDataController(w http.ResponseWriter, r *http.Request) {
	processorInfo := getProcessorData()
	json.NewEncoder(w).Encode(processorInfo)
}

func postProcessorDataController(w http.ResponseWriter, r *http.Request) {
	processorInfo := getProcessorData()

	postBody := preparePostBody(processorInfo)

	requestBody, err := json.Marshal(postBody)
	dealwithErr(err)

	body := postRequest("/processors", requestBody)

	fmt.Println(string(body))
	w.Write(body)
}

func getProcessorData() ProcessorInfo {
	cpuStat, err := cpu.Info()
	dealwithErr(err)
	percentage, err := cpu.Percent(0, true)
	dealwithErr(err)

	processorInfo := ProcessorInfo{
		CPUIndex:      int64(cpuStat[0].CPU),
		VendorID:      cpuStat[0].VendorID,
		Family:        cpuStat[0].Family,
		NumberOfCores: int64(cpuStat[0].Cores),
		ModelName:     cpuStat[0].ModelName,
		Speed:         strconv.FormatFloat(cpuStat[0].Mhz, 'f', 2, 64),
	}

	for idx, cpupercent := range percentage {
		processorInfo.CurrentCPUUtilization = append(processorInfo.CurrentCPUUtilization, CPUUtilization{Index: idx, Utilization: strconv.FormatFloat(cpupercent, 'f', 2, 64) + "%"})
	}

	return processorInfo
}

// RunningProccessesInfo Structure
type RunningProccessesInfo struct {
	Total         uint64        `json:"total"`
	Running       int           `json:"running"`
	ProcessesList []ProcessInfo `json:"processesList"`
}

// ProcessInfo Structure
type ProcessInfo struct {
	Pid  int    `json:"pid"`
	Name string `json:"name"`
}

func getRunningProcessesController(w http.ResponseWriter, r *http.Request) {
	runningProcesses := getRunningProcesses()
	json.NewEncoder(w).Encode(runningProcesses)
}

func postRunningProcessesController(w http.ResponseWriter, r *http.Request) {
	runningProcesses := getRunningProcesses()

	postBody := preparePostBody(runningProcesses)

	requestBody, err := json.Marshal(postBody)
	dealwithErr(err)

	body := postRequest("/runningProcesses", requestBody)

	fmt.Println(string(body))
	w.Write(body)
}

func getRunningProcesses() RunningProccessesInfo {
	infoStat, err := host.Info()
	dealwithErr(err)

	miscStat, err := load.Misc()
	dealwithErr(err)

	processList, err := ps.Processes()
	dealwithErr(err)

	runningProcesses := RunningProccessesInfo{
		Total:   infoStat.Procs,
		Running: miscStat.ProcsRunning,
	}

	for x := range processList {
		var process ps.Process
		process = processList[x]

		runningProcesses.ProcessesList = append(runningProcesses.ProcessesList, ProcessInfo{Pid: process.Pid(), Name: process.Executable()})
	}

	return runningProcesses
}

// UsersLog structure
type UsersLog struct {
	Users []UserInfo `json:"activeUsers"`
}

// UserInfo structure
type UserInfo struct {
	Username    string `json:"username"`
	Application string `json:"application"`
	Date        string `json:"date"`
	Time        string `json:"time"`
}

func getCurrentUsersController(w http.ResponseWriter, r *http.Request) {
	users := getCurrentUsers()
	json.NewEncoder(w).Encode(users)
}

func postCurrentUsersController(w http.ResponseWriter, r *http.Request) {
	users := getCurrentUsers()

	postBody := preparePostBody(users)

	requestBody, err := json.Marshal(postBody)
	dealwithErr(err)

	body := postRequest("/users", requestBody)

	fmt.Println(string(body))
	w.Write(body)
}

func getCurrentUsers() UsersLog {
	out, err := exec.Command("who").Output()
	dealwithErr(err)

	userInfoArray := strings.Fields(string(out))

	var users []UserInfo

	for i := 0; i < len(userInfoArray)/5; i++ {

		limit := i * 5
		username := strings.Join(userInfoArray[limit:limit+1], "")
		application := strings.Join(userInfoArray[limit+1:limit+2], "")
		date := strings.Join(userInfoArray[limit+2:limit+4], "")
		time := strings.Join(userInfoArray[limit+4:limit+5], "")

		user := UserInfo{
			Username:    username,
			Application: application,
			Date:        date,
			Time:        time,
		}
		users = append(users, user)
	}
	return UsersLog{Users: users}
}

// OSInfo structure
type OSInfo struct {
	Runtime  string `json:"runtime"`
	Name     string `json:"name"`
	Platform string `json:"platform"`
}

func getOSDataController(w http.ResponseWriter, r *http.Request) {

	osInfo := getOSData()
	json.NewEncoder(w).Encode(osInfo)
}

func postOSDataController(w http.ResponseWriter, r *http.Request) {
	osInfo := getOSData()

	postBody := preparePostBody(osInfo)

	requestBody, err := json.Marshal(postBody)
	dealwithErr(err)

	body := postRequest("/os", requestBody)

	fmt.Println(string(body))
	w.Write(body)
}

func getOSData() OSInfo {
	hostStat, err := host.Info()
	dealwithErr(err)

	osInfo := OSInfo{
		Runtime:  runtime.GOOS,
		Name:     hostStat.OS,
		Platform: hostStat.Platform,
	}
	return osInfo
}

func getHostInfo() *host.InfoStat {
	hostStat, err := host.Info()
	dealwithErr(err)

	return hostStat
}

func index(w http.ResponseWriter, r *http.Request) {

	hostInfo := getHostInfo()

	msg := "Agente de la máquina " + hostInfo.Hostname + "\n\n"
	msg = msg + "Consultas disponibles:\n"
	msg = msg + "* Información de procesador: /processor\n"
	msg = msg + "* Información de procesos: /runningProcesses\n"
	msg = msg + "* Información de procesos: /users\n"
	msg = msg + "* Información de SO: /os\n"
	msg = msg + "\n"
	msg = msg + "Envíos al API disponibles:\n"
	msg = msg + "* Información de procesador: /log/processor\n"
	msg = msg + "* Información de procesos: /log/runningProcesses\n"
	msg = msg + "* Información de procesos: /log/users\n"
	msg = msg + "* Información de SO: /log/os\n"
	w.Write([]byte(msg))
}

func handshakeAPI() {
	postBody := PostBody{
		AgetUID: agentUID,
	}

	requestBody, err := json.Marshal(postBody)
	dealwithErr(err)

	body := postRequest("/agents", requestBody)

	fmt.Println(string(body))
}

func main() {
	mux := http.NewServeMux()

	// Set Mux Server Routes
	mux.HandleFunc("/", index)
	mux.HandleFunc("/processor", getProcessorDataController)
	mux.HandleFunc("/runningProcesses", getRunningProcessesController)
	mux.HandleFunc("/users", getCurrentUsersController)
	mux.HandleFunc("/os", getOSDataController)

	mux.HandleFunc("/log/processor", postProcessorDataController)
	mux.HandleFunc("/log/runningProcesses", postRunningProcessesController)
	mux.HandleFunc("/log/users", postCurrentUsersController)
	mux.HandleFunc("/log/os", postOSDataController)

	// Get host Data for the welcome log
	hostInfo := getHostInfo()
	fmt.Println("Agent Hostname:" + hostInfo.Hostname)
	fmt.Println("Agent UID:" + agentUID)
	fmt.Println("Listening at: localhost:" + port + "; API: " + api)

	// Call API for Agent register if not yet registered
	handshakeAPI()

	// Run Mux HTTP Server
	http.ListenAndServe(":"+port, mux)

}
