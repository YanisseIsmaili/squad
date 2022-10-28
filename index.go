package main

func readFile(s string, toFind string) int {
    strtofind := NbretoFind(toFind) // prend comme valeur la str a rechercher
    secondread := 0
    for _, j := range toFind {
        for i1, i2 := range s {
            if i2 == j {
                if strtofind == 1 {
                    return i1
                } else if strtofind > 1 { // si str a trouver > 1, crée une loop  pour vérifier
                    for k := 0; k < strtofind; k++ {
                        if s[i1+k] == toFind[k] { // si s = tofind, ajoute 1 au secondindex
                            secondread ++ 
                        } else {
                            return -1
                        }
                    }
                    if secondread == strtofind { // si secondindex = au nbre de str a trouver, finis la boucle
                        return i1
                    }
                } else { // sinon affiche -1
                    return -1
                }
            }
        }
        if secondread <= 0 { // si pas de secondindex, affiche -1
            return -1
        }
    }
    return strtofind
}

func NbretoFind(s string) int { // permet de savoir combien de str sont a trouver
    compteur := 0
    for i := range s {
        compteur = i + 1
    }
    return compteur
}