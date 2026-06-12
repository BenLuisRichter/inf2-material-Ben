package aufgabe4

import "slices"

// AUFGABENSTELLUNG: Vervollständigen Sie die vorgegebene Funktion.
// KONTEXT: In der Datei graph.go finden Sie eine Implementierung eines Knotens in einem
//          Graphen. Die zu implementierende Funktion ist eine Methode dieses Structs.
// MAX. PUNKTE: 10

// ReachableNodes soll eine Liste aller von n aus erreichbaren Knoten liefern.
func (n *Node) ReachableNodes() []*Node {
	alreadyVisited := []*Node{} //um Duplikate zu vermeiden, speichern wir hier alle bereits besuchten Knoten
	toVisit := []*Node{n}

	//Grundsätzlicher Ansatz einer Breitensuche
	//Eine Liste von Knoten pflegen, die noch abzuarbeiten sind
	//Die Suche endet, wenn diese Liste leer ist
	for len(toVisit) > 0 {
		first := toVisit[0]
		toVisit = toVisit[1:]
		if !slices.Contains(alreadyVisited, first) {
			alreadyVisited = append(alreadyVisited, first)
			toVisit = append(toVisit, first.neighbours...)
		}
	}

	return alreadyVisited
}

//Klausurrelevante Themen
//Algorithmen (Dijkstra, Floyd, A*)
//ob einer von den algorithmen sinnvoll für eine Aufgabe ist
//Hashmaps Funktionsweise und Komplexität

/* Tiefensuche, rekursiv
func (n *Node) ReachableNodes2() []*Node {
	visitedNodes := []*Node{}
	result := []*Node{n}
	x := 0

	for len(visitedNodes) != len(result) {
		for _, node := range n.ReachableIn(x) {
			//...
	return visitedNodes
}

func (n *Node) ReachableNodesIn(x int) []*Node {
	visitedNodes := []*Node{n}
	if x == 0 {
		return visitedNodes
	}

	rfn := []*Node{}
	for _, nb := range n.neighbours {
		rfn = append(rfn, nb.ReachableNodesIn(x-1)...)
	}
	rfn = append(rfn, n)

	return visitedNodes
}
*/
