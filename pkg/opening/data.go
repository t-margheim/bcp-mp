package opening

import "github.com/t-margheim/bcp-mp/internal/calendar"

var (
	Openings = map[calendar.Key][]Opening{
		calendar.SeasonAdvent: {
			{
				Text:     "Watch, for you know not when the master of the house will come, in the evening, or at midnight, or at cockcrow, or in the morning; lest he come suddenly and find you asleep.",
				Citation: "Mark 13:35, 36",
			},
			{
				Text:     "In the wilderness prepare the way of the Lord, make straight in the desert a highway for our God.",
				Citation: "Isaiah 40:3",
			},
			{
				Text:     "The glory of the Lord shall be revealed, and all flesh shall see it together.",
				Citation: "Isaiah 40:5",
			},
		},
		calendar.SeasonChristmas: {
			{
				Text:     "Behold, I bring you good news of a great joy, which will come to all the people; for unto you is born this day in the city of David, a Savior, who is Christ the Lord.",
				Citation: "Luke 2:10, 11",
			},
			{
				Text:     "Behold, the dwelling of God is with mankind. He will dwell with them, and they shall be his people, and God himself will be with them, and be their God.",
				Citation: "Revelation 21:3",
			},
		},
		calendar.SeasonEaster: {
			{
				Text:     "On this day the Lord has acted; we will rejoice and be glad in it.",
				Citation: "Psalm 118:24",
			},
			{
				Text:     "Thanks be to God, who gives us the victory through our Lord Jesus Christ.",
				Citation: "1 Corinthians 15:57",
			},
			{
				Text:     "If then you have been raised with Christ, seek the things that are above, where Christ is, seated at the right hand of God.",
				Citation: "Colossians 3:1",
			},
			{
				Text:     "Christ has entered, not into a sanctuary made with hands, a copy of the true one, but into heaven itself, now to appear in the presence of God on our behalf.",
				Citation: "Hebrews 9:24",
			},
			{
				Text:     "You shall receive power when the Holy Ghost has come upon you; and you shall be my witness in Jerusalem, and in all Judea, and Samaria, and to the ends of the earth.",
				Citation: "Acts 1:8",
			},
		},
		calendar.SeasonEpiphany: {
			{
				Text:     "Nations shall come to your light, and kings to the brightness of your rising.",
				Citation: "Isaiah 60: 3",
			},
			{
				Text:     "I will give you as a light to the nations, that my salvation may reach to the end of the earth.",
				Citation: "Isaiah 49: 6b",
			},
			{
				Text:     "From the rising of the sun to its setting my Name shall be great among the nations, and in every place incense shall be offered to my Name, and a pure offering: for my Name shall be great among the nations, says the Lord of hosts.",
				Citation: "Malachi 1: 11",
			},
		},
		calendar.SeasonHolyWeek: {
			{
				Text:     "All we like sheep have gone astray; we have turned every one to his own way; and the Lord has laid on him the iniquity of us all.",
				Citation: "Isaiah 53:6",
			},
			{
				Text:     "Is it nothing to you, all you who pass by? Look and see if there is any sorrow like my sorrow which was brought upon me, whom the Lord hath afflicted.",
				Citation: "Lamentations 1:12",
			},
		},
		calendar.SeasonLent: {

			{
				Text:     "If we say we have no sin, we deceive ourselves, and the truth is not in us; but if we confess our sins, God who is faithful and just, will forgive our sins and cleanse us from all unrighteousness.",
				Citation: "1 John 1:8,9",
			},
			{
				Text:     "Rend your hearts and not your garments. Return to the Lord your God, for he is gracious and merciful, slow to anger and abounding in steadfast love, and repents of evil.",
				Citation: "Joel 2:13",
			},
			{
				Text:     "I will arise and go to my father, and I will say to him, \"Father, I have sinned against heaven, and before you; I am no longer worthy to be called your son.\"",
				Citation: "Luke 15:18,19",
			},
			{
				Text:     "To the Lord our God belong mercy and forgiveness, because we have rebelled against him and have not obeyed the voice of the Lord our God by following his laws which he set before us.",
				Citation: "Daniel 9:9,10",
			},
			{
				Text:     "Jesus said, \"If anyone will come after me, let him deny himself, and take up his cross, and follow me.\"",
				Citation: "Mark 8:34",
			},
		},
		calendar.SeasonOrdinary: {
			{
				Text:     "Grace to you and peace from God our Father and from the Lord Jesus Christ.",
				Citation: "Philippians 1:2",
			},
			{
				Text:     "I was glad when they said to me, \"Let us go to the house of the Lord.\"",
				Citation: "Psalm 122:1",
			},
			{
				Text:     "Let the words of my mouth and the meditation of my heart be acceptable in your sight, O Lord, my strength and my redeemer.",
				Citation: "Psalm 19:14",
			},
			{
				Text:     "Send out your light and your truth, that they may lead me, and bring me to your holy hill and to your dwelling.",
				Citation: "Psalm 43:3",
			},
			{
				Text:     "The Lord is in his holy temple; let all the earth keep silence before him.",
				Citation: "Habakkuk 2:20",
			},
			{
				Text:     "The hour is coming, and now is, when the true worshipers will worship the Father in spirit and in truth, for such the Father seeks to worship him.",
				Citation: "John 4:23",
			},
			{
				Text:     "Isaiah 57:15",
				Citation: "Thus says the high and lofty One who inhabits eternity, whose name is Holy, \"I dwell in the high and holy place and also with the one who has a contrite and humble spirit, to revive the spirit of the humble and to revive the heart of the contrite.\"",
			},
		},
		calendar.OpenAllSaints: {
			{
				Text:     "We give thanks to the Father, who has made us worthy to share in the inheritance of the saints in light.",
				Citation: "Colossians 1:12",
			},
			{
				Text:     "You are no longer strangers and sojourners, but fellow citizens with the saints and members of the household of God.",
				Citation: "Ephesians 2:19",
			},
			{
				Text:     "Their sound has gone out into all lands, and their message to the ends of the world.",
				Citation: "Psalm 19:4",
			},
		},
		calendar.OpenTrinitySunday: {
			{
				Text:     "Holy, holy, holy is the Lord God Almighty, who was, and is, and is to come.",
				Citation: "Revelation 4:8",
			},
		},
	}
)
