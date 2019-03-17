package closings

var closings = []string{
	"The grace of our Lord Jesus Christ, and the love of God, and the fellowship of the Holy Spirit, be with us all evermore.",
	"May the God of hope fill us with all joy and peace in believing through the power of the Holy Spirit.",
	"Glory to God whose power, working in us, can do infinitely more than we can ask or imagine: Glory to him from generation to generation in the Church, and in Christ Jesus for ever and ever. ",
}

func Get(iterator int) string {
	return closings[iterator%len(closings)]
}
