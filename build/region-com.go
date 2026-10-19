package main

var regionCOM = regionDef{
	Domain: "sh-development.com",
	Lang:   "en",

	Text: pageText{
		Login: "log in",
		About: "A player for long audio files. It remembers where you stopped and lets you set the playback start time by hand.",

		TabPlayer:  "Player",
		TabLibrary: "Library",
		Apps:       "apps",
		Readme:     "readme",
		Logout:     "log out",

		Volume:   "Volume",
		Previous: "Previous",
		Next:     "Next",
		IdleMsg:  "Pick a track from the Library tab",

		JS: jsText{
			StorageDown:  "Storage is unavailable,\ntry again later",
			AutoplayOn:   "Autoplay: on",
			AutoplayOff:  "Autoplay: off",
			LibraryError: "Could not load library.",
		},
	},
}
