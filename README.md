# word-search
Determine whether a word is in a wordsearh or not

## Ejercicio 3: Alphabet Soup

Para realizar este ejercicio se debe modificar el archivo WordSearcher.java del paquete com.baufest.ingreso.alphabetSoup .

### Objetivo

El objetivo de este ejercicio es implementar una función que determine si una palabra está en una sopa de letras.

### Reglas
- Las palabras pueden estar dispuestas direcciones horizontal o vertical, _no_ en diagonal.
- Las palabras pueden estar orientadas en cualquier sentido, esto es, de derecha a izquierda o viceversa, y de arriba
para abajo o viceversa.
- El cambio de dirección puede estar a media palabra, de modo que, por ejemplo, parte de la palabra
esté horizontal y de izquierda a derecha, parte esté vertical y de arriba hacia abajo, y otra parte horizontal
de derecha a la izquierda.

A continuación se muestran un par de ejemplos:

**Horizontal**

X X X X X X X  
P A L A B R A  
X X X X X X X   
X X X X X X X   
X X X X X X X

**Vertical**

A X X X X X X  
R X X X X X X   
B X X X X X X  
A X X X X X X   
L X X X X X X  
A X X X X X X  
P X X X X X X  

**Horizontal y Vertical**

X X P X X X X X    
X X A X X X X X    
X X L X X X X X   
X X A B R A X X    
X X X X X X X X   

X X X X X X X X  
X A R B X X X X  
X X X A X X X X   
X X X L A P X X  
X X X X X X X X   
