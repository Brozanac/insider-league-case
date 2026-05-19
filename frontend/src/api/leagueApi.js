import axios from "axios";

const API_BASE_URL = "http://localhost:8080";

export const initializeLeague = () => {
  return axios.post(`${API_BASE_URL}/league/init`);
};

export const getLeagueTable = () => {
  return axios.get(`${API_BASE_URL}/league/table`);
};

export const getMatches = () => {
  return axios.get(`${API_BASE_URL}/matches`);
};

export const playWeek = (week) => {
  return axios.post(`${API_BASE_URL}/league/play/week/${week}`);
};

export const playAll = () => {
  return axios.post(`${API_BASE_URL}/league/play/all`);
};

export const getPredictions = () => {
  return axios.get(`${API_BASE_URL}/league/predictions`);
};

export const updateMatchResult = (matchId, homeGoals, awayGoals) => {
  return axios.put(`${API_BASE_URL}/matches/${matchId}`, {
    home_goals: Number(homeGoals),
    away_goals: Number(awayGoals),
  });
};