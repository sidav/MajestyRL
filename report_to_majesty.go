package main 

func reportToPlayer(text string, f *faction) {
	plrname := "Unknown One"
	if f != nil {
		plrname = f.name
	}
	log.AppendMessage(plrname + ", " + text)
}
