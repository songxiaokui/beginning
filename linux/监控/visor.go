package main

import (
	"bufio"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// 邮件配置信息
type MailConfig struct {
	SMTPServer string
	SMTPPort   string
	Sender     string
	Receiver   string
	Username   string
	Password   string
}

// 获取 CPU 使用情况
func getCPUUsage() (float64, error) {
	// 获取 /proc/stat 内容
	file, err := os.Open("/proc/stat")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "cpu ") {
			// cpu 负载数据
			fields := strings.Fields(line)
			user, _ := strconv.Atoi(fields[1])
			nice, _ := strconv.Atoi(fields[2])
			system, _ := strconv.Atoi(fields[3])
			idle, _ := strconv.Atoi(fields[4])
			// 计算 CPU 使用率
			total := user + nice + system + idle
			idlePercentage := float64(idle) / float64(total)
			return (1 - idlePercentage) * 100, nil
		}
	}
	return 0, fmt.Errorf("could not find CPU stats")
}

// 获取内存使用情况
func getMemoryUsage() (float64, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var total, free, buffers, cached int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "MemTotal") {
			total, _ = strconv.Atoi(strings.Fields(line)[1])
		}
		if strings.HasPrefix(line, "MemFree") {
			free, _ = strconv.Atoi(strings.Fields(line)[1])
		}
		if strings.HasPrefix(line, "Buffers") {
			buffers, _ = strconv.Atoi(strings.Fields(line)[1])
		}
		if strings.HasPrefix(line, "Cached") {
			cached, _ = strconv.Atoi(strings.Fields(line)[1])
		}
	}

	used := total - free - buffers - cached
	return float64(used) / float64(total) * 100, nil
}

// 获取占比最大的 5 个进程
func getTopProcesses() ([]string, error) {
	cmd := exec.Command("ps", "-eo", "pid,%cpu,%mem,comm", "--sort=-%cpu", "--no-headers", "-n", "5")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var processes []string
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		processes = append(processes, scanner.Text())
	}
	return processes, nil
}

// 获取磁盘使用情况
func getDiskUsage() (float64, string, error) {
	cmd := exec.Command("df", "-h")
	output, err := cmd.Output()
	if err != nil {
		return 0, "", err
	}

	var usagePercentage float64
	var diskDetails string
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "/dev/") { // 检查文件系统行
			fields := strings.Fields(line)
			usedPercentageStr := fields[4]
			usedPercentageStr = strings.TrimSuffix(usedPercentageStr, "%")
			usedPercentage, err := strconv.ParseFloat(usedPercentageStr, 64)
			if err != nil {
				return 0, "", err
			}
			usagePercentage = usedPercentage
			diskDetails += line + "\n"
		}
	}
	return usagePercentage, diskDetails, nil
}

// 发送邮件
func sendEmail(mailConfig MailConfig, subject, body string) error {
	auth := smtp.PlainAuth("", mailConfig.Username, mailConfig.Password, mailConfig.SMTPServer)
	msg := []byte("To: " + mailConfig.Receiver + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	return smtp.SendMail(mailConfig.SMTPServer+":"+mailConfig.SMTPPort, auth, mailConfig.Sender, []string{mailConfig.Receiver}, msg)
}

func main() {
	mailConfig := MailConfig{
		SMTPServer: "smtp.example.com",     // 邮件服务器地址
		SMTPPort:   "587",                  // 邮件服务器端口
		Sender:     "sender@example.com",   // 发件人邮箱
		Receiver:   "receiver@example.com", // 收件人邮箱
		Username:   "sender@example.com",   // 邮箱用户名
		Password:   "your_password",        // 邮箱密码
	}

	cpuThreshold := 80.0
	diskThreshold := 90.0
	var highCPUTime time.Time
	var highDiskTime time.Time

	for {
		// 获取 CPU 使用情况
		cpuUsage, err := getCPUUsage()
		if err != nil {
			log.Printf("获取 CPU 使用情况失败: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// 获取内存使用情况
		memUsage, err := getMemoryUsage()
		if err != nil {
			log.Printf("获取内存使用情况失败: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// 获取磁盘使用情况
		diskUsage, diskDetails, err := getDiskUsage()
		if err != nil {
			log.Printf("获取磁盘使用情况失败: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// 打印 CPU、内存和磁盘使用情况
		fmt.Printf("CPU 使用率: %.2f%%, 内存使用率: %.2f%%, 磁盘使用率: %.2f%%\n", cpuUsage, memUsage, diskUsage)

		// 检查 CPU 使用是否超过 80% 且持续超过 1 分钟
		if cpuUsage > cpuThreshold {
			if highCPUTime.IsZero() {
				highCPUTime = time.Now()
			} else if time.Since(highCPUTime) > time.Minute {
				// 如果超过 1 分钟，获取进程信息并发送邮件
				topProcesses, err := getTopProcesses()
				if err != nil {
					log.Printf("获取进程信息失败: %v", err)
					time.Sleep(5 * time.Second)
					continue
				}

				body := fmt.Sprintf("CPU 使用率超过 %.2f%%，持续超过 1 分钟。占比最大的 5 个进程:\n", cpuThreshold)
				for _, process := range topProcesses {
					body += process + "\n"
				}

				// 发送邮件
				err = sendEmail(mailConfig, "CPU 使用率告警", body)
				if err != nil {
					log.Printf("发送邮件失败: %v", err)
				}

				// 重置计时器
				highCPUTime = time.Time{}
			}
		} else {
			highCPUTime = time.Time{}
		}

		// 检查磁盘使用是否超过 90%
		if diskUsage > diskThreshold {
			if highDiskTime.IsZero() {
				highDiskTime = time.Now()
			} else if time.Since(highDiskTime) > time.Minute {
				// 如果超过 1 分钟，发送邮件并报告磁盘使用情况
				body := fmt.Sprintf("磁盘使用率超过 %.2f%%，持续超过 1 分钟。详细信息:\n%s", diskThreshold, diskDetails)

				// 发送邮件
				err := sendEmail(mailConfig, "磁盘使用率告警", body)
				if err != nil {
					log.Printf("发送邮件失败: %v", err)
				}

				// 重置计时器
				highDiskTime = time.Time{}
			}
		} else {
			highDiskTime = time.Time{}
		}

		time.Sleep(5 * time.Second)
	}
}
