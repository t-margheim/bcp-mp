package invitatory

import (
	"html/template"

	"github.com/t-margheim/bcp-mp/pkg/calendar"
)

var (
	easter = Entry{
		Name: `Christ our Passover`,
		Content: `<p>Alleluia. <br/>
		Christ our Passover has been sacrificed for us; * <br/>
		&nbsp;&nbsp;&nbsp;&nbsp;therefore let us keep the feast, <br/>
		Not with old leaven, the leaven of malice and evil, * <br/>
		&nbsp;&nbsp;&nbsp;&nbsp;but with the unleavened bread of sincerity and truth. Alleluia. </p>
		
		<p>Christ being raised from the dead will never die again; * <br/>
		&nbsp;&nbsp;&nbsp;&nbsp;death no longer has dominion over him. <br/>
		The death that he died, he died to sin, once for all; * <br/>
		&nbsp;&nbsp;&nbsp;&nbsp;but the life he lives, he lives to God. <br/>
		So also consider yourselves dead to sin, * <br/>
		&nbsp;&nbsp;&nbsp;&nbsp;and alive to God in Jesus Christ our Lord. Alleluia.</p>
		
		<p>Christ has been raised from the dead, * <br/>
		&nbsp;&nbsp;&nbsp;&nbsp;the first fruits of those who have fallen asleep. <br/>
		For since by a man came death, * <br/>
		&nbsp;&nbsp;&nbsp;&nbsp;by a man has come also the resurrection of the dead. <br/>
		For as in Adam all die, * <br/>
		&nbsp;&nbsp;&nbsp;&nbsp;so in Christ shall all be made alive. Alleluia.</p>
		`,
	}

	options = []Entry{
		{
			Name: "Venite",
			Content: `<p>Come, let us sing to the Lord; * <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;let us shout for joy to the Rock of our salvation. <br/>
			Let us come before his presence with thanksgiving * <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;and raise a loud shout to him with psalms. </p>
			
			<p>For the Lord is a great God, * <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;and a great King above all gods. <br/>
			In his hand are the caverns of the earth, * <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;and the heights of the hills are his also. <br/>
			The sea is his, for he made it, * <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;and his hands have molded the dry land.</p>
			
			<p>Come, let us bow down, and bend the knee, * <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;and kneel before the Lord our Maker. <br/>
			For he is our God, <br/>
			and we are the people of his pasture and the sheep of his hand. *<br/>
			&nbsp;&nbsp;&nbsp;&nbsp;Oh, that today you would hearken to his voice!</p>
			`,
		},
		{
			Name: "Jubilate",
			Content: `<p>Be joyful in the Lord, all you lands; * <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;serve the Lord with gladness <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;and come before his presence with a song. </p>
			<p>Know this: The Lord himself is God; * <br/>
&nbsp;&nbsp;&nbsp;&nbsp;he himself has made us, and we are his; <br/>
&nbsp;&nbsp;&nbsp;&nbsp;we are his people and the sheep of his pasture.</p>

<p>Enter his gates with thanksgiving; <br/>
go into his courts with praise; * <br/>
&nbsp;&nbsp;&nbsp;&nbsp;give thanks to him and call upon his Name. </p>

<p>For the Lord is good; <br/>
his mercy is everlasting; * <br/>
&nbsp;&nbsp;&nbsp;&nbsp;and his faithfulness endures from age to age. </p>
`,
		},
	}
)

func Get(keys calendar.KeyChain) Entry {
	if keys.Season == calendar.SeasonEaster {
		return easter
	}

	return options[keys.Iterator%len(options)]
}

type Entry struct {
	Name    string
	Content template.HTML
}
