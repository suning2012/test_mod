package testmod

import (
    "errors"
    "fmt" 
) 

// Hi returns a friendly greeting in language lang
func Hi(name, lang string) (string, error) {
    switch lang {
    case "en":
        return fmt.Sprintf("en 1111111111   Hi, %s!", name), nil
    case "pt":
        return fmt.Sprintf("pt 222222222    Oi, %s!", name), nil
    case "es":
        return fmt.Sprintf("es 333333333    ¡Hola, %s!", name), nil
    case "fr":
        return fmt.Sprintf("fr 444444444    Bonjour, %s!", name), nil
    default:
        return "", errors.New("unknown language")
    }
}

