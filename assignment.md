## Assignment


### Notes:
```
* Document all steps performed in each task. Update README.md with steps to setup your project.  We should be able to set up the project while reading the README.md.
* Make sure all the files you create have the exact same names as provided.
* Don't commit all of your work as a single commit, commit it as you finish each part, so that we can see the work as you built it up.
* The solution should work on a different machine, which has docker, without any modifications.
* If you find any mistake or confusion in the tasks, please get it cleared.  If you make assumptions, you should include that in your documentation.
* Read these notes carefully, that is an important part to solve the assignment.
```

### Build Events Collector

Events collector is a program to capture and view events of any backend application.  (Similar to Google Analytics)

> Requirements
1. Setup a PostgreSQL database.  Create required tables.  Structure data in relational form (rows and columns) instead of storing JSON in single column.
2. Write REST API in Golang to:

   a. Capture single event in single request and store in DB.
   
   b. View all events.
   

3. REST API has to run on a custom port `9999`.
4. Do not use any available web server like Apache, Nginx or any other for this purpose.
5. Update your documentation with all details you did in above steps.
