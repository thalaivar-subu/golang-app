# golang-app
- Front End -> ReactJS
- Back End -> GoLang

Built a web UI per below. 
1) It has a text box to enter a URL. After URL submission, it will:
    - Grabs the content of URL
    - Strips all HTML tag thus only the real-content left
    - Gets the count of all words and presented in tabular manner

2) Excel Colum Finder 
  - It has 3 entries:
  - Column to start, number of rows and number of columns
  - After the submit, it will:
  - Finds excel column started from entered column and displayed the column per below
  ie.1) if enter A, 3 and 2 then it will return A, B, C, D, E, F and display ABC and DEF on each row
  ie.2) if enter AZ, 3 and 3 then it will return AZ, BA, BB, BC, BD, BE, BF, BG, BH and display AZ BA BB and BC BD BE and BF BG BH.

3) CRUD API's using Golang.
- Insert, Update, Delete & View in User Details Using Golang with Mysql.

4) Sample Service Using Golang
  - Creates an Ajax call that receives data from the client in json format in its request body (call type POST request). 
  - Input in the form of clock and date like (Request). The output is JSON with a date & end day of date
  Request:
  {Date: "2020-02-18T01:50"}
  Result:
  {Date: "2020-02-18T01:50", LastDayOfMonth:29}

5) Goroutine.
- Creates an ajax call to “Find each number either Prime or Not equal to ‘N’ Input”. Used Goroutine to perform parallel calculations of each number.
