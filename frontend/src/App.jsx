import { useEffect, useState } from "react";
import "./App.css";

import {
  getLeagueTable,
  getMatches,
  getPredictions,
  initializeLeague,
  playAll,
  playWeek,
  updateMatchResult,
} from "./api/leagueApi";

function App() {
  const [table, setTable] = useState([]);
  const [matches, setMatches] = useState([]);
  const [predictions, setPredictions] = useState([]);
  const [selectedWeek, setSelectedWeek] = useState(1);
  const [message, setMessage] = useState("");
  const [loading, setLoading] = useState(false);

  const refreshData = async () => {
    try {
      setLoading(true);

      const tableRes = await getLeagueTable();
      const matchesRes = await getMatches();
      const predictionRes = await getPredictions();

      setTable(tableRes.data);
      setMatches(matchesRes.data);
      setPredictions(predictionRes.data);
    } catch (error) {
      console.error(error);
      setMessage("Could not fetch data. Make sure backend is running.");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    refreshData();
  }, []);

  const handleInit = async () => {
    try {
      await initializeLeague();
      setMessage("League initialized successfully.");
      await refreshData();
    } catch (error) {
      setMessage(error.response?.data?.error || "Could not initialize league.");
    }
  };

  const handlePlayWeek = async () => {
    try {
      await playWeek(selectedWeek);
      setMessage(`Week ${selectedWeek} played successfully.`);
      await refreshData();
    } catch (error) {
      setMessage(error.response?.data?.error || "Could not play selected week.");
    }
  };

  const handlePlayAll = async () => {
    try {
      await playAll();
      setMessage("All remaining matches played successfully.");
      await refreshData();
    } catch (error) {
      setMessage(error.response?.data?.error || "Could not play all matches.");
    }
  };

  const handleUpdateMatch = async (matchId, homeGoals, awayGoals) => {
    try {
      await updateMatchResult(matchId, homeGoals, awayGoals);
      setMessage("Match result updated successfully.");
      await refreshData();
    } catch (error) {
      setMessage(error.response?.data?.error || "Could not update match.");
    }
  };

  const teamNameById = table.reduce((acc, team) => {
    acc[team.team_id] = team.team_name;
    return acc;
  }, {});

  const filteredMatches = matches.filter(
    (match) => match.week === Number(selectedWeek)
  );

  return (
    <div className="app">
      <header className="header">
        <h1>Insider League Simulation</h1>
        <p>4-team Premier League-style simulation dashboard</p>
      </header>

      {message && <div className="message">{message}</div>}
      {loading && <div className="message">Loading...</div>}

      <section className="controls">
        <button onClick={handleInit}>Initialize League</button>

        <select
          value={selectedWeek}
          onChange={(event) => setSelectedWeek(Number(event.target.value))}
        >
          {[1, 2, 3, 4, 5, 6].map((week) => (
            <option key={week} value={week}>
              Week {week}
            </option>
          ))}
        </select>

        <button onClick={handlePlayWeek}>Play Selected Week</button>
        <button onClick={handlePlayAll}>Play All League</button>
      </section>

      <main className="grid">
        <section className="card table-card">
          <h2>League Table</h2>
          <LeagueTable table={table} />
        </section>

        <section className="card">
          <h2>Week {selectedWeek} Matches</h2>
          <MatchList
            matches={filteredMatches}
            teamNameById={teamNameById}
            onUpdateMatch={handleUpdateMatch}
          />
        </section>

        <section className="card">
          <h2>Championship Predictions</h2>
          <PredictionList predictions={predictions} />
        </section>
      </main>
    </div>
  );
}

function LeagueTable({ table }) {
  if (table.length === 0) {
    return <p>No table data yet. Click Initialize League.</p>;
  }

  return (
    <table>
      <thead>
        <tr>
          <th>Team</th>
          <th>P</th>
          <th>W</th>
          <th>D</th>
          <th>L</th>
          <th>GF</th>
          <th>GA</th>
          <th>GD</th>
          <th>Pts</th>
        </tr>
      </thead>

      <tbody>
        {table.map((team) => (
          <tr key={team.team_id}>
            <td>{team.team_name}</td>
            <td>{team.played}</td>
            <td>{team.won}</td>
            <td>{team.drawn}</td>
            <td>{team.lost}</td>
            <td>{team.goals_for}</td>
            <td>{team.goals_against}</td>
            <td>{team.goal_difference}</td>
            <td>
              <strong>{team.points}</strong>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}

function MatchList({ matches, teamNameById, onUpdateMatch }) {
  const [scores, setScores] = useState({});

  const handleChange = (matchId, field, value) => {
    setScores((previousScores) => ({
      ...previousScores,
      [matchId]: {
        ...previousScores[matchId],
        [field]: value,
      },
    }));
  };

  if (matches.length === 0) {
    return <p>No matches available. Initialize the league first.</p>;
  }

  return (
    <div className="matches">
      {matches.map((match) => {
        const homeGoals =
          scores[match.id]?.home_goals ?? match.home_goals ?? 0;

        const awayGoals =
          scores[match.id]?.away_goals ?? match.away_goals ?? 0;

        return (
          <div className="match" key={match.id}>
            <div>
              <strong>Match #{match.id}</strong>
              <p>
                {teamNameById[match.home_team_id] ||
                  `Team ${match.home_team_id}`}{" "}
                vs{" "}
                {teamNameById[match.away_team_id] ||
                  `Team ${match.away_team_id}`}
              </p>
              <p>Status: {match.played ? "Played" : "Not Played"}</p>
            </div>

            <div className="score-editor">
              <input
                type="number"
                min="0"
                value={homeGoals}
                onChange={(event) =>
                  handleChange(match.id, "home_goals", event.target.value)
                }
              />

              <span>-</span>

              <input
                type="number"
                min="0"
                value={awayGoals}
                onChange={(event) =>
                  handleChange(match.id, "away_goals", event.target.value)
                }
              />

              <button
                onClick={() => onUpdateMatch(match.id, homeGoals, awayGoals)}
              >
                Update
              </button>
            </div>
          </div>
        );
      })}
    </div>
  );
}

function PredictionList({ predictions }) {
  if (predictions.length === 0) {
    return <p>No predictions yet.</p>;
  }

  return (
    <div className="predictions">
      {predictions.map((prediction) => (
        <div className="prediction" key={prediction.team_id}>
          <span>{prediction.team_name}</span>
          <strong>{prediction.probability.toFixed(1)}%</strong>
        </div>
      ))}
    </div>
  );
}

export default App;