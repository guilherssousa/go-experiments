package main

import (
	"encoding/json"
	"net"
	"os"
	"runtime"
	"strings"
)

type HostsFile struct {
	Entries []HostsEntry
}

type HostsEntry struct {
	Host    string `json:"host,omitempty"`
	Ip      net.IP `json:"ip,omitempty"`
	Enabled bool   `json:"enabled"`
	Line    int    `json:"line"`
	Comment string `json:"comment,omitempty"`
}

func (h HostsFile) Length() int {
	return len(h.Entries)
}

func (h HostsFile) ToJSON() []byte {
	object, _ := json.Marshal(h.Entries)

	return object
}

func (h HostsFile) GetByIp(ip any) HostsEntry {
	var valid_ip net.IP

	switch ip.(type) {
	case net.IP:
		valid_ip = ip.(net.IP)
	case string:
		valid_ip = net.ParseIP(ip.(string))
	}

	if valid_ip == nil {
		panic("GetByIp: Invalid IP")
	}

	for _, entry := range h.Entries {
		if entry.Ip.Equal(valid_ip) {
			return entry
		}
	}

	return HostsEntry{}
}

func (h HostsFile) Remove(index int) {
	h.Entries = append(h.Entries[:index], h.Entries[index+1:]...)
}

func (h HostsFile) Set(ip any, host string, enabled bool) {
	var valid_ip net.IP

	switch ip.(type) {
	case net.IP:
		valid_ip = ip.(net.IP)
	case string:
		valid_ip = net.ParseIP(ip.(string))
	default:
		panic("Set: the ip argument must be of type net.IP or string")
	}

	if valid_ip == nil {
		panic("Set: Invalid IP")
	}

	h.Entries = append(h.Entries, HostsEntry{
		Ip:      valid_ip,
		Host:    host,
		Enabled: enabled,
		Line:    h.Length() + 1,
	})
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
			Line: index,
		}

		if line == "" {
			continue
		}

		if line[0] != '#' {
			entry.Enabled = true
		} else {
			line = line[1:]
		}

		line = strings.TrimSpace(line)

		ip, host, comment := parse_line(line)

		if comment {
			if line == "" {
				continue
			}
			entry.Comment = line
			entry.Enabled = false
		} else {
			entry.Ip = ip
			entry.Host = host
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
		Entries: parse_hosts_file(buffer),
	}
}

func main() {
	hosts := load_hosts_file()

	hosts_json := hosts.ToJSON()

	os.WriteFile("hosts.json", hosts_json, 0666)
}
