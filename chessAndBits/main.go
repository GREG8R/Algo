package main

func main() {
	kingPath := "conditions/0.BITS/1.Bitboard_king/"
	horsePath := "conditions/0.BITS/2.Bitboard_knight/"
	fenPath := "conditions/0.BITS/3.Bitboard_FEN/"
	lineItemsPath := "conditions/0.BITS/4.Bitboard_lineItems/"

	RunTest(kingPath, King)
	RunTest(horsePath, Knight)
	RunTest(fenPath, FEN)
	RunTest(lineItemsPath, LineItems)
}
