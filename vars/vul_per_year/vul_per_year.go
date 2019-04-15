package main

import (
	"fmt"

	"github.com/cbelk/vars/pkg/varsapi"
	//"github.com/wcharczuk/go-chart"
)

// Name the report
var Name = "Vulnerabilities by Year"

// GenerateReport sorted by year
func GenerateReport() (string, error) {
	// html creates the table that will show vulnerabilities by year
	html := `
            <div class="table-responsive">
            <table class="table table-striped table-dark table-bordered table-hover" id="num-table">
                <thead>
                    <tr>
                        <th scope="col" class="col-3" onclick="sortTable('num-table', 0)">Year</th>
                        <th scope="col" class="col-3" onclick="sortTable('num-table', 1)">Count</th>
                    </tr>
                </thead>
                <tbody>
                    %s
                </tbody>
            </table>
            </div>
            `
	// Get all vulnerabilities
	vulns, err := varsapi.GetVulnerabilities()
	if !varsapi.IsNilErr(err) {
		return "", err
	}

	count := make(map[int]int)

	for _, v := range vulns {
		year, _, _ := v.Dates.Initiated.Date()
		_, ok := count[year]
		if ok {
			count[year]++
		} else {
			count[year] = 1
		}
	}
	s := ""
	for y, c := range count {
		s += fmt.Sprintf("<tr><td>%d</td><td>%d</td></tr>", y, c)
	}
	return fmt.Sprintf(html, s), nil
}
