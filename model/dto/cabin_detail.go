package dto

type CabinDetail struct {
	ServiceClass 	string 	
	Deck			string
	FirstRowNumber	int8
	LastRowNumber	int8

	CabinLayout		[]string
}