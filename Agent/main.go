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

	ps "github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
)

var port string = os.Args[1]
var api string = os.Args[2]

func dealwithErr(err error) {
	if err != nil {
		fmt.Println(err)
		//os.Exit(-1)
	}
}

// ProcessorInfo Structure
type ProcessorInfo struct {
	CPUIndex              int64            `json:"index"`
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

// PostBody Structure
type PostBody struct {
	AgentName string      `json:"agentName"`
	Data      interface{} `json:"data"`
}

func postProcessorDataController(w http.ResponseWriter, r *http.Request) {
	processorInfo := getProcessorData()
	hostInfo := getHostInfo()
	postBody := PostBody{
		AgentName: hostInfo.Hostname,
		Data:      processorInfo,
	}
	requestBody, err := json.Marshal(postBody)
	dealwithErr(err)

	body := postRequest("/processors", requestBody)

	fmt.Println(string(body))
	// json.NewEncoder(w).Encode(processorInfo)
}

func postRequest(endpoint string, requestBody []byte) []byte {
	resp, err := http.Post(api+endpoint, "application/json", bytes.NewBuffer(requestBody))
	dealwithErr(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	dealwithErr(err)

	return body
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

func getCurrentUsers() []UserInfo {
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
	return users
}

// OSInfo structure
type OSInfo struct {
	Runtime  string
	Name     string
	Platform string
}

func getSODataController(w http.ResponseWriter, r *http.Request) {

	osInfo := getSOData()
	json.NewEncoder(w).Encode(osInfo)
}

func getSOData() OSInfo {
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
	msg = msg + "* Información de SO: /so\n"
	msg = msg + "\n"
	msg = msg + "Envíos al API disponibles:\n"
	msg = msg + "* Información de procesador: /send/processor\n"
	msg = msg + "* Información de procesos: /send/runningProcesses\n"
	msg = msg + "* Información de procesos: /send/users\n"
	msg = msg + "* Información de SO: /send/so\n"
	w.Write([]byte(msg))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/processor", getProcessorDataController)
	mux.HandleFunc("/runningProcesses", getRunningProcessesController)
	mux.HandleFunc("/users", getCurrentUsersController)
	mux.HandleFunc("/so", getSODataController)

	mux.HandleFunc("/send/processor", postProcessorDataController)
	//mux.HandleFunc("/send/runningProcesses", postRunningProcessesController)
	//mux.HandleFunc("/send/users", postCurrentUsersController)
	//mux.HandleFunc("/send/so", postSODataController)

	fmt.Println("Listening at: localhost:" + port + "; API: " + api)
	http.ListenAndServe(":"+port, mux)

}
