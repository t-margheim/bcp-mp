package prayers

var suffrages = []Prayer{
	{
		Name: "A",
		Content: `<p>
		V.&nbsp;&nbsp;&nbsp;&nbsp;Show us your mercy, O Lord; <br/>
		R.&nbsp;&nbsp;&nbsp;&nbsp;And grant us your salvation. <br/>
		V.&nbsp;&nbsp;&nbsp;&nbsp;Clothe your ministers with righteousness; <br/>
		R.&nbsp;&nbsp;&nbsp;&nbsp;Let your people sing with joy. <br/>
		V.&nbsp;&nbsp;&nbsp;&nbsp;Give peace, O Lord, in all the world; <br/>
		R.&nbsp;&nbsp;&nbsp;&nbsp;For only in you can we live in safety.<br/>
		V.&nbsp;&nbsp;&nbsp;&nbsp;Lord, keep this nation under your care; <br/>
		R.&nbsp;&nbsp;&nbsp;&nbsp;And guide us in the way of justice and truth. <br/>
		V.&nbsp;&nbsp;&nbsp;&nbsp;Let your way be known upon earth; <br/>
		R.&nbsp;&nbsp;&nbsp;&nbsp;Your saving health among all nations. <br/>
		V.&nbsp;&nbsp;&nbsp;&nbsp;Let not the needy, O Lord, be forgotten; <br/>
		R.&nbsp;&nbsp;&nbsp;&nbsp;Nor the hope of the poor be taken away. <br/>
		V.&nbsp;&nbsp;&nbsp;&nbsp;Create in us clean hearts, O God; <br/>
		R.&nbsp;&nbsp;&nbsp;&nbsp;And sustain us with your Holy Spirit. </p>
		`,
	},
	{
		Name: "B",
		Content: `<p>
		V.&nbsp;&nbsp;&nbsp;&nbsp;Save your people, Lord, and bless your inheritance; <br/>
		R.&nbsp;&nbsp;&nbsp;&nbsp;Govern them and uphold them, now and always.<br/>
		V.&nbsp;&nbsp;&nbsp;&nbsp;Day by day we bless you; <br/>
		R.&nbsp;&nbsp;&nbsp;&nbsp;We praise your name for ever. <br/>
		V.&nbsp;&nbsp;&nbsp;&nbsp;Lord, keep us from all sin today; <br/>
		R.&nbsp;&nbsp;&nbsp;&nbsp;Have mercy upon us, Lord, have mercy. <br/>
		V.&nbsp;&nbsp;&nbsp;&nbsp;Lord, show us your love and mercy; <br/>
		R.&nbsp;&nbsp;&nbsp;&nbsp;For we put our trust in you. <br/>
		V.&nbsp;&nbsp;&nbsp;&nbsp;In you, Lord, is our hope; <br/>
		R.&nbsp;&nbsp;&nbsp;&nbsp;And we shall never hope in vain. </p>
		`,
	},
}

func GetSuffrage(iterator int) Prayer {
	return suffrages[iterator%len(suffrages)]
}
