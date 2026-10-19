package main

var regionRU = regionDef{
	Domain: "sh-development.ru",
	Lang:   "ru",

	Text: pageText{
		Login: "войти",
		About: "Плеер для длинных аудиофайлов. Запоминает место остановки и позволяет вручную указать время начала воспроизведения.",

		TabPlayer:  "Плеер",
		TabLibrary: "Библиотека",
		Apps:       "apps",
		Readme:     "readme",
		Logout:     "выйти",

		Volume:   "Громкость",
		Previous: "Предыдущий",
		Next:     "Следующий",
		IdleMsg:  "Выберите трек во вкладке «Библиотека»",

		JS: jsText{
			StorageDown:  "Яндекс Хранилище недоступно,\nпопробуйте позже",
			AutoplayOn:   "Автовоспроизведение: вкл",
			AutoplayOff:  "Автовоспроизведение: выкл",
			LibraryError: "Не удалось загрузить библиотеку.",
		},
	},
}
