package repository

import (
	"encoding/json"
	"os"

	"hacktown-cli/internal/domain"
)

type HacktownDataRepository struct {
	data map[string]interface{}
}

func NewHacktownDataRepository() *HacktownDataRepository {
	repo := &HacktownDataRepository{}
	repo.loadData()
	return repo
}

func (r *HacktownDataRepository) loadData() {
	file, err := os.ReadFile("data/hacktown.json")
	if err != nil {
		r.data = make(map[string]interface{})
		return
	}

	err = json.Unmarshal(file, &r.data)
	if err != nil {
		r.data = make(map[string]interface{})
	}
}

func (r *HacktownDataRepository) GetHacktownInfo() string {
	jsonData, err := json.MarshalIndent(r.data, "", "  ")
	if err != nil {
		return ""
	}
	return string(jsonData)
}