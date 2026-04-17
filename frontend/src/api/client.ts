//shape of game returned from api
export interface Game {
    appid: number
    name: string
}

//calls go backend which calls steam
export async function searchGames(query: string): Promise<Game[]> {
    const response = await fetch(`/api/games/search?q=${encodeURIComponent(query)}`)

    if (!response.ok) {
        throw new Error('Failed to fetch games')
    }

    return response.json()
}  
