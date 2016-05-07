package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	// "atlas-api/config/db"
	"atlas-api/config/schema"
	"atlas-api/helpers"
)

// ProjectReq ...
type ProjectReq struct {
	Name           string `json:"name"`
	Client         string `json:"client"`
	SlackChannel   string `json:"slackChannel"`
	StartDate      string `json:"startDate"`
	OrganizationID int    `json:"organizationId"`
}

// CreateProject will create a new project
func CreateProject(rw http.ResponseWriter, req *http.Request) {
	var projectReq ProjectReq
	var project schema.Project

	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		helper.ProjectResponse(rw, req, 200, project, nil)
		return
	}
	if err := req.Body.Close(); err != nil {
		helper.ProjectResponse(rw, req, 200, project, nil)
		return
	}

	if err := json.Unmarshal(body, &projectReq); err != nil {
		helper.ProjectResponse(rw, req, 200, project, nil)
		return
	}

	project.Name = projectReq.Name
	project.Client = projectReq.Client
	project.SlackChannel = projectReq.SlackChannel
	project.StartDate = projectReq.StartDate
	project.OrganizationID = projectReq.OrganizationID

	// if err := db.DB.Create(&project).Error; err != nil {
	//
	// 	err = helper.HandleError(rw, req, 400, err)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	return
	// }

	helper.ProjectResponse(rw, req, 200, project, nil)
}
