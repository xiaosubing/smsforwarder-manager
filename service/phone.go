package service

import (
	"smsforwarder-manager/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"os"

	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var scanResults map[string][]int

func GetPhones(c *gin.Context) {

	startScan()

	if len(scanResults) == 0 {
		fmt.Println("没有发现开放指定端口的设备")
		return
	}

	// 对IP地址进行排序
	var ips []string
	for ip := range scanResults {
		ips = append(ips, ip)
	}
	sort.Strings(ips)

	for _, ip := range ips {
		ports := scanResults[ip]

		url := fmt.Sprintf("http://%s:%v/config/query", ip, ports[0])
		payload := map[string]string{
			"data":      "",
			"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
			"sign":      "",
		}
		payloadStr, _ := json.Marshal(payload)
		ret := HttpPost(url, string(payloadStr))

		models.InsertPhoneData(ret, ip)

	}

	// return 信息
	phones := models.QueryPhoneData("", 0)
	var phoneDataRet []string
	for _, phone := range phones {
		tmp := phone.Alias
		if phone.Alias == "" {
			tmp = phone.Phone
		}
		phoneDataRet = append(phoneDataRet, tmp)
	}
	c.JSON(200, phoneDataRet)
}

func startScan() {
	// 获取本机IP地址
	localIP, err := getLocalIP()
	if err != nil {
		fmt.Printf("获取本机IP失败: %v\n", err)
		os.Exit(1)
	}

	// 计算C段IP范围
	cSegment := getCSegment(localIP)

	//ports := []int{5000}
	ports := []int{801, 5000}
	scanResults = scanCSegmentPorts(cSegment, ports)

	// 打印扫描结果
	//printScanResults(scanResults, ports)
}

// 获取本机IP地址
func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", fmt.Errorf("没有找到有效的IP地址")
}

// 获取C段网络地址 (192.168.1.xxx -> 192.168.1)
func getCSegment(ip string) string {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return ip
	}
	return strings.Join(parts[:3], ".")
}

// 扫描C段中的指定端口
func scanCSegmentPorts(cSegment string, ports []int) map[string][]int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	results := make(map[string][]int)

	for i := 1; i <= 254; i++ {
		ip := fmt.Sprintf("%s.%d", cSegment, i)
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			openPorts := checkPorts(ip, ports)
			if len(openPorts) > 0 {
				mu.Lock()
				results[ip] = openPorts
				mu.Unlock()
			}
		}(ip)
	}
	wg.Wait()
	return results
}

// 检查IP上的多个端口
func checkPorts(ip string, ports []int) []int {
	var openPorts []int
	for _, port := range ports {
		if isPortOpen(ip, port) {
			openPorts = append(openPorts, port)
		}
	}
	return openPorts
}

// 检查单个端口是否开放
func isPortOpen(ip string, port int) bool {
	address := net.JoinHostPort(ip, strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

// 打印扫描结果
func printScanResults(results map[string][]int, ports []int) {
	if len(results) == 0 {
		fmt.Println("没有发现开放指定端口的设备")
		return
	}

	// 对IP地址进行排序
	var ips []string
	for ip := range results {
		ips = append(ips, ip)
	}
	sort.Strings(ips)

	fmt.Printf("\n扫描结果 (开放端口 %v):\n", ports)
	for _, ip := range ips {
		ports := results[ip]
		fmt.Printf("%-15s 开放端口: %v\n", ip, ports)
	}
}
