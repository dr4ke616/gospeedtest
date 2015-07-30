package go_speedtest

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Speedtest struct {
	FileLocation string
	Verbos       bool
}

func (s *Speedtest) MakeRequest() (*http.Request, error) {
	req, err := http.NewRequest("GET", s.FileLocation, nil)
	if err != nil {
		return nil, err
	}

	// Set HTTP Request User-Agent & Referer header
	req.Header.Set("User-Agent", "Mozilla/5.0")
	return req, nil
}

func (s *Speedtest) RunSpeedtest(request *http.Request, client *http.Client) (int, error) {
	t1 := time.Now()

	resp, err := client.Do(request)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}

	t2 := time.Now()

	return s.CalculateRate(len(body), t1, t2), nil
}

func (s *Speedtest) CalculateRate(data_size int, t1 time.Time, t2 time.Time) int {

	// Conver t2.Sub(t1) value to seconds unit.
	elapsed_second_time := int(t2.Sub(t1) / 1000000000)

	// Convert the size of the file from bytes to Maga Bits
	size_magabits := ((data_size * 8) / 1024 / 1024)

	if s.Verbos {
		log.Printf("File Size: %dMb", size_magabits)
		log.Printf("File Size: %dMB", size_magabits/8)
		log.Printf("Elapsed Time: %ds", elapsed_second_time)
	}

	// Divide magabits by time it took to get Mbps
	return size_magabits / elapsed_second_time
}

func (s *Speedtest) Start() int {
	req, err := s.MakeRequest()
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}

	var rate int = -1
	rate, err = s.RunSpeedtest(req, client)
	if err != nil || rate == -1 {
		log.Fatalln(err)
	}

	if s.Verbos {
		log.Printf("The download rate is %d Mbps.\n", rate)
	}

	return rate
}
