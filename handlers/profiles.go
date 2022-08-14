package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	profiledto "waysbucks/dto/profile"
	dto "waysbucks/dto/result"
	"waysbucks/models"
	"waysbucks/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlersProfile struct {
	ProfileRepository repositories.ProfileRepository
}

func HandlerProfile(ProfileRepository repositories.ProfileRepository) *handlersProfile {
	return &handlersProfile{ProfileRepository}
}

func (h *handlersProfile) FindProfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	profiles, err := h.ProfileRepository.FindProfiles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: profiles}
	json.NewEncoder(w).Encode(response)
}

func (h *handlersProfile) GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	profiles, err := h.ProfileRepository.GetProfile(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: convertResponseProfile(profiles)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlersProfile) CreateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(profiledto.CreateProfile)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "hai"}
		json.NewEncoder(w).Encode(response)
		return
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "hello"}
		json.NewEncoder(w).Encode(response)
		return
	}

	profile := models.Profile{
		Phone:      request.Phone,
		Address:    request.Address,
		City:       request.City,
		PostalCode: request.PostalCode,
		Image:      request.Image,
		UserID:     request.UserID,
	}

	data, err := h.ProfileRepository.CreateProfile(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: convertResponseProfile(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlersProfile) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(profiledto.UpdateProfile)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	profile, err := h.ProfileRepository.GetProfile(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	if request.Phone != "" {
		profile.Phone = request.Phone
	}
	if request.Address != "" {
		profile.Address = request.Address
	}
	if request.City != "" {
		profile.City = request.City
	}
	if request.PostalCode != 0 {
		profile.PostalCode = request.PostalCode
	}
	if request.Image != "" {
		profile.Image = request.Image
	}

	data, err := h.ProfileRepository.UpdateProfile(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: convertResponseProfile(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlersProfile) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	profile, err := h.ProfileRepository.GetProfile(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	fmt.Println(profile)

	data, err := h.ProfileRepository.DeleteProfile(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func convertResponseProfile(u models.Profile) profiledto.ProfileResponse {
	fmt.Println(u)
	return profiledto.ProfileResponse{
		Image:      u.Image,
		Phone:      u.Phone,
		Address:    u.Address,
		PostalCode: u.PostalCode,
		City:       u.City,
		UserID:     u.UserID,
		User:       u.User,
	}
}
