**Project TravelBlog**
<br>
This is project using Backend Golang (Framework GORM+GIN)<br>
<br>
### Tutorial Setup
- Step 1: Clone code in gitlab
- Step 2: Go to the project and: 
    - `go run main.go`
### Folder Structure Project
```
`DTPT-GOLANG`
|__ common (shared components between features)
|
|__ module
|   |__ article
|       |__ business
|       |__ model
|       |__ storage
|       |__ transport
|           |__gin
|   |__ user
|       |__ business
|       |__ model
|       |__ storage
|       |__ transport
|           |__gin
|   |__ comment
|       |__ business
|       |__ model
|       |__ storage
|       |__ transport
|           |__gin
|__ gin_form
|__ go.mod
|__ go.som
|__ log.log
|__ main.go
```