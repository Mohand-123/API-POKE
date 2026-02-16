# 🌟 Pokémon Tracker

Application web moderne pour explorer et gérer votre collection de Pokémon en utilisant l'API PokeBuild.

## ✨ Fonctionnalités

- **Liste complète** : Parcourez tous les Pokémon avec pagination
- **Recherche avancée** : Trouvez des Pokémon par nom avec suggestions en temps réel
- **Filtrage par type** : Filtrez les Pokémon par catégories (Feu, Eau, Plante, etc.)
- **Système de favoris** : Sauvegardez vos Pokémon préférés
- **Design moderne** : Interface sombre et élégante avec animations fluides
- **Responsive** : Fonctionne sur tous les appareils

## 🛠️ Technologies utilisées

- **Backend** : Go (Golang)
- **Frontend** : HTML5, CSS3, JavaScript
- **API** : PokeBuild API
- **Font** : Poppins (Google Fonts)

## 🚀 Installation

1. Clonez le dépôt ou téléchargez les fichiers

2. Naviguez dans le dossier du projet :
```bash
cd "API POKEMONES/src"
```

3. Lancez le serveur :
```bash
go run cmd/main.go
```

4. Ouvrez votre navigateur et accédez à :
```
http://localhost:8080
```

## 📁 Structure du projet

```
API POKEMONES/
├── src/
│   ├── cmd/
│   │   └── main.go              # Point d'entrée de l'application
│   ├── controllers/
│   │   ├── pokemons.controllers.go
│   │   └── errors.controller.go
│   ├── models/
│   │   ├── pokemons.model.go
│   │   └── errors.model.go
│   ├── services/
│   │   └── pokemons.service.go  # Appels API
│   ├── routes/
│   │   ├── main.routes.go
│   │   ├── pokemons.routes.go
│   │   └── errors.router.go
│   └── helper/
│       ├── templates.helper.go
│       └── favorites.helper.go
├── templates/
│   ├── home.html                # Page d'accueil
│   ├── list_pokemons.html       # Liste des Pokémon
│   ├── search.html              # Page de recherche
│   ├── categories.html          # Filtrage par type
│   ├── favorites.html           # Pokémon favoris
│   └── error.html               # Page d'erreur
├── assets/
│   └── dark-theme.css           # Styles CSS
└── favoris.json                 # Stockage des favoris
```

## 🎮 Utilisation

### Page d'accueil
- Accédez aux différentes sections via les boutons de navigation

### Recherche
- Tapez le nom d'un Pokémon dans la barre de recherche
- Les suggestions apparaissent automatiquement
- Cliquez sur un résultat pour voir les détails

### Catégories
- Entrez un type de Pokémon (ex: "Feu", "Eau", "Plante")
- Cliquez sur "Filtrer" pour voir tous les Pokémon de ce type

### Favoris
- Ajoutez des Pokémon à vos favoris depuis la page de détails
- Gérez votre collection dans la section "Favoris"
- Supprimez des favoris en cliquant sur le bouton "×"


## 📝 API

L'application utilise l'API PokeBuild pour récupérer les données des Pokémon.

Endpoints utilisés :
- Liste des Pokémon : `/pokemon`
- Détails d'un Pokémon : `/pokemon/{id}`
- Recherche : `/pokemon?name={query}`


