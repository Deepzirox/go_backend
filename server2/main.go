// Cargar staticos router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
// JSOn response: json.NewEncoder(w).Encode(listofjson)
// Json req: _ = json.NewDecoder(req.Body).Decode(&tempPerson)
package main

import (
	KEY "./keyGenerator"
)

func main() {
	hashes := []uint{1, 2, 1, 2}
	println(KEY.KeyGen(hashes))

}
