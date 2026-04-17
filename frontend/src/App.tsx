import { useState } from 'react'
import { searchGames, type Game } from './api/client'

function App() {
  const [query, setQuery] = useState<string>('')
  const [games, setGames] = useState<Game[]>([])
  const [loading, setLoading] = useState<boolean>(false)
  const [error, setError] = useState<string | null>(null)

  async function handleSearch(e: React.FormEvent) {
    // prevent the browser from refreshing the page on form submit
    e.preventDefault()
    if (!query.trim()) return

    setLoading(true)
    setError(null)

    try {
      const results = await searchGames(query)
      setGames(results)
    } catch {
      setError('Something went wrong. Is the backend running?')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div>
      <h1>Gameboxd</h1>
      <p>Track the games you've played.</p>

      <form onSubmit={handleSearch}>
        <input
          type="text"
          value={query}
          onChange={e => setQuery(e.target.value)}
          placeholder="Search for a game..."
        />
        <button type="submit" disabled={loading}>
          {loading ? 'Searching...' : 'Search'}
        </button>
      </form>

      {error && <p>{error}</p>}

      <ul>
        {games.map(game => (
          <li key={game.appid}>
            {game.name} (AppID: {game.appid})
          </li>
        ))}
      </ul>
    </div>
  )
}

export default App