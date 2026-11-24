
├── main.go              # démarrage du serveur + routes
├── api/                 # tout ce qui parle à l'API (client, structs, cache)
│   └── client.go
├── handlers/            # une fonction par page
├── logic/               # recherche, filtres, pagination (Go pur, sans HTTP)
├── favorites/           # lecture/écriture du fichier JSON
├── templates/           # header.html, footer.html, pages...
├── static/              # style.css, images
├── data/favorites.json
├── README.md
└── .gitignore
