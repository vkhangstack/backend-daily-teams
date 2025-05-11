package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/vkhangstack/dlt/internal/core/domain/dto"
	"github.com/vkhangstack/dlt/internal/core/domain/model"
	"github.com/vkhangstack/dlt/internal/core/ports"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type DailyService struct {
	repo ports.DailyRepository
}

func NewDailyService(repo ports.DailyRepository) *DailyService {
	return &DailyService{repo: repo}
}

func (d *DailyService) CreateTask(payload *dto.CreateTaskDto, userId uint64) (*model.Task, error) {
	return d.repo.CreateTask(payload, userId)
}

func (d *DailyService) UpdateTask(payload *dto.UpdateTaskDto, userId uint64) error {
	return d.repo.UpdateTask(payload, userId)
}
func (d *DailyService) DeleteTask(id uint64, userId uint64) error {
	return d.repo.DeleteTask(id, userId)
}
func (d *DailyService) ListTasks(userId uint64) ([]*model.Task, error) {
	return d.repo.ListTasks(userId)
}
func (d *DailyService) DailyTask(userId uint64, payload *dto.DailyTaskDto) string {
	todayDaily := d.repo.CheckDaily(userId)
	if todayDaily == true {
		return "Today is daily task"
	}

	content := ""
	if payload.Content == "" {
		content = "<p><strong>Yesterday:</strong></p>"
		// Get yesterday's date in UTC
		yesterday := time.Now().AddDate(0, 0, -1).Truncate(24 * time.Hour)
		today := time.Now().Truncate(24 * time.Hour)
		result, err := d.repo.ListTasks(userId)
		if err != nil {
			return "Something went wrong that we didn't want."
		}
		// Find the event with `end` matching yesterday
		for i, e := range result {
			startTime, err := time.Parse(time.RFC3339, e.End)
			if err != nil {
				fmt.Printf("invalid start time: %v\n", err)
				continue
			}

			if time.Now().Weekday() == time.Monday {
				yesterday = time.Now().AddDate(0, 0, -3).Truncate(24 * time.Hour)
			}

			if startTime.Truncate(24 * time.Hour).Equal(yesterday) {
				if !strings.Contains(result[i].Content, "<ul>") && !strings.Contains(result[i].Content, "<p>") {
					result[i].Content = "<ul><li>" + result[i].Content + "</li></ul>"
				}
				content += result[i].Content
			}

			if startTime.Truncate(24 * time.Hour).Equal(today) {
				if !strings.Contains(result[i].Content, "<ul>") {
					result[i].Content = "<ul><li>" + result[i].Content + "</li></ul>"
				}
				content += "<p><strong>Today:</strong></p>" + result[i].Content
			}
		}
	} else {
		content = payload.Content
	}

	teamId := os.Getenv("TEAMS_TEAM_ID")
	chanelId := os.Getenv("TEAMS_CHANNEL_ID")
	teamsUrl := os.Getenv("TEAMS_URL")

	if teamsUrl == "" && chanelId == "" && teamId == " " {
		return "Something went wrong that we didn't want."
	}
	////GET /teams/{team-id}/channels/{channel-id}/messages/{message-id}
	//url := "https://graph.microsoft.com/v1.0/teams/teamId/channels/chanelId/messages"
	//url := "https://graph.microsoft.com/v1.0/chats/19:0c20f454-b4a4-4827-8587-c5c42a9c45fa_d6df20d9-94ad-47b4-a774-cfd1e4369cdf@unq.gbl.spaces/messages"

	client := &http.Client{}
	url := fmt.Sprintf("https://graph.microsoft.com/v1.0/teams/%s/channels/%s/messages", teamId, chanelId)

	if teamsUrl != "" {
		url = teamsUrl
	}

	// send teams
	var reqBody struct {
		Body struct {
			Content string `json:"content"`
		} `json:"body"`
	}
	reqBody.Body.Content = content

	jsonBytes, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+payload.AccessToken)

	resp, err := client.Do(req)

	if err != nil {
		return "Send teams failed!"
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("failed to close response body: %v\n", err)
		}
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	_ = fmt.Errorf("error %d", resp.StatusCode)
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return "Send teams failed!"
	}
	err = d.repo.CreateDaily(userId)
	if err != nil {
		return "Create daily failed!"
	}
	return ""
}
