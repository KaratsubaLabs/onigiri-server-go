package api

type route struct {
	name string
	method map[string]Handler
}

var routeSchema = []route {
	{
		name: "/login",
		method: map[string]Handler {
			"POST": postLogin,
		},
	},
	{
		name: "/music/play",
		method: map[string]Handler {
			"POST": postPlayMusic,
		},
	},
	{
		name: "/music/queue",
		method: map[string]Handler {
			"POST": postPlayMusic,
		},
	},
	{
		name: "/music/youtube",
		method map[string]Handler {
			"POST": postYoutube,
			"Delete": deleteYoutube,
		},
	},
	{
		name: "/music/spotify/account",
		method map[string]Handler {
			"POST": postSpotifyAccount,
			"DELETE": deleteSpotifyAccount,
		},
	},
	{
		name: "/music/spotify/playlist",
		method map[string]Handler {
			"POST": postSpotifyPlaylist,
			"DELETE": deleteSpotifyPlaylist,
		},
	},
}

