package main

import (
	"fmt"
	//"time"
	//"github.com/wcharczuk/go-chart"
	//"github.com/cbelk/vars/pkg/varsapi"
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
                        <th scope="col" class="col-3" onclick="sortTable('num-table', 0)">2018</th>
                        <th scope="col" class="col-3" onclick="sortTable('num-table', 1)">2019</th>
                    </tr>
                </thead>
                <tbody>
                    %s
                </tbody>
            </table>
            </div>
            `
	a := 0
	b := 0

	s := fmt.Sprintf("<tr><td>%d</td><td>%d</td></tr>", a, b)
	return fmt.Sprintf(html, s), nil
}
