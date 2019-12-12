package main 

func reportToPlayer(text string, f *faction) {
	plrname := "Unknown One"
	if f != nil {
		plrname = f.name
	}
	LOG.AppendMessage(plrname + ", " + text)
}
