package main

const (
	templateString = `<h1>Morning Prayer for {{.Date}}</h1>
	<h2>{{.Title}}</h2>
	<h2>Opening Sentence</h2>
	<p>{{.Opening.Text}}</p>
	
	<p>{{.Opening.Citation}}</p>
	
	<h2>Confession of Sin</h2>
	
	<p>
		Let us confess our sins against God and our neighbor.
	</p>
	
	<p>
		Most merciful God,<br>
		we confess that we have sinned against you<br>
		in thought, word, and deed,<br>
		by what we have done,<br>
		and by what we have left undone.<br>
		We have not loved you with our whole heart;<br>
		we have not loved our neighbors as ourselves.<br>
		We are truly sorry and we humbly repent.<br>
		For the sake of your Son Jesus Christ,<br>
		have mercy on us and forgive us;<br>
		that we may delight in your will,<br>
		and walk in your ways,<br>
		to the glory of your Name. Amen.<br>
	</p>
	
	<p>
		Almighty God have mercy on you, forgive you all your sins<br>
		through our Lord Jesus Christ, strengthen you in all<br>
		goodness, and by the power of the Holy Spirit keep you in<br>
		eternal life. Amen.<br>
	</p>
	
	<h2>The Invitatory and the Psalter</h2>
	
	Lord, open our lips.<br>
	And our mouth shall proclaim your praise.<br>
	
	Glory to the Father, and to the Son, and to the Holy Spirit:<br>
	as it was in the beginning, is now, and will be for ever. Amen.<br>
	
	<h3>{{.Invitatory.Name}}</h3>
	
	<p>
		{{.Invitatory.Content}}
	</p>
	
	<h3>{{.Psalms.Reference}}</h3>
	
	
	
	{{.Psalms.Body}}
	
	Glory to the Father, and to the Son, and to the Holy Spirit: *<br>
	as it was in the beginning, is now, and will be for ever. Amen.<br>
	
	<h2>The Lessons</h2>
	
	<h3>{{.Lesson1.Reference}}</h3>
	<p>
		{{.Lesson1.Body}}
	</p>
	
	<h3>{{.Canticle1.EnglishTitle}}</h3>
	<em>{{.Canticle1.LatinTitle}}</em><br />
	{{.Canticle1.Content}}
	
	
	<h3>{{.Lesson2.Reference}}</h3>
	<p>
		{{.Lesson2.Body}}
	</p>
	
	<h3>{{.Canticle2.EnglishTitle}}</h3>
	<em>{{.Canticle2.LatinTitle}}</em><br />
	{{.Canticle2.Content}}
	
	<h3>{{.Gospel.Reference}}</h3>
	<p>
		{{.Gospel.Body}}
	</p>
	
	
	<h2>The Apostles' Creed</h2>
	
	<p>
		<p>I believe in God, the Father almighty, <br />
			&nbsp;&nbsp;&nbsp;&nbsp;creator of heaven and earth; <br />
			I believe in Jesus Christ, his only Son, our Lord. <br />
			&nbsp;&nbsp;&nbsp;&nbsp;He was conceived by the power of the Holy Spirit <br />
			&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;and born of the Virgin Mary. <br />
			&nbsp;&nbsp;&nbsp;&nbsp;He suffered under Pontius Pilate, <br />
			&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;was crucified, died, and was buried. <br />
			&nbsp;&nbsp;&nbsp;&nbsp;He descended to the dead. <br />
			&nbsp;&nbsp;&nbsp;&nbsp;On the third day he rose again. <br />
			&nbsp;&nbsp;&nbsp;&nbsp;He ascended into heaven, <br />
			&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;and is seated at the right hand of the Father. <br />
			&nbsp;&nbsp;&nbsp;&nbsp;He will come again to judge the living and the dead. <br />
			I believe in the Holy Spirit, <br />
			&nbsp;&nbsp;&nbsp;&nbsp;the holy catholic Church, <br />
			&nbsp;&nbsp;&nbsp;&nbsp;the communion of saints, <br />
			&nbsp;&nbsp;&nbsp;&nbsp;the forgiveness of sins<br />
			&nbsp;&nbsp;&nbsp;&nbsp;the resurrection of the body, <br />
			&nbsp;&nbsp;&nbsp;&nbsp;and the life everlasting. Amen. </p>
	
	</p>
	
	<h2>The Prayers</h2>
	
	<p>
		Our Father in heaven, <br />
		&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;hallowed be your Name, <br />
		&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;your kingdom come, <br />
		&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;your will be done, <br />
		&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;on earth as in
		heaven. <br />
		Give us today our daily bread. <br />
		Forgive us our sins <br />
		&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;as we forgive those <br />
		&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;who sin against us.
		<br />
		Save us from the time of trial, <br />
		&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;and deliver us from evil. <br />
		For the kingdom, the power, <br />
		&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;and the glory are yours, <br />
		&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;now and for ever. Amen.
	</p>
	
	<h3>Suffrage {{.Suffrage.Name}}</h3>
	
	<p>{{.Suffrage.Content}}</p>
	
	<h3>{{.Collect.Name}}</h3>
	
	<p>{{.Collect.Content}}</p>
	
	<h3>For Mission</h3>
	
	<p>{{.Mission.Content}}</p>
	
	{{range .Prayers}}
	<h3>{{.Name}}</h3>
	<p>{{.Content}}</p>
	{{end}}
	
	<h3>The General Thanksgiving</h3>
	
	<p>
		Almighty God, Father of all mercies, <br />
		we your unworthy servants give you humble thanks <br />
		for all your goodness and loving-kindness <br />
		to us and to all whom you have made. <br />
		We bless you for our creation, preservation, <br />
		and all the blessings of this life; <br />
		but above all for your immeasurable love <br />
		in the redemption of the world by our Lord Jesus Christ; <br />
		for the means of grace, and for the hope of glory. <br />
		And, we pray, give us such an awareness of your mercies, <br />
		that with truly thankful hearts we may show forth your praise, <br />
		not only with our lips, but in our lives, <br />
		by giving up our selves to your service, <br />
		and by walking before you <br />
		in holiness and righteousness all our days; <br />
		through Jesus Christ our Lord, <br />
		to whom, with you and the Holy Spirit, <br />
		be honor and glory throughout all ages. <em>Amen.</em>
	</p>
	
	<h3>A Prayer of St Chrysostom</h3>
	
	<p>
		Almighty God, you have given us grace at this time with one
		accord to make our common supplication to you; and you
		have promised through your well-beloved Son that when two
		or three are gathered together in his Name you will be in the
		midst of them: Fulfill now, O Lord, our desires and petitions
		as may be best for us; granting us in this world knowledge of
		your truth, and in the age to come life everlasting. <em>Amen.</em>
	</p>
	
	<p>
		<em>
			{{.Closing}}
		</em>
	</p>`
)
