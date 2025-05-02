package main 


//creando un type de posiciones para definir que funciones existen
type Movement interface {
    Arriba()
    Abajo()
    Izquierda()
    Derecha()
    House()
}

// creando el lugar donde almacenan cordenadas
type Coordenadas struct {
    x int
    y int
    houses int
}

// Creando las funciones que tiene que hacer 
func (c *Coordenadas) Arriba(){
    c.y += 1
}

// abajo 
func (c *Coordenadas ) Abajo(){
    c.y -= 1
}


// izquierda
func (c *Coordenadas) Izquierda(){
    c.x -= 1
}


// derecha
func (c *Coordenadas) Derecha(){
    c.x += 1
}
// house
func (c *Coordenadas) House(){
    c.houses += 1
}

//Mover arriba 
func MoverArriba(p Movement){
    p.Arriba()
}

//Mover abajo
func MoverAbajo(p Movement){
    p.Abajo()
}
// Mover Izquierda
func MoverIzquierda(p Movement){
    p.Izquierda()
}
//Mover Derecha
func MoverDerecha(p Movement){
    p.Derecha()
}
// Contar casas
func ContarCasas (p Movement){
    p.House()
}


