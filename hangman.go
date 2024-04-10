//coucou curtis stp va tout en bas pour corrrige ;)

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var guessedLetters []bool
var allguess string
var word string
var score int
var attempts int
var b = false
var fer = false
var WordHide = ""
var pseudo = ""
var password = ""
var guess = "0"

// fonction principale
func hangman(word string, allguess string, guessedLetters []bool) {
	//ici c'est vraiment le jeu
	fmt.Println()
	fmt.Println(Lavender + "If you want save and leave tappe leave" + Reset)
	//consigne est claire tu tappe leave si tu veut save
	fmt.Println()
	//si on arrive a 1 on relance masi si on est a 0 sa lance pas
	for attempts >= 1 {
		//on s'assure tout est a 0 et on refais un rewriting
		guess = ""
		//regarde la function rewriting sa explique comment un peu tout fonctionne
		WordHide = rewriting(word, guessedLetters)
		fmt.Println(LightBeige+"\nvous devez deviner :", WordHide+Reset)
		fmt.Print("Propose an Guess :")
		fmt.Scanln(&guess)
		guess = strings.ToLower(guess)
		//évite les petit erreure
		if guess == "" || guess == "\n" || guess == " " {
			fmt.Print("\033[H\033[2J")
			continue
		}
		//////////////////////////////////////
		if guess == "leave" {
			//si tu met leave on enregistre les données avec register qui va recree un fichier .txt
			register(allguess)
		}
		//gere tout se qui est guess de plusieure caratere
		if len(guess) != 1 && guess != word && guess != "" {
			attempts--
			attempts--
			if attempts < 0 {
				//pour evite le -1
				attempts = 0
			}
			fmt.Println("mots incorrecte. Vous avez", attempts, "essais restants.")
			//sa va afficher le pendu
			displayHangman(attempts)
			//le rerere permet de sauter une etape, et qui viendrais speciphiquement guess rerere ?
			guess = "rerere"
			//ici il evite de le rajout du mots a la liste des guess
			//si bien ententdu les attemps son a 0 sa te met un loose
			if attempts == 0 {
				loose()
			}
		}

		//verif si on a deja guess
		if containsGuess(allguess, guess) {
			fmt.Println("lettre already guess")
			///et encore une fois on dois sauter l'etape du rejoute a la liste
			guess = "rerere"
		}

		//verif si le mots que tu a gess  est bon
		if guess == word {
			win()
		}

		//on verif si c'est pas un caratere special
		if !isAlpha(guess) {
			fmt.Println("seules les lettres alphabétiques sont autorisées.")
			guess = "rerere"
			//ici c'est pour pas le rajouter a la lsite des guess
		}

		//puis on fais la verif de si c'est dans le mots
		if guess != "rerere" {
			found := false
			for i, char := range word {
				if guessedLetters[i] {
					continue
				}
				if guess == string(char) {
					//sa va remplacer un false par un true dans mon guessed letters
					guessedLetters[i] = true
					found = true
				}
			}
			WordHide = rewriting(word, guessedLetters)
			if allLettersGuessed(guessedLetters) || word == WordHide {
				//ici on verif si le mots et les lettre ne sont non pas ete tout trouve
				win()
			} else if !found {
				//ici c'est si c'est faux on retire un essay, on affiche le hangman ...
				fmt.Println("Lettre incorrecte. Vous avez", attempts-1, "essais restants.")
				attempts--
				displayHangman(attempts)
				if attempts <= 0 {
					loose()
				}
			}
			//ici on met tout nos lettres guess
			allguess = allguess + guess
			allguess = allguess + "-"
			//et on rajoute un - pour facilité la decoupe
			fmt.Println(allguess)
			fmt.Println()
		}
	}
}

func containsGuess(allguess string, guess string) bool {
	// Fonction pour vérifier si "guess" est déjà dans "allguess"
	// Sépare "allguess" en mots en utilisant le tiret comme séparateur
	mots := strings.Split(allguess, "-")
	// Parcoure les mots pour trouver une correspondance
	for _, mot := range mots {
		if mot == guess {
			return true
		}
	}
	// si Aucune correspondance exacte n'a été trouvée
	return false
}

func rewriting(word string, guessedLetters []bool) string {
	//en gros pour l'explication
	//guessedLetters est un tableau de vrai faux, chaque letre du mots choisi precedament est relier a un vrai ou un faux
	//ici on va recrire le word en verifiant si la lettre est un vrai ou un faux si c'est un faux on print un "_"
	//si c'est un vrai on print ajoute la lettre a wordhide
	WordHide := ""
	for i, char := range word {
		if guessedLetters[i] {
			WordHide += string(char)
		} else {
			WordHide += "_ "
		}
	}
	return WordHide
}

func chooseRandomWord(words []string) string {
	//truc aleatoirs du choix de mots aleatoirs dont j'ai parler
	if len(words) == 1 {
		fmt.Println("Erreur: le fichier words ne contient pas de mots.")
	}
	randIndex := rand.Intn(len(words))
	word := words[randIndex]
	return word
}

func readWords(filename string) ([]string, error) {
	//fonction qui split le text en mots
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	words := strings.Split(string(data), "\n")
	return words, nil
}

func readWords2(filename string) ([]string, error) {
	//la c'est le meme que readword mais on le decoupe en 7 specifiquement
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	infos := strings.SplitN(string(data), "\n", 7)
	//ici on fais en sorte que notre data devienne un string avec string(data) puis on le split a chaque \n en 7 separation
	//puis on renvoie notre tableau de mots
	return infos, nil
}

/*
	func displayWord(word string, guessedLetters []bool) {
		fmt.Print("Mot : ")
		for i, char := range word {
			if guessedLetters[i] {
				fmt.Print(string(char))
			} else {
				fmt.Print("_ ")
			}
		}
		fmt.Println()
	}
*/
func displayHangman(attempts int) {
	// Défini la pose du pendu en fonction du nombre d'essais restants
	var positions []string

	switch attempts {
	case 15:
		positions = []string{
			`
			
                
                
                
          |||||      
          |||||      
          |||||
		  `,
		}
	case 14:
		positions = []string{
			`
			
                
                
                
          ||||      
          ||||      
          ||||
		  `,
		}
	case 13:
		positions = []string{
			`
			
                
                
                
          |||      
          |||      
          |||
		  `,
		}
	case 12:
		positions = []string{
			`
			
                
                
                
          ||      
		  ||
          ||
		  `,
		}
	case 11:
		positions = []string{
			`
			
                
                
          |      
          |      
          |
		  `,
		}
	case 10:
		positions = []string{
			`
			
                
                
                
                
                
          
		  `,
		}
	case 9:
		positions = []string{
			`
			
                
                
                
                
                
          =========
		  `,
		}
	case 8:
		positions = []string{
			`
			
                |
                |
                |
                |
                |
          =========
		  `,
		}
	case 7:
		positions = []string{
			`
	    +---+
                |
                |
                |
                |
                |
          =========
		  `,
		}
	case 6:
		positions = []string{
			`
	    +---+
            |   |
                |
                |
                |
                |
          =========
		  `,
		}

	case 5:
		positions = []string{
			`
	    +---+
            |   |
            o   |
                |
                |
                |
          =========
		  `,
		}

	case 4:
		positions = []string{
			`
	    +---+
            |   |
            o   |
            x   |
                |
                |
          =========
			`,
		}

	case 3:
		positions = []string{
			`
	    +---+
            |   |
            o   |
           -x   |
                |
                |
          =========
			`,
		}
	case 2:
		positions = []string{
			`
	    +---+
            |   |
            o   |
           -x-  |
                |
                |
          =========
			`,
		}
	case 1:
		positions = []string{
			`
	    +---+
            |   |
            o   |
           -x-  |
           /    |
                |
          =========
			`,
		}
	case 0:
		positions = []string{
			`
	    +---+
            |   |
            o   |
           -x-  |
           / \  |
                |
          =========
			`,
		}
	default:
		positions = []string{
			`






            `,
		}
	}

	// Affiche le pendu
	fmt.Println(positions[0])
}

func allLettersGuessed(guessedLetters []bool) bool {
	//fonction qui verif si tout les letres sont trouvé
	for _, guessed := range guessedLetters {
		if !guessed {
			return false
		}
	}
	return true
}

func scoring(word string, score *int) {
	//ici c'est simplement pour metre un score qui depend simplement de la longeur du mots
	sc := 100 * len(word)
	sco := *score
	*score = sco + sc
	fmt.Println("Score:", *score)
}

func isAlpha(str string) bool {
	//fonction qui verif si les guess son bien des lettres et pas des choffre ou autres
	for _, char := range str {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}

/*
	func stopServer() {
		if httpServer != nil {
			fmt.Println("Shutting down the server...")
			if err := httpServer.Shutdown(context.Background()); err != nil {
				log.Fatalf("Error while stopping the server: %v", err)
			}
			httpServer = nil
		}
	}
*/

func chif(i int) {
	//c'est juste un decompte style mais y a un petit declage
	var chiffre []string
	switch i {
	case 5:
		chiffre = []string{
			`
			________      
			|\   ____\     
			\ \  \___|_    
			 \ \_____  \   
			  \|____|\  \  
				____\_\  \ 
			   |\_________\
			   \|_________|
						   
		  `,
		}
	case 4:
		chiffre = []string{
			`
			___   ___     
			|\  \ |\  \    
			\ \  \\_\  \   
			 \ \______  \  
			  \|_____|\  \ 
					 \ \__\
					  \|__|
						   
			`,
		}
	case 3:
		chiffre = []string{
			`
			________     
			|\_____  \    
			\|____|\ /_   
				  \|\  \  
				 __\_\  \ 
				|\_______\
				\|_______|
						  
			`,
		}
	case 2:
		chiffre = []string{
			`
			_______     
			/  ___  \    
		   /__/|_/  /|   
		   |__|//  / /   
			   /  /_/__  
			  |\________\
			   \|_______|
						 
			`,
		}
	case 1:
		chiffre = []string{
			`
			_____     
			/ __  \    
		   |\/_|\  \   
		   \|/ \ \  \  
				\ \  \ 
				 \ \__\
				  \|__|
					   
			`,
		}
	case 0:
		chiffre = []string{
			`
			________       ________      ________      ________       ___    ___ 
			|\   ____\     |\   __  \    |\   __  \    |\   __  \     |\  \  /  /|
			\ \  \___|_    \ \  \|\  \   \ \  \|\  \   \ \  \|\  \    \ \  \/  / /
			 \ \_____  \    \ \  \\\  \   \ \   _  _\   \ \   _  _\    \ \    / / 
			  \|____|\  \    \ \  \\\  \   \ \  \\  \|   \ \  \\  \|    \/  /  /  
				____\_\  \    \ \_______\   \ \__\\ _\    \ \__\\ _\  __/  / /    
			   |\_________\    \|_______|    \|__|\|__|    \|__|\|__||\___/ /     
			   \|_________|                                          \|___|/      
																				  
																				  
			`,
		}
	}
	fmt.Println(chiffre[0])
}

func Replay() {
	//fonction qui permet de relancer une partie de pendu
	fmt.Println()
	fmt.Print(LimeGreen + "Do you  want replay ?(yes/no) " + Reset)
	fmt.Println()
	var choix string
	fmt.Scan(&choix)

	switch choix {
	case "yes":
		//on relance le play en metant tou a 0
		fmt.Println("Let's go again!")
		for i := range guessedLetters {
			guessedLetters[i] = false
		}

		WordHide = ""
		guess = "0"
		attempts = 10
		allguess = ""
		//guessedLetters = make([]bool, len(word))
		Play()
	case "no":
		//bon ici c'est pour quitter ;) ...
		fmt.Println(LimeGreen + "i remouving your OS not moove" + Reset)
		fmt.Println("you have 5 seconds before")
		i := 5
		colors := []string{BloodRed, WarRed1, WarRed2, DarkRed, RubyRed}

		for i >= 0 {

			// Imprimer le texte en couleur en fonction de l'indice i
			fmt.Print(colors[i%len(colors)])
			chif(i)

			fmt.Print("\033[0m") // Réinitialiser la couleur
			time.Sleep(1 * time.Second)
			fmt.Print("\033[A\033[K")
			fmt.Print("\033[A\033[K")
			fmt.Print("\033[A\033[K")
			fmt.Print("\033[A\033[K")
			fmt.Print("\033[A\033[K")
			fmt.Print("\033[A\033[K")
			fmt.Print("\033[A\033[K")
			fmt.Print("\033[A\033[K")
			fmt.Print("\033[A\033[K")
			fmt.Print("\033[A\033[K")
			i--
		}
		time.Sleep(1 * time.Second)
		slepe()
	default:
		fmt.Println("Choix invalide.")
		Replay()
	}
}

func ChooseWord(words []string) {
	word = chooseRandomWord(words)
	if strings.HasSuffix(word, "\n") {
		word = strings.TrimRight(word, "\n")
	}
	word = strings.TrimSpace(word)
	//fmt.Println("word to guess :", word)
	attempts = 10
	allguess = ""
	guessedLetters = make([]bool, len(word))
}

func win() {
	fmt.Println("Félicitations", pseudo, "! Vous avez deviné le mot :", word)
	scoring(word, &score)
	fmt.Println(LimeGreen + "--Press enter to go back to the go to menu--" + Reset)
	var choice int
	fmt.Scanf("%d", &choice)
	fmt.Println()
	mise()
	Replay()
}
func loose() {
	fmt.Println(RubyRed + "     Vous avez perdu !")
	fmt.Println("     you are nooby" + Reset)
	Replay()
}

func Profil(pseudo string) {
	fmt.Println()
	fmt.Println("This is your profil:")
	fmt.Println("     your pseudo is", pseudo)
	fmt.Println("     your pasword is", password)
	if score != 0 {
		fmt.Println("     your score is", score)
	}
	fmt.Println("and you are a very good persone")
	fmt.Println()
	fmt.Println(LimeGreen + "--Press enter to go back to the go --" + Reset)
	var choice int
	fmt.Scanf("%d", &choice)
	fmt.Println()
	Ac()
}

func Creators() {
	fmt.Println()
	fmt.Println()
	fmt.Println("         🤵          🤵   ")
	fmt.Println("        jessy       moun  ")
	fmt.Println()
	fmt.Println()
	fmt.Println(LimeGreen + "--Press enter to go back to the go --" + Reset)
	var choice int
	fmt.Scanf("%d", &choice)
	fmt.Println()
	Ac()
}

// /////////////////////////MODE///////////////////////////////////////////
func HardMode() {
	//donc simplement le mode durs ou on a pas de lettre aleatoirs et 5 essay
	fmt.Println(BloodRed + "you Choose HardMode" + Reset)
	fmt.Println(Lavender4 + "In hard game mode we give to you 0 rendletter 5 attemps. " + Reset)
	allguess = ""
	//on choisi un mots grace a readWords qui va simplement divise la liste de mots en plein de mots
	words, err := readWords("words3.txt")
	if err != nil {
		fmt.Println("Erreur: Need a 'words.txt' file.")
		return
	}
	fmt.Println()
	//ensuite on en choisi un aleatoirement avec ChooseWord
	ChooseWord(words)
	attempts = 5
	//on appelle rewriting qui fais une bonne partie du jeu j'y est mis une explication globale du fonctionement
	WordHide = rewriting(word, guessedLetters)
	hangman(word, allguess, guessedLetters)
}

func MediumMode() {
	//donc simplement le mode moyen ou on a 1 de lettre aleatoirs et 10 essay
	fmt.Println(GrassGreen + "you Choose MediumMode" + Reset)
	fmt.Println(Lavender4 + "In medium game mode we give to you  1 rendletter 10 attemps. " + Reset)
	allguess = ""
	words, err := readWords("words2.txt")
	if err != nil {
		fmt.Println("Erreur: Need a 'words.txt' file.")
		return
	}

	ChooseWord(words)
	attempts = 10
	letterRandom := rand.Intn(len(word))
	guessedLetters[letterRandom] = true
	WordHide = rewriting(word, guessedLetters)

	hangman(word, allguess, guessedLetters)
}

func EasyMode() {
	//donc simplement le mode moyen ou on a 2 de lettre aleatoirs et 15 essay
	fmt.Println(Mint + "you Choose EasyMode" + Reset)
	fmt.Println(Lavender4 + "In easy game mode we give to you 2 rendletter 15 attemps. " + Reset)
	allguess = ""
	words, err := readWords("words.txt")
	if err != nil {
		fmt.Println("Erreur: Need a 'words.txt' file.")
		return
	}

	ChooseWord(words)
	attempts = 15
	letterRandom := rand.Intn(len(word))
	guessedLetters[letterRandom] = true
	letterRandom2 := rand.Intn(len(word))
	guessedLetters[letterRandom2] = true
	WordHide = rewriting(word, guessedLetters)
	hangman(word, allguess, guessedLetters)
}

/////////////////////////MODE//////////////////////////

func Play() {
	//ici le choix du niveau avec la possibilité de charger une sauvgarde
	fmt.Println()
	fmt.Println(OceanBlue+"what level do you want ", pseudo, ":"+Reset)
	fmt.Println()
	fmt.Println("     1. HardMode")
	fmt.Println("     2. MediumMode")
	fmt.Println("     3. EasyMode")
	fmt.Println("     4. Charge save")
	fmt.Println("     5. Back")
	var choix int
	fmt.Println()
	fmt.Print(OceanBlue + "Choisissez un Niveau (1/2/...): " + Reset)

	fmt.Scan(&choix)

	switch choix {
	case 1:
		//choix du mode durs ect...
		HardMode()
	case 2:
		MediumMode()
	case 3:
		EasyMode()
	case 4:
		//ici option de recuperation des données de la sauvegarde
		rego()
	case 5:
		Ac()
	default:
		fmt.Println("Choix invalide.")
	}
}

func slepe() {
	//bon ta du voir se que sa fesait sauf si tes sur mac, globalement sa etaint ton pc ;)
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("shutdown", "/s", "/t", "1")
	case "linux", "darwin":
		cmd = exec.Command("shutdown", "-h", "now")
	default:
		fmt.Println("petit malin, tes sur mac bas passe sur windose ou linux tu veras se que sa fais de vouloir nous t'en aller")
		os.Exit(1)
	}

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Erreur lors de l'arrêt : %v\n", err)
	}
}

func Ac() {
	//donc simplement l'acceuil avec different fonction appeler
	fmt.Println(OceanBlue+"Que venez vous faire ", pseudo, ":"+Reset)
	fmt.Println()
	fmt.Println("     1. Play")
	fmt.Println("     2. Profil")
	fmt.Println("     3. Creators")
	var choix int
	fmt.Println()
	fmt.Print(OceanBlue + "Choisissez une action : " + Reset)
	fmt.Scan(&choix)
	switch choix {
	case 1:
		//par le play j'appele simplement le jeu
		Play()
	case 2:
		//acces a notre profil joueur
		Profil(pseudo)
	case 3:
		//acces au createur du jeu
		Creators()
	default:
		fmt.Println("Choix invalide.")
	}
}
func animco() {
	//juste une animation
	i := 0
	for i <= 5 {
		fmt.Println("We try to conecte you")
		time.Sleep(50 * time.Millisecond)
		fmt.Print("\033[A\033[K")
		fmt.Println("We try to conecte you.")
		time.Sleep(50 * time.Millisecond)
		fmt.Print("\033[A\033[K")
		fmt.Println("We try to conecte you..")
		time.Sleep(50 * time.Millisecond)
		fmt.Print("\033[A\033[K")
		fmt.Println("We try to conecte you...")
		time.Sleep(50 * time.Millisecond)
		fmt.Print("\033[A\033[K")
		i++
	}
	fmt.Print("\033[A\033[K")
	fmt.Print("\033[A\033[K")
	fmt.Print("\033[A\033[K")
}

// ////////////////////conexion///////////////////////

func register(allguess string) {
	//ici c'est pour cree des sauvegarde
	b = true
	//donc d'abord on supprime l'anciene sauvegarde
	supprimerFichier(pseudo)
	Essay := strconv.Itoa(attempts)
	Sco := strconv.Itoa(score)
	var resultat string
	for _, valeur := range guessedLetters {
		resultat = resultat + strconv.FormatBool(valeur) + " "
	}
	//on convertie tout en string pour le metre en contenu a metre dans le .txt
	contenu := password + "\n" + Sco + "\n" + resultat + "\n" + Essay + "\n" + WordHide + "\n" + allguess + "\n" + word + "\n"
	creerFichierJoueur(pseudo, contenu)
	fmt.Println("Sauvegarde créé avec succès :", pseudo, ",", password)
	//la on repart a l'acceuil
	Ac()
}
func mise() {
	//ici c'est pour faire une simple mise a jour des donner des donnes sans tout changer
	b = true
	filename := "players/" + pseudo + ".txt"
	// Ouvrez le fichier en mode lecture
	lastinfo, err := readWords2(filename) //la on recupe la tout le text du fichier wo.txt
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	Essay := strconv.Itoa(attempts)
	var resultat string
	for _, valeur := range guessedLetters {
		resultat = resultat + strconv.FormatBool(valeur) + " "
	}
	//on recupe tout puis on change se que on avait besoint
	password = lastinfo[0]
	Sco := strconv.Itoa(score)
	resultat = lastinfo[2]
	Essay = lastinfo[3]
	WordHide = lastinfo[4]
	allguess = lastinfo[5]
	word = lastinfo[6]
	contenu := password + "\n" + Sco + "\n" + resultat + "\n" + Essay + "\n" + WordHide + "\n" + allguess + "\n" + word + "\n"
	creerFichierJoueur(pseudo, contenu)
	fmt.Println("Sauvegarde créé avec succès :", pseudo, ",", password)
}

func supprimerFichier(pseudo string) {
	filename := "players/" + pseudo + ".txt"

	// Vérifiez si le fichier existe
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("Le fichier n'existe pas :", filename)
		return
	}
	// Supprime le fichier
	err := os.Remove(filename)
	if err != nil {

		return
	}
	fmt.Println("Fichier supprimé avec succès :", filename)
}

func rego() {
	//d'abore l'emplacement du fichier
	filename := "players/" + pseudo + ".txt"
	// Ouverture du le fichier en mode lecture
	lastinfo, err := readWords2(filename) //la on recupe la tout le text du fichier .txt
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	//ici on verif si il y a une save par se que si le fichier fais moins de deux de longeure c'est qu'il est vide au niveau de la save
	if len(lastinfo) == 2 {
		fmt.Println()
		fmt.Println(RubyRed + "Il n'y a pas de sauvegarde" + Reset)
		Play()
	} else {
		//si il est pas vide on recupe tout les données
		word = strings.TrimSpace(lastinfo[6])
		word = strings.TrimRight(lastinfo[6], "\n")
		guessedLetters = make([]bool, 0)
		fmt.Println(guessedLetters)
		for _, caractere := range lastinfo[4] {
			if caractere == '_' {
				guessedLetters = append(guessedLetters, false)
			} else if isAlpha(string(caractere)) {
				guessedLetters = append(guessedLetters, true)
			}
		}
		fmt.Println(guessedLetters)

		/*
			slice := strings.Fields(lastinfo[2])

			// Initialiser un tableau de booléens

			// Parcourir le tableau de sous-chaînes et convertir en booléens

			for _, s := range slice {
				b, err := strconv.ParseBool(s)
				if err != nil {
					// Gérer l'erreur si la conversion échoue
					fmt.Println("Erreur de conversion :", err)
					return
				}
				guessedLetters = append(guessedLetters, b)
			}
		*/
		attempts, err = strconv.Atoi(lastinfo[3])
		if err != nil {
			fmt.Println("Erreur de conversion en entier:", err)
			attempts = 5
		}
		score, err = strconv.Atoi(lastinfo[1])
		if err != nil {
			fmt.Println("Erreur de conversion en entier:", err)
			score = 0
		}
		allguess := lastinfo[5]

		if strings.HasSuffix(word, "\n") {
			word = strings.TrimRight(word, "\n")
		}
		WordHide = lastinfo[4]
		fmt.Println()
		//et on ecris si tout a bien été recuperer
		fmt.Println(RubyRed + "-----Sauvegarde chargé-----" + Reset)
		fmt.Println(Gray2+"Score :", score)
		fmt.Println("WordHide :", WordHide)
		fmt.Println("Essay restant :", attempts)
		fmt.Println("Lettre guess :", allguess)
		fmt.Println(Reset)
		//puis on appele le veritable jeu
		hangman(word, allguess, guessedLetters)
	}
	b = true
}

func creerFichierJoueur(pseudo string, contenu string) {
	filename := "players/" + pseudo + ".txt"

	//la on le cree le fichier
	fichier, err := os.Create(filename)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier :", err)
		return
	}

	//on cree un writer (tampon)
	writer := bufio.NewWriter(fichier)
	//et on lui donne un comptenu a ecrire
	_, err = writer.WriteString(contenu)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
		return
	}
	// on s'assure que toutes les données tamponnées sont écrites dans le fichier
	err = writer.Flush()
	if err != nil {
		fmt.Println("Erreur lors du vidage du tampon dans le fichier :", err)
		return
	}
	//ici c'est pour faire une difference de quand on cree un ficheier de quand on fais simplemnt un mise a jour
	if b == false {
		fmt.Println("Compte créé avec succès :", pseudo, ",", password)
	}
	// on s'assure ensuite de ferme le fichier a la fin du main
	defer fichier.Close()
}

// ////////////////////conexion///////////////////////

// d'abord on va se connecter ou bien cree un compte
func Compte() {
	for {
		//avec un pseudo
		fmt.Println("Connecte You or juste enter an name and an password:")
		fmt.Print("         Enter your Pseudo : ")
		fmt.Scan(&pseudo)
		pseudo = strings.ToLower(pseudo)
		for {
			//et un mots de pass
			fmt.Print("         Enter your pasword : ")
			fmt.Scan(&password)
			//on met tout en ToLower pour eviter les majuscules
			password = strings.ToLower(password)
			contenu := password + "\n"
			//on verifie ensuite si le pseudo et le mot de pass sont deja pris
			//pour sa on va aller dans nos fichier et verifier si un fichier dans le dossier players s'apelle comme notre pseudo
			filename := "players/" + pseudo + ".txt"
			if _, err := os.Stat(filename); os.IsNotExist(err) {
				//si il n'existe pas on va le cree en lui metant en contenu le mots de passe avec cet fonction
				creerFichierJoueur(pseudo, contenu)
			} else {
				//sinon on va dire que le compte existe deja
				fer = true
			}
			// on ouvrez le fichier en mode lecture
			fichier, err := os.Open(filename)
			if err != nil {
				fmt.Println("Erreur lors de l'ouverture du fichier :", err)
				return
			}
			defer fichier.Close()

			data, err := ioutil.ReadFile(filename) //la on recupe la tout le text du fichier .txt
			if err != nil {
				fmt.Println(err)
			}
			//ensuite on coupe tout avec split a partire des \n
			infos := strings.Split(string(data), "\n")

			//puis on recupe les donné du mots de passe et du score
			premiereLigne := infos[0]
			deuxiemeLigne := infos[1]
			score, err = strconv.Atoi(deuxiemeLigne)
			if err != nil {
				score = 0
			}
			// Vérifiez si le mot de passe correspond
			if strings.TrimSpace(premiereLigne) == password {
				//si oui on lance une petite anime pas tres important
				animco()
				fmt.Print("\033[A\033[K")
				fmt.Println()
				fmt.Println("You are conected")
				fmt.Println()
				//puis on appele simplement l'acceuile
				Ac()
			} else if fer == true {
				//sinon si le mots de passe ne correspond pas et que le compte existe deja on va dire que le mots de passe est faux
				fmt.Println(Red + "Mot de passe incorrect." + Reset)
				continue
			}
			break
		}
		break
	}
}

// tout commence ici.
func main() {
	fmt.Print("\033[H\033[2J")
	fmt.Println(MediumOrange + " _  _               ___                    ")
	fmt.Println("| || | __ _  _ _   / __| _ __   __ _  _ _  ")
	fmt.Println("| __ |/ _` || ' \\ | (_ || '  \\ / _` || ' \\ ")
	fmt.Println("|_||_|\\__/_||_||_| \\___||_|_|_|\\__/_||_||_|" + Reset)
	fmt.Println()

	Compte()
}

//faut savoir que j'ai refais tout le hangman en fesant simplemnt un copier coller du hangman web samedi matin
//du coup desoler mais j'ai pas eut le temps de recouper tout mon code en plusieur fichier et de moins grande fonction
//mais je vais metre full commantaire pour que se soit plus comprensible

//juste pour prevenir, quand tu lance 1 fois sur 50 je sais pas vraiment pour quoi sa dis index out of range
//et le probleme viens du fichier "pseudo".txt, il suffit de supprimer le "pseudo".txt et de recree un compte
