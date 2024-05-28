package main

import (
   "bufio"
   "os"
    "fmt"
    )
    
func romnummer (x1 string) int {

    switch x1 {
	case "I":
 		return 1
	case "II":
		return 2
	case "III":
		return 3
	case "IV":
		return 4
	case "V":
	    return 5
    case "VI":
        return 6 
    case "VII":
        return 7
    case "VIII":
        return 8
    case "IX":
        return 9
    case "X":
        return 10
    default: return 101
    }
}
        
func nummerrom (result int) string {
    g:=0
    c := ""
    if result < 1 { panic("Результат меньше еденицы!") }
    
    if result / 100 > 0 { 
        c = c + "C"
        result = result%100 }
        
    if result / 50 > 0 {
        c = c + "L"
        result = result%50
    }
    
    if result / 10 > 0 {
        for g=1; g <= result / 10; g++ {
        c = c + "X"
        }
        result = result%10
    }
    
    if result / 5 > 0 {
        c = c + "V"
        result = result - 5
    }
   
    if result > 0 { 
        for g=1; g <= result; g++ {
        c = c + "I" 
            }
        }
        
    return c
}
    
func arabic (count int, s string) {

    j:=0
    a:=0 
    x:=0 
    y:=0 
    o:=""
    result :=0
    
    for j<count {
        
        if s[j] >= 48 && s[j] <= 57 && j < 2 { 
            x = x*10 + int(s[j] - 48)
        } else {
            if s[j] >= 48 && s[j] <= 57 && j > 1 {
                y = y*10 + int(s[j] - 48)
            }
	        if (string(s[j]) == "+" || string(s[j]) == "-" || string(s[j]) == "*" || string(s[j]) == "/") {
	            o = string(s[j])
	            a++
	            
	        } else  {
	            
	            if string(s[j]) == "I" || string(s[j]) == "V" || string(s[j]) == "X" { panic("Используются одновременно разные системы счисления.")
	            
	        } 
        }
    } 
    j++
}

    if x < 0 || x >10 || y < 0 || y > 10 { panic("Неккоректный ввод данных") 
    }
    
    if a > 1 { panic("Формат математической операции не удовлетворяет заданию ") 
    }
    
    result = ( calc (x, y, o))
    
    if result <= 100 { fmt.Print(result)
    } else {
   
   return
 }
}


func rom (count int, s string) {

    k:=0
    a:=0
    b:=0
    o:=""
    x1:=""
    y1:=""
    
    for k<count {
        
        if ( string(s[k]) == "X" || string(s[k]) == "I" || string(s[k]) == "V")  && b == 0 { 
            x1 = x1 + string(s[k])
        } else {
            if ( string(s[k]) == "X" || string(s[k]) == "I" || string(s[k]) == "V") && b >= 1 { 
            y1 = y1 + string(s[k])
        } else {
	        if (string(s[k]) == "+" || string(s[k]) == "-" || string(s[k]) == "*" || string(s[k]) == "/") {
	            o = string(s[k])
	            a++
	            
	        } else {
	            if string(s[k]) == " " { b++ 
	
	        } else {
	            if s[k] >= 48 && s[k] <= 57 { panic("Используются одновременно разные системы счисления.")
	        }
	       }
        }
    } 
} 
    k++
}

    if a > 1 { panic("Формат математической операции не удовлетворяет заданию.") 
    }
    
    x := (romnummer(x1))
    y := (romnummer(y1))
    
   if x > 10 || y > 10 { panic("Некорректный ввод данных.")
    }
    
    
    result := ( calc (x, y, o))
    
    if result < 0 { panic("В римской системе нет отрицательных чисел.")
        
        } else {
            if result == 101 { panic("Некорректный ввод данных.")
        } else { 
            fmt.Println( nummerrom(result))
            
        }
    }
}


 func calc (x, y int, o string) int {

	switch o {
	case "+":
 		return (x + y)
	case "-":
		return (x - y)
	case "*":
		return (x * y)
	case "/":
		if y == 0 {
			panic("На ноль делить нельзя!")
		}	
			return (x / y)
		}
	return 101
	
}

func main() {
	var x, y int
	var s, o string
	
    scanner := bufio.NewScanner(os.Stdin)
        _ = scanner.Scan()
    	s = scanner.Text() 
	
	var count int = len(s) 
	
	if count < 3 { panic( "Строка не является математической операцией.")
    } 
	
	
    if string(s[0]) == "I" || string(s[0]) == "V" || string(s[0]) == "X" {

                rom(count, s)
                
                return
            } else {
                    if s[0] > 48 && s[0] <= 57 { 
                        arabic(count, s)
                        return
                } else {
                        panic("Неккоректный ввод данных")
            }
    
}
	        fmt.Println( calc (x, y, o))
    
}
