package main

import (
	"net"
	"os"
	"runtime"
	"strings"
)

type HostsFile struct {
	entries []HostsEntry
}

func (h HostsFile) Length() int {
	return len(h.entries)
}

func (h HostsFile) GetByIp(ip string) HostsEntry {
	valid_ip := net.ParseIP(ip)

	if valid_ip == nil {
		panic("GetByIp: Invalid IP")
	}

	for _, entry := range h.entries {
		if entry.ip.Equal(valid_ip) {
			return entry
		}
	}

	return HostsEntry{}
}

func (h HostsFile) Remove(index int) {
	h.entries = append(h.entries[:index], h.entries[index+1:]...)
}

func (h HostsFile) Set(ip string, host string, enabled bool) {
	valid_ip := net.ParseIP(ip)

	if valid_ip == nil {
		panic("Set: Invalid IP")
	}

	h.entries = append(h.entries, HostsEntry{
		ip:      valid_ip,
		host:    host,
		enabled: enabled,
		line:    h.Length() + 1,
	})
}

type HostsEntry struct {
	host    string
	ip      net.IP
	enabled bool
	line    int
	comment string
}

var HOSTS_FILE_PATH = map[string]string{
	"windows": "C:\\Windows\\System32\\drivers\\etc\\hosts",
	"linux":   "/etc/hosts",
	"darwin":  "/etc/hosts",
}

func parse_line(line string) (net.IP, string, bool) {
	host_ip_pair := strings.Split(line, " ")

	if len(host_ip_pair) == 2 {
		ip := net.ParseIP(host_ip_pair[0])
		host := host_ip_pair[1]

		if ip != nil {
			return ip, host, false
		}
	}

	return nil, "", true
}

func parse_hosts_file(hosts_file []byte) []HostsEntry {
	hosts := make([]HostsEntry, 0)

	// split file into lines
	lines := strings.Split(string(hosts_file), "\n")

	for index, line := range lines {
		entry := HostsEntry{
			line: index,
		}

		if line == "" {
			continue
		}

		if line[0] != '#' {
			entry.enabled = true
		} else {
			line = line[1:]
		}

		line = strings.TrimSpace(line)

		ip, host, comment := parse_line(line)

		if comment {
			entry.comment = line
			entry.enabled = false
		} else {
			entry.ip = ip
			entry.host = host
		}

		hosts = append(hosts, entry)
	}
	return hosts
}

func load_hosts_file() HostsFile {
	os_name := runtime.GOOS

	hosts_file_path, ok := HOSTS_FILE_PATH[os_name]

	if !ok {
		panic("Unsupported OS")
	}

	hosts_file, err := os.Open(hosts_file_path)

	if err != nil {
		panic(err)
	}

	defer hosts_file.Close()

	// read file size
	file_info, err := hosts_file.Stat()

	if err != nil {
		panic(err)
	}

	file_size := file_info.Size()

	// read file to a line buffer

	buffer := make([]byte, file_size)

	_, err = hosts_file.Read(buffer)

	if err != nil {
		panic(err)
	}

	return HostsFile{
		entries: parse_hosts_file(buffer),
	}
}

func main() {
	hosts := load_hosts_file()

	hosts.Set("192.168.0.1", "localhost", true)

	hosts.GetByIp("192.168.0.14")
}
