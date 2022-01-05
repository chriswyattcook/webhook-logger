package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
)

type APIResponse struct {
	Message string                    `json:"message"`
	Squawk  WorkflowJobWebhookPayload `json:"squawk"`
}

func workflowJobSquawk(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}

	var wj WorkflowJobWebhookPayload
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&wj)

	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}

	squawkResponse(w, wj)
}

func workflowJob(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}

	var wj WorkflowJobWebhookPayload
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&wj)

	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}

	// we only want to log completed/skipped/failed workflows
	// TODO: change this to the actions that matter
	// if wj.Action == "in_progress" {
	// 	logger.WithFields()
	// }

	// parse only info that is needed from the payload
	wjData := logrus.Fields{
		"action":           wj.Action,
		"repo.id":          wj.Repository.ID,
		"repo.name":        wj.Repository.Name,
		"workflow.id":      wj.WorkflowJob.ID,
		"workflow.runId":   wj.WorkflowJob.RunID,
		"workflow.headSha": wj.WorkflowJob.HeadSha,
		"conclusion":       wj.WorkflowJob.Conclusion,
		"startedAt":        wj.WorkflowJob.StartedAt,
		"completedAt":      wj.WorkflowJob.CompletedAt,
		"totalSteps":       len(wj.WorkflowJob.Steps),
		"steps":            wj.WorkflowJob.Steps,
	}

	logger.WithFields(wjData).Info("log received")
}

func squawkResponse(w http.ResponseWriter, wj WorkflowJobWebhookPayload) {
	data := APIResponse{
		Message: "success",
		Squawk:  wj,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
