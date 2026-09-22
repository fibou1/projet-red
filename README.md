# Projet RED

> Un mini-jeu RPG en ligne de commande, developpe en Go.
>
> Creation de personnage, inventaire, economie, marchand et premiers elements de combat.

[![Go](https://img.shields.io/badge/Go-1.22.2-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![CI](https://github.com/fibou1/projet-red/actions/workflows/go.yml/badge.svg)](https://github.com/fibou1/projet-red/actions/workflows/go.yml)
[![License](https://img.shields.io/badge/license-educational-lightgrey)](#licence)

## Vue d'ensemble

Projet RED est un jeu CLI realise dans le cadre d'un projet pedagogique. Le joueur
cree un personnage, choisit sa classe, gere ses points de vie et son inventaire,
puis interagit avec le menu principal et le marchand.

Le projet sert aussi de terrain d'experimentation pour une chaine de livraison
simple : Git, GitHub, GitHub Actions et verification automatique du lancement.

## Quick start

### Prerequis

- Go `1.22.2` ou plus recent
- Git
- Un terminal compatible avec Go

### Installation

```bash
git clone https://github.com/fibou1/projet-red.git
cd projet-red
```

### Lancement interactif

```bash
go run .
```
avec docker 

```bash
docker build -t projet-red .
docker run -it projet-red
```

Le jeu demande successivement une classe, un nom, puis une action dans le menu.

### Lancement reproductible

Cette commande fournit automatiquement les entrees minimales au programme :

```bash
printf '1\nNova\n6\n' | go run .
```

Elle choisit la classe `1`, cree le personnage `Nova`, puis selectionne `6` pour
quitter proprement. C'est le scenario utilise par la CI.

## Gameplay actuel

### Creation du personnage

| Classe | PV maximum | PV de depart |
| --- | ---: | ---: |
| Alien | 100 | 50 |
| Humain | 80 | 40 |
| Reptile | 120 | 60 |

Le personnage commence au niveau 1, avec le sort de base `coup de poing` et
100 coins.

### Menu principal

| Choix | Action |
| ---: | --- |
| 1 | Afficher les informations du personnage |
| 2 | Acceder a l'inventaire |
| 3 | Ouvrir le marchand |
| 4 | Ouvrir le forgeron |
| 5 | Lancer l'entrainement |
| 6 | Quitter |

### Objets et economie

Le projet contient actuellement des objets avec un nom, une description, une
quantite et un prix :

- Potion de vie
- Arme nucleaire
- Veste d'Agent Special
- Casque
- Super bottes
- Potion de poison
- Virus alien

Le marchand permet d'acheter des objets, de deduire leur prix et de respecter la
capacite maximale de l'inventaire. La potion de vie restaure jusqu'au maximum de
points de vie du personnage.

## Architecture du depot

```text
.
|-- .github/workflows/go.yml  # Pipeline GitHub Actions
|-- main.go                   # Point d'entree
|-- Character.go              # Personnage et creation
|-- mainmenu.go               # Navigation principale
|-- inventaire.go             # Objets et gestion de l'inventaire
|-- marchand.go               # Boutique et achats
|-- monstre.go                # Structure Monster
|-- attack.go                 # Effet de poison temporise
|-- go.mod                    # Module et version Go
`-- README.md                 # Documentation du projet
```

## Pipeline DevOps

Le workflow `.github/workflows/go.yml` applique une verification minimale a
chaque modification importante :

```text
push / pull request
	|
	v
checkout du depot
	|
	v
installation de Go depuis go.mod
	|
	v
printf des entrees -> go run .
	|
	v
job vert si le programme se termine sans erreur
```

### Declencheurs

- Push sur `main`
- Pull Request vers `main`

### Consulter le resultat

1. Ouvrir le depot GitHub.
2. Aller dans l'onglet **Actions**.
3. Selectionner **Go Run Check**.
4. Ouvrir le dernier workflow pour lire les logs.

Le resultat est consultable depuis un navigateur ou l'application GitHub sur
telephone.

## Commandes de developpement

```bash
# Lancer le jeu
go run .

# Verifier la compilation et les tests Go presents
go test ./...

# Detecter certaines erreurs de code Go
go vet ./...

# Formater les fichiers Go
gofmt -w *.go

# Verifier les espaces et erreurs de patch Git
git diff --check
```

Avant un push, le controle local recommande est :

```bash
gofmt -w *.go
go test ./...
go vet ./...
git diff --check
```

## Workflow Git recommande

```bash
git status
git add .
git commit -m "Decrit clairement le changement"
git push
```

Ne jamais ajouter un token GitHub dans le code, le README ou l'URL du remote.
Utiliser SSH ou le gestionnaire d'identifiants du systeme.

## Etat du projet

### Disponible

- Creation interactive du personnage
- Classes et calcul des PV de depart
- Affichage des informations
- Inventaire, ajout et suppression d'objets
- Limite de capacite de l'inventaire
- Potion de vie
- Monnaie et marchand
- Potion de poison temporisee
- Structure de monstre et gestion de la mort
- Verification de lancement par GitHub Actions

### Prochaines evolutions

- Relier completement l'option Inventaire aux actions utilisables
- Finaliser le menu Forgeron et la fabrication d'equipements
- Ajouter l'equipement tete, torse et pieds avec bonus de PV
- Implementer l'augmentation de capacite de l'inventaire
- Finaliser le combat tour par tour contre le gobelin
- Ajouter les sorts offensifs, l'experience, l'initiative et le mana
- Ajouter des tests Go lorsque les fonctions metier seront stabilisees

## Licence

Projet pedagogique realise dans le cadre de l'immersion. Aucune licence de
distribution n'est definie pour le moment.
# Projet RED

Mini-jeu en ligne de commande (CLI) developpe en Go dans le cadre du Projet RED.
Le joueur cree un personnage, consulte son inventaire, utilise des objets et peut
acceder a differents espaces du jeu comme le marchand.

## Installation

### Prerequis

- Go 1.22 ou une version plus recente
- Git

### Recuperer le projet

```bash
git clone https://github.com/fibou1/projet-red.git
cd projet-red
```

## Lancer le jeu

Depuis la racine du projet :

```bash
go run .
```

Le programme demande d'abord :

1. une classe : `1` alien, `2` humain ou `3` reptile ;
2. le nom du personnage ;
3. une option dans le menu principal.

## Menu principal

| Choix | Fonction |
| --- | --- |
| 1 | Afficher les informations du personnage |
| 2 | Acceder a l'inventaire |
| 3 | Ouvrir le marchand |
| 4 | Ouvrir le forgeron |
| 5 | Lancer l'entrainement |
| 6 | Quitter le jeu |

## Fonctionnalites actuelles

- Creation interactive d'un personnage.
- Trois classes avec des points de vie maximum differents.
- Niveau initial a 1 et points de vie actuels a 50 % des points de vie maximum.
- Inventaire avec limite de capacite.
- Affichage et ajout d'objets.
- Suppression d'objets de l'inventaire.
- Utilisation d'une potion de vie avec restauration limitee aux PV maximum.
- Systeme de monnaie avec 100 coins de depart.
- Marchand avec achat d'objets et deduction du prix.
- Debut de systeme de poison et de monstre.
- Fonction de resurrection a 50 % des PV maximum lorsque le personnage est vaincu.
- Verification automatique du lancement avec GitHub Actions.

## Objets disponibles

Le marchand contient actuellement plusieurs objets, notamment :

- Potion de vie
- Arme nucleaire
- Veste d'Agent Special
- Casque
- Super bottes
- Potion de poison
- Virus alien

Chaque objet possede un nom, une description, une quantite et un prix.

## Verification locale

Pour verifier que le projet compile et se lance :

```bash
printf '1\nNova\n6\n' | go run .
```

Cette commande simule le choix de la classe `1`, le nom `Nova`, puis le choix `6`
pour quitter proprement le programme.

## GitHub Actions

Le fichier `.github/workflows/go.yml` lance automatiquement `go run .` :

- apres un push sur `main` ;
- lors d'une Pull Request vers `main`.

Le workflow fournit les entrees necessaires au programme et verifie qu'il se
lance sans erreur. Le resultat est disponible dans l'onglet **Actions** du depot
GitHub.

## Organisation du code

- `main.go` : point d'entree du programme.
- `Character.go` : personnage, creation, affichage et gestion des PV.
- `inventaire.go` : objets, inventaire, ajout, suppression et potion de vie.
- `mainmenu.go` : menu principal et navigation.
- `marchand.go` : marchand et achats.
- `monstre.go` : structure du monstre.
- `attack.go` : effet de poison avec temporisation.
- `.github/workflows/go.yml` : verification GitHub Actions.

## Fonctionnalites a poursuivre

Les prochaines etapes prevues sont notamment :

- finaliser le menu d'inventaire et l'utilisation des objets depuis ce menu ;
- implementer la fabrication d'equipements et le forgeron ;
- ajouter les equipements et leurs bonus de PV ;
- finaliser l'augmentation de capacite de l'inventaire ;
- implementer le combat tour par tour complet contre le gobelin ;
- ajouter les sorts, l'experience, l'initiative et le mana.