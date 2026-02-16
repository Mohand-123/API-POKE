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

## 📋 Prérequis

- Go 1.16 ou supérieur
- Connexion Internet (pour l'API PokeBuild)

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

## 🎨 Design

L'application utilise un thème sombre moderne avec :
- Fond noir (#0f0f23)
- Effets de glassmorphisme
- Animations fluides et interactives
- Dégradés violets et rouges
- Font Poppins pour une typographie moderne

## 🔧 Configuration

### Port du serveur
Le serveur démarre par défaut sur le port 8080. Pour changer le port, modifiez le fichier `cmd/main.go`.

### Cache
Les headers anti-cache sont configurés pour éviter les problèmes de mise en cache. Si vous voyez une ancienne version :
- Appuyez sur **Ctrl+Shift+R** pour un hard refresh
- Ou ouvrez en navigation privée

## 📝 API

L'application utilise l'API PokeBuild pour récupérer les données des Pokémon.

Endpoints utilisés :
- Liste des Pokémon : `/pokemon`
- Détails d'un Pokémon : `/pokemon/{id}`
- Recherche : `/pokemon?name={query}`

## 🐛 Dépannage

### Le serveur ne démarre pas
- Vérifiez que Go est installé : `go version`
- Assurez-vous d'être dans le bon dossier : `cd src`

### Les styles ne s'affichent pas
- Videz le cache du navigateur (Ctrl+Shift+Delete)
- Rechargez la page avec Ctrl+Shift+R
- Essayez en navigation privée

### Les favoris ne se sauvegardent pas
- Vérifiez que le fichier `favoris.json` existe
- Assurez-vous d'avoir les permissions d'écriture

## 🤝 Contribution

Les contributions sont les bienvenues ! N'hésitez pas à :
- Signaler des bugs
- Proposer des nouvelles fonctionnalités
- Améliorer le design
- Optimiser le code

## 📄 Licence

Ce projet est open source et disponible sous licence MIT.

## 👨‍💻 Auteur

Créé avec ❤️ pour les fans de Pokémon

---

**Note** : Ce projet utilise l'API PokeBuild qui peut avoir des limites de taux. Utilisez-le de manière responsable.
# API-POKE
