package services

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/i-ceu/go-project-manager/internal/config"
	"github.com/i-ceu/go-project-manager/internal/enums"
	"github.com/i-ceu/go-project-manager/internal/helpers"
	"github.com/i-ceu/go-project-manager/internal/models"
	"github.com/i-ceu/go-project-manager/internal/requests"
	"gorm.io/gorm"
)

func generateProjectTag(projectName string) string {
	splitName := strings.Split(projectName, " ")
	if len(splitName) > 1 {
		accronymArray := []string{}
		for _, v := range splitName {
			accronymArray = append(accronymArray, v[0:1])
		}
		projectAccronym := strings.Join(accronymArray, "")
		projectTag := strings.ToUpper(projectAccronym)
		return projectTag
	}

	projectTag := strings.ToUpper(projectName[:3])

	return projectTag
}

func CreateProject(req *requests.CreateProjectRequest, teamId string) (*models.Project, error) {

	startDate, err := time.Parse(enums.Date_format, req.StartDate)
	if err != nil {
		return nil, errors.New("invalid date format")
	}

	dueDate, err := time.Parse(enums.Date_format, req.DeliveryDate)
	if err != nil {
		return nil, errors.New("invalid date format")
	}

	var team models.Team
	if err = config.DB.First(&team, "id = ?", teamId).Error; err != nil {
		return nil, errors.New("team not found")
	}
	projectTag := generateProjectTag(req.Title)

	project := models.Project{
		Title:        req.Title,
		Tag:          projectTag,
		Description:  req.Description,
		Status:       req.Status,
		TeamID:       teamId,
		StartDate:    &startDate,
		DeliveryDate: &dueDate,
		WorkFlow:     models.StringWorkflow(req.WorkFlow),
	}

	test := config.DB.Create(&project)
	if test.Error != nil {
		return nil, test.Error
	}

	return &project, nil
}

func GetProject(projectId *string, userIds []string, search *string) (*models.Project, error) {
	var project models.Project

	result := config.DB.Preload("Tasks", func(db *gorm.DB) *gorm.DB {
		// Always preload AssignedTo relationship
		db = db.Preload("AssignedTo")

		// Handle user ID filtering
		if len(userIds) > 0 {
			var regularIDs []string
			hasUnassigned := false

			for _, id := range userIds {
				if id == "unassigned" {
					hasUnassigned = true
				} else if id != "" {
					regularIDs = append(regularIDs, id)
				}
			}

			// Build WHERE clause for user IDs
			switch {
			case hasUnassigned && len(regularIDs) > 0:
				db = db.Where("assigned_to_id IS NULL OR assigned_to_id IN (?)", regularIDs)
			case hasUnassigned:
				db = db.Where("assigned_to_id IS NULL")
			case len(regularIDs) > 0:
				db = db.Where("assigned_to_id IN (?)", regularIDs)
			}
		}

		// Handle search if provided
		if search != nil && *search != "" {
			searchTerm := "%" + strings.ToLower(*search) + "%"
			db = db.Where(
				"LOWER(title) LIKE ? OR LOWER(description) LIKE ?",
				searchTerm,
				searchTerm,
			)
		}

		return db
	}).
		First(&project, projectId)

	if result.Error != nil {
		return nil, errors.New("no project with this Id")
	}

	return &project, nil
}

func GetAllProjects(teamId *string) (*[]models.Project, error) {
	var projects []models.Project

	result := config.DB.Where("team_id =?", teamId).Find(&projects)
	if result.Error != nil {
		return nil, errors.New("error retrieving projects")
	}

	return &projects, nil
}

func GenerateTasks(projectId *string, userId string) (string, error) {
	var project models.Project

	proj := config.DB.Where("id =?", projectId).Find(&project)

	if proj.Error != nil {
		return " ", errors.New("no project with this Id")
	}

	base_prompt, err := os.ReadFile("query-method.md")
	if err != nil {
		return "fail", err
	}

	prompt := string(base_prompt) + " " + project.Title + ". " + project.Description

	groqResponse, err := helpers.CallGroqAPI(prompt)
	if err != nil {
		return "fail", err
	}

	taskInserts, err := helpers.ParseGroqResponse(groqResponse)
	if err != nil {
		return "fail", err
	}

	numWorkers := 10

	bulkInsertTasks(taskInserts, numWorkers, userId, project.ID)
	config.DB.Model(&project).Update("tasks_ai_generated", true)
	return "Success", nil

}

func ResetProject(projectId *string) (string, error) {
	var project models.Project

	err := config.DB.Table("tasks").Where("project_id = ?", projectId).Delete(nil)
	if err.Error != nil {
		return "fail", err.Error
	}

	err = config.DB.Model(&project).Where("id =?", projectId).Update("tasks_ai_generated", false)
	if err.Error != nil {
		return "fail", err.Error
	}

	return "Success", nil
}

func insertTasksRecord(record requests.GroqTasksResponse, userId string, projectId string) (string, error) {
	taskRequest := requests.CreateTaskRequest{
		Title:       record.TaskTitle,
		Description: record.TaskDescription,
	}

	_, err := CreateTask(&taskRequest, userId, projectId)
	if err != nil {
		return "", err
	}
	return "Creating Tasks in progress", nil
}

func bulkInsertTasks(records []requests.GroqTasksResponse, numWorkers int, userId string, projectId string) {
	jobs := make(chan requests.GroqTasksResponse, len(records))
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for record := range jobs {
				fmt.Printf("Worker %d processing: %s\n", workerID, record.TaskTitle)
				insertTasksRecord(record, userId, projectId)
			}
		}(i)
	}

	for _, record := range records {
		jobs <- record
	}
	close(jobs)

	wg.Wait()
}
